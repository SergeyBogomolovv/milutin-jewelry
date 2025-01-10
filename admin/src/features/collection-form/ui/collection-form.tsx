'use client'
import { useForm } from 'react-hook-form'
import { NewCollectionFields, newCollectionSchema } from '../model/schema'
import { zodResolver } from '@hookform/resolvers/zod'
import { ChangeEventHandler, useRef, useState } from 'react'
import { Form, FormItem, FormLabel } from '@/shared/ui/form'
import { Button } from '@/shared/ui/button'
import { Paperclip } from 'lucide-react'
import Image from 'next/image'
import { createCollection } from '../api/create-collection'
import { toast } from 'sonner'
import {
  Dialog,
  DialogClose,
  DialogContent,
  DialogFooter,
  DialogHeader,
  DialogTitle,
  DialogTrigger,
} from '@/shared/ui/dialog'
import { FormInputField } from '@/shared/ui/form-input-field'
import HiddenInput from '@/shared/ui/hidden-input'

export function CollectionForm({ children }: { children: React.ReactNode }) {
  const [open, setOpen] = useState(false)
  const [imagePreview, setImagePreview] = useState<string | null>(null)

  const fileInputRef = useRef<HTMLInputElement | null>(null)

  const form = useForm<NewCollectionFields>({
    resolver: zodResolver(newCollectionSchema),
    defaultValues: {
      title: '',
      description: '',
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

  const onSubmit = async (data: NewCollectionFields) => {
    try {
      await createCollection(data)
      toast.success('Коллекция создана')
      setImagePreview(null)
      form.reset()
      setOpen(false)
    } catch (error) {
      toast.error('Произошла ошибка создания коллекции')
    }
  }

  return (
    <Dialog open={open} onOpenChange={setOpen}>
      <DialogTrigger asChild>{children}</DialogTrigger>
      <DialogContent className='max-h-[90vh] overflow-scroll'>
        <DialogHeader>
          <DialogTitle>Новая коллекция</DialogTitle>
        </DialogHeader>
        <Form {...form}>
          <form className='flex flex-col gap-4' onSubmit={form.handleSubmit(onSubmit)}>
            <FormInputField
              control={form.control}
              name='title'
              label='Название'
              placeholder='Название'
              description='Название новой коллекции.'
            />
            <FormInputField
              control={form.control}
              name='description'
              label='Описание'
              placeholder='Описание'
              description='Необязательное поле.'
            />
            <FormItem className='w-full flex flex-col gap-2'>
              <FormLabel>Изображение</FormLabel>
              {imagePreview && (
                <Image
                  className='w-[70%] mx-auto rounded-md'
                  width={500}
                  height={500}
                  src={imagePreview}
                  alt='Загруженное изображение'
                />
              )}
              <HiddenInput ref={fileInputRef} handleImageChange={handleImageChange} />
              <Button type='button' variant='outline' onClick={() => fileInputRef.current?.click()}>
                <Paperclip />
                {imagePreview ? 'Изменить изображение' : 'Прикрепить изображение'}
              </Button>
            </FormItem>
            <DialogFooter className='flex items-center gap-2'>
              <Button className='w-full' disabled={form.formState.isSubmitting} type='submit'>
                {form.formState.isSubmitting ? 'Создание...' : 'Создать'}
              </Button>
              <DialogClose asChild>
                <Button className='w-full' type='button' variant='secondary'>
                  Отмена
                </Button>
              </DialogClose>
            </DialogFooter>
          </form>
        </Form>
      </DialogContent>
    </Dialog>
  )
}
