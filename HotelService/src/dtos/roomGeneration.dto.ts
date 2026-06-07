import {z} from 'zod';

export const RoomGeneratonRequestSchema = z.object({
    roomTypeId: z.number().positive(),
    hotelId: z.number().positive(),
    startDate: z.iso.datetime(),
    endDate: z.iso.datetime(),
    scheduleType:z.enum(['IMMEDIATE', 'SCHEDULED']).default('IMMEDIATE'),
    scheduleAt: z.iso.datetime().optional(),
    priceOverride: z.number().positive().optional(),
    batchSize: z.number().positive().default(100),
});

export const RoomGenerationJobSchema = z.object({
    roomTypeId: z.number().positive(),
    hotelId: z.number().positive(),
    startDate: z.iso.datetime(),
    endDate: z.iso.datetime(),
    priceOverride: z.number().positive().optional(),
    batchSize: z.number().positive().default(100),
    correlationId: z.string().optional(),
})

export type RoomGeneratonRequestSchemaType = z.infer<typeof RoomGeneratonRequestSchema>;
export type RoomGenerationJobSchemaType = z.infer<typeof RoomGenerationJobSchema>;

export interface RoomGenerationResponse {
    success: boolean;
    totalRoomsGenerated:number;
    totalDatesGenerated:number;
    errors?: string[];
    jobId?:string;
}
