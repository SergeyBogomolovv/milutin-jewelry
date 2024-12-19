import { z } from 'zod'

export const collectionSchema = z.object({
  id: z.number(),
  title: z.string(),
  description: z.nullable(z.string()),
  image_id: z.string(),
})

export type Collection = z.infer<typeof collectionSchema>
