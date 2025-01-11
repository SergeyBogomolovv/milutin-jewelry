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
  title: 'Milutin Jewellery',
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
