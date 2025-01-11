'use client'
import { default as NextImage, ImageProps } from 'next/image'
import { IMAGE_URL } from '../constants'
import { useState } from 'react'

export function Image({ src, alt, ...props }: ImageProps) {
  const [loaded, setLoaded] = useState(false)
  return (
    <NextImage
      {...props}
      alt={alt}
      src={`${IMAGE_URL}/${src}${loaded ? '' : '_low'}.jpg`}
      blurDataURL='/placeholder.jpg'
      placeholder='blur'
      onLoadingComplete={() => setLoaded(true)}
    />
  )
}
