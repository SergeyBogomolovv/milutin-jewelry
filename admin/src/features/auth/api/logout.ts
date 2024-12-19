'use server'
import { cookies } from 'next/headers'

export const logout = async () => {
  const c = await cookies()
  c.delete('auth_token')
}
