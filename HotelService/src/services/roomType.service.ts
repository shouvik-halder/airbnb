import { CreateRoomTypeDTO } from "../dtos/roomType.dto";
import HotelRepository from "../repositories/hotel.repository";
import RoomTypeRepository from "../repositories/roomType.repository";
import { BadRequestError } from "../utils/errors/app.error";

const roomTypeRepository = new RoomTypeRepository();
const hotelRepository = new HotelRepository();

export async function createRoomTypeByHotelIdService(
  createRoomTypeData: CreateRoomTypeDTO,
) {
  const roomtype = await roomTypeRepository.create(createRoomTypeData);
  return roomtype;
}


export async function getAllRoomTypesByHotelIdService(hotel_id:number) {
  const hotel = await hotelRepository.findById(hotel_id);
  if(!hotel || hotel.deletedAt !== null)
    throw new BadRequestError(`No hotel found with id: ${hotel_id}`);
    const roomtypes = await roomTypeRepository.findAllByHotelId(hotel_id);
    return roomtypes;
    
}

export async function updateRoomTypeByHotelIdService(id:number, hotel_id:number, updateRoomTypedata: Partial<CreateRoomTypeDTO>){
  const hotel = await hotelRepository.findById(hotel_id);
  if(!hotel || hotel.deletedAt !== null)
    throw new BadRequestError(`No hotel found with id: ${hotel_id}`);
    const roomtype = await roomTypeRepository.findById(id);
    if(!roomtype || roomtype.hotel_id !== hotel_id)
        throw new BadRequestError(`No room type found for id ${id} in hotel with id: ${hotel_id}`);
    const updatedRoomtype = await roomTypeRepository.update(id, updateRoomTypedata);
    return updatedRoomtype;
}