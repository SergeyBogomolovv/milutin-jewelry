'use server'
import { fetchWithAuth } from '@/shared/lib/fetcher'
import { revalidateTag } from 'next/cache'

export async function deleteBanner(id: number) {
  await fetchWithAuth(`/banners/delete/${id}`, { method: 'DELETE' })
  revalidateTag('banners')
}
