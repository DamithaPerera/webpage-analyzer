package utils

import (
	"net/url"
	"strings"
	"webpage-analyzer/internal/analyzer"

	"golang.org/x/net/html"
)

func ExtractHTMLVersion(doc *html.Node, result *analyzer.AnalysisResult) {
	result.HTMLVersion = "HTML5"
}

func ExtractTitle(doc *html.Node, result *analyzer.AnalysisResult) {
	var findTitle func(*html.Node)
	findTitle = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "title" && n.FirstChild != nil {
			result.Title = n.FirstChild.Data
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			findTitle(c)
		}
	}
	findTitle(doc)
}

func CountHeadings(doc *html.Node, result *analyzer.AnalysisResult) {
	var countHeadings func(*html.Node)
	countHeadings = func(n *html.Node) {
		if n.Type == html.ElementNode && strings.HasPrefix(n.Data, "h") && len(n.Data) == 2 {
			result.HeadingCounts[n.Data]++
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			countHeadings(c)
		}
	}
	countHeadings(doc)
}

func AnalyzeLinks(baseURL string, doc *html.Node, result *analyzer.AnalysisResult) {
	var countLinks func(*html.Node)
	parsedBaseURL, _ := url.Parse(baseURL)

	countLinks = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "a" {
			for _, attr := range n.Attr {
				if attr.Key == "href" {
					linkURL, err := url.Parse(attr.Val)
					if err == nil {
						if linkURL.Host == "" || linkURL.Host == parsedBaseURL.Host {
							result.InternalLinks++
						} else {
							result.ExternalLinks++
						}
					}
				}
			}
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			countLinks(c)
		}
	}
	countLinks(doc)
}

func CheckForLoginForm(doc *html.Node, result *analyzer.AnalysisResult) {
	var findLoginForm func(*html.Node)
	findLoginForm = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "input" {
			for _, attr := range n.Attr {
				if attr.Key == "type" && attr.Val == "password" {
					result.HasLoginForm = true
					return
				}
			}
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			findLoginForm(c)
		}
	}
	findLoginForm(doc)
}