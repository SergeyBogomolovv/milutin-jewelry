'use client'
import { Banner } from '@/entities/banner'
import { useIsMobile } from '@/shared/hooks/use-mobile'
import S3Image from '@/shared/ui/s3-image'
import { Carousel as AntCarousel } from 'antd'
import Image from 'next/image'
import Link from 'next/link'

export function Carousel({ banners }: { banners?: Banner[] }) {
  const isMobile = useIsMobile()

  if (!banners?.length) {
    return (
      <AntCarousel autoplay autoplaySpeed={3000}>
        <Image
          className='w-full object-cover sm:aspect-[16/7] aspect-square'
          src='/placeholder.jpg'
          width={1200}
          height={700}
          alt='Banner'
          priority={true}
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
                  width={1200}
                  height={700}
                  alt={banner.id.toString()}
                  priority
                />
              </Link>
            ) : (
              <S3Image
                className='w-full object-cover aspect-auto'
                key={banner.id}
                src={banner.image_id}
                width={1200}
                height={700}
                alt={banner.id.toString()}
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
                  alt={banner.id.toString()}
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
                alt={banner.id.toString()}
                priority
              />
            ),
          )}
    </AntCarousel>
  )
}
