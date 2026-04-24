export type CreateRoomTypeDTO = {
    hotel_id:number;
    name:string;
    description?:string;
    max_occupancy:number;
    room_count:number;
}