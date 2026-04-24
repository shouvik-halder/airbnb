import logger from "../config/logger.config";
import { CreateHotelDTO, UpdateHotelDTO } from "../dtos/hotel.dto";
import HotelRepository from "../repositories/hotel.repository";
// import { createHotel, getAllHotels, getHotelById, softDeleteHotel, updateHotel } from "../repositories/hotel.repository";

const hotelRepository = new HotelRepository();
export async function createHotelService(hotelData:CreateHotelDTO) {
    logger.info(`Creating hotel with name: ${hotelData.name}`);
    const hotel = await hotelRepository.create(hotelData);    
    return hotel;
}

export async function getHotelByIdService(id:number) {
    logger.info(`Finding hotel with id: ${id}`);
    const hotel = await hotelRepository.findById(id);
    return hotel;
}

export async function getAllHotelsService() {
    logger.info(`Get all hotels from database`);
    const hotels = await hotelRepository.findAll();
    return hotels;
}

export async function updateHotelService(id:number, updateHotelData:UpdateHotelDTO) {
    const hotel = await hotelRepository.update(id, updateHotelData);
    return hotel;
}

export async function softDeleteHotelService(id:number) {
    logger.info(`Deleting hotel with id: ${id}`);
    await hotelRepository.softDeleteHotel(id);
    return true;
}