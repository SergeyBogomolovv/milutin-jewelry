import Image, { ImageProps } from 'next/image'
import { IMAGE_URL } from '../lib/constants'

export default function S3Image({ src, alt, ...props }: ImageProps) {
  return (
    <Image
      {...props}
      alt={alt}
      src={`${IMAGE_URL}/${src}_high.jpg`}
      blurDataURL={`${IMAGE_URL}/${src}_low.jpg`}
      placeholder='blur'
    />
  )
}
