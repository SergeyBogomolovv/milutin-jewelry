import { SidebarMenu, SidebarMenuButton, SidebarMenuItem } from '@/shared/ui/sidebar'
import { ScrollText } from 'lucide-react'
import { getCollections } from '@/entities/collection'
import Link from 'next/link'

export default async function CollectionsGroup() {
  try {
    const collections = await getCollections()
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
  } catch (error) {
    return <span>Ошибка загрузки коллекций</span>
  }
}
