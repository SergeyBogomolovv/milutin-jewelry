import { z } from 'zod'

export const loginSchema = z.object({
  code: z.string().length(6, { message: 'Код должен содержать строго 6 символов' }),
})

export type LoginFields = z.infer<typeof loginSchema>
