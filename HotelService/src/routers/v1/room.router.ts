import express from "express";
import { findAvailableRoomByTypeAndDateController, generateRoomsController, updateBookingIdToBookedRoomsController } from "../../controllers/room.controller";
import { validateRequestBody } from "../../validators/index.validator";
import { RoomGeneratonRequestSchema } from "../../dtos/roomGeneration.dto";
import { AvailableRoomByTypeAndDateRangeSchema, UpdateBookingIdToBookedRoomsSchema } from "../../validators/room.validator";
const RoomRouter = express.Router();

RoomRouter.post('/generate', validateRequestBody(RoomGeneratonRequestSchema), generateRoomsController);

RoomRouter.get('/available', validateRequestBody(AvailableRoomByTypeAndDateRangeSchema), findAvailableRoomByTypeAndDateController);

RoomRouter.post('/book', validateRequestBody(UpdateBookingIdToBookedRoomsSchema), updateBookingIdToBookedRoomsController);

export default RoomRouter;