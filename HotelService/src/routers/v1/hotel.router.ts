import express from 'express';
import { createHotelController, deleteHotelController, getAllHotelsController, getHotelByIdController, updateHotelController } from '../../controllers/hotel.controller';
import { validateRequestBody, validateRequestParams } from '../../validators/index.validator';
import { createHotelSchema, getHotelByIdSchema, updateHotelSchema } from '../../validators/hotel.validator';

const HotelRouter = express.Router();

HotelRouter.post('/', validateRequestBody(createHotelSchema), createHotelController);

HotelRouter.get('/:id', validateRequestParams(getHotelByIdSchema), getHotelByIdController);

HotelRouter.get('/',getAllHotelsController);

HotelRouter.post('/:id', validateRequestBody(updateHotelSchema), validateRequestParams(getHotelByIdSchema), updateHotelController);

HotelRouter.delete('/:id', deleteHotelController);

export default HotelRouter;