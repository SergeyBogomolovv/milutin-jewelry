import type { Metadata } from 'next'
import { Inter } from 'next/font/google'
import './globals.css'
import { Toaster } from '@/shared/ui/sonner'

const poiret = Inter({
  variable: '--inter',
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
      <body className={`${poiret.className} antialiased`}>
        {children}
        <Toaster />
      </body>
    </html>
  )
}
