import { prisma } from "../prisma/client";
import { Prisma } from "../prisma/generated/client";

export async function Createbooking(bookingData:Prisma.BookingCreateInput){
    const booking = await prisma.booking.create({
        data:bookingData
    });

    return booking;
}

