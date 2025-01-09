'use client'
import { useForm } from 'react-hook-form'
import { NewBannerFields, newBannerSchema } from '../model/schema'
import { zodResolver } from '@hookform/resolvers/zod'
import { ChangeEventHandler, useRef, useState } from 'react'
import {
  Form,
  FormControl,
  FormDescription,
  FormField,
  FormItem,
  FormLabel,
  FormMessage,
} from '@/shared/ui/form'
import { Button } from '@/shared/ui/button'
import { Paperclip } from 'lucide-react'
import Image from 'next/image'
import { Select, SelectContent, SelectItem, SelectTrigger, SelectValue } from '@/shared/ui/select'
import {
  Dialog,
  DialogClose,
  DialogContent,
  DialogFooter,
  DialogHeader,
  DialogTitle,
  DialogTrigger,
} from '@/shared/ui/dialog'
import HiddenInput from '@/shared/ui/hidden-input'
import { Collection } from '@/entities/collection'
import { createBanner } from '../api/create-banner'
import { toast } from 'sonner'

export function BannerForm({
  children,
  collections,
}: {
  children: React.ReactNode
  collections: Collection[]
}) {
  const [open, setOpen] = useState(false)
  const [imagePreview, setImagePreview] = useState<string | null>(null)
  const [mobileImagePreview, setMobileImagePreview] = useState<string | null>(null)

  const fileInputRef = useRef<HTMLInputElement | null>(null)
  const mobileFileInputRef = useRef<HTMLInputElement | null>(null)

  const form = useForm<NewBannerFields>({
    resolver: zodResolver(newBannerSchema),
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

  const handleMobileImageChange: ChangeEventHandler<HTMLInputElement> = (event) => {
    const file = event.target.files?.[0]
    if (file) {
      const imageUrl = URL.createObjectURL(file)
      setMobileImagePreview(imageUrl)
      form.setValue('mobile_image', file)
      return () => URL.revokeObjectURL(imageUrl)
    }
  }

  const onSubmit = async (data: NewBannerFields) => {
    const result = await createBanner(data)
    if (!result.success) {
      toast.error(`Ошибка создания баннера: ${result.error || 'Неизвестная ошибка'}`)
      return
    }
    toast.success('Баннер успешно создан!')
    setImagePreview(null)
    form.reset()
    setOpen(false)
  }

  return (
    <Dialog open={open} onOpenChange={setOpen}>
      <DialogTrigger asChild>{children}</DialogTrigger>
      <DialogContent className='max-h-[90vh] overflow-scroll'>
        <DialogHeader>
          <DialogTitle>Новый баннер</DialogTitle>
        </DialogHeader>
        <Form {...form}>
          <form className='flex flex-col gap-4' onSubmit={form.handleSubmit(onSubmit)}>
            <FormField
              control={form.control}
              name='collection_id'
              render={({ field }) => (
                <FormItem>
                  <FormLabel>Коллекция</FormLabel>
                  <Select onValueChange={field.onChange} defaultValue={field.value}>
                    <FormControl>
                      <SelectTrigger>
                        <SelectValue placeholder='Выберите коллекцию' />
                      </SelectTrigger>
                    </FormControl>
                    <SelectContent>
                      {collections?.map((collection) => (
                        <SelectItem key={collection.id} value={String(collection.id)}>
                          {collection.title}
                        </SelectItem>
                      ))}
                    </SelectContent>
                  </Select>
                  <FormDescription>
                    Выберите, если хотите чтобы баннер вел на страницу коллекции.
                  </FormDescription>
                  <FormMessage />
                </FormItem>
              )}
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
              <FormDescription>
                Изображение, которое будет отображаться на устройствах с большим экраном.
              </FormDescription>
            </FormItem>

            <FormItem className='w-full flex flex-col gap-2'>
              <FormLabel>Изображение на мобильных устройствах</FormLabel>
              {mobileImagePreview && (
                <Image
                  className='w-[70%] mx-auto rounded-md'
                  width={500}
                  height={500}
                  src={mobileImagePreview}
                  alt='Загруженное изображение'
                />
              )}
              <HiddenInput ref={mobileFileInputRef} handleImageChange={handleMobileImageChange} />
              <Button
                type='button'
                variant='outline'
                onClick={() => mobileFileInputRef.current?.click()}
              >
                <Paperclip />
                {mobileImagePreview ? 'Изменить изображение' : 'Прикрепить изображение'}
              </Button>
              <FormDescription>
                Изображение, которое будет отображаться на более мелких устройствах.
              </FormDescription>
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
