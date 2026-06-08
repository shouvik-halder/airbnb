import { AvailableRoomByTypeAndDateRangeDTO, BookingIdToBookedRoomsDTO } from "../dtos/room.dto";
import RoomRepository from "../repositories/room.repository";
import { sequelize } from "../db/models/sequelize";
import { BadRequestError } from "../utils/errors/app.error";

const roomRepository = new RoomRepository();

export async function findAvailableRoomByTypeAndDateService(
  data: AvailableRoomByTypeAndDateRangeDTO,
) {
  const { roomTypeId, hotelId, checkIn, checkOut } = data;
  const checkInDate = new Date(checkIn);
  const checkOutDate = new Date(checkOut);

  const totalDays = Math.ceil((checkOutDate.getTime() - checkInDate.getTime()) / (1000 * 60 * 60 * 24));
  const availableRooms =
    await roomRepository.findAvailableRoomsByTypeAndDateRange(
      roomTypeId,
      hotelId,
      checkInDate,
      checkOutDate,
      totalDays,
    );
  return availableRooms;
}

export async function updateBookingIdToBookedRoomsService(data:BookingIdToBookedRoomsDTO){
    const{bookingId, roomTypeId, hotelId, checkIn, checkOut} = data;
    const checkInDate = new Date(checkIn);
    const checkOutDate = new Date(checkOut);
    const totalDays = Math.ceil((checkOutDate.getTime() - checkInDate.getTime()) / (1000 * 60 * 60 * 24));

    // Use database transaction to ensure atomic booking operation
    // This prevents race conditions where multiple concurrent requests try to book the same room
    const result = await sequelize.transaction(async (transaction) => {
      // Query available rooms within transaction context
      const availableRooms = await roomRepository.findAvailableRoomsByTypeAndDateRange(
        roomTypeId,
        hotelId,
        checkInDate,
        checkOutDate,
        totalDays,
        transaction
      );
      
      if (!availableRooms || availableRooms.length === 0) {
        throw new BadRequestError("No available rooms found for the requested dates");
      }

      // Book the first available room within the transaction
      // If concurrent request already booked it, the update will return 0 rows
      const roomNumber = availableRooms[0].getDataValue('room_number');
      const roomsUpdated = await roomRepository.updateBookingIdToBookedRoomsWithLock(
        transaction,
        bookingId,
        roomTypeId,
        roomNumber,
        hotelId,
        checkInDate,
        checkOutDate,
      );

      if (!roomsUpdated || roomsUpdated === 0) {
        throw new BadRequestError("Failed to book room - it was claimed by another request. Please retry.");
      }

      return { roomsBooked: roomsUpdated, roomNumber };
    });

    return result;
}
