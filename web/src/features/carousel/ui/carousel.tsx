import { Carousel as AntCarousel } from 'antd'
import Image from 'next/image'

const slides = [
  {
    src: 'https://storage.yandexcloud.net/mikhail-milutin/slider/alicafon1full.jpg',
    placeholder: 'https://storage.yandexcloud.net/mikhail-milutin/slider/alicafon1.jpg',
  },
  {
    src: 'https://storage.yandexcloud.net/mikhail-milutin/slider/alicafonfull.jpg',
    placeholder: 'https://storage.yandexcloud.net/mikhail-milutin/slider/alicafon.jpg',
  },
  {
    src: 'https://storage.yandexcloud.net/mikhail-milutin/slider/alicafon2full.JPG',
    placeholder: 'https://storage.yandexcloud.net/mikhail-milutin/slider/alicafon2.jpg',
  },
  {
    src: 'https://storage.yandexcloud.net/mikhail-milutin/slider/alicafon3full.JPG',
    placeholder: 'https://storage.yandexcloud.net/mikhail-milutin/slider/alicafon3.jpg',
  },
  {
    src: 'https://storage.yandexcloud.net/mikhail-milutin/slider/alicafon4full.JPG',
    placeholder: 'https://storage.yandexcloud.net/mikhail-milutin/slider/alicafon4.JPG',
  },
]

export function Carousel() {
  return (
    <AntCarousel autoplay autoplaySpeed={3000}>
      {slides.map(({ src, placeholder }) => (
        <Image
          className='w-full object-cover sm:aspect-auto aspect-square'
          key={src}
          src={src}
          width={1200}
          height={700}
          blurDataURL={placeholder}
          placeholder='blur'
          alt='Алиса в стране чудес'
          priority={true}
        />
      ))}
    </AntCarousel>
  )
}
