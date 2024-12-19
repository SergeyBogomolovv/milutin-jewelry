'use server'
import { cookies } from 'next/headers'
import { API_URL } from '../constants'

export async function fetcher(path: string, options: RequestInit = {}) {
  const c = await cookies()
  return fetch(`${API_URL}${path}`, {
    ...options,
    headers: {
      Authorization: `Bearer ${c.get('auth_token')?.value}`,
      'Content-Type': 'application/json',
      ...options.headers,
    },
    credentials: 'include',
  })
}
