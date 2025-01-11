'use client'
import { Banner } from '@/entities/banner'
import { useIsMobile } from '@/shared/hooks/use-mobile'
import S3Image from '@/shared/ui/s3-image'
import { Carousel as AntCarousel } from 'antd'
import Image from 'next/image'
import Link from 'next/link'
import placeholder from '@/assets/placeholder.jpg'

export function Carousel({ banners }: { banners?: Banner[] }) {
  const isMobile = useIsMobile()

  if (!banners?.length) {
    return (
      <AntCarousel autoplay autoplaySpeed={3000}>
        <Image
          className='w-full object-cover sm:aspect-[16/7] aspect-[16/10]'
          src={placeholder}
          width={800}
          height={350}
          alt='Banner'
          priority
        />
      </AntCarousel>
    )
  }

  return (
    <AntCarousel autoplay autoplaySpeed={3000}>
      {!isMobile
        ? banners.map((banner) =>
            banner.collection_id ? (
              <Link href={`/${banner.collection_id}`} key={banner.id}>
                <S3Image
                  className='w-full object-cover aspect-auto'
                  src={banner.image_id}
                  width={1000}
                  height={450}
                  alt={banner.image_id}
                  priority
                />
              </Link>
            ) : (
              <S3Image
                className='w-full object-cover aspect-auto'
                key={banner.id}
                src={banner.image_id}
                width={1000}
                height={450}
                alt={banner.image_id}
                priority
              />
            ),
          )
        : banners.map((banner) =>
            banner.collection_id ? (
              <Link href={`/${banner.collection_id}`} key={banner.id}>
                <S3Image
                  className='w-full object-cover aspect-auto'
                  src={banner.mobile_image_id}
                  width={700}
                  height={700}
                  alt={banner.image_id}
                  priority
                />
              </Link>
            ) : (
              <S3Image
                className='w-full object-cover aspect-auto'
                key={banner.id}
                src={banner.mobile_image_id}
                width={700}
                height={700}
                alt={banner.image_id}
                priority
              />
            ),
          )}
    </AntCarousel>
  )
}
