import { bookAvailableRooms, getAvailableRooms } from "../apis/hotel.api";
import { redLock } from "../config/redis.config";
import { CreateBookingDTO } from "../dtos/booking.dto";
import prisma from "../prisma/client";
import {
  ConfirmBookingWithLock,
  Createbooking,
  CreateIdempotencyKey,
  FinalizeIdempotencyKeyWithLock,
  GetIdempotencyKeyWithLock,
} from "../repositories/booking.repository";
import { BadRequestError, NotFoundError } from "../utils/errors/app.error";
import { GenerateIdempotencyKey } from "../utils/generateIdempotencyKey";

export async function CreateBookingService(BookingData: CreateBookingDTO) {
  const ttl = 60000;
  const bookingResource = `booking:hotel:${BookingData.hotelId}:roomType:${BookingData.roomTypeId}`;

  const avaiableRooms = await getAvailableRooms(BookingData.roomTypeId, BookingData.hotelId, BookingData.checkInDate, BookingData.checkOutDate);

  if(!avaiableRooms){
    throw new BadRequestError(`No rooms available for checkin ${BookingData.checkInDate} and checkout ${BookingData.checkOutDate}`);
  }
  
  const lock = await redLock.acquire([bookingResource], ttl);
  try {
    const booking = await Createbooking(BookingData);

    const idempotencyKey = await CreateIdempotencyKey(
      GenerateIdempotencyKey(),
      booking.id,
    );

    return {
      booking,
      idempotencyKey,
    };
  } catch (error) {
    throw new BadRequestError((error as Error).message);
  } finally {
    await lock.release().catch(() => {}); // Ensure lock is released even on error
  }
}

export async function ConfirmBookingService(idempotencykey: string) {
  return await prisma.$transaction(async (tx) => {
    const idempotencyKey = await GetIdempotencyKeyWithLock(tx, idempotencykey);

    if (!idempotencyKey) {
      throw new NotFoundError("Idempotency key not found");
    }

    if (idempotencyKey.finalized) {
      throw new BadRequestError("Booking already finalized");
    }

    if (!idempotencyKey.bookingId) {
      throw new BadRequestError(
        "No booking associated with this idempotency key",
      );
    }

    const booking = await ConfirmBookingWithLock(tx, idempotencyKey.bookingId);
    
    // Call external API BEFORE finalizing idempotency key to ensure idempotent behavior
    // If this fails, the transaction rolls back and idempotency key remains unfinalzed for retry
    const response = await bookAvailableRooms(booking.id, booking.roomTypeId, booking.hotelId, booking.checkInDate.toISOString().split("T")[0], booking.checkOutDate.toISOString().split("T")[0]);
    console.log(response);
    // Only finalize after API call succeeds
    await FinalizeIdempotencyKeyWithLock(tx, idempotencykey);
    return booking;
  });
}
