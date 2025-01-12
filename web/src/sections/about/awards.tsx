import { PreviewGroup } from '@/shared/ui/image'
import { AwardCard } from './award-card'

export function Awards() {
  return (
    <PreviewGroup>
      <section className='lg:flex gap-10 items-center grid md:grid-cols-2 w-11/12 sm:w-full'>
        <AwardCard
          src='/award-1.jpg'
          description='В 2022 году стал призером выставки «Гохран России», в номинации «Использование нетрадиционных материалов в авторских работах»'
        />
        <AwardCard
          src='/award-2.jpg'
          description='В 2021 году стал призером ювелирной выставки J – 1, в номинации «Лучшее ювелирное искусство»'
        />
        <AwardCard
          src='/award-4.jpg'
          description='С 2020 года состоит в «Международной Академии творчества'
        />
        <AwardCard
          src='/award-3.jpg'
          description='В 2019 году стал призером выставки «Гохран России», в номинации «Ювелирные техники: традиции и мастерство»'
        />
      </section>
    </PreviewGroup>
  )
}
