import { API_URL } from '../constants'
import { getToken } from '../utils/auth'

export async function fetchWithAuth(path: string, opts: RequestInit = {}) {
  const token = await getToken()
  return fetch(`${API_URL}${path}`, {
    ...opts,
    headers: {
      Authorization: `Bearer ${token}`,
      ...opts.headers,
    },
  })
}

export async function fetcher(path: string, opts: RequestInit = {}) {
  return fetch(API_URL + path, opts)
}
