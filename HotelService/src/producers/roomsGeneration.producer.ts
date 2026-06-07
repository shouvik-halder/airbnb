import logger from "../config/logger.config";
import { RoomGeneratonRequestSchemaType, RoomGenerationJobSchemaType } from "../dtos/roomGeneration.dto";
import { getCorrelationId } from "../helpers/request.helper";
import { roomsGenerationQueue } from "../queues/roomsGeneration.queue";
import { BadRequestError } from "../utils/errors/app.error";

export const ROOMS_GENERATION_PAYLOAD = `rooms-generation-payload`;

export const addRoomsGenerationToQueue = async (jobdata:RoomGeneratonRequestSchemaType) => {
    let delay = 0;

    if(jobdata.scheduleType === "SCHEDULED"){
        if(!jobdata.scheduleAt){
            throw new BadRequestError("scheduleAt is required for scheduled room generation");
        }

        delay = new Date(jobdata.scheduleAt).getTime() - Date.now();
        if(delay <= 0){
            throw new BadRequestError("scheduleAt must be in the future");
        }
    }

    const payload: RoomGenerationJobSchemaType = {
        roomTypeId: jobdata.roomTypeId,
        hotelId: jobdata.hotelId,
        startDate: jobdata.startDate,
        endDate: jobdata.endDate,
        priceOverride: jobdata.priceOverride,
        batchSize: jobdata.batchSize,
        correlationId: getCorrelationId(),
    };

    const job = await roomsGenerationQueue.add(ROOMS_GENERATION_PAYLOAD, payload, {
        delay,
        attempts: 3,
        backoff: {
            type: "exponential",
            delay: 5000,
        },
        removeOnComplete: true,
        removeOnFail: false,
    });

    logger.info(`Room generation job ${job.id} added to queue for hotel ${jobdata.hotelId} and room type ${jobdata.roomTypeId}`);

    return job;
}
