export default async function CollectionPage({ params }: { params: Promise<{ id: string }> }) {
  const { id } = await params

  return <main>CollectionsPage {id}</main>
}
