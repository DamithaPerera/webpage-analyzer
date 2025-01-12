package analyzer

import (
	"errors"
	"net/http"
	"webpage-analyzer/internal/models"
	"webpage-analyzer/internal/services"
	"webpage-analyzer/internal/utils"

	"golang.org/x/net/html"
)

// Analyze fetches and analyzes a webpage, returning its metadata and Go routines to process tasks concurrently.
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

	result := &models.AnalysisResult{}
	done := make(chan error, 4)

	// Process tasks concurrently using Go routines
	go func() {
		result.HTMLVersion = utils.DetectHTMLVersion(doc)
		done <- nil
	}()

	go func() {
		result.Title = utils.ExtractTitle(doc)
		done <- nil
	}()

	go func() {
		result.Headings = utils.CountHeadings(doc)
		done <- nil
	}()

	go func() {
		internal, external, inaccessible := utils.AnalyzeLinks(url, doc)
		result.InternalLinks = internal
		result.ExternalLinks = external
		result.InaccessibleLinks = inaccessible
		done <- nil
	}()

	for i := 0; i < 4; i++ {
		if err := <-done; err != nil {
			return nil, err
		}
	}

	result.HasLoginForm = utils.CheckForLoginForm(doc)

	return result, nil
}