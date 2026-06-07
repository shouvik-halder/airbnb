import { Worker} from 'bullmq';
import { ROOMS_GENERATION_QUEUE } from '../queues/roomsGeneration.queue';
import logger from '../config/logger.config';
import { BadRequestError } from '../utils/errors/app.error';
import { RoomGenerationJobSchemaType } from '../dtos/roomGeneration.dto';
import { ROOMS_GENERATION_PAYLOAD } from '../producers/roomsGeneration.producer';
import { generateRoomsService } from '../services/roomGenerations.service';
import { getRedisClient } from '../config/redis.config';
import { runWithCorrelationId } from '../helpers/request.helper';

export const setupRoomGenerationWorker = () =>{
    const roomsGenerator = new Worker<RoomGenerationJobSchemaType>(ROOMS_GENERATION_QUEUE, async (job)=>{
        return runWithCorrelationId(job.data.correlationId, async () => {
            if(job.name !== ROOMS_GENERATION_PAYLOAD){
                logger.error(`Unknown job payload ${job.name} in rooms generation queue`);
                throw new BadRequestError(`Unknown job type ${job.name}`);
            }

            const payload = job.data;
            logger.info(`Processing room generation for hotel ${payload.hotelId} and room type ${payload.roomTypeId}`);

            const result = await generateRoomsService(job.data);

            return result;
        });
    },
    {
        connection: getRedisClient()
    });

    roomsGenerator.on("completed", (job)=>{
        runWithCorrelationId(job.data.correlationId, () => {
            logger.info(`Room generation job completed for hotel ${job.data.hotelId} and room type ${job.data.roomTypeId}`);
        });
    });
    
    roomsGenerator.on("failed", (job, err)=>{
        runWithCorrelationId(job?.data.correlationId, () => {
            logger.error(`Room generation job failed for hotel ${job?.data.hotelId} and room type ${job?.data.roomTypeId} with error ${err.message}`);
        });
    });

    return roomsGenerator;
}
