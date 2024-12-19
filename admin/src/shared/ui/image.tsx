'use client'
import Image, { ImageProps } from 'next/image'
import { IMAGE_URL } from '../constants'
import { useState } from 'react'

export function CustomImage({ src, alt, ...props }: ImageProps) {
  const [isLoaded, setLoaded] = useState(false)

  return (
    <Image
      {...props}
      alt={alt}
      src={`${IMAGE_URL}/${src}${isLoaded ? '_high.jpg' : '_low.jpg'}`}
      blurDataURL='/placeholder.jpg'
      placeholder='blur'
      onLoad={() => setLoaded(true)}
    />
  )
}
