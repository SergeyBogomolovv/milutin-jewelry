import { SidebarMenu, SidebarMenuButton, SidebarMenuItem } from '@/shared/ui/sidebar'
import { use } from 'react'
import { NotepadText } from 'lucide-react'
import { getCollections } from '@/entities/collection'
import Link from 'next/link'

export default function CollectionsGroup() {
  const collections = use(getCollections())

  return (
    <SidebarMenu>
      {collections.map((collection) => (
        <SidebarMenuItem key={collection.id}>
          <SidebarMenuButton asChild>
            <Link href={`/collections/${collection.id}`}>
              <NotepadText />
              <span>{collection.title}</span>
            </Link>
          </SidebarMenuButton>
        </SidebarMenuItem>
      ))}
    </SidebarMenu>
  )
}