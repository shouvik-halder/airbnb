import {z} from "zod";

export const pingSchema = z.object({
    fname: z.string(),
    lname: z.string()
});