// This file controls the configuration of the project.
// You can add any configuration related code here and export it to be used in other files.

import dotenv from "dotenv";

type ServerConfigType = {
  PORT: number;
};

type DBConfigType = {
  DATABASE_URL: string;
  DATABASE_USER: string;
  DATABASE_PASSWORD: string;
  DATABASE_NAME: string;
  DATABASE_HOST: string;
  DATABASE_PORT: number;
};

type RedisServerConfigType = {
  REDIS_HOST: string;
  REDIS_PORT: number;
  REDIS_URL: string;
};

function getConfig() {
  dotenv.config();
}
getConfig();
export const serverConfig: ServerConfigType = {
  PORT: Number(process.env.PORT) || 3000,
};

export const dbConfig: DBConfigType = {
  DATABASE_URL: process.env.DATABASE_URL ?? "",
  DATABASE_HOST: process.env.DATABASE_HOST ?? "",
  DATABASE_USER: process.env.DATABASE_USER ?? "",
  DATABASE_PORT: Number(process.env.DATABASE_PORT) || 3307,
  DATABASE_PASSWORD: process.env.DATABASE_PASSWORD ?? "",
  DATABASE_NAME: process.env.DATABASE_NAME ?? "",
};

export const redisConfig: RedisServerConfigType = {
    REDIS_HOST: process.env.REDIS_HOST ?? "",
    REDIS_PORT: Number(process.env.REDIS_PORT) || 6379,
    REDIS_URL: process.env.REDIS_URL ?? ""
}