// This file controls the configuration of the project.
// You can add any configuration related code here and export it to be used in other files.

import dotenv from "dotenv";

type ServerConfigType = {
  PORT: number;
};

type RedisServerConfigType = {
  REDIS_HOST: string;
  REDIS_PORT: number;
  REDIS_URL: string;
};

type NodeMailerConfigType = {
  SMTP_USER: string;
  SMTP_PASS: string;
};

function getConfig() {
  dotenv.config();
}
getConfig();
export const serverConfig: ServerConfigType = {
  PORT: Number(process.env.PORT) || 3000,
};

export const redisConfig: RedisServerConfigType = {
  REDIS_HOST: process.env.REDIS_HOST ?? "",
  REDIS_PORT: Number(process.env.REDIS_PORT) || 6379,
  REDIS_URL: process.env.REDIS_URL ?? "",
};

export const nodemailerConfig: NodeMailerConfigType = {
  SMTP_USER: process.env.SMTP_USER ?? "",
  SMTP_PASS: process.env.SMTP_PASS ?? "",
};