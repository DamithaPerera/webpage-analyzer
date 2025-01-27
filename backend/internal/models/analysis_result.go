package models

// AnalysisResult represents the analysis result for a webpage
type AnalysisResult struct {
	HTMLVersion       string         `json:"html_version"`
	Title             string         `json:"title"`
	Headings          map[string]int `json:"headings"`
	InternalLinks     int            `json:"internal_links"`
	ExternalLinks     int            `json:"external_links"`
	InaccessibleLinks int            `json:"inaccessible_links"`
	HasLoginForm      bool           `json:"has_login_form"`
	MissingLabels     int            `json:"missing_labels"`
	InvalidHref       int            `json:"invalid_href"`
}
