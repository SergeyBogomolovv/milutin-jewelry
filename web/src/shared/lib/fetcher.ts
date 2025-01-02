import { API_URL } from './constants'

export function fetcher(path: string, opts: RequestInit = {}) {
  return fetch(`${API_URL}${path}`, opts)
}
