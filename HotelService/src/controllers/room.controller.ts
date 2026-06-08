import { Request, Response, NextFunction } from "express";
import { StatusCodes } from "http-status-codes";
import { addRoomsGenerationToQueue } from "../producers/roomsGeneration.producer";
import { findAvailableRoomByTypeAndDateService, updateBookingIdToBookedRoomsService } from "../services/room.service";

export async function generateRoomsController(req:Request, res:Response, next:NextFunction){
    const job = await addRoomsGenerationToQueue(req.body);

    res.status(StatusCodes.ACCEPTED).json({
            message:"room generation job accepted",
            success: true,
            data: {jobId: job.id}
        });
}

export async function findAvailableRoomByTypeAndDateController(req:Request, res:Response, next:NextFunction) {
    const rooms = await findAvailableRoomByTypeAndDateService(req.body);

    res.status(StatusCodes.OK).json({
        message:"available rooms fetched successfully",
        success: true,
        data: rooms
    });
}

export async function updateBookingIdToBookedRoomsController(req:Request, res:Response, next:NextFunction) {
    const roomsUpdated = await updateBookingIdToBookedRoomsService(req.body);
    res.status(StatusCodes.OK).json({
        message:"rooms booked successfully",
        success: true,
        data: roomsUpdated
    });
}
