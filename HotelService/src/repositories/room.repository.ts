import { Op } from "sequelize";
import Room, { RoomStatus } from "../db/models/room";
import { BaseRepository } from "./base.repository";
import { sequelize } from "../db/models/sequelize";

class RoomRepository extends BaseRepository<Room> {
  constructor() {
    super(Room);
  }

  async findByRoomTypeIdAndDate(
    room_type_id: number,
    hotel_id: number,
    date: Date,
  ) {
    return await this.model.findOne({
      where: {
        room_type_id,
        hotel_id,
        status: RoomStatus.available,
        deletedAt: null,
        statusDate: date,
      },
    });
  }

  async findByRoomAssignment(
    room_type_id: number,
    hotel_id: number,
    room_number: string,
    statusDate: Date,
  ) {
    return await this.model.findOne({
      where: {
        room_type_id,
        hotel_id,
        room_number,
        statusDate,
        deletedAt: null,
      },
    });
  }

  async bulkCreateRooms(rooms: Array<Partial<Room>>) {
    return await this.model.bulkCreate(rooms as any[]);
  }

  async findLatestStatusDate(
    roomTypeId: number,
    hotelId: number,
  ): Promise<Date | null> {
    const latestStatusDate = await this.model.max("statusDate", {
      where: {
        room_type_id: roomTypeId,
        hotel_id: hotelId,
        deletedAt: null,
      },
    });

    return latestStatusDate
      ? new Date(latestStatusDate as string | number | Date)
      : null;
  }

  async findExistingRooms(
    roomTypeId: number,
    hotelId: number,
    startDate: Date,
    endDate: Date,
  ) {
    return this.model.findAll({
      where: {
        room_type_id: roomTypeId,
        hotel_id: hotelId,
        statusDate: {
          [Op.between]: [startDate, endDate],
        },
      },
      attributes: ["room_number", "statusDate"],
    });
  }

  async findAvailableRoomsByTypeAndDateRange(
    roomTypeId: number,
    hotelId: number,
    checkIn: Date,
    checkOut: Date,
    totalDays: number,
    transaction?: any,
  ) {
    return await this.model.findAll({
      attributes: ["room_number"],
      where: {
        room_type_id: roomTypeId,
        hotel_id: hotelId,
        status: RoomStatus.available,
        booking_id: null,
        statusDate: {
          [Op.gte]: checkIn,
          [Op.lt]: checkOut,
        },
      },
      group: ["room_number"],
      having: sequelize.literal(`COUNT(*) = ${totalDays}`),
      ...(transaction && { transaction }),
    });
  }

async updateBookingIdToBookedRooms(
  bookingId: number,
  roomTypeId: number,
  roomNumber:string,
  hotelId:number,
  checkIn: Date,
  checkOut: Date
) {
  const [updatedRows] = await this.model.update(
    {
      booking_id: bookingId,
      status: RoomStatus.booked
    },
    {
      where: {
        hotel_id:hotelId,
        room_type_id: roomTypeId,
        room_number:roomNumber,
        booking_id: null,
        status: RoomStatus.available,
        statusDate: {
          [Op.gte]: checkIn,
          [Op.lt]: checkOut
        }
      }
    }
  );

  return updatedRows;
}

// Add method with explicit transaction and row-level locking
async updateBookingIdToBookedRoomsWithLock(
  transaction: any,
  bookingId: number,
  roomTypeId: number,
  roomNumber: string,
  hotelId: number,
  checkIn: Date,
  checkOut: Date
) {
  const [updatedRows] = await this.model.update(
    {
      booking_id: bookingId,
      status: RoomStatus.booked
    },
    {
      where: {
        hotel_id: hotelId,
        room_type_id: roomTypeId,
        room_number: roomNumber,
        booking_id: null,
        status: RoomStatus.available,
        statusDate: {
          [Op.gte]: checkIn,
          [Op.lt]: checkOut
        }
      },
      transaction
    }
  );

  return updatedRows;
}
}

export default RoomRepository;
