import Link from "next/link"

type Work = {
  id: number
  title: string
  slug: string
  summary: string
  thumbnailUrl: string
}

async function fetchWorks(): Promise<Work[]> {
  const apiBaseUrl =
    process.env.API_BASE_URL ?? "http://backend:18080"

  const res = await fetch(`${apiBaseUrl}/api/works`, {
    cache: "no-store",
  })

  if (!res.ok) {
    throw new Error("作品一覧の取得に失敗しました")
  }

  return res.json()
}

export default async function WorksPage() {
  const works = await fetchWorks()

  return (
    <main className="mx-auto max-w-4xl p-8">
      <h1 className="mb-6 text-3xl font-bold">Works</h1>

      <div className="grid gap-4">
        {works.map((work) => (
          <Link
            key={work.slug}
            href={`/works/${work.slug}`}
            className="block rounded-lg border p-4 shadow-sm transition hover:bg-gray-50"
          >
            <h2 className="text-xl font-semibold">{work.title}</h2>
            <p className="mt-1 text-sm text-gray-500">{work.slug}</p>
            <p className="mt-3">{work.summary}</p>
          </Link>
        ))}
      </div>
    </main>
  )
}