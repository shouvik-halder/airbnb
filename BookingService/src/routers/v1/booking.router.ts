import express from "express";
import { validateRequestBody, validateRequestParams } from "../../validators/index.validator";
import { CreateBookingSchema, ConfirmBookingSchema } from "../../validators/booking.validator";
import { CreateBookingController, ConfirmBookingController } from "../../controllers/booking.controller";

const BookingRouter = express.Router();

BookingRouter.post('/', validateRequestBody(CreateBookingSchema), CreateBookingController);

BookingRouter.post('/confirm/:idempotencyKey', validateRequestParams(ConfirmBookingSchema),ConfirmBookingController)

export default BookingRouter;