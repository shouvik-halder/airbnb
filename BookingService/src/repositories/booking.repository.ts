import  prisma  from "../prisma/client";
import { Prisma } from "../prisma/generated/client";

export async function Createbooking(bookingData:Prisma.BookingCreateInput){
    const booking = await prisma.booking.create({
        data:bookingData
    });

    return booking;
}

export async function CreateIdempotencyKey(key:string, bookingId:number) {
    const idempotencyKey = await prisma.idempotencyKey.create({
        data:{
            key,
            booking:{
                connect:{
                    id:bookingId
                }
            }
        }
    })
    return idempotencyKey;
}

export async function GetIdempotencyKey(key:string) {
    const idempotencyKey = await prisma.idempotencyKey.findUnique({
        where:{
            key:key
        }
    });

    return idempotencyKey;
}

export async function GetBookingById(bookingId:number) {
    const booking = await prisma.booking.findUnique({
        where:{
            id:bookingId
        }
    });

    return booking;
}

export async function FinalizeIdempotencyKey(key:string) {
    const idempotencyKey = await prisma.idempotencyKey.update({
        where:{
            key:key
        },
        data:{
            finalized:true
        }
    });

    return idempotencyKey;
}

export async function ConfirmBooking(bookingId:number) {
    const booking = await prisma.booking.update({
        where:{
            id:bookingId
        },
        data:{
            bookingStatus: "CONFIRMED"
        }
    });

    return booking;
}

export async function CancellBooking(bookingId:number) {
    const booking = await prisma.booking.update({
        where:{
            id:bookingId
        },
        data:{
            bookingStatus: "CANCELLED"
        }
    });

    return booking;
}