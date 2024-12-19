'use server'
import { fetcher } from '@/shared/api/fetcher'
import { revalidateTag } from 'next/cache'

export async function deleteCollection(id: number) {
  await fetcher(`/collections/delete/${id}`, { method: 'DELETE' })
  revalidateTag('collections')
}
