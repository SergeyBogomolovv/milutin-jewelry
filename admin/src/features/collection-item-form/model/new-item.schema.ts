import { z } from 'zod'

export const newItemSchema = z.object({
  title: z.string().min(1, { message: 'Название не может быть пустым' }),
  description: z.optional(z.string()),
  image: typeof window === 'undefined' ? z.any() : z.instanceof(File),
})

export type NewItemFields = z.infer<typeof newItemSchema>
