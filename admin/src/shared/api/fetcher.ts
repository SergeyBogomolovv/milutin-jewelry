import { Schema } from 'zod'

export async function fetcher<T>(
  url: string,
  schema: Schema<T>,
  options: RequestInit = {},
): Promise<T> {
  const { credentials = 'include', body, headers } = options

  const res = await fetch(url, {
    ...options,
    headers: {
      'Content-Type': 'application/json',
      ...headers,
    },
    body: body ? JSON.stringify(body) : undefined,
    credentials,
  })

  if (!res.ok) {
    const errorBody = await res.text()
    throw new Error(`Request failed with status ${res.status}: ${errorBody}`)
  }

  const json = await res.json()

  return schema.parse(json)
}
