import cron from "node-cron";
import {v4 as uuidv4} from "uuid";
import logger from "../config/logger.config";
import { runWithCorrelationId } from "../helpers/request.helper";
import { ROOMS_GENERATION_PAYLOAD } from "../producers/roomsGeneration.producer";
import { roomsGenerationQueue } from "../queues/roomsGeneration.queue";
import HotelRepository from "../repositories/hotel.repository";
import RoomRepository from "../repositories/room.repository";
import RoomTypeRepository from "../repositories/roomType.repository";

const DAILY_ROOM_GENERATION_CRON = "0 0 * * *";
const DAILY_ROOM_GENERATION_TIMEZONE =
  process.env.ROOM_GENERATION_CRON_TIMEZONE || "UTC";

const hotelRepository = new HotelRepository();
const roomRepository = new RoomRepository();
const roomTypeRepository = new RoomTypeRepository();

function getUTCDateKey(date: Date) {
  return date.toISOString().split("T")[0];
}

function getNextUTCDateISOString(date = new Date()) {
  return new Date(
    Date.UTC(
      date.getUTCFullYear(),
      date.getUTCMonth(),
      date.getUTCDate() + 1,
    ),
  ).toISOString();
}

export async function enqueueDailyRoomGeneration() {
  const targetDate = getNextUTCDateISOString();
  const targetDateKey = getUTCDateKey(new Date(targetDate));
  const hotels = await hotelRepository.findAll();
  let jobsQueued = 0;
  let categoriesSkipped = 0;

  logger.info(
    `Starting daily room generation scheduler for target date ${targetDateKey}`,
  );

  for (const hotel of hotels) {
    const roomTypes = await roomTypeRepository.findActiveByHotelId(hotel.id);

    if (roomTypes.length === 0) {
      logger.info(
        `Skipping daily room generation for hotel ${hotel.id} because no room types were found`,
      );
      continue;
    }

    for (const roomType of roomTypes) {
      const latestStatusDate = await roomRepository.findLatestStatusDate(
        roomType.id,
        hotel.id,
      );
      const latestStatusDateKey = latestStatusDate
        ? getUTCDateKey(new Date(latestStatusDate))
        : null;

      if (latestStatusDateKey && latestStatusDateKey >= targetDateKey) {
        categoriesSkipped++;
        logger.info(
          `Skipping daily room generation for hotel ${hotel.id} and room type ${roomType.id}; latest room date is ${latestStatusDateKey}`,
        );
        continue;
      }

      const jobId = `daily-room-generation-${hotel.id}-${roomType.id}-${targetDateKey}`;

      await roomsGenerationQueue.add(
        ROOMS_GENERATION_PAYLOAD,
        {
          roomTypeId: roomType.id,
          hotelId: hotel.id,
          startDate: targetDate,
          endDate: targetDate,
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
        `Queued daily room generation job ${jobId} for hotel ${hotel.id} and room type ${roomType.id}`,
      );
    }
  }

  logger.info(
    `Daily room generation scheduler queued ${jobsQueued} jobs and skipped ${categoriesSkipped} room types for target date ${targetDateKey}`,
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
            `Daily room generation scheduler failed: ${(error as Error).message}`,
          );
        }
      });
    },
    {
      timezone: DAILY_ROOM_GENERATION_TIMEZONE,
    },
  );

  logger.info(
    `Daily room generation scheduler has been started with cron ${DAILY_ROOM_GENERATION_CRON} in timezone ${DAILY_ROOM_GENERATION_TIMEZONE}`,
  );

  return task;
}
