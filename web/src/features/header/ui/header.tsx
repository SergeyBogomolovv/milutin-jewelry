'use client'
import Link from 'next/link'
import Image from 'next/image'
import { ClipboardCopyIcon, UserRound } from 'lucide-react'
import { toast } from 'sonner'
import { Contacts } from '@/features/contacts'
import { Button } from '@/shared/ui/button'
import Logo from '@/assets/logo.svg'
import contacts from '@/assets/contacts.json'

export function Header() {
  const handleClick = async () => {
    await navigator.clipboard.writeText(contacts.phone)
    toast.success('Скопировано в буфер обмена')
  }

  return (
    <div className='bg-zinc-950 flex gap-4 items-end justify-between lg:pb-7 pb-6 pt-3 xl:px-12 lg:px-10 md:px-9 px-5'>
      <Link href='/'>
        <Image src={Logo} alt='Михаил Милютин' className='xl:w-72 lg:w-60 md:w-56 w-48' />
      </Link>
      <div className='flex gap-2'>
        <Contacts>
          <Button variant='ghost' className='text-lg flex font-bold'>
            <UserRound />
            <p className='hidden sm:block'>Контакты</p>
          </Button>
        </Contacts>
        <Button variant='ghost' className='text-lg hidden sm:flex font-bold' onClick={handleClick}>
          <ClipboardCopyIcon />
          {contacts.phoneText}
        </Button>
      </div>
    </div>
  )
}
