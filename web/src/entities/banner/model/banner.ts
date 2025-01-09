import { z } from 'zod'

export const bannerSchema = z.object({
  id: z.number(),
  image_id: z.string(),
  mobile_image_id: z.string(),
  collection_id: z.optional(z.number()),
})

export type Banner = z.infer<typeof bannerSchema>
