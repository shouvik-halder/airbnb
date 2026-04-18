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
  const booking = await Createbooking(BookingData);

  const idempotencyKey = await CreateIdempotencyKey(
    GenerateIdempotencyKey(),
    booking.id,
  );

  return {
    booking,
    idempotencyKey,
  };
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
    await FinalizeIdempotencyKeyWithLock(tx, idempotencykey);

    return booking;
  });
}