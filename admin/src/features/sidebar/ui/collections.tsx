import { SidebarMenu, SidebarMenuButton, SidebarMenuItem } from '@/shared/ui/sidebar'
import { use } from 'react'
import { ScrollText } from 'lucide-react'
import { getCollections } from '@/entities/collection'
import Link from 'next/link'

export default function CollectionsGroup() {
  const { data: collections, success } = use(getCollections())

  return (
    <SidebarMenu>
      {!success && <span>Ошибка загрузки коллекций</span>}
      {collections &&
        collections.map((collection) => (
          <SidebarMenuItem key={collection.id}>
            <SidebarMenuButton asChild>
              <Link href={`/collections/${collection.id}`}>
                <ScrollText />
                <span>{collection.title}</span>
              </Link>
            </SidebarMenuButton>
          </SidebarMenuItem>
        ))}
    </SidebarMenu>
  )
}
