'use client'
import { logout } from '@/features/auth'
import { DropdownMenuContent, DropdownMenuItem } from '@/shared/ui/dropdown-menu'
import { useRouter } from 'next/navigation'

export default function LogoutDropdown() {
  const router = useRouter()
  return (
    <DropdownMenuContent side='top' className='w-[--radix-popper-anchor-width]'>
      <DropdownMenuItem
        onClick={async () => {
          await logout()
          router.refresh()
        }}
      >
        <span>Выйти</span>
      </DropdownMenuItem>
    </DropdownMenuContent>
  )
}
