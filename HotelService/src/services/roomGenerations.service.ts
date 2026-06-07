import { RoomGenerationJobSchemaType } from "../dtos/roomGeneration.dto";
import RoomTypeRepository from "../repositories/roomType.repository";
import { NotFoundError, BadRequestError } from "../utils/errors/app.error";
import RoomRepository from "../repositories/room.repository";
import RoomType from "../db/models/roomTypes";
import logger from "../config/logger.config";

const roomTypeRepository = new RoomTypeRepository();
const roomRepository = new RoomRepository();

const FLOORS_PER_CATEGORY = 10;
const ROOMS_PER_FLOOR_PER_CATEGORY = 10;

function addDays(date: Date, days: number) {
  const nextDate = new Date(date);
  nextDate.setDate(nextDate.getDate() + days);
  return nextDate;
}

function buildRoomNumber(floor: number, roomOnFloor: number) {
  return `${floor}${roomOnFloor.toString().padStart(2, "0")}`;
}

export async function generateRoomsService(
  jobData: RoomGenerationJobSchemaType,
) {
  const { roomTypeId, hotelId, startDate, endDate, priceOverride } = jobData;
  const batchSize = jobData.batchSize ?? 100;

  const startDateTime = new Date(startDate);
  const endDateTime = new Date(endDate);

  const roomType = await roomTypeRepository.findById(roomTypeId);

  if (!roomType) {
    throw new NotFoundError(`Room Type with id ${roomTypeId} not found`);
  }

  if (roomType.hotel_id !== hotelId) {
    throw new BadRequestError(
      `Room Type with id ${roomTypeId} does not belong to hotel ${hotelId}`,
    );
  }

  if (startDateTime > endDateTime) {
    throw new BadRequestError("Start date must be before end date");
  }

  if (startDateTime < new Date()) {
    throw new BadRequestError("Start date must be in the future");
  }
  const totalDays = Math.floor(
    (endDateTime.getTime() - startDateTime.getTime()) / (1000 * 3600 * 24),
  ) + 1;
  if (totalDays > 180) {
    throw new BadRequestError(
      "Cannot generate rooms for more than 180 days at a time",
    );
  }
  let totalRoomsToGenerate = 0;
  const roomsPerDate = FLOORS_PER_CATEGORY * ROOMS_PER_FLOOR_PER_CATEGORY;
  const datesPerBatch = Math.max(1, Math.floor(batchSize / roomsPerDate));
  const currentDate = new Date(startDateTime);

  logger.info(
    `Starting room generation for Room Type ID: ${roomTypeId} in Hotel ID: ${hotelId} from ${startDateTime.toISOString()} to ${endDateTime.toISOString()} with price override: ${priceOverride ?? "none"}`,
  );

  while (currentDate <= endDateTime) {
    const batchStartDate = new Date(currentDate);
    const batchEndDate = addDays(batchStartDate, datesPerBatch - 1);

    if (batchEndDate > endDateTime) {
      batchEndDate.setTime(endDateTime.getTime());
    }
    logger.info(
      `Checking existing rooms for batch from ${batchStartDate.toISOString()} to ${batchEndDate.toISOString()}`,
    );
    const { roomsGenerated } = await processDateBatch(
      roomType,
      hotelId,
      batchStartDate,
      batchEndDate,
      priceOverride,
    );
    currentDate.setTime(addDays(batchEndDate, 1).getTime());
    totalRoomsToGenerate += roomsGenerated;
  }

  return {
    roomsGenerated: totalRoomsToGenerate,
    datesGenerated: totalDays,
  };
}

export async function processDateBatch(
  roomType: RoomType,
  hotelId: number,
  startDate: Date,
  endDate: Date,
  priceOverride?: number,
) {
  let roomsGenerated = 0;
  let datesGenerated = 0;

  const existingRooms = await roomRepository.findExistingRooms(
    roomType.id,
    hotelId,
    startDate,
    endDate,
  );

  const existingRoomSet = new Set(
    existingRooms.map(
      (room: any) =>
        `${room.room_number}|${
          new Date(room.statusDate).toISOString().split("T")[0]
        }`,
    ),
  );

  const roomsToCreate: any[] = [];
  let currentDate = new Date(startDate);

  while (currentDate <= endDate) {
    const dateKey = currentDate.toISOString().split("T")[0];

    for (let floor = 1; floor <= FLOORS_PER_CATEGORY; floor++) {
      for (
        let roomOnFloor = 1;
        roomOnFloor <= ROOMS_PER_FLOOR_PER_CATEGORY;
        roomOnFloor++
      ) {
        const roomNumber = buildRoomNumber(floor, roomOnFloor);
        const roomKey = `${roomNumber}|${dateKey}`;

        if (!existingRoomSet.has(roomKey)) {
          roomsToCreate.push({
            hotel_id: hotelId,
            room_type_id: roomType.id,
            room_number: roomNumber,
            floor,
            status: "available",
            statusDate: new Date(currentDate),
            price: priceOverride ?? Number(roomType.base_price),
          });

          roomsGenerated++;
        }
      }
    }

    datesGenerated++;
    currentDate = addDays(currentDate, 1);
  }

  if (roomsToCreate.length > 0) {
    await roomRepository.bulkCreateRooms(roomsToCreate);
  }

  return {
    roomsGenerated,
    datesGenerated,
  };
}
