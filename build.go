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

type Page struct {
	Title       string    `yaml:"title"`
	Description string    `yaml:"description"`
	Date        time.Time `yaml:"date"`
	Course      string    `yaml:"course"`
	Unit        string    `yaml:"unit"`
	Slug        string
	Content     template.HTML
	URL         string
}

func main() {
	// Clear and recreate public directory
	os.RemoveAll("public")
	os.MkdirAll("public", 0755)

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
		} else if strings.HasPrefix(rel, "posts/") {
			slug := strings.TrimSuffix(filepath.Base(path), ".md")
			outPath = filepath.Join("public/posts", slug, "index.html")
			templateName = "post"
			page.Slug = slug
			page.URL = "/posts/" + slug + "/"
			posts = append(posts, page)
		}

		if outPath != "" {
			os.MkdirAll(filepath.Dir(outPath), 0755)
			renderTemplate(tmpl, templateName, outPath, page)
		}

		return nil
	})

	// Sort posts by date descending
	sort.Slice(posts, func(i, j int) bool {
		return posts[i].Date.After(posts[j].Date)
	})

	// Generate index page
	renderTemplate(tmpl, "home", "public/index.html", map[string]any{
		"Posts": posts,
	})

	// Copy styles directory
	copyDir("styles", "public/styles")

	println("Build complete!")
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
