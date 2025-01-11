import Image, { ImageProps } from 'next/image'
import { IMAGE_URL } from '../constants'

export function CustomImage({ src, alt, ...props }: ImageProps) {
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
