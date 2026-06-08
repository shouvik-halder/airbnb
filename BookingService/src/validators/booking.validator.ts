import { z } from "zod";

export const CreateBookingSchema = z.object({
  userId: z.number().min(1, "userId is required"),
  hotelId: z.number().positive("hotelId is required"),
  roomTypeId: z.number().positive("roomTypeId is required"),
  bookingPrice: z.number().positive("bookingPrice must be greater than 0"),
  totalGuests: z.number().int().positive("totalGuests must be a positive integer"),
  checkInDate: z.iso.date("checkInDate must be a valid ISO datetime string"),
  checkOutDate: z.iso.date("checkOutDate must be a valid ISO datetime string"),
});

export const ConfirmBookingSchema = z.object({
    idempotencyKey: z.uuid("Invalid idempotency key format")
})
