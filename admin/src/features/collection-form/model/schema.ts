import { z } from 'zod'

export const newCollectionSchema = z.object({
  title: z.string().min(1, { message: 'Название не может быть пустым' }),
  description: z.optional(z.string()),
  image: typeof window === 'undefined' ? z.any() : z.instanceof(File),
})

export type NewCollectionFields = z.infer<typeof newCollectionSchema>
