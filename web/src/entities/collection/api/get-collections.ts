import { fetcher } from '@/shared/lib/fetcher'
import { z } from 'zod'
import { collectionSchema } from '../model/schema'

export async function getCollections() {
  const res = await fetcher('/collections/all', {
    cache: 'force-cache',
    next: { revalidate: 300 },
  })
  const json = await res.json()
  return z.array(collectionSchema).parse(json)
}
