import { Contacts } from '@/features/contacts'
import { Button } from '@/shared/ui/button'
import { UserRound } from 'lucide-react'

export function Footer() {
  return (
    <div className='bg-zinc-950 flex items-center justify-center mt-10 py-6 px-4 sm:gap-10 gap-4'>
      <p className='sm:text-xl text-lg text-muted-foreground tracking-wide font-bold'>
        Посещение только по
        <br />
        предварительной записи.
      </p>
      <Contacts>
        <Button className='rounded-full aspect-square sm:size-12 size-10' aria-label='Контакты'>
          <UserRound />
        </Button>
      </Contacts>
    </div>
  )
}
