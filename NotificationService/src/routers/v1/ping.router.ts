import express from "express";
import { PingController } from "../../controllers/ping.controller";
import { validateRequestBody } from "../../validators/index.validator";
import { pingSchema } from "../../validators/ping.validator";

const PingRouter = express.Router();

PingRouter.get("/", validateRequestBody(pingSchema), PingController);

PingRouter.get('/:user_id/comments',validateRequestBody(pingSchema), PingController);

PingRouter.get("/health", (req, res)=>{
    res.status(200).send('OK');
})

export default PingRouter;