import express  from "express";
import PingRouter from "./ping.router";

const v1Router = express.Router();

v1Router.use('/ping', PingRouter);

export default v1Router;