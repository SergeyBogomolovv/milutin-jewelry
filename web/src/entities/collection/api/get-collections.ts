'use server'
import { fetcher } from '@/shared/lib/fetcher'
import { z } from 'zod'
import { collectionSchema } from '../model/schema'

export async function getCollections() {
  try {
    const res = await fetcher('/collections/all', {
      next: { tags: ['collections'], revalidate: 60 },
      cache: 'force-cache',
    })
    const json = await res.json()
    const { data, success } = z.array(collectionSchema).safeParse(json)
    return { data, success }
  } catch (error) {
    return { success: false, data: undefined }
  }
}
