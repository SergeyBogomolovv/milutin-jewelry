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
  title: {
    default: 'Milutin Jewellery',
    template: '%s | Milutin Jewellery',
  },
  description:
    'Михаил Милютин — Художник, создатель драгоценностей. Произведения Михаила Милютина – синтез ювелирного мастерства и художественной фантазии.',
  openGraph: {},
  keywords: [
    'Михаил Милютин',
    'Милютин',
    'Mikhail Milutin',
    'milutin jewellery',
    'milutin',
    'jewellery',
    'jewelry',
    'ювелирная мастерская',
    'ювелир',
    'ювелирские изделия',
    'ювелирные украшения',
    'изделия ручной работы',
    'эксклюзивные украшения',
    'ювелирное искусство',
    'авторские изделия',
    'дизайн украшений',
    'традиции ювелирного дела',
    'драгоценности ручной работы',
    'изделия из драгоценных камней',
    'индивидуальный дизайн украшений',
    'эксклюзивные подарки',
    'украшения',
  ],
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
