'use server'
import { fetcher } from '@/shared/lib/fetcher'
import { z } from 'zod'
import { itemSchema } from '../model/collection-item'

export async function getItemsByCollection(id: string) {
  try {
    const res = await fetcher(`/items/collection/${id}`, {
      next: { tags: ['items', id] },
      cache: 'force-cache',
    })
    const data = await res.json()
    return z.array(itemSchema).safeParse(data)
  } catch (error) {
    return { success: false, data: undefined }
  }
}
