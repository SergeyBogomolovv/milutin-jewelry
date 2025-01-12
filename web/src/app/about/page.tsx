import type { Metadata } from 'next'
import Image from 'next/image'
import mikhail from '@/assets/mikhail.jpg'
import { Separator } from '@/shared/ui/separator'
import { PreviewGroup } from '@/shared/ui/image'
import AchievmentCard from '@/sections/about/achievment-card'

export const metadata: Metadata = {
  title: 'Об Авторе | Milutin Jewellery',
}

export default function AboutPage() {
  return (
    <main className='grow xl:w-9/12 lg:w-11/12 sm:w-9/12 w-11/12 mx-auto flex flex-col md:gap-10 gap-6 items-center md:py-10 py-6'>
      <h1 className='text-5xl text-center tracking-wide'>Об авторе</h1>
      <section className='grid lg:grid-cols-2 mx-auto gap-4'>
        <p className='lg:hidden text-center text-lg'>
          Художник, создатель драгоценностей. Произведения Михаила Милютина – синтез ювелирного
          мастерства и художественной фантазии.
        </p>
        <Image src={mikhail} alt='Михаил Милютин' className='rounded-lg m-auto' />
        <div className='text-center tracking-wider flex flex-col justify-between text-lg gap-3'>
          <p>
            В основе профессиональной деятельности Михаила Милютина - скрупулезно изученное наследие
            великих мастеров прошлого, сильнейшая школа ювелирного дела, долгий витиеватый путь в
            мире ювелирного искусства и по сей день не угасающая страсть к профессии.
          </p>
          <p>
            Находясь в постоянном поиске и изучая необъятное разнообразие стилей, техник исполнения
            и новых технологий, Михаил не занимается их механическим воспроизведением и
            копированием. Пропуская получаемую информацию через призму собственного художественного
            видения, мастер стремится создавать каждое украшение как самодостаточный предмет
            ювелирного искусства.
          </p>
          <p>
            Особое влияние на творчество Михаила Милютина оказало великое наследие ювелирной школы
            Карла Фаберже, а ведущими интересами в настоящее время являются эксперименты с оптикой и
            инновационной полимерной керамикой.
          </p>
        </div>
      </section>
      <Separator />
      <h2 className='text-4xl text-center tracking-wide'>Достижения</h2>
      <PreviewGroup>
        <section className='lg:flex gap-10 items-center grid md:grid-cols-2 w-11/12 sm:w-full'>
          <AchievmentCard
            src='/award-1.jpg'
            description='В 2022 году стал призером выставки «Гохран России», в номинации
              «Использование нетрадиционных материалов в авторских работах»'
          />
          <AchievmentCard
            src='/award-2.jpg'
            description='В 2021 году стал призером ювелирной выставки J – 1, в номинации
              «Лучшее ювелирное искусство»'
          />
          <AchievmentCard
            src='/award-4.jpg'
            description='С 2020 года состоит в «Международной Академии творчества'
          />
          <AchievmentCard
            src='/award-3.jpg'
            description='В 2019 году стал призером выставки «Гохран России», в номинации
              «Ювелирные техники: традиции и мастерство»'
          />
        </section>
      </PreviewGroup>
    </main>
  )
}
