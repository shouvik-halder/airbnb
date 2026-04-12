import { Request, Response, NextFunction } from "express";
import { AppError } from "../../utils/errors/app.error";

export const genericErrorMiddleware = (err:Error, req: Request, res: Response, next:NextFunction)=>{
    res.json({
        name:err.name,
        success:false,
        message:err.message
    });
}


export const appErrorMiddleware = (err:AppError, req: Request, res: Response, next:NextFunction)=>{
    res.status(err.statusCode).json({
        name:err.name,
        success:err.success,
        message:err.message
    });
}