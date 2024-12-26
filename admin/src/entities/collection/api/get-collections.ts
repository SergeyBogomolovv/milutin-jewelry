'use server'
import { collectionSchema } from '@/entities/collection'
import { fetcher } from '@/shared/lib/fetcher'
import { z } from 'zod'

export async function getCollections() {
  try {
    const res = await fetcher('/collections/all', {
      next: { tags: ['collections'] },
      cache: 'force-cache',
    })
    const data = await res.json()
    return z.array(collectionSchema).safeParse(data)
  } catch (error) {
    return { success: false, data: undefined }
  }
}
