import { SidebarMenu, SidebarMenuButton, SidebarMenuItem } from '@/shared/ui/sidebar'
import { Gem, Newspaper } from 'lucide-react'
import Link from 'next/link'

const links = [
  {
    title: 'Коллекции',
    url: '/collections',
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
