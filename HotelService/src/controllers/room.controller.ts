import { Request, Response, NextFunction } from "express";
import { StatusCodes } from "http-status-codes";
import { addRoomsGenerationToQueue } from "../producers/roomsGeneration.producer";

export async function generateRooms(req:Request, res:Response, next:NextFunction){
    const job = await addRoomsGenerationToQueue(req.body);

    res.status(StatusCodes.ACCEPTED).json({
            message:"room generation job accepted",
            success: true,
            data: {jobId: job.id}
        });
}
