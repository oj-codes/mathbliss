# mathbliss

A static site for plain English math explanations. Built with a custom Go-based static site generator. No Hugo, Jekyll, or other frameworks. Markdown in, HTML out.

## Build locally

```bash
go run build.go
```

This generates the site in `/public`. Serve it locally with:

```bash
static-server public
```

after installing the static-server package:

```bash
go install github.com/eliben/static-server@latest
```

Then visit `http://localhost:8080`.

## Project structure

```
content/
  posts/       # Blog posts (markdown)
  about.md     # About page
templates/     # HTML templates (base, home, post, category, 404)
styles/        # CSS
static/        # Static assets (favicon, robots.txt, etc.)
public/        # Generated site (git-ignored)
build.go       # The static site generator
categories.yaml
```

## Deployment

Push to `main` triggers GitHub Actions (`.github/workflows/ci.yaml`), which runs the build and deploys `/public` to GitHub Pages.

## Adding content

Create a new `.md` file in `content/posts/`:

```markdown
---
title: "Post Title"
description: "A short description for SEO and previews."
date: 2025-01-15T12:00:00-04:00
categories: [Calculus]
unit: 1
---

Your content here.
```

**Frontmatter fields:**
- `title` – Post title
- `description` – Short description for meta tags
- `date` – Publication date (ISO format)
- `categories` – Array of category names (must match `categories.yaml`)
- `unit` – Optional number for ordering posts within a category
