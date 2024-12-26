import { z } from 'zod'

export const updateItemSchema = z.object({
  title: z.optional(z.string()),
  description: z.optional(z.string()),
  image: typeof window === 'undefined' ? z.optional(z.any()) : z.optional(z.instanceof(File)),
})

export type UpdateItemFields = z.infer<typeof updateItemSchema>
