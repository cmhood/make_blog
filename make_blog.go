// Copyright (c) 2024 Charles Hood <chood@chood.net>
//
// This program is free software: you can redistribute it and/or modify it under
// the terms of the GNU General Public License as published by the Free Software
// Foundation, either version 3 of the License, or (at your option) any later
// version.
//
// This program is distributed in the hope that it will be useful, but WITHOUT
// ANY WARRANTY; without even the implied warranty of MERCHANTABILITY or FITNESS
// FOR A PARTICULAR PURPOSE. See the GNU General Public License for more
// details.
//
// You should have received a copy of the GNU General Public License along with
// this program. If not, see <https://www.gnu.org/licenses/>.

package main

import (
	"fmt"
	"os"
	"text/template"
	"path/filepath"
	"strings"
	"sort"
	"time"
)

type article = struct {
	Title string
	Date  articleDate
	Path  string
	Body  string
}

type articleDate = struct {
	Iso8601 string
	Weekday string
	Long    string
	Rfc2822 string
}

const articlesDir = "articles"
const tmplDir = "tmpl"
const outputDir = "public_html"

var indexTemplate = template.Must(template.ParseFiles(tmplDir + "/index.tmpl"))
var rssTemplate = template.Must(template.ParseFiles(tmplDir + "/rss.tmpl"))
var articleTemplate = template.Must(template.ParseFiles(tmplDir + "/article.tmpl"))

var allArticles = []article{}

func main() {
	fmt.Println("Generating pages...")

	entries, err := os.ReadDir(articlesDir)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	articlePaths := []string{}
	for _, ent := range entries {
		name := ent.Name()
		if name[0] == '.' {
			continue
		}
		articlePaths = append(articlePaths, filepath.Join(articlesDir, name))
	}
	sort.Sort(sort.Reverse(sort.StringSlice(articlePaths)))

	for _, path := range articlePaths {
		loadArticle(path)
	}

	generateIndex(indexTemplate, "index.html")
	generateIndex(rssTemplate, "rss.xml")
	for _, article := range allArticles {
		generateArticle(article)
	}

	fmt.Println("Done!")
}

func loadArticle(path string) {
	parts := strings.Split(filepath.Base(path), "_")
	date := parts[0]
	slug := strings.Join(parts[2:], "_")

	bytes, err := os.ReadFile(path)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	lines := strings.Split(string(bytes), "\n")
	title := lines[0]
	body := strings.Join(lines[2:], "\n")

	articlePath := filepath.Join(strings.Replace(date, "-", "/", -1), slug)
	article := article{title, getArticleDate(date), articlePath, body}
	allArticles = append(allArticles, article)
}

func getArticleDate(iso8601 string) articleDate {
	t, err := time.Parse("2006-01-02", iso8601)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	weekday := t.Weekday().String()
	long := t.Format("January 2, 2006")
	rfc2822 := t.Format("Mon, 02 Jan 2006 00:00:00 +0000")
	return articleDate{iso8601, weekday, long, rfc2822}
}

func generateIndex(template *template.Template, path string) {
	fullPath := filepath.Join(outputDir, path)
	fmt.Println(fullPath)
	os.MkdirAll(filepath.Dir(fullPath), 0755)
	file, err := os.Create(fullPath)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	err = template.Execute(file, allArticles)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func generateArticle(article article) {
	fullPath := filepath.Join(outputDir, article.Path)
	fmt.Println(fullPath)
	os.MkdirAll(filepath.Dir(fullPath), 0755)
	file, err := os.Create(fullPath)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	err = articleTemplate.Execute(file, article)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
