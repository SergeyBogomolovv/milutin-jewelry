'use client'
import { ImageProps, Image as AntImage } from 'antd'
import { Eye } from 'lucide-react'
import { IMAGE_URL } from '../lib/constants'

export function Image({ src, ...props }: ImageProps) {
  return (
    <AntImage
      src={`${IMAGE_URL}/${src}.jpg`}
      fallback='/placeholder.jpg'
      placeholder={<AntImage src={`${IMAGE_URL}/${src}_low.jpg`} {...props} />}
      preview={{
        mask: (
          <div className='flex items-center gap-1'>
            <Eye size={16} />
            Развернуть
          </div>
        ),
      }}
      {...props}
    />
  )
}

export { default as PreviewGroup } from 'antd/es/image/PreviewGroup'
