import express  from "express";
import PingRouter from "./ping.router";
import HotelRouter from "./hotel.router";
import RoomRouter from "./room.router";

const v1Router = express.Router();

v1Router.use('/ping', PingRouter);
v1Router.use('/hotel', HotelRouter);
v1Router.use('/room', RoomRouter);

export default v1Router;