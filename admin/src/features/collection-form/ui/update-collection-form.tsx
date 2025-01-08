'use client'
import {
  Dialog,
  DialogClose,
  DialogContent,
  DialogFooter,
  DialogHeader,
  DialogTitle,
  DialogTrigger,
} from '@/shared/ui/dialog'
import { zodResolver } from '@hookform/resolvers/zod'
import { useForm } from 'react-hook-form'
import { UpdateCollectionFields, updateCollectionSchema } from '../model/update.schema'
import { Form, FormControl, FormField, FormItem, FormLabel, FormMessage } from '@/shared/ui/form'
import { Button } from '@/shared/ui/button'
import { updateCollection } from '../api/update-collection'
import { toast } from 'sonner'
import { Collection } from '@/entities/collection'
import { CustomImage } from '@/shared/ui/image'
import { Paperclip } from 'lucide-react'
import { ChangeEventHandler, useRef, useState } from 'react'
import Image from 'next/image'
import { Textarea } from '@/shared/ui/textarea'
import HiddenInput from '@/shared/ui/hidden-input'
import { FormInputField } from '@/shared/ui/form-input-field'

export function UpdateCollectionForm({
  collection,
  children,
}: {
  collection: Collection
  children: React.ReactNode
}) {
  const [open, setOpen] = useState(false)
  const [imagePreview, setImagePreview] = useState<string | null>(null)

  const fileInputRef = useRef<HTMLInputElement | null>(null)

  const form = useForm<UpdateCollectionFields>({
    resolver: zodResolver(updateCollectionSchema),
    defaultValues: {
      title: collection.title,
      description: collection.description || '',
    },
  })

  const handleImageChange: ChangeEventHandler<HTMLInputElement> = (event) => {
    const file = event.target.files?.[0]
    if (file) {
      const imageUrl = URL.createObjectURL(file)
      setImagePreview(imageUrl)
      form.setValue('image', file)
      return () => URL.revokeObjectURL(imageUrl)
    }
  }

  const onSubmit = async (data: UpdateCollectionFields) => {
    const result = await updateCollection(data, collection.id)
    if (!result.success) {
      toast.error(`Ошибка обновления коллекции: ${result.error || 'Неизвестная ошибка'}`)
      return
    }
    setImagePreview(null)
    setOpen(false)
    toast.success('Коллекция обновлена')
  }

  return (
    <Dialog open={open} onOpenChange={setOpen}>
      <DialogTrigger asChild>{children}</DialogTrigger>
      <DialogContent className='max-h-[90vh] overflow-scroll'>
        <DialogHeader>
          <DialogTitle>Изменить информацию</DialogTitle>
        </DialogHeader>
        <Form {...form}>
          <form className='flex flex-col gap-3' onSubmit={form.handleSubmit(onSubmit)}>
            <div className='flex flex-col w-full md:gap-2 gap-4 justify-between'>
              <FormInputField
                control={form.control}
                name='title'
                label='Название'
                placeholder='Название'
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
              <FormItem className='w-full flex flex-col gap-2'>
                <FormLabel>Изображение</FormLabel>
                {imagePreview ? (
                  <Image
                    className='w-full object-cover rounded-md'
                    width={500}
                    height={500}
                    src={imagePreview}
                    alt='Загруженная картинка'
                  />
                ) : (
                  <CustomImage
                    className='w-full object-cover rounded-md'
                    src={collection.image_id}
                    alt={collection.title}
                    width={500}
                    height={500}
                  />
                )}
                <HiddenInput ref={fileInputRef} handleImageChange={handleImageChange} />
              </FormItem>
              <Button
                className='w-full'
                type='button'
                variant='outline'
                onClick={() => fileInputRef.current?.click()}
              >
                <Paperclip />
                Изменить
              </Button>
              <DialogFooter className='flex items-center gap-2'>
                <Button className='w-full' disabled={form.formState.isSubmitting} type='submit'>
                  {form.formState.isSubmitting ? 'Сохранение...' : 'Сохранить'}
                </Button>
                <DialogClose asChild>
                  <Button className='w-full' type='button' variant='secondary'>
                    Отмена
                  </Button>
                </DialogClose>
              </DialogFooter>
            </div>
          </form>
        </Form>
      </DialogContent>
    </Dialog>
  )
}
