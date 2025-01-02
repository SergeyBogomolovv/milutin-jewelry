'use client'
import { Sheet, SheetContent, SheetHeader, SheetTitle, SheetTrigger } from '@/shared/ui/sheet'
import { PropsWithChildren } from 'react'
import { Separator } from '@/shared/ui/separator'
import { FaMap, FaTelegram, FaVk, FaWhatsapp } from 'react-icons/fa6'
import { FaMapMarkerAlt, FaPhoneAlt } from 'react-icons/fa'
import { IoMail } from 'react-icons/io5'
import { Clickable } from './clickable'
import { Linkable } from './linkable'
import content from '@/assets/contacts.json'

export function Contacts({ children }: PropsWithChildren) {
  return (
    <Sheet>
      <SheetTrigger asChild>{children}</SheetTrigger>
      <SheetContent side='top'>
        <SheetHeader>
          <SheetTitle className='tracking-wide text-center text-3xl'>Контакты</SheetTitle>
        </SheetHeader>
        <div className='flex flex-col gap-5 mt-5'>
          <Separator />
          <div className='flex flex-col lg:flex-row justify-between xl:w-[70%] lg:w-[90%] lg:mx-auto lg:gap-5 gap-3'>
            <div className='space-y-2'>
              <Clickable Icon={IoMail} text={content.mail1} value={content.mail1} />
              <Clickable Icon={IoMail} text={content.mail2} value={content.mail2} />
            </div>
            <Separator className='lg:hidden' />
            <div className='space-y-2'>
              <Clickable Icon={FaWhatsapp} text='WhatsApp' value={content.whatsapp} />
              <Clickable Icon={FaPhoneAlt} text={content.phoneText} value={content.phone} />
            </div>
            <Separator className='lg:hidden' />
            <div className='space-y-2'>
              <Linkable Icon={FaTelegram} text='Telegram' url={content.telegram} />
              <Linkable Icon={FaVk} text='ВКонтакте' url={content.vk} />
            </div>
            <Separator className='lg:hidden' />
            <div className='space-y-2'>
              <Clickable Icon={FaMap} text={content.addressFirst} value={content.addressFull} />
              <Linkable
                Icon={FaMapMarkerAlt}
                text={content.addressSecond}
                url={content.addressLink}
              />
            </div>
          </div>
          <Separator />
          <p className='text-center font-bold md:text-xl text-lg'>
            Посещение только по предварительной записи.
          </p>
        </div>
      </SheetContent>
    </Sheet>
  )
}
