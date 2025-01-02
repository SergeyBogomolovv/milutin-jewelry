import { Image } from '@/shared/ui/image'
import contents from '../assets/main.json'
import { Contacts } from '@/features/contacts'
import { Button } from '@/shared/ui/button'

export default function Home() {
  return (
    <main>
      <Image
        src={'collections/e061afa3-2286-4c04-8419-57beb169c48a'}
        alt='Mikhail'
        width={500}
        height={500}
      />
      <Contacts>
        <Button>Контакты</Button>
      </Contacts>
      <p className='text-3xl'>{contents.aboutText}</p>
    </main>
  )
}
