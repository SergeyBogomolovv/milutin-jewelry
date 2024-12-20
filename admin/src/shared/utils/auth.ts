'use server'
import { jwtVerify } from 'jose'
import { JWT_SECRET } from '../constants'
import { cookies } from 'next/headers'

const secret = new TextEncoder().encode(JWT_SECRET)

export async function checkAuth(token: string): Promise<boolean> {
  try {
    const decoded = await jwtVerify(token, secret)
    return !!decoded
  } catch (error) {
    return false
  }
}

export async function getToken() {
  const c = await cookies()
  return c.get('auth_token')?.value
}
