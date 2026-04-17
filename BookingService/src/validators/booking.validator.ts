import { z } from "zod";

export const CreateBookingSchema = z.object({
  userId: z.number().min(1, "userId is required"),
  hotelId: z.number().min(1, "hotelId is required"),
  bookingPrice: z.number().positive("bookingPrice must be greater than 0"),
  totalGuests: z.number().int().positive("totalGuests must be a positive integer"),
});

export const ConfirmBookingSchema = z.object({
    idempotencyKey: z.uuid("Invalid idempotency key format")
})
