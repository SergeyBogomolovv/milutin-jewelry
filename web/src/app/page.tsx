import Image from 'next/image'
import contents from '../assets/main.json'
import mikhail from '../assets/mikhail.jpg'

export default function Home() {
  return (
    <main>
      <Image src={mikhail} alt='Mikhail' />
      <p className='text-3xl'>{contents.title}</p>
    </main>
  )
}
