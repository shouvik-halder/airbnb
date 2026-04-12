import { StatusCodes } from "http-status-codes";

export interface AppError extends Error {
  statusCode: number;
  success: boolean;
}

export abstract class BaseError extends Error implements AppError {
  statusCode: number;
  success: boolean;

  constructor(message: string, statusCode: number) {
    super(message);
    this.message = message;
    this.statusCode = statusCode;
    this.success = false;

  }
}

export class InternalServerError extends BaseError {
  constructor(message = "Internal Server Error") {
    super(message, StatusCodes.INTERNAL_SERVER_ERROR);
    this.name = "InternalServerError";
  }
}
export class BadRequestError extends BaseError {
  constructor(message = "Bad Request") {
    super(message, StatusCodes.BAD_REQUEST);
    this.name = "BadRequestError";
  }
}
export class UnauthorizedError extends BaseError {
  constructor(message = "Unauthorized") {
    super(message, StatusCodes.UNAUTHORIZED);
    this.name = "UnauthorizedError";
  }
}
export class ForbiddenError extends BaseError {
  constructor(message = "Forbidden") {
    super(message, StatusCodes.FORBIDDEN);
    this.name = "ForbiddenError";
  }
}
export class NotFoundError extends BaseError {
  constructor(message = "Resource Not Found") {
    super(message, StatusCodes.NOT_FOUND);
    this.name = "NotFoundError";
  }
}
export class ValidationError extends BaseError {
  constructor(message = "Validation Failed") {
    super(message, StatusCodes.UNPROCESSABLE_ENTITY);
    this.name = "ValidationError";
  }
}
