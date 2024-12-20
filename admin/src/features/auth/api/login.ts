'use server'
import { cookies } from 'next/headers'
import { LoginFields, tokenReponse } from '../model/login-schema'
import { fetcher } from '@/shared/lib/fetcher'

export const login = async (body: LoginFields): Promise<boolean> => {
  try {
    const c = await cookies()
    const res = await fetcher('/auth/login', { method: 'POST', body: JSON.stringify(body) })
    if (!res.ok) {
      return false
    }
    const data = await res.json()
    const { token } = tokenReponse.parse(data)
    c.set('auth_token', token, {
      sameSite: 'strict',
      expires: new Date(Date.now() + 1000 * 60 * 60 * 24),
      path: '/',
      secure: true,
      httpOnly: true,
    })
    return true
  } catch (error) {
    return false
  }
}
