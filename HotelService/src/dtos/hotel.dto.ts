export type CreateHotelDTO = {
    name:string;
    address:string;
    location:string;
    rating?:number;
    ratingCount?:number;
}

export type UpdateHotelDTO = {
    name?:string;
    rating?:number;
    ratingCount?:number;
}