package analyzer

import (
	"errors"
	"net/http"
	"webpage-analyzer/internal/utils"
	"golang.org/x/net/html"
)

type AnalysisResult struct {
	HTMLVersion       string         `json:"html_version"`
	Title             string         `json:"title"`
	HeadingCounts     map[string]int `json:"heading_counts"`
	InternalLinks     int            `json:"internal_links"`
	ExternalLinks     int            `json:"external_links"`
	InaccessibleLinks int            `json:"inaccessible_links"`
	HasLoginForm      bool           `json:"has_login_form"`
}

func Analyze(url string, client http.Client) (*AnalysisResult, error) {
	resp, err := client.Get(url)
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

	result := &AnalysisResult{
		HeadingCounts: make(map[string]int),
	}

	utils.ExtractHTMLVersion(doc, result)
	utils.ExtractTitle(doc, result)
	utils.CountHeadings(doc, result)
	utils.AnalyzeLinks(url, doc, result)
	utils.CheckForLoginForm(doc, result)

	return result, nil
}
