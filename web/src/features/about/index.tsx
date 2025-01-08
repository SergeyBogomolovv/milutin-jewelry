'use client'
import { useIsMobile } from '@/shared/hooks/use-mobile'
import AboutMobile from './ui/mobile'
import AboutDesktop from './ui/desktop'

export function About() {
  const isMobile = useIsMobile()
  return isMobile ? <AboutMobile /> : <AboutDesktop />
}
