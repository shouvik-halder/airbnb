import z from "zod";

export const createHotelSchema = z.object({
  name: z.string().min(1),
  address: z.string().min(1),
  location: z.string().min(1),
  rating: z.number().optional(),
  ratingCount: z.number().optional(),
});

export const getHotelByIdSchema = z.object({
    id: z.coerce.number().min(1)
})

export const updateHotelSchema = z.object({
  name:z.string().min(1).optional(),
  rating:z.number().min(1).optional(),
  ratingCount: z.number().min(1).optional()
})
