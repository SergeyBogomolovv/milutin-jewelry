'use server'
import { collectionSchema } from '@/entities/collection'
import { fetcher } from '@/shared/api/fetcher'
import { z } from 'zod'

export async function getCollections() {
  try {
    const res = await fetcher('/collections/all', { next: { tags: ['collections'] } })
    if (!res.ok) {
      return []
    }
    const data = await res.json()
    return z.array(collectionSchema).parse(data)
  } catch (error) {
    return []
  }
}
