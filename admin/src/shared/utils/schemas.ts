import { z } from 'zod'

export const messageSchema = z.object({
  message: z.string(),
})

export type MessageResponse = z.infer<typeof messageSchema>
