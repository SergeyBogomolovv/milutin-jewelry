'use server'
import { fetcher } from '@/shared/lib/fetcher'
import { revalidateTag } from 'next/cache'

interface DeleteBannerResponse {
  success: boolean
  error?: string
}

export async function deleteBanner(id: number): Promise<DeleteBannerResponse> {
  try {
    const res = await fetcher(`/banners/delete/${id}`, { method: 'DELETE' })

    if (!res.ok) {
      const errorMessage = await res.text()
      return { success: false, error: errorMessage || 'Ошибка при удалении баннера' }
    }

    revalidateTag('banners')
    return { success: true }
  } catch (error) {
    return { success: false, error: (error as Error).message || 'Произошла ошибка при удалении' }
  }
}
