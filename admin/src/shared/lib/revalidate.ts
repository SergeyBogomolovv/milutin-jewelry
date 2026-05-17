import { revalidateTag } from 'next/cache'

export function revalidateCacheTag(tag: string) {
  const revalidate = revalidateTag as (tag: string, profile?: string) => void
  revalidate(tag, 'max')
}
