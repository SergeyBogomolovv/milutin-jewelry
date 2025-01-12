import './globals.css'
import type { Metadata } from 'next'
import { Poiret_One } from 'next/font/google'
import { Toaster } from '@/shared/ui/sonner'
import { Header } from '@/features/header'
import { Footer } from '@/features/footer'
import { AntdRegistry } from '@ant-design/nextjs-registry'
import '@ant-design/v5-patch-for-react-19'

const poiretOne = Poiret_One({
  weight: '400',
  subsets: ['cyrillic', 'latin'],
})

export const metadata: Metadata = {
  metadataBase: new URL('https://milutin-jewellery.com'),
  title: {
    default: 'Milutin Jewellery',
    template: '%s | Milutin Jewellery',
  },
  description:
    'Михаил Милютин — Художник, создатель драгоценностей. Произведения Михаила Милютина – синтез ювелирного мастерства и художественной фантазии.',
  openGraph: {
    siteName: 'Milutin Jewellery',
    title: 'Milutin Jewellery',
    description:
      'Михаил Милютин — Художник, создатель драгоценностей. Произведения Михаила Милютина – синтез ювелирного мастерства и художественной фантазии.',
    url: 'https://milutin-jewellery.com',
    locale: 'ru_RU',
    type: 'website',
  },
  keywords: [
    'Михаил Милютин',
    'Mikhail Milutin',
    'milutin jewellery',
    'milutin',
    'jewellery',
    'jewelry',
    'ювелирная мастерская',
    'украшения ручной работы',
    'эксклюзивные украшения',
    'ювелирное искусство',
    'уникальные ювелирные изделия',
    'авторские украшения',
    'дизайн украшений',
    'традиции ювелирного дела',
    'драгоценности ручной работы',
    'украшения из драгоценных камней',
    'индивидуальный дизайн украшений',
    'эксклюзивные подарки',
    'украшения',
    'ювелир',
  ],
  robots: {
    index: true,
    follow: true,
  },
}

export default function RootLayout({
  children,
}: Readonly<{
  children: React.ReactNode
}>) {
  return (
    <html lang='ru'>
      <body className={`${poiretOne.className} antialiased min-h-screen flex flex-col`}>
        <Toaster />
        <AntdRegistry>
          <Header />
          {children}
          <Footer />
        </AntdRegistry>
      </body>
    </html>
  )
}
