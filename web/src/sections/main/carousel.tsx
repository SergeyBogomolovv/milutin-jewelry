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
          className='w-full object-cover md:aspect-[21/8] aspect-[4/3]'
          src='/placeholder.jpg'
          width={1200}
          height={800}
          alt='Banner'
          priority
        />
      </AntCarousel>
    )
  }

  return (
    <AntCarousel autoplay autoplaySpeed={3000}>
      {banners.map((banner) =>
        banner.collection_id ? (
          <Link href={`/${banner.collection_id}`} key={banner.id}>
            <S3Image
              className='w-full object-cover md:aspect-[21/8] aspect-[4/3]'
              src={isMobile ? banner.mobile_image_id : banner.image_id}
              width={1200}
              height={800}
              alt={banner.image_id}
              priority
            />
          </Link>
        ) : (
          <S3Image
            className='w-full object-cover md:aspect-[21/8] aspect-[4/3]'
            key={banner.id}
            src={isMobile ? banner.mobile_image_id : banner.image_id}
            width={1200}
            height={800}
            alt={banner.image_id}
            priority
          />
        ),
      )}
    </AntCarousel>
  )
}
