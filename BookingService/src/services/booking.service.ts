import { CreateBookingDTO } from "../dtos/booking.dto";
import { ConfirmBooking, Createbooking, CreateIdempotencyKey, FinalizeIdempotencyKey, GetIdempotencyKey } from "../repositories/booking.repository";
import { BadRequestError, NotFoundError } from "../utils/errors/app.error";
import { GenerateIdempotencyKey } from "../utils/generateIdempotencyKey";

export async function CreateBookingService(BookingData:CreateBookingDTO) {

    const booking = await Createbooking(BookingData);

    const idempotencyKey = await CreateIdempotencyKey(GenerateIdempotencyKey(), booking.id)

    return {
        booking,
        idempotencyKey
    }
}

export async function ConfirmBookingService(idempotencykey:string) {
    const idempotencyKey = await GetIdempotencyKey(idempotencykey);

    if(!idempotencyKey){
        throw new NotFoundError("Idempotency key not found");
    }

    if(idempotencyKey.finalized){
        throw new BadRequestError("Booking already finalized");
    }

    if(!idempotencyKey.bookingId){
        throw new BadRequestError("No booking associated with this idempotency key");
    }

    const booking = await ConfirmBooking(idempotencyKey.bookingId);
    await FinalizeIdempotencyKey(idempotencykey);

    return booking;
}