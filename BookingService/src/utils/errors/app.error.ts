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
    super(message, 500);
    this.name = "InternalServerError";
  }
}
export class BadRequestError extends BaseError {
  constructor(message = "Bad Request") {
    super(message, 400);
    this.name = "BadRequestError";
  }
}
export class UnauthorizedError extends BaseError {
  constructor(message = "Unauthorized") {
    super(message, 401);
    this.name = "UnauthorizedError";
  }
}
export class ForbiddenError extends BaseError {
  constructor(message = "Forbidden") {
    super(message, 403);
    this.name = "ForbiddenError";
  }
}
export class NotFoundError extends BaseError {
  constructor(message = "Resource Not Found") {
    super(message, 404);
    this.name = "NotFoundError";
  }
}
export class ValidationError extends BaseError {
  constructor(message = "Validation Failed") {
    super(message, 422);
    this.name = "ValidationError";
  }
}