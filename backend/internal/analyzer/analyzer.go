package analyzer

import (
	"errors"
	"net/http"
	"webpage-analyzer/internal/models"
	"webpage-analyzer/internal/services"
	"webpage-analyzer/internal/utils"

	"golang.org/x/net/html"
)

// Analyze fetches and analyzes a webpage, returning its metadata.
func Analyze(url string, client services.HTTPClient) (*models.AnalysisResult, error) {
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

	result := &models.AnalysisResult{
		HTMLVersion: utils.DetectHTMLVersion(doc),
		Title:       utils.ExtractTitle(doc),
		Headings:    utils.CountHeadings(doc),
	}

	internal, external, inaccessible := utils.AnalyzeLinks(url, doc)
	result.InternalLinks = internal
	result.ExternalLinks = external
	result.InaccessibleLinks = inaccessible

	result.HasLoginForm = utils.CheckForLoginForm(doc)

	return result, nil
}