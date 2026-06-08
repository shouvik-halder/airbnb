export type AvailableRoomByTypeAndDateRangeDTO = {
    roomTypeId: number;
    hotelId: number;
    checkIn: Date;
    checkOut: Date;
}

export type BookingIdToBookedRoomsDTO = {
  bookingId: number,
  roomTypeId: number;
  hotelId:number,
  checkIn: Date,
  checkOut: Date
}