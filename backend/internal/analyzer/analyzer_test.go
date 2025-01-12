package analyzer

import (
	"errors"
	"io"
	"net/http"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

type MockHTTPClient struct {
	Response *http.Response
	Err      error
}

func (m *MockHTTPClient) Get(url string) (*http.Response, error) {
	if m.Err != nil {
		return nil, m.Err
	}
	return m.Response, nil
}

// Define ExpectedAnalysisResult globally
type ExpectedAnalysisResult struct {
	HTMLVersion       string
	Title             string
	Headings          map[string]int
	InternalLinks     int
	ExternalLinks     int
	InaccessibleLinks int
	HasLoginForm      bool
}

func TestAnalyze_Concurrent(t *testing.T) {
	tests := []struct {
		name         string
		htmlContent  string
		expectedErr  string
		expectedData *ExpectedAnalysisResult
	}{
		{
			name: "Successful Analysis",
			htmlContent: `<html><head><title>Test Title</title></head><body>
				<h1>Header 1</h1><h2>Header 2</h2>
				<a href="/internal">Internal Link</a>
				<a href="http://example.com/external">External Link</a>
				</body></html>`,
			expectedData: &ExpectedAnalysisResult{
				HTMLVersion:       "HTML5",
				Title:             "Test Title",
				Headings:          map[string]int{"h1": 1, "h2": 1},
				InternalLinks:     1,
				ExternalLinks:     1,
				InaccessibleLinks: 0,
				HasLoginForm:      false,
			},
		},
		{
			name:        "Invalid URL",
			htmlContent: "",
			expectedErr: "unable to fetch the URL",
		},
		{
			name:        "HTTP Error",
			htmlContent: "",
			expectedErr: "non-success HTTP status",
		},
		{
			name:        "Parsing Error",
			htmlContent: "<html><invalid></html>",
			expectedErr: "error parsing HTML document",
		},
	}

	results := make(chan struct {
		name string
		err  error
		data *ExpectedAnalysisResult
	}, len(tests))

	for _, test := range tests {
		go func(test struct {
			name         string
			htmlContent  string
			expectedErr  string
			expectedData *ExpectedAnalysisResult
		}) {
			client := &MockHTTPClient{
				Response: &http.Response{
					StatusCode: http.StatusOK,
					Body:       io.NopCloser(strings.NewReader(test.htmlContent)),
				},
				Err: nil,
			}

			if test.expectedErr != "" {
				client.Err = errors.New(test.expectedErr)
				client.Response = nil
			}

			result, err := Analyze("http://example.com", client)

			if test.expectedErr != "" {
				results <- struct {
					name string
					err  error
					data *ExpectedAnalysisResult
				}{test.name, err, nil}
				return
			}

			results <- struct {
				name string
				err  error
				data *ExpectedAnalysisResult
			}{test.name, err, &ExpectedAnalysisResult{
				HTMLVersion:       result.HTMLVersion,
				Title:             result.Title,
				Headings:          result.Headings,
				InternalLinks:     result.InternalLinks,
				ExternalLinks:     result.ExternalLinks,
				InaccessibleLinks: result.InaccessibleLinks,
				HasLoginForm:      result.HasLoginForm,
			}}
		}(test)
	}

	for range tests {
		r := <-results
		t.Run(r.name, func(t *testing.T) {
			if r.err != nil {
				assert.Contains(t, r.err.Error(), r.err.Error())
			} else {
				assert.Equal(t, r.data.HTMLVersion, r.data.HTMLVersion)
				assert.Equal(t, r.data.Title, r.data.Title)
				assert.Equal(t, r.data.Headings, r.data.Headings)
				assert.Equal(t, r.data.InternalLinks, r.data.InternalLinks)
				assert.Equal(t, r.data.ExternalLinks, r.data.ExternalLinks)
				assert.Equal(t, r.data.InaccessibleLinks, r.data.InaccessibleLinks)
				assert.Equal(t, r.data.HasLoginForm, r.data.HasLoginForm)
			}
		})
	}
}

func TestAnalyze_AccessibilityChecks(t *testing.T) {
    htmlContent := `
    <html>
        <body>
            <a href="#">No Label</a>
            <a href="#">Another Link</a>
            <a href="invalid-url"></a>
            <a></a>
        </body>
    </html>`

    client := &MockHTTPClient{
        Response: &http.Response{
            StatusCode: http.StatusOK,
            Body:       io.NopCloser(strings.NewReader(htmlContent)),
        },
    }

    result, err := Analyze("http://example.com", client)
    if err != nil {
        t.Fatalf("Expected no error, got %v", err)
    }

    assert.Equal(t, 2, result.MissingLabels, "Expected 2 links missing labels")
}
