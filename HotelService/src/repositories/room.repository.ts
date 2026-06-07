import { Op } from "sequelize";
import Room from "../db/models/room";
import { BaseRepository } from "./base.repository";

class RoomRepository extends BaseRepository<Room>{
    constructor(){
        super(Room);
    }

    async findByRoomTypeIdAndDate(room_type_id:number, hotel_id:number, date:Date){

        return await this.model.findOne({
            where:{
                room_type_id,
                hotel_id,
                status:"available",
                deletedAt:null,
                statusDate:date
            }
        })
    }

    async findByRoomAssignment(room_type_id:number, hotel_id:number, room_number:string, statusDate:Date){
        return await this.model.findOne({
            where:{
                room_type_id,
                hotel_id,
                room_number,
                statusDate,
                deletedAt:null
            }
        });
    }

    async bulkCreateRooms(rooms: Array<Partial<Room>>){
        return await this.model.bulkCreate(rooms as any[]);
    }

    async findLatestStatusDate(roomTypeId: number, hotelId: number): Promise<Date | null> {
        const latestStatusDate = await this.model.max("statusDate", {
            where: {
                room_type_id: roomTypeId,
                hotel_id: hotelId,
                deletedAt: null
            }
        });

        return latestStatusDate ? new Date(latestStatusDate as string | number | Date) : null;
    }

    async findExistingRooms(
    roomTypeId: number,
    hotelId: number,
    startDate: Date,
    endDate: Date
) {
    return Room.findAll({
        where: {
            room_type_id: roomTypeId,
            hotel_id: hotelId,
            statusDate: {
                [Op.between]: [startDate, endDate]
            }
        },
        attributes: ["room_number", "statusDate"]
    });
}
}

export default RoomRepository;
