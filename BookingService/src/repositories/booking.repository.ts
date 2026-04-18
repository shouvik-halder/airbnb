import { validate as isValidUUID } from "uuid";
import  prisma  from "../prisma/client";
import { IdempotencyKey, Prisma } from "../prisma/generated/client";
import { BadRequestError } from "../utils/errors/app.error";

export async function Createbooking(bookingData:Prisma.BookingCreateInput){
    const booking = await prisma.booking.create({
        data:bookingData
    });

    return booking;
}

export async function CreateIdempotencyKey(key:string, bookingId:number) {
    const idempotencyKey = await prisma.idempotencyKey.create({
        data:{
            idemKey:key,
            booking:{
                connect:{
                    id:bookingId
                }
            }
        }
    })
    return idempotencyKey;
}

export async function GetIdempotencyKeyWithLock(tx:Prisma.TransactionClient, key:string) {
    
    if(!isValidUUID(key))
        throw new BadRequestError("Invalid idempotency key format");
    const idempotencyKey = await tx.$queryRaw<IdempotencyKey[]>`SELECT * FROM IdempotencyKey WHERE idemKey = ${key} FOR UPDATE`;


    return idempotencyKey[0];
}

export async function GetBookingById(bookingId:number) {
    const booking = await prisma.booking.findUnique({
        where:{
            id:bookingId
        }
    });

    return booking;
}

export async function FinalizeIdempotencyKeyWithLock(tx:Prisma.TransactionClient,key:string) {
    const idempotencyKey = await tx.idempotencyKey.update({
        where:{
            idemKey:key
        },
        data:{
            finalized:true
        }
    });

    return idempotencyKey;
}

export async function ConfirmBookingWithLock(tx:Prisma.TransactionClient, bookingId:number) {
    const booking = await tx.booking.update({
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