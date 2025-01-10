import { fetcher } from '@/shared/lib/fetcher'
import { z } from 'zod'
import { itemSchema } from '../model/schema'

export async function getItems(collectionId: string) {
  try {
    const res = await fetcher(`/items/collection/${collectionId}`)
    const json = await res.json()
    const { success, data } = z.array(itemSchema).safeParse(json)
    return { success, data }
  } catch (error) {
    return { success: false, data: undefined }
  }
}
