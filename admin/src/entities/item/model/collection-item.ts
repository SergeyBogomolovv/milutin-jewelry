import { z } from 'zod'

export const itemSchema = z.object({
  id: z.number(),
  collection_id: z.number(),
  title: z.optional(z.string()),
  description: z.optional(z.string()),
  image_id: z.string(),
  created_at: z.string(),
})

export type Item = z.infer<typeof itemSchema>
