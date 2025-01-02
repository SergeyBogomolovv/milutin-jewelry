'use client'
import { default as NextImage, ImageProps } from 'next/image'
import { useState } from 'react'
import { IMAGE_URL } from '../lib/constants'

export function Image({ src, alt, ...props }: ImageProps) {
  const [isLoaded, setLoaded] = useState(false)

  return (
    <NextImage
      {...props}
      alt={alt}
      src={`${IMAGE_URL}/${src}${isLoaded ? '_high.jpg' : '_low.jpg'}`}
      blurDataURL='/placeholder.jpg'
      placeholder='blur'
      onLoad={() => setLoaded(true)}
    />
  )
}
