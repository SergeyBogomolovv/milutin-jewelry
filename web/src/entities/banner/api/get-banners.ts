'use server'
import { fetcher } from '@/shared/lib/fetcher'
import { bannerSchema } from '../model/banner'
import { z } from 'zod'

export async function getBanners() {
  try {
    const res = await fetcher('/banners/all', {
      next: { tags: ['banners'], revalidate: 60 },
      cache: 'force-cache',
    })
    const data = await res.json()
    return z.array(bannerSchema).safeParse(data)
  } catch (error) {
    return { success: false, data: undefined }
  }
}
