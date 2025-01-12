import { fetcher } from '@/shared/lib/fetcher'
import { bannerSchema } from '../model/banner'
import { z } from 'zod'

export async function getBanners() {
  const res = await fetcher('/banners/all', { cache: 'force-cache', next: { revalidate: 300 } })
  const data = await res.json()
  return z.array(bannerSchema).parse(data)
}
