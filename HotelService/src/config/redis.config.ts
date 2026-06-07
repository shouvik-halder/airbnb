import { Redis } from "ioredis";
import { redisConfig } from ".";

function connectToRedis() {
  try {
    
    let connection: Redis;
    return () => {
      if (!connection) {
        connection = new Redis({
            host:redisConfig.REDIS_HOST,
            port:redisConfig.REDIS_PORT,
            maxRetriesPerRequest:null
        });
      }
      return connection;
    };
  } catch (error) {
    throw new Error("Failed to connect to Redis");
  }
}

export const getRedisClient = connectToRedis();

