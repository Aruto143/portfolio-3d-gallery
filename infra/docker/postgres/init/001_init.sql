CREATE TABLE works (
  id BIGSERIAL PRIMARY KEY,
  title VARCHAR(255) NOT NULL,
  slug VARCHAR(255) NOT NULL UNIQUE,
  summary TEXT NOT NULL,
  thumbnail_url VARCHAR(500),
  created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

INSERT INTO works (title, slug, summary, thumbnail_url)
VALUES
  (
    '3D Portfolio Gallery',
    '3d-portfolio-gallery',
    'Three.js を使った3Dギャラリー',
    'https://example.com/thumb1.png'
  ),
  (
    'Task Management API',
    'task-management-api',
    'Clean Architecture を意識したAPI設計のサンプル',
    'https://example.com/thumb2.png'
  );