import { z } from 'zod'

export const loginSchema = z.object({
  email: z.string().email({ message: 'Введите корректный email' }),
  password: z.string().min(8, { message: 'Пароль должен содержать минимум 8 символов' }),
})

export type LoginFields = z.infer<typeof loginSchema>

export const tokenReponse = z.object({
  token: z.string(),
})
