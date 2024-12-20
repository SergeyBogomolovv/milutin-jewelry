'use server'
import { fetcher } from '@/shared/lib/fetcher'

export async function sendCode() {
  try {
    const res = await fetcher('/auth/send-code', { method: 'POST' })
    if (!res.ok) {
      return false
    }
    return true
  } catch (error) {
    return false
  }
}
