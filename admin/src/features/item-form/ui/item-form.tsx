'use client'
import { useForm } from 'react-hook-form'
import { zodResolver } from '@hookform/resolvers/zod'
import { useRef, useState } from 'react'
import { Form, FormItem, FormLabel } from '@/shared/ui/form'
import { Button } from '@/shared/ui/button'
import { Paperclip } from 'lucide-react'
import Image from 'next/image'
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
import { NewItemFields, newItemSchema } from '../model/new-item.schema'
import { createItem } from '../api/create-item'
import { FormInputField } from '@/shared/ui/form-input-field'
import HiddenInput from '@/shared/ui/hidden-input'

export function ItemForm({ children, id }: { children: React.ReactNode; id: string }) {
  const [open, setOpen] = useState(false)
  const [imagePreview, setImagePreview] = useState<string | null>(null)

  const fileInputRef = useRef<HTMLInputElement | null>(null)

  const form = useForm<NewItemFields>({
    resolver: zodResolver(newItemSchema),
    defaultValues: {
      title: '',
      description: '',
    },
  })

  const handleImageChange = (event: React.ChangeEvent<HTMLInputElement>) => {
    const file = event.target.files?.[0]
    if (file) {
      const imageUrl = URL.createObjectURL(file)
      setImagePreview(imageUrl)
      form.setValue('image', file)
      return () => URL.revokeObjectURL(imageUrl)
    }
  }

  const onSubmit = async (data: NewItemFields) => {
    const ok = await createItem(data, id)
    if (!ok) {
      toast.error('Ошибка создания украшения')
      return
    }
    setImagePreview(null)
    form.reset()
    setOpen(false)
  }

  return (
    <Dialog open={open} onOpenChange={setOpen}>
      <DialogTrigger asChild>{children}</DialogTrigger>
      <DialogContent className='max-h-[90vh] overflow-scroll'>
        <DialogHeader>
          <DialogTitle>Новое украшение</DialogTitle>
        </DialogHeader>
        <Form {...form}>
          <form className='flex flex-col gap-4' onSubmit={form.handleSubmit(onSubmit)}>
            <FormInputField
              control={form.control}
              name='title'
              label='Название'
              placeholder='Название'
              description='Необязательное поле.'
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
                  alt='Uploaded Image'
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
