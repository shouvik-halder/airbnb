import { NextFunction, Request, Response } from "express";
import { ZodObject} from "zod";
import { ValidationError } from "../utils/errors/app.error";
import logger from "../config/logger.config";

export const validateRequestBody = (schema: ZodObject) =>{
    return async (req:Request, res:Response, next:NextFunction) =>{
        try {
            logger.info("Validating request body")
            await schema.parseAsync(req.body);
            next();
        } catch (error) {
            logger.error("request body is invalid");
            throw new ValidationError("Validation failed");
        }
    }
}

