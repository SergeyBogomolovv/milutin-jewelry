'use client'
import { ImageProps, Image as AntImage } from 'antd'
import { Eye } from 'lucide-react'
import { IMAGE_URL } from '../lib/constants'

interface Props extends ImageProps {
  title?: string
  description?: string
}

export function Image({ src, title, description, ...props }: Props) {
  return (
    <AntImage
      loading='lazy'
      src={`${IMAGE_URL}/${src}_high.jpg`}
      placeholder={<AntImage preview={false} src={`${IMAGE_URL}/${src}_low.jpg`} {...props} />}
      preview={{
        mask: (
          <div className='flex items-center gap-1'>
            <Eye size={16} />
            Развернуть
          </div>
        ),
        // TODO: поправить
        toolbarRender: () => (
          <div className='flex flex-col gap-4 bg-background/60 py-2 px-6 items-center text-white rounded-lg'>
            <h3 className='font-bold text-lg'>{title}</h3>
            <p>{description}</p>
          </div>
        ),
      }}
      {...props}
    />
  )
}

export { default as PreviewGroup } from 'antd/es/image/PreviewGroup'
