'use server'
import jwt from 'jsonwebtoken'
import { JWT_SECRET } from '../constants'

export function checkAuth(token: string): boolean {
  try {
    const decoded = jwt.verify(token, JWT_SECRET)
    return !!decoded
  } catch (error) {
    return false
  }
}
