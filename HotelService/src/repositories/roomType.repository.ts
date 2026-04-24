import RoomType from "../db/models/roomTypes";
import { NotFoundError } from "../utils/errors/app.error";
import { BaseRepository } from "./base.repository";

class RoomTypeRepository extends BaseRepository<RoomType>{
    constructor(){
        super(RoomType);
    }

    async findAllByHotelId(hotel_id:number): Promise<RoomType[]> {
        const roomTypes =  await this.model.findAll({
            where:{
                hotel_id,
                deletedAt:null
            }
        });

        if(!roomTypes || roomTypes.length === 0)
            throw new NotFoundError(`No room types found for hotel with id: ${hotel_id}`);

        return roomTypes;
    }

    async softDeleteByHotelId(id:number, hotel_id:number){
        const roomType = await this.model.findOne({
            where:{
                id,
                hotel_id
            }
        });

        if(!roomType)
            throw new NotFoundError(`No room type found for id: ${id} in Hotel with id: ${hotel_id}`);

        roomType.deletedAt = new Date();
        await roomType.save();
        return roomType;
    }

}

export default RoomTypeRepository;

