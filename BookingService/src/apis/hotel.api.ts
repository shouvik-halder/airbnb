import axios from 'axios';
import { apiConfig } from '../config';

export const getAvailableRooms = async(roomTypeId:number, hotelId:number, checkIn:string, checkOut:string)=>{
    const payload = {
  roomTypeId,
  hotelId,
  checkIn,
  checkOut
};
    const response = await axios({
        method:'get',
        url:`${apiConfig.HOTEL_SERVER_BASE_URL}/room/available`,
        data:payload
    });

    return response.data;
}


export const bookAvailableRooms = async (
  bookingId: number,
  roomTypeId: number,
  hotelId:number,
  checkIn: string,
  checkOut: string) =>{

    const payload = {
  bookingId,
  roomTypeId,
  hotelId,
  checkIn,
  checkOut
};
    const response = await axios({
        method:'post',
        url:`${apiConfig.HOTEL_SERVER_BASE_URL}/room/book`,
        data:payload
    });

    return response.data;
  }