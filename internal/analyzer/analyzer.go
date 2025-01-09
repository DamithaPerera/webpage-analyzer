package analyzer

import (
    "errors"
    "net/http"
    "golang.org/x/net/html"
)

type AnalysisResult struct {
    Title string `json:"title"`
}

func Analyze(url string) (*AnalysisResult, error) {
    resp, err := http.Get(url)
    if err != nil {
        return nil, errors.New("unable to fetch the URL")
    }
    defer resp.Body.Close()

    if resp.StatusCode != http.StatusOK {
        return nil, errors.New("non-success HTTP status received: " + http.StatusText(resp.StatusCode))
    }

    doc, err := html.Parse(resp.Body)
    if err != nil {
        return nil, errors.New("error parsing HTML document")
    }

    title := extractTitle(doc)

    return &AnalysisResult{
        Title: title,
    }, nil
}

func extractTitle(doc *html.Node) string {
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