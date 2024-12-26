'use client'
import { useForm } from 'react-hook-form'
import { zodResolver } from '@hookform/resolvers/zod'
import { useRef, useState } from 'react'
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
import { Input } from '@/shared/ui/input'
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
import { NewItemFields } from '../model/new-item.schema'
import { CollectionItem } from '@/entities/collection-item'
import { UpdateItemFields, updateItemSchema } from '../model/update-item.schema'
import { updateCollectionItem } from '../api/update-item'
import { CustomImage } from '@/shared/ui/image'

export function UpdateCollectionItemForm({
  children,
  collectionItem,
}: {
  children: React.ReactNode
  collectionItem: CollectionItem
}) {
  const [open, setOpen] = useState(false)
  const form = useForm<UpdateItemFields>({
    resolver: zodResolver(updateItemSchema),
    defaultValues: {
      title: collectionItem.title || '',
      description: collectionItem.description || '',
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

  const onSubmit = async (data: NewItemFields) => {
    const ok = await updateCollectionItem(data, String(collectionItem.id))
    if (!ok) {
      toast.error('Ошибка обновления коллекции')
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
          <DialogTitle>Изменить украшение</DialogTitle>
        </DialogHeader>
        <Form {...form}>
          <form className='flex flex-col gap-4' onSubmit={form.handleSubmit(onSubmit)}>
            <FormField
              control={form.control}
              name='title'
              render={({ field }) => (
                <FormItem className='w-full'>
                  <FormLabel>Название</FormLabel>
                  <FormControl>
                    <Input placeholder='Название' {...field} />
                  </FormControl>
                  <FormDescription>Необязательное поле.</FormDescription>
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
                    <Input placeholder='Описание' {...field} />
                  </FormControl>
                  <FormDescription>Необязательное поле.</FormDescription>
                  <FormMessage />
                </FormItem>
              )}
            />
            <FormItem className='w-full flex flex-col gap-2'>
              <FormLabel>Изображение</FormLabel>
              {imagePreview ? (
                <Image
                  className='w-[70%] mx-auto rounded-md'
                  width={500}
                  height={500}
                  src={imagePreview}
                  alt='Uploaded Image'
                />
              ) : (
                <CustomImage
                  className='w-[70%] mx-auto rounded-md'
                  src={collectionItem.image_id}
                  alt='collection image'
                  width={500}
                  height={500}
                />
              )}
              <FormControl>
                <input
                  type='file'
                  accept='image/*'
                  ref={fileInputRef}
                  hidden
                  onChange={handleImageChange}
                  aria-hidden='true'
                  tabIndex={-1}
                />
              </FormControl>
              <Button type='button' variant='outline' onClick={() => fileInputRef.current?.click()}>
                <Paperclip />
                Изменить изображение
              </Button>
            </FormItem>
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
          </form>
        </Form>
      </DialogContent>
    </Dialog>
  )
}
