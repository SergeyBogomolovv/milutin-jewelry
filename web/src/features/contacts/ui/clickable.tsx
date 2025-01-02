import { Button } from '@/shared/ui/button'
import { IconType } from 'react-icons'
import { toast } from 'sonner'

interface Props {
  Icon: IconType
  text: string
  value: string
}

export function Clickable({ Icon, text, value }: Props) {
  return (
    <Button
      variant='link'
      className='flex gap-2 tracking-wider font-bold text-md'
      onClick={async () => {
        await navigator.clipboard.writeText(value)
        toast.success('Скопировано в буфер обмена')
      }}
    >
      <Icon />
      {text}
    </Button>
  )
}
