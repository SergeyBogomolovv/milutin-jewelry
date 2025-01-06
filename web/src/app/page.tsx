import contents from '../assets/main.json'
import { Contacts } from '@/features/contacts'
import { Button } from '@/shared/ui/button'
import { Image } from '@/shared/ui/image'

export default function Home() {
  return (
    <main>
      <Image
        src={'collections/21eb3162-61bd-482a-8f73-9aef0028bfaa'}
        alt='Mikhail'
        title='Кольцо кролик'
        description='Золото, рубины, радирование.'
      />

      <Image src={'collections/21eb3162-61bd-482a-8f73-9aef0028bfaa'} alt='Mikhail' />
      <Image src={'collections/21eb3162-61bd-482a-8f73-9aef0028bfaa'} alt='Mikhail' />
      <Image src={'collections/21eb3162-61bd-482a-8f73-9aef0028bfaa'} alt='Mikhail' />
      <Contacts>
        <Button>Контакты</Button>
      </Contacts>
      <p className='text-3xl'>{contents.aboutText}</p>
    </main>
  )
}
