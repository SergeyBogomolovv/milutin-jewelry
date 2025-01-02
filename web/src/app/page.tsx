import Image from 'next/image'
import contents from '../assets/main.json'
import mikhail from '../assets/mikhail.jpg'
import { Contacts } from '@/features/contacts'
import { Button } from '@/shared/ui/button'

export default function Home() {
  return (
    <main>
      <Image src={mikhail} alt='Mikhail' />
      <Contacts>
        <Button>Контакты</Button>
      </Contacts>
      <p className='text-3xl'>{contents.aboutText}</p>
    </main>
  )
}
