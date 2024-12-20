'use server'
import { fetcher } from '@/shared/lib/fetcher'
import { InfoFields } from '../model/info-schema'
import { revalidateTag } from 'next/cache'

export const updateCollection = async (data: InfoFields, id: number) => {
  try {
    const formData = new FormData()
    if (data.title) formData.append('title', data.title)
    if (data.description) formData.append('description', data.description)
    if (data.image) formData.append('image', data.image)
    const res = await fetcher(`/collections/update/${id}`, {
      method: 'PUT',
      body: formData,
    })
    if (!res.ok) {
      return false
    }
    revalidateTag('collections')
    return true
  } catch (error) {
    return false
  }
}
