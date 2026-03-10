type Work = {
  id: number
  title: string
  slug: string
  summary: string
  thumbnailUrl: string
}

type WorksDetailPageProps = {
  params: Promise<{
    slug: string
  }>
}

async function fetchWorkBySlug(slug: string): Promise<Work> {
  const res = await fetch(
    `http://host.docker.internal:18080/api/works/${slug}`,
    {
      cache: "no-store",
    }
  )

  if (res.status === 404) {
    throw new Error("NOT_FOUND")
  }

  if (!res.ok) {
    throw new Error("作品詳細の取得に失敗しました")
  }

  return res.json()
}

export default async function WorksDetailPage({
  params,
}: WorksDetailPageProps) {
  const { slug } = await params
  const work = await fetchWorkBySlug(slug)

  return (
    <main className="mx-auto max-w-3xl p-8">
      <p className="mb-2 text-sm text-gray-500">slug: {work.slug}</p>
      <h1 className="mb-6 text-3xl font-bold">{work.title}</h1>

      <section className="rounded-lg border p-6 shadow-sm">
        <h2 className="mb-3 text-lg font-semibold">Summary</h2>
        <p>{work.summary}</p>
      </section>
    </main>
  )
}