package main

import (
	"fmt"
	"os"
	"strings"

	"golang.org/x/net/html"
)

// Create a struct type to match the result links
// Open and read the html file
// use the x/net/html package to parse the html content
// For each html tag identified as "a", extract link and text

type Anchor struct {
	Href string
	Text string
}

func extractAllTexts(node *html.Node) string {
	var textString string
	if node.Type == html.TextNode {
		textString = node.Data
	}
	for c := node.FirstChild; c != nil; c = c.NextSibling {
		textString += extractAllTexts(c)
	}
	return strings.Trim(textString, "\n")
}

func parseAnchorLinks(node *html.Node) []Anchor {

	var links []Anchor

	if node.Type == html.ElementNode && node.Data == "a" {
		//extract href
		//extract all texts including subtexts
		var href string
		for _, attr := range node.Attr {
			if attr.Key == "href" {
				href = attr.Val
			}
		}
		links = append(links, Anchor{Href: href, Text: extractAllTexts(node)})
	}
	for c := node.FirstChild; c != nil; c = c.NextSibling {
		links = append(links, parseAnchorLinks(c)...) 
	}

	return links
}

func main() {
	path := "./ex3.html"
	file, err := os.Open(path)
	if err != nil {
		fmt.Printf("Error opening file %v\n", err)
		os.Exit(1)
	}
	defer file.Close()

	doc, err := html.Parse(file)
	if err != nil {
		fmt.Printf("Error parsing html file %v\n", err)
	}

	fmt.Println(parseAnchorLinks(doc))
}
