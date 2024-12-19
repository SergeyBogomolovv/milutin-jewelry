'use server'
import { cookies } from 'next/headers'
import { LoginFields, tokenReponse } from '../model/login-schema'
import { fetcher } from '@/shared/api/fetcher'

export const login = async (data: LoginFields): Promise<boolean> => {
  const c = await cookies()
  try {
    const res = await fetcher('/auth/login', { method: 'POST', body: JSON.stringify(data) })
    if (res.status === 400) {
      return false
    }
    const json = await res.json()
    const { token } = tokenReponse.parse(json)
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
