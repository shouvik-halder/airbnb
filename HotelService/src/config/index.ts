// This file controls the configuration of the project. 
// You can add any configuration related code here and export it to be used in other files.

import dotenv from "dotenv";

type ServerConfigType = {
    PORT: number
}

type DBConfigType = {
  MYSQL_DATABASE: string;
  MYSQL_USER: string;
  MYSQL_PASSWORD: string;
};

type RedisConfigType = {
  REDIS_HOST: string;
  REDIS_PORT: number;
  REDIS_URL: string;
};


type SchedulerConfigType = {
  ROOM_GENERATION_CRON_TIMEZONE:string;
  DAILY_ROOM_GENERATION_CRON:string;

}
function getConfig() {
    dotenv.config();
}
getConfig();
export const serverConfig: ServerConfigType = {
    PORT: Number(process.env.PORT) || 3000
};

export const dbConfig:DBConfigType = {
    MYSQL_DATABASE: process.env.MYSQL_DATABASE || 'hotelservicedb',
  MYSQL_USER: process.env.MYSQL_DATABASE || 'user',
  MYSQL_PASSWORD: process.env.MYSQL_DATABASE || 'user123'
}

export const redisConfig: RedisConfigType = {
    REDIS_HOST: process.env.REDIS_HOST ?? "",
    REDIS_PORT: Number(process.env.REDIS_PORT) || 6381,
    REDIS_URL: process.env.REDIS_URL ?? "",
    };

export const schedulerConfig: SchedulerConfigType = {
  ROOM_GENERATION_CRON_TIMEZONE: process.env.ROOM_GENERATION_CRON_TIMEZONE ?? "UTC",
  DAILY_ROOM_GENERATION_CRON: process.env.DAILY_ROOM_GENERATION_CRON ?? "0 0 * * *",
}