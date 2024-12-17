import { Schema } from 'zod'
import { API_URL } from '../constants'

export async function fetcher<T>(
  path: string,
  schema: Schema<T>,
  options: RequestInit = {},
): Promise<T> {
  const { credentials = 'include', headers } = options

  const res = await fetch(`${API_URL}${path}`, {
    ...options,
    headers: {
      'Content-Type': 'application/json',
      ...headers,
    },
    credentials,
  })

  if (!res.ok) {
    const errorBody = await res.text()
    throw new Error(`Request failed with status ${res.status}: ${errorBody}`)
  }

  const json = await res.json()

  return schema.parse(json)
}
