package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gomarkdown/markdown"
	"github.com/gomarkdown/markdown/html"
	"github.com/gomarkdown/markdown/parser"
)

func handler(w http.ResponseWriter, r *http.Request, file_path string) {
	markdown_content, err := os.ReadFile(file_path)
	if err != nil {
		fmt.Fprintf(w, "Error reading: %s", file_path)
		return
	}

	md := []byte(markdown_content)
	html := mdToHTML(md)
	fmt.Fprintf(w, "<html><head></head><body><div style=\"width: 800px;\">%s</div></body></html>", html)
}

func mdToHTML(md []byte) []byte {
	// create markdown parser with extensions
	extensions := parser.CommonExtensions | parser.AutoHeadingIDs | parser.NoEmptyLineBeforeBlock
	p := parser.NewWithExtensions(extensions)
	doc := p.Parse(md)

	// create HTML renderer with extensions
	htmlFlags := html.CommonFlags | html.HrefTargetBlank
	opts := html.RendererOptions{Flags: htmlFlags}
	renderer := html.NewRenderer(opts)

	return markdown.Render(doc, renderer)
}

func main() {
	if len(os.Args) != 2 {
		fmt.Printf("Usage: ./mdviewer <path-to-file.md>\n")
		os.Exit(1)
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		handler(w, r, os.Args[1])
	})
	serverAddress := "127.0.0.1:8888"
	fmt.Printf("Running mdviewer at %s", serverAddress)
	log.Fatal(http.ListenAndServe(serverAddress, nil))
}
