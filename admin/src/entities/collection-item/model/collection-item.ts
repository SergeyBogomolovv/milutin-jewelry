import { z } from 'zod'

export const collectionItemSchema = z.object({
  id: z.number(),
  collection_id: z.number(),
  title: z.string(),
  description: z.nullable(z.string()),
  image_id: z.string(),
})

export type CollectionItem = z.infer<typeof collectionItemSchema>
