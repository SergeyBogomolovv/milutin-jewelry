'use client'
import {
  Dialog,
  DialogContent,
  DialogDescription,
  DialogHeader,
  DialogTitle,
  DialogTrigger,
} from '@/shared/ui/dialog'
import { default as NextImage, ImageProps } from 'next/image'
import { useState } from 'react'
import { VisuallyHidden } from '@radix-ui/react-visually-hidden'
import { IMAGE_URL } from '../lib/constants'
import { cn } from '../lib/utils'

interface Props extends ImageProps {
  className?: string
  previewClassName?: string
}
export function Image({ src, alt, className, previewClassName, ...props }: Props) {
  const [isLoaded, setLoaded] = useState(false)

  return (
    <Dialog>
      <DialogTrigger asChild>
        <NextImage
          {...props}
          alt={alt}
          src={`${IMAGE_URL}/${src}${isLoaded ? '_high.jpg' : '_low.jpg'}`}
          blurDataURL='/placeholder.jpg'
          placeholder='blur'
          onLoad={() => setLoaded(true)}
          className={cn('rounded-lg cursor-pointer', className)}
        />
      </DialogTrigger>
      <DialogContent className='p-0 m-0 border-none bg-transparent flex items-center justify-center shadow-none'>
        <VisuallyHidden>
          <DialogTitle>{alt}</DialogTitle>
          <DialogDescription>{alt}</DialogDescription>
        </VisuallyHidden>
        <DialogHeader className='max-h-[90vh] overflow-auto flex justify-center'>
          <NextImage
            {...props}
            alt={alt}
            src={`${IMAGE_URL}/${src}${isLoaded ? '_high.jpg' : '_low.jpg'}`}
            blurDataURL='/placeholder.jpg'
            placeholder='blur'
            onLoad={() => setLoaded(true)}
            className={cn('object-cover relative rounded-lg max-h-[90vh]', previewClassName)}
          />
        </DialogHeader>
      </DialogContent>
    </Dialog>
  )
}
