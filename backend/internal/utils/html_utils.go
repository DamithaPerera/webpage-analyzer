package utils

import (
	"net/url"

	"golang.org/x/net/html"
)

func DetectHTMLVersion(doc *html.Node) string {
	// Assume HTML5 by default
	return "HTML5"
}

//finds and returns the title of the document
func ExtractTitle(doc *html.Node) string {
	var title string
	var findTitle func(*html.Node)
	findTitle = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "title" && n.FirstChild != nil {
			title = n.FirstChild.Data
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			findTitle(c)
		}
	}
	findTitle(doc)
	return title
}

//counts the occurrences of heading tags
func CountHeadings(doc *html.Node) map[string]int {
	counts := make(map[string]int)
	var count func(*html.Node)
	count = func(n *html.Node) {
		if n.Type == html.ElementNode && len(n.Data) == 2 && n.Data[0] == 'h' {
			counts[n.Data]++
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			count(c)
		}
	}
	count(doc)
	return counts
}

//analyzes the links in the document and classifies them as internal, external, or inaccessible
func AnalyzeLinks(baseURL string, doc *html.Node) (int, int, int) {
	internal, external, inaccessible := 0, 0, 0
	base, _ := url.Parse(baseURL)

	var countLinks func(*html.Node)
	countLinks = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "a" {
			for _, attr := range n.Attr {
				if attr.Key == "href" {
					hrefURL, err := url.Parse(attr.Val)
					if err == nil {
						if hrefURL.Host == "" || hrefURL.Host == base.Host {
							internal++
						} else {
							external++
						}
					} else {
						inaccessible++
					}
				}
			}
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			countLinks(c)
		}
	}
	countLinks(doc)
	return internal, external, inaccessible
}

//checks if the document contains a login form with a password field
func CheckForLoginForm(doc *html.Node) bool {
	var found bool
	var checkForm func(*html.Node)
	checkForm = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "input" {
			for _, attr := range n.Attr {
				if attr.Key == "type" && attr.Val == "password" {
					found = true
					return
				}
			}
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			checkForm(c)
		}
	}
	checkForm(doc)
	return found
}