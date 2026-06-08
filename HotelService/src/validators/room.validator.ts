import z from "zod";

export const AvailableRoomByTypeAndDateRangeSchema = z.object({
    roomTypeId: z.number().positive(),
    hotelId: z.number().positive(),
    checkIn: z.iso.date(),
    checkOut: z.iso.date(),
});

export const UpdateBookingIdToBookedRoomsSchema = z.object({
    bookingId: z.number().positive(),
  roomTypeId: z.number().positive(),
  hotelId:z.number().positive(),
  checkIn: z.iso.date(),
  checkOut: z.iso.date()
})