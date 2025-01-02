import { z } from 'zod'

export const newItemSchema = z.object({
  title: z.optional(z.string()),
  description: z.optional(z.string()),
  image: typeof window === 'undefined' ? z.any() : z.instanceof(File),
})

export type NewItemFields = z.infer<typeof newItemSchema>
