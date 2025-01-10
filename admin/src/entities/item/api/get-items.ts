'use server'
import { fetcher } from '@/shared/lib/fetcher'
import { z } from 'zod'
import { itemSchema } from '../model/collection-item'

export async function getItemsByCollection(id: string) {
  const res = await fetcher(`/items/collection/${id}`, { next: { tags: ['items'] } })
  const data = await res.json()
  return z.array(itemSchema).parse(data)
}
