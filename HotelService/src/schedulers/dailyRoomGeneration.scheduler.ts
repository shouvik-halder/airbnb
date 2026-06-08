import cron from "node-cron";
import { v4 as uuidv4 } from "uuid";
import logger from "../config/logger.config";
import { runWithCorrelationId } from "../helpers/request.helper";
import { ROOMS_GENERATION_PAYLOAD } from "../producers/roomsGeneration.producer";
import { roomsGenerationQueue } from "../queues/roomsGeneration.queue";
import HotelRepository from "../repositories/hotel.repository";
import RoomRepository from "../repositories/room.repository";
import RoomTypeRepository from "../repositories/roomType.repository";
import { schedulerConfig } from "../config";

const DAILY_ROOM_GENERATION_CRON =
  schedulerConfig.DAILY_ROOM_GENERATION_CRON;

const DAILY_ROOM_GENERATION_TIMEZONE =
  schedulerConfig.ROOM_GENERATION_CRON_TIMEZONE;

const hotelRepository = new HotelRepository();
const roomRepository = new RoomRepository();
const roomTypeRepository = new RoomTypeRepository();

function getUTCDateKey(date: Date) {
  return date.toISOString().split("T")[0];
}

function getNextUTCDate() {
  const now = new Date();

  return new Date(
    Date.UTC(
      now.getUTCFullYear(),
      now.getUTCMonth(),
      now.getUTCDate() + 1,
    ),
  );
}

function addUTCDateDays(date: Date, days: number) {
  return new Date(
    Date.UTC(
      date.getUTCFullYear(),
      date.getUTCMonth(),
      date.getUTCDate() + days,
    ),
  );
}

export async function enqueueDailyRoomGeneration() {
  const hotels = await hotelRepository.findAll();

  let jobsQueued = 0;

  logger.info(
    `Starting daily room generation scheduler for ${hotels.length} hotels`,
  );

  for (const hotel of hotels) {
    const roomTypes = await roomTypeRepository.findActiveByHotelId(
      hotel.id,
    );

    if (roomTypes.length === 0) {
      logger.info(
        `Skipping hotel ${hotel.id} because no room types were found`,
      );
      continue;
    }

    for (const roomType of roomTypes) {
      const latestStatusDate =
        await roomRepository.findLatestStatusDate(
          roomType.id,
          hotel.id,
        );

      let generationDate: Date;

      if (!latestStatusDate) {
        generationDate = getNextUTCDate();

        logger.info(
          `No availability found for hotel ${hotel.id} room type ${roomType.id}. Generating inventory for ${getUTCDateKey(
            generationDate,
          )}`,
        );
      } else {
        generationDate = addUTCDateDays(
          new Date(latestStatusDate),
          1,
        );

        logger.info(
          `Latest availability for hotel ${hotel.id} room type ${
            roomType.id
          } is ${getUTCDateKey(
            new Date(latestStatusDate),
          )}. Generating inventory for ${getUTCDateKey(
            generationDate,
          )}`,
        );
      }

      const generationDateKey =
        getUTCDateKey(generationDate);

      const jobId =
        `daily-room-generation-${hotel.id}-${roomType.id}-${generationDateKey}`;

      await roomsGenerationQueue.add(
        ROOMS_GENERATION_PAYLOAD,
        {
          roomTypeId: roomType.id,
          hotelId: hotel.id,
          startDate: generationDate.toISOString(),
          endDate: generationDate.toISOString(),
          batchSize: 100,
          correlationId: uuidv4(),
        },
        {
          jobId,
          attempts: 3,
          backoff: {
            type: "exponential",
            delay: 5000,
          },
          removeOnComplete: true,
          removeOnFail: false,
        },
      );

      jobsQueued++;

      logger.info(
        `Queued room generation job ${jobId} for hotel ${hotel.id} room type ${roomType.id}`,
      );
    }
  }

  logger.info(
    `Daily room generation scheduler completed. Total jobs queued: ${jobsQueued}`,
  );
}

export function setupDailyRoomGenerationScheduler() {
  const task = cron.schedule(
    DAILY_ROOM_GENERATION_CRON,
    () => {
      runWithCorrelationId(undefined, async () => {
        try {
          await enqueueDailyRoomGeneration();
        } catch (error) {
          logger.error(
            `Daily room generation scheduler failed: ${
              (error as Error).message
            }`,
          );
        }
      });
    },
    {
      timezone: DAILY_ROOM_GENERATION_TIMEZONE,
    },
  );

  logger.info(
    `Daily room generation scheduler started with cron "${DAILY_ROOM_GENERATION_CRON}" in timezone "${DAILY_ROOM_GENERATION_TIMEZONE}"`,
  );

  return task;
}