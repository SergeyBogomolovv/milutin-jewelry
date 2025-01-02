import { z } from 'zod'

export const collectionSchema = z.object({
  id: z.number(),
  title: z.string(),
  description: z.optional(z.string()),
  image_id: z.string(),
  created_at: z.string(),
})

export type Collection = z.infer<typeof collectionSchema>
