import { Request, Response, NextFunction } from "express";
import {v4 as uuidv4} from "uuid";
import { asyncLocalStorage } from "../helpers/request.helper";

export const attachCorrelationId = (req:Request, res:Response, next:NextFunction)=>{
    const correlationId = req.header("x-correlation-id") || uuidv4();
    req.headers['x-correlation-id'] = correlationId;
    res.setHeader("X-Correlation-Id", correlationId);
    asyncLocalStorage.run({correlationId:correlationId},()=>{
        next();
    });
}
