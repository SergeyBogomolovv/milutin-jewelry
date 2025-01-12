'use client'
import Image, { ImageProps } from 'next/image'
import { IMAGE_URL } from '../lib/constants'
import { useState } from 'react'

export default function S3Image({ src, alt, ...props }: ImageProps) {
  const [loaded, setLoaded] = useState(false)
  return (
    <Image
      {...props}
      alt={alt}
      src={`${IMAGE_URL}/${src}${loaded ? '' : '_low'}.jpg`}
      blurDataURL='/placeholder.jpg'
      placeholder='blur'
      onLoad={() => setLoaded(true)}
    />
  )
}
