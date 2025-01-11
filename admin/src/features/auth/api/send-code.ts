'use server'
import { fetcher } from '@/shared/lib/fetcher'

export async function sendCode() {
  await fetcher('/auth/send-code', { method: 'POST' })
}
