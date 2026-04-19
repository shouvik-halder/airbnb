import IORedis, { Redis } from "ioredis";
import Redlock from "redlock";
import { redisConfig } from ".";


export const redisClient = new IORedis(redisConfig.REDIS_URL);

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

export const redLock = new Redlock([redisClient as any], {
  driftFactor: 0.01,
  retryCount: 10,
  retryDelay: 200,
  retryJitter: 200,
});