import { NextFunction, Request, Response } from 'express';
import { createRoomTypeByHotelIdService, getAllRoomTypesByHotelIdService, updateRoomTypeByHotelIdService } from '../services/roomType.service';
import { StatusCodes } from 'http-status-codes';

export async function getAllRoomtypesByHotelIdController(req:Request, res:Response, next:NextFunction) {
    const roomTypesResponse = await getAllRoomTypesByHotelIdService(Number(req.params.hotel_id));
    res.status(StatusCodes.OK).json({
            message:"hotel fetched successfully",
            success:true,
            data:roomTypesResponse
        });

}

export async function createRoomTypeByHotelIdController(req:Request, res:Response, next:NextFunction) {
    const roomTypeResponse = await createRoomTypeByHotelIdService(req.body);
    res.status(StatusCodes.CREATED).json({
        message:"room type created successfully",
        success: true,
        data: roomTypeResponse
    })
}

export async function updateRoomTypeByHotelIdController(req:Request, res:Response, next:NextFunction) {
    const roomTypeResponse = await updateRoomTypeByHotelIdService(Number(req.params.id), Number(req.params.hotel_id), req.body);
    res.status(StatusCodes.OK).json({
        message:"room type updated successfully",
        success: true,
        data: roomTypeResponse
    })
}