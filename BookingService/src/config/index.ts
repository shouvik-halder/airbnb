// This file controls the configuration of the project. 
// You can add any configuration related code here and export it to be used in other files.

import dotenv from "dotenv";

type ServerConfigType = {
    PORT: number
}

function getConfig() {
    dotenv.config();
}
getConfig();
export const serverConfig: ServerConfigType = {
    PORT: Number(process.env.PORT) || 3000
};