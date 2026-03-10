import Link from "next/link"

export default function HomePage() {
  return (
    <main className="mx-auto max-w-5xl px-6 py-16">
      <section className="mb-12">
        <p className="mb-3 text-sm text-gray-500">Portfolio</p>
        <h1 className="mb-4 text-4xl font-bold tracking-tight">
          portfolio-3d-gallery
        </h1>
        <p className="max-w-3xl leading-7 text-gray-700">
          Next.js / Go / PostgreSQL を用いて構築している、作品一覧・作品詳細を表示する
          Webアプリケーションです。機能だけでなく、構成・設計・開発基盤の整理も重視して実装しています。
        </p>
      </section>

      <section className="mb-12 grid gap-4 md:grid-cols-2">
        <Link
          href="/works"
          className="rounded-xl border p-6 shadow-sm transition hover:bg-gray-50"
        >
          <h2 className="mb-2 text-xl font-semibold">Works</h2>
          <p className="text-sm text-gray-600">
            作品一覧と詳細ページを確認できます。
          </p>
        </Link>

        <a
          href="https://github.com/Aruto143/portfolio-3d-gallery"
          target="_blank"
          rel="noreferrer"
          className="rounded-xl border p-6 shadow-sm transition hover:bg-gray-50"
        >
          <h2 className="mb-2 text-xl font-semibold">GitHub Repository</h2>
          <p className="text-sm text-gray-600">
            ソースコードと構成を確認できます。
          </p>
        </a>
      </section>

      <section className="grid gap-4 md:grid-cols-3">
        <article className="rounded-xl border p-5">
          <h3 className="mb-2 font-semibold">Frontend</h3>
          <p className="text-sm text-gray-600">
            Next.js / React / TypeScript / Tailwind CSS
          </p>
        </article>

        <article className="rounded-xl border p-5">
          <h3 className="mb-2 font-semibold">Backend</h3>
          <p className="text-sm text-gray-600">
            Go / Gin / sqlx / pgx
          </p>
        </article>

        <article className="rounded-xl border p-5">
          <h3 className="mb-2 font-semibold">Infrastructure</h3>
          <p className="text-sm text-gray-600">
            Docker Compose / Makefile / GitHub Actions
          </p>
        </article>
      </section>
    </main>
  )
}