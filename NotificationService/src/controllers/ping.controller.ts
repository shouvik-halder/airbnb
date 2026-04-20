import { Request, Response } from "express";
import { InternalServerError } from "../utils/errors/app.error";
import logger from "../config/logger.config";

export const PingController = (req:Request, res:Response) => {

    try {
        logger.info(PingController.name)
        console.log('Request body', req.body);
        console.log('Request params', req.params);
        console.log('Request query', req.query);
        // throw new Error("Error in the method")
        res.send("pong");
        
    } catch (error) {
        throw new InternalServerError((error as Error).message);
    }
}