import { Button } from '@/shared/ui/button'
import Link from 'next/link'
import { IconType } from 'react-icons'

interface Props {
  Icon: IconType
  text: string
  url: string
}

export function Linkable({ Icon, text, url }: Props) {
  return (
    <Link target='_blank' href={url}>
      <Button variant='link' className='flex gap-2 tracking-wider font-bold text-md'>
        <Icon />
        {text}
      </Button>
    </Link>
  )
}
