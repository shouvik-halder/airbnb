import { Sequelize } from "sequelize";

export const sequelize = new Sequelize({
    username: process.env.MYSQL_USER,
    password: process.env.MYSQL_PASSWORD,
    database: process.env.MYSQL_DATABASE,
    host: "127.0.0.1",
    dialect: "mysql",
    logging:true
})