package main

import (
	"bytes"
	"html/template"
	"io"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"time"

	"github.com/yuin/goldmark"
	"gopkg.in/yaml.v3"
)

const baseURL = "https://mathbliss.com"

type Category struct {
	Name        string `yaml:"name"`
	Slug        string `yaml:"slug"`
	Description string `yaml:"description"`
	Title       string `yaml:"-"`
	URL         string `yaml:"-"`
	FullURL     string `yaml:"-"`
	Image       string `yaml:"-"`
	Type        string `yaml:"-"` // "website"
	Posts       []Page `yaml:"-"`
}

type Page struct {
	Title       string    `yaml:"title"`
	Description string    `yaml:"description"`
	Date        time.Time `yaml:"date"`
	Course      string    `yaml:"course"`
	Unit        int       `yaml:"unit"`
	Categories  []string  `yaml:"categories"`
	Image       string    `yaml:"image"` // Can be set in frontmatter
	Slug        string
	Content     template.HTML
	URL         string
	FullURL     string // Absolute URL for SEO
	Type        string // "article" or "website"
}

// Site holds global data passed to all templates
type Site struct {
	BaseURL    string
	Categories []Category
}

func main() {
	// Clear and recreate public directory
	os.RemoveAll("public")
	os.MkdirAll("public", 0755)

	// Load categories config
	categories := loadCategories("categories.yaml")
	categoryMap := make(map[string]*Category)
	for i := range categories {
		categories[i].Title = categories[i].Name
		categories[i].URL = "/categories/" + categories[i].Slug + "/"
		categories[i].FullURL = baseURL + categories[i].URL
		categories[i].Image = baseURL + "/images/og-default.png"
		categories[i].Type = "website"
		categoryMap[strings.ToLower(categories[i].Name)] = &categories[i]
	}

	site := Site{BaseURL: baseURL, Categories: categories}

	// Load templates
	tmpl := template.Must(template.ParseGlob("templates/*.html"))

	// Collect all posts for the index
	var posts []Page

	// Process all markdown files
	filepath.Walk("content", func(path string, info os.FileInfo, err error) error {
		if err != nil || info.IsDir() || !strings.HasSuffix(path, ".md") {
			return nil
		}

		page := parsePage(path)

		// Determine output path
		rel, _ := filepath.Rel("content", path)
		var outPath string
		var templateName string

		if rel == "about.md" {
			outPath = "public/about/index.html"
			templateName = "page"
			page.URL = "/about/"
			page.Type = "website"
		} else if strings.HasPrefix(rel, "posts/") {
			slug := strings.TrimSuffix(filepath.Base(path), ".md")
			outPath = filepath.Join("public/posts", slug, "index.html")
			templateName = "post"
			page.Slug = slug
			page.URL = "/posts/" + slug + "/"
			page.Type = "article"
		}

		if outPath != "" {
			page.FullURL = baseURL + page.URL
			if page.Image == "" {
				page.Image = baseURL + "/images/og-default.png"
			} else if strings.HasPrefix(page.Image, "/") {
				page.Image = baseURL + page.Image
			}
			os.MkdirAll(filepath.Dir(outPath), 0755)
			renderTemplate(tmpl, templateName, outPath, map[string]any{
				"Page": page,
				"Site": site,
			})

			// Add posts to collection after FullURL is set
			if templateName == "post" {
				posts = append(posts, page)
				for _, catName := range page.Categories {
					if cat, ok := categoryMap[strings.ToLower(catName)]; ok {
						cat.Posts = append(cat.Posts, page)
					}
				}
			}
		}

		return nil
	})

	// Sort posts by date descending
	sort.Slice(posts, func(i, j int) bool {
		return posts[i].Date.After(posts[j].Date)
	})

	// Sort category posts by unit (if set) then date
	for i := range categories {
		sort.Slice(categories[i].Posts, func(a, b int) bool {
			pa, pb := categories[i].Posts[a], categories[i].Posts[b]
			if pa.Unit != pb.Unit {
				return pa.Unit < pb.Unit
			}
			return pa.Date.Before(pb.Date)
		})
	}

	// Generate category pages
	for _, cat := range categories {
		outPath := filepath.Join("public/categories", cat.Slug, "index.html")
		os.MkdirAll(filepath.Dir(outPath), 0755)
		renderTemplate(tmpl, "category", outPath, map[string]any{
			"Category": cat,
			"Site":     site,
		})
	}

	// Generate index page
	renderTemplate(tmpl, "home", "public/index.html", map[string]any{
		"Posts":       posts,
		"Site":        site,
		"Title":       "",
		"Description": "Math explanations that actually make sense. Plain English guides to calculus and CS math.",
		"URL":         "/",
		"FullURL":     baseURL + "/",
		"Image":       baseURL + "/images/og-default.png",
		"Type":        "website",
	})

	// Generate 404 page
	renderTemplate(tmpl, "404", "public/404.html", map[string]any{
		"Site":        site,
		"Title":       "Page Not Found",
		"Description": "This page doesn't exist.",
		"URL":         "/404.html",
		"FullURL":     baseURL + "/404.html",
		"Image":       baseURL + "/images/og-default.png",
		"Type":        "website",
	})

	// Generate sitemap
	var sitemapURLs []string
	sitemapURLs = append(sitemapURLs, baseURL+"/")
	sitemapURLs = append(sitemapURLs, baseURL+"/about/")
	for _, cat := range categories {
		sitemapURLs = append(sitemapURLs, cat.FullURL)
	}
	for _, post := range posts {
		sitemapURLs = append(sitemapURLs, post.FullURL)
	}
	generateSitemap("public/sitemap.xml", sitemapURLs)

	// Copy static assets
	copyDir("styles", "public/styles")
	copyDir("static", "public")

	println("Build complete!")
}

func loadCategories(path string) []Category {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil
	}
	var categories []Category
	yaml.Unmarshal(data, &categories)
	return categories
}

func parsePage(path string) Page {
	data, _ := os.ReadFile(path)
	content := string(data)

	var page Page

	// Parse YAML frontmatter
	if strings.HasPrefix(content, "---") {
		parts := strings.SplitN(content[3:], "---", 2)
		if len(parts) == 2 {
			yaml.Unmarshal([]byte(parts[0]), &page)
			content = strings.TrimSpace(parts[1])
		}
	}

	// Convert markdown to HTML
	var buf bytes.Buffer
	goldmark.Convert([]byte(content), &buf)
	page.Content = template.HTML(buf.String())

	return page
}

func renderTemplate(tmpl *template.Template, name, outPath string, data any) {
	f, err := os.Create(outPath)
	if err != nil {
		println("Error creating", outPath, err.Error())
		return
	}
	defer f.Close()

	if err := tmpl.ExecuteTemplate(f, name, data); err != nil {
		println("Error rendering", name, err.Error())
	}
}

func copyDir(src, dst string) {
	filepath.Walk(src, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return nil
		}

		rel, _ := filepath.Rel(src, path)
		target := filepath.Join(dst, rel)

		if info.IsDir() {
			os.MkdirAll(target, 0755)
		} else {
			copyFile(path, target)
		}
		return nil
	})
}

func copyFile(src, dst string) {
	in, err := os.Open(src)
	if err != nil {
		return
	}
	defer in.Close()

	os.MkdirAll(filepath.Dir(dst), 0755)
	out, err := os.Create(dst)
	if err != nil {
		return
	}
	defer out.Close()

	io.Copy(out, in)
}

func generateSitemap(path string, urls []string) {
	f, err := os.Create(path)
	if err != nil {
		println("Error creating sitemap:", err.Error())
		return
	}
	defer f.Close()

	f.WriteString(`<?xml version="1.0" encoding="UTF-8"?>
<urlset xmlns="http://www.sitemaps.org/ns/sitemap/0.9">
`)
	for _, url := range urls {
		f.WriteString("  <url>\n    <loc>" + url + "</loc>\n  </url>\n")
	}
	f.WriteString("</urlset>\n")
}
