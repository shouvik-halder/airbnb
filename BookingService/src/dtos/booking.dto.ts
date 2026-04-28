export type CreateBookingDTO = {
    userId: number,
    hotelId: number,
    roomTypeId: number,
    bookingPrice: number,
    totalGuests: number
}