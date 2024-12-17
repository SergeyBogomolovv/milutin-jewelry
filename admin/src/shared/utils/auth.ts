'use server'
import { jwtVerify } from 'jose'
import { JWT_SECRET } from '../constants'

const secret = new TextEncoder().encode(JWT_SECRET)

export async function checkAuth(token: string): Promise<boolean> {
  try {
    const decoded = await jwtVerify(token, secret)
    return !!decoded
  } catch (error) {
    console.log(error)
    return false
  }
}
