import { NextFunction, Request, Response } from "express";
import { CreateBookingService, ConfirmBookingService } from "../services/booking.service";
import { StatusCodes } from "http-status-codes";

export const CreateBookingController = async (req:Request, res:Response, next:NextFunction) => {
    const bookingResponse = await CreateBookingService(req.body);

    res.status(StatusCodes.CREATED).json({
        message:"booking created",
        success:true,
        data:bookingResponse
    });
}

export const ConfirmBookingController = async (req:Request, res:Response, next:NextFunction) => {
    const bookingResponse = await ConfirmBookingService(req.params.idempotencyKey as string);

    res.status(StatusCodes.OK).json({
        message:"Booking Finalized",
        success:true,
        data:bookingResponse
    })
}