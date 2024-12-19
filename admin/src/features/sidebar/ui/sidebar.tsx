import {
  Sidebar,
  SidebarContent,
  SidebarFooter,
  SidebarGroup,
  SidebarGroupContent,
  SidebarGroupLabel,
  SidebarMenu,
  SidebarMenuButton,
  SidebarMenuItem,
} from '@/shared/ui/sidebar'
import { DropdownMenu, DropdownMenuTrigger } from '@/shared/ui/dropdown-menu'
import { ChevronUp, User2 } from 'lucide-react'
import CollectionsGroup from './collections'
import { GiJewelCrown } from 'react-icons/gi'
import { FaClipboardList } from 'react-icons/fa'
import { IoNewspaper } from 'react-icons/io5'
import { Suspense } from 'react'
import Link from 'next/link'
import LogoutDropdown from './logout-dropdown'

const links = [
  {
    title: 'Коллекции',
    url: '/collections',
    icon: FaClipboardList,
  },
  {
    title: 'Украшения',
    url: '/collection-items',
    icon: GiJewelCrown,
  },
  {
    title: 'Статьи',
    url: '/posts',
    icon: IoNewspaper,
  },
]

export function AppSidebar() {
  return (
    <Sidebar>
      <SidebarContent>
        <SidebarGroup>
          <SidebarGroupLabel>Навигация</SidebarGroupLabel>
          <SidebarGroupContent>
            <SidebarMenu>
              {links.map((link) => (
                <SidebarMenuItem key={link.title}>
                  <SidebarMenuButton asChild>
                    <Link href={link.url}>
                      <link.icon />
                      <span>{link.title}</span>
                    </Link>
                  </SidebarMenuButton>
                </SidebarMenuItem>
              ))}
            </SidebarMenu>
          </SidebarGroupContent>
        </SidebarGroup>
        <SidebarGroup>
          <SidebarGroupLabel>Коллекци</SidebarGroupLabel>
          <SidebarGroupContent>
            <Suspense fallback={<span>Загрузка...</span>}>
              <CollectionsGroup />
            </Suspense>
          </SidebarGroupContent>
        </SidebarGroup>
      </SidebarContent>
      <SidebarFooter>
        <SidebarMenu>
          <SidebarMenuItem>
            <DropdownMenu>
              <DropdownMenuTrigger asChild>
                <SidebarMenuButton>
                  <User2 /> Админ
                  <ChevronUp className='ml-auto' />
                </SidebarMenuButton>
              </DropdownMenuTrigger>
              <LogoutDropdown />
            </DropdownMenu>
          </SidebarMenuItem>
        </SidebarMenu>
      </SidebarFooter>
    </Sidebar>
  )
}
