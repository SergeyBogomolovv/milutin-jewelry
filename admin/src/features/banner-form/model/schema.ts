import { z } from 'zod'

export const newBannerSchema = z.object({
  collection_id: z.optional(z.string()),
  image: typeof window === 'undefined' ? z.any() : z.instanceof(File),
  mobile_image: typeof window === 'undefined' ? z.any() : z.instanceof(File),
})

export type NewBannerFields = z.infer<typeof newBannerSchema>
