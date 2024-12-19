import { SidebarMenu, SidebarMenuButton, SidebarMenuItem } from '@/shared/ui/sidebar'
import { Gem, Newspaper, ScrollText } from 'lucide-react'
import Link from 'next/link'

const links = [
  {
    title: 'Коллекции',
    url: '/collections',
    icon: ScrollText,
  },
  {
    title: 'Украшения',
    url: '/collection-items',
    icon: Gem,
  },
  {
    title: 'Статьи',
    url: '/posts',
    icon: Newspaper,
  },
]

export default function NavigationGroup() {
  return (
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
  )
}
