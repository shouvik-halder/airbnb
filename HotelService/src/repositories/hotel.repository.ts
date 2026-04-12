import logger from "../config/logger.config";
import Hotel from "../db/models/hotel";
import { CreateHotelDTO, UpdateHotelDTO } from "../dtos/hotel.dto";
import { NotFoundError } from "../utils/errors/app.error";

export async function createHotel(hotelData:CreateHotelDTO) {
    const hotel = await Hotel.create({
        ...hotelData
    });

    logger.info('Hotel data created successfully');
    return hotel;
}

export async function getHotelById(id:number) {
    const hotel = await Hotel.findByPk(id);
    if(!hotel){
        logger.error(`No hotels were found with this ${id}`);
        throw new NotFoundError(`No hotels were found with this ${id}`);
    }
    logger.info(`Hotel found ${hotel.id}`, hotel.dataValues);
    return hotel;
}

export async function getAllHotels() {
    logger.info('Fetching all hotels from database')
    const hotels = await Hotel.findAll({
    where: {
        deletedAt: null
    }
});

    if(!hotels || hotels.length == 0){
        logger.error('No hotels found');
        throw new NotFoundError('No hotels found');
    }

    return hotels;
}


export async function updateHotel(id:number, updateHotelData:UpdateHotelDTO) {
    logger.info(`Updating info for id:${id}`);
    try {
        const hotel = await Hotel.findByPk(id);
        if(!hotel || hotel.deletedAt!=null){
            throw new NotFoundError(`There is no Hotel with this ${hotel?.id} or the hotel is deleted`);
        }
        await hotel.update({
          ...(updateHotelData.name !== undefined && {
            name: updateHotelData.name,
          }),
          ...(updateHotelData.rating !== undefined && {
            rating: updateHotelData.rating,
          }),
          ...(updateHotelData.ratingCount !== undefined && {
            ratingCount: updateHotelData.ratingCount,
          }),
        });

    } catch (error) {
        logger.error("Error updating hotel:", error);
        throw error;
    }
}

export async function softDeleteHotel(id:number) {
    const hotel = await Hotel.findByPk(id);
    if(!hotel){
        logger.error(`No hotels were found with this ${id}`);
        throw new NotFoundError(`No hotels were found with this ${id}`);
    }

    hotel.deletedAt = new Date();
    await hotel.save();
    logger.info(`Hotel ${hotel.id} soft deleted`);
    return hotel;
}


