import type { Metadata } from 'next'
import { Montserrat } from 'next/font/google'
import './globals.css'

const monsterrat = Montserrat({
  variable: '--monsterrat',
  subsets: ['latin'],
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
      <body className={`${monsterrat.className} antialiased`}>{children}</body>
    </html>
  )
}
