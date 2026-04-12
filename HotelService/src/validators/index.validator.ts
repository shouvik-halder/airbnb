import { NextFunction, Request, Response } from "express";
import { ZodObject, ZodError} from "zod";
import { ValidationError } from "../utils/errors/app.error";
import logger from "../config/logger.config";

// export const validateRequestBody = (schema: ZodObject) =>{
//     return async (req:Request, res:Response, next:NextFunction) =>{
//         try {
//             logger.info("Validating request body")
//             await schema.parseAsync(req.body);
//             next();
//         } catch (error) {
//             logger.error("request body is invalid");
//             throw new ValidationError("Validation failed");
//         }
//     }
// }

// export const validateRequestParams = (schema: ZodObject) =>{
//     return async (req:Request, res:Response, next:NextFunction) =>{
//         try {
//             logger.info("Validating request params")
//             await schema.parseAsync(req.params);
//             next();
//         } catch (error) {
//             logger.error("request params is invalid");
//             throw new ValidationError("Validation failed");
//         }
//     }
// }

// export const validateRequestQuery = (schema: ZodObject) =>{
//     return async (req:Request, res:Response, next:NextFunction) =>{
//         try {
//             logger.info("Validating request query")
//             await schema.parseAsync(req.query);
//             next();
//         } catch (error) {
//             logger.error("request query is invalid");
//             throw new ValidationError("Validation failed");
//         }
//     }
// }

const validate = (schema: ZodObject, source: "body" | "params" | "query") => {
    return async (req: Request, res: Response, next: NextFunction) => {
        try {
            logger.info(`Validating request ${source}`);
            req[source] = await schema.parseAsync(req[source]);
            next();
        } catch (error) {
            logger.error(`Invalid request ${source}`);
            if (error instanceof ZodError) {
                throw new ValidationError(`Validation failed-${error.message}`);
            }
        }
    };
};

export const validateRequestBody = (schema: ZodObject) =>
    validate(schema, "body");

export const validateRequestParams = (schema: ZodObject) =>
    validate(schema, "params");

export const validateRequestQuery = (schema: ZodObject) =>
    validate(schema, "query");

