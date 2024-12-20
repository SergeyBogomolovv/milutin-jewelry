import { API_URL } from '../constants'
import { getToken } from '../utils/auth'

export async function fetcher(path: string, opts: RequestInit = {}) {
  const token = await getToken()
  return fetch(`${API_URL}${path}`, {
    ...opts,
    headers: {
      Authorization: `Bearer ${token}`,
      ...opts.headers,
    },
  })
}
