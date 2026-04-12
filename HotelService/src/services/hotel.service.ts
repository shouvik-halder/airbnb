import { CreateHotelDTO, UpdateHotelDTO } from "../dtos/hotel.dto";
import { createHotel, getAllHotels, getHotelById, softDeleteHotel, updateHotel } from "../repositories/hotel.repository";

export async function createHotelService(hotelData:CreateHotelDTO) {
    const hotel = await createHotel(hotelData);    
    return hotel;
}

export async function getHotelByIdService(id:number) {
    const hotel = await getHotelById(id);
    return hotel;
}

export async function getAllHotelsService() {
    const hotels = await getAllHotels();
    return hotels;
}

export async function updateHotelService(id:number, updateHotelData:UpdateHotelDTO) {
    const hotel = await updateHotel(id, updateHotelData);
    return hotel;
}

export async function softDeleteHotelService(id:number) {
    await softDeleteHotel(id);
    return true;
}