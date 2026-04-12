import { NextFunction, Request, Response } from "express";
import { createHotelService, getAllHotelsService, getHotelByIdService, softDeleteHotelService, updateHotelService } from "../services/hotel.service";
import { StatusCodes } from "http-status-codes";

export async function createHotelController(req:Request, res:Response, next:NextFunction) {
    const hotelResponse = await createHotelService(req.body);

    res.status(StatusCodes.CREATED).json({
        message:"hotel created",
        success:true,
        data:hotelResponse
    });
}

export async function getHotelByIdController(req:Request, res:Response, next:NextFunction) {
    const hotelResponse = await getHotelByIdService(Number(req.params.id));

    res.status(StatusCodes.OK).json({
        message:"hotel fetched successfully",
        success:true,
        data:hotelResponse
    });
}

export async function getAllHotelsController(req:Request, res:Response, next:NextFunction) {
    const hotels = await getAllHotelsService();
    res.status(StatusCodes.OK).json({
        message:`${hotels.length} hotels found`,
        success:true,
        data:hotels
    })
}

export async function updateHotelController(req:Request, res:Response, next:NextFunction) {
    const hotel = await updateHotelService(Number(req.params.id), req.body);
    res.status(StatusCodes.OK).json({
        message:`hotel data updated successfully`,
        success:true,
        data:hotel
    })
}

export async function deleteHotelController(req:Request, res:Response, next:NextFunction) {
    const hotel = await softDeleteHotelService(Number(req.params.id));
    res.status(200).json({
        message:`hotel data deleted successfully`,
        success:true,
        data:hotel
    })
}