import express  from "express";
import PingRouter from "./ping.router";
import HotelRouter from "./hotel.router";

const v1Router = express.Router();

v1Router.use('/ping', PingRouter);
v1Router.use('/hotel', HotelRouter);

export default v1Router;