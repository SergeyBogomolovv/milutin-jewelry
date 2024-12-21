'use client'
import { zodResolver } from '@hookform/resolvers/zod'
import { useForm } from 'react-hook-form'
import { InfoFields, infoSchema } from '../model/info-schema'
import { Form, FormControl, FormField, FormItem, FormLabel, FormMessage } from '@/shared/ui/form'
import { Button } from '@/shared/ui/button'
import { updateCollection } from '../api/update-collection'
import { toast } from 'sonner'
import { Input } from '@/shared/ui/input'
import { Collection } from '@/entities/collection'
import { CustomImage } from '@/shared/ui/image'
import { Paperclip } from 'lucide-react'
import { useRef, useState } from 'react'
import Image from 'next/image'
import { Textarea } from '@/shared/ui/textarea'

export function CollectionInfo({ collection }: { collection: Collection }) {
  const form = useForm<InfoFields>({
    resolver: zodResolver(infoSchema),
    defaultValues: {
      title: collection.title,
      description: collection.description || '',
    },
  })

  const [imagePreview, setImagePreview] = useState<string | null>(null)
  const fileInputRef = useRef<HTMLInputElement | null>(null)

  const handleImageChange = (event: React.ChangeEvent<HTMLInputElement>) => {
    const file = event.target.files?.[0]
    if (file) {
      const imageUrl = URL.createObjectURL(file)
      setImagePreview(imageUrl)
      form.setValue('image', file)
    }
  }

  const onSubmit = async (data: InfoFields) => {
    const ok = await updateCollection(data, collection.id)
    if (!ok) {
      toast.error('Ошибка обновления коллекции')
      return
    }
    toast.success('Коллекция обновлена')
  }

  return (
    <Form {...form}>
      <form className='flex md:flex-row md:gap-6 gap-4' onSubmit={form.handleSubmit(onSubmit)}>
        <div className='flex flex-col items-center gap-2'>
          {imagePreview ? (
            <Image
              className='md:size-44 sm:size-40 size-36 object-cover rounded-md'
              width={500}
              height={500}
              src={imagePreview}
              alt='Uploaded Image'
            />
          ) : (
            <CustomImage
              className='md:size-44 sm:size-40 size-36 object-cover rounded-md'
              src={collection.image_id}
              alt='Collection Image'
              width={500}
              height={500}
            />
          )}
          <input
            type='file'
            accept='image/*'
            ref={fileInputRef}
            hidden
            onChange={handleImageChange}
            aria-hidden='true'
            tabIndex={-1}
          />
          <Button
            className='w-full'
            type='button'
            variant='outline'
            onClick={() => fileInputRef.current?.click()}
          >
            <Paperclip />
            Изменить
          </Button>
        </div>
        <div className='flex flex-col w-full md:gap-2 gap-4 justify-between'>
          <FormField
            control={form.control}
            name='title'
            render={({ field }) => (
              <FormItem className='w-full'>
                <FormLabel>Название</FormLabel>
                <FormControl>
                  <Input placeholder='Название' {...field} />
                </FormControl>
                <FormMessage />
              </FormItem>
            )}
          />
          <FormField
            control={form.control}
            name='description'
            render={({ field }) => (
              <FormItem className='w-full'>
                <FormLabel>Описание</FormLabel>
                <FormControl>
                  <Textarea className='resize-none' placeholder='Описание' {...field} />
                </FormControl>
                <FormMessage />
              </FormItem>
            )}
          />
          <Button className='w-fit' disabled={form.formState.isSubmitting} type='submit'>
            {form.formState.isSubmitting ? 'Сохранение...' : 'Сохранить'}
          </Button>
        </div>
      </form>
    </Form>
  )
}
