import { z } from 'zod'

export const updateCollectionSchema = z.object({
  title: z.string().min(1, { message: 'Название не может быть пустым' }),
  description: z.optional(z.string()),
  image: typeof window === 'undefined' ? z.optional(z.any()) : z.optional(z.instanceof(File)),
})

export type UpdateCollectionFields = z.infer<typeof updateCollectionSchema>
