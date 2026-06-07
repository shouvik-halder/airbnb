import { Queue } from "bullmq";
import { getRedisClient } from "../config/redis.config";

export const ROOMS_GENERATION_QUEUE = `queue-rooms-generation`;

export const roomsGenerationQueue = new Queue(ROOMS_GENERATION_QUEUE, {
    connection: getRedisClient(),
});