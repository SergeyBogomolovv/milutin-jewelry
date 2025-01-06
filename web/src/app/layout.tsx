import type { Metadata } from 'next'
import { Poiret_One } from 'next/font/google'
import './globals.css'
import { Toaster } from '@/shared/ui/sonner'
import { Header } from '@/features/header'
import { AntdRegistry } from '@ant-design/nextjs-registry'
import '@ant-design/v5-patch-for-react-19'

const poiretOne = Poiret_One({
  weight: '400',
  subsets: ['cyrillic', 'latin'],
})

export const metadata: Metadata = {
  title: 'Главная | Milutin Jewellery',
  description: 'Generated by create next app',
}

export default function RootLayout({
  children,
}: Readonly<{
  children: React.ReactNode
}>) {
  return (
    <html lang='ru'>
      <body className={`${poiretOne.className} antialiased`}>
        <Toaster />
        <AntdRegistry>
          <Header />
          {children}
        </AntdRegistry>
      </body>
    </html>
  )
}
