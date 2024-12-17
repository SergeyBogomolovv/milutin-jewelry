import type { Metadata } from 'next'
import { Poiret_One } from 'next/font/google'
import './globals.css'

const poiret = Poiret_One({
  variable: '--monsterrat',
  subsets: ['latin'],
  weight: ['400'],
})

export const metadata: Metadata = {
  title: 'Milutin Jewelry Admin',
  description: 'Admin app for milutin jewelry',
}

export default function RootLayout({
  children,
}: Readonly<{
  children: React.ReactNode
}>) {
  return (
    <html lang='ru'>
      <body className={`${poiret.className} antialiased`}>{children}</body>
    </html>
  )
}
