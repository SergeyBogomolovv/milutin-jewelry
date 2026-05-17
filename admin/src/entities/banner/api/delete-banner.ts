'use server'
import { fetchWithAuth } from '@/shared/lib/fetcher'
import { revalidateCacheTag } from '@/shared/lib/revalidate'

export async function deleteBanner(id: number) {
  await fetchWithAuth(`/banners/delete/${id}`, { method: 'DELETE' })
  revalidateCacheTag('banners')
}
