import "dotenv/config";
import { PrismaMariaDb } from "@prisma/adapter-mariadb";
import { PrismaClient } from "./generated/client";
import { dbConfig } from "../config";

const adapter = new PrismaMariaDb({
  host: dbConfig.DATABASE_HOST,
  user: dbConfig.DATABASE_USER,
  port: dbConfig.DATABASE_PORT,
  password: dbConfig.DATABASE_PASSWORD,
  database: dbConfig.DATABASE_NAME,
  connectionLimit: 5,
});
const prisma = new PrismaClient({ adapter });

export default prisma;