import { SidebarMenu, SidebarMenuButton, SidebarMenuItem } from '@/shared/ui/sidebar'
import Link from 'next/link'
import { use } from 'react'
import { GiJeweledChalice } from 'react-icons/gi'
import { getCollections } from '../api/get-collections'

export default function CollectionsGroup() {
  const collections = use(getCollections())

  return (
    <SidebarMenu>
      {collections.map((collection) => (
        <SidebarMenuItem key={collection.id}>
          <SidebarMenuButton asChild>
            <Link href={`/collections/${collection.id}`}>
              <GiJeweledChalice />
              <span>{collection.title}</span>
            </Link>
          </SidebarMenuButton>
        </SidebarMenuItem>
      ))}
    </SidebarMenu>
  )
}
