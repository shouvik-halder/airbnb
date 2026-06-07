import express from "express";
import { generateRooms } from "../../controllers/room.controller";
import { validateRequestBody } from "../../validators/index.validator";
import { RoomGeneratonRequestSchema } from "../../dtos/roomGeneration.dto";
const RoomRouter = express.Router();

RoomRouter.post('/generate', validateRequestBody(RoomGeneratonRequestSchema), generateRooms);

export default RoomRouter;