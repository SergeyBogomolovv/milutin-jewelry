'use server'
import { fetcher } from '@/shared/lib/fetcher'
import { bannerSchema } from '../model/banner'
import { z } from 'zod'

export async function getBanners() {
  try {
    const res = await fetcher('/banners/all', { next: { tags: ['banners'] } })
    const data = await res.json()
    return z.array(bannerSchema).safeParse(data)
  } catch (error) {
    return { success: false, data: undefined }
  }
}
