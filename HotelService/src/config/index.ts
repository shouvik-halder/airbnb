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