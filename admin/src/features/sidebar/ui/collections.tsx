import { SidebarMenu, SidebarMenuButton, SidebarMenuItem } from '@/shared/ui/sidebar'
import { ScrollText } from 'lucide-react'
import { getCollections } from '@/entities/collection'
import Link from 'next/link'
import { use } from 'react'

export default function CollectionsGroup() {
  const collections = use(getCollections().catch(() => []))

  return (
    <SidebarMenu>
      {collections.map((collection) => (
        <SidebarMenuItem key={collection.id}>
          <SidebarMenuButton asChild>
            <Link href={`/${collection.id}`}>
              <ScrollText />
              <span>{collection.title}</span>
            </Link>
          </SidebarMenuButton>
        </SidebarMenuItem>
      ))}
    </SidebarMenu>
  )
}
