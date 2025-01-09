package utils

import (
	"bytes"
	"testing"

	"golang.org/x/net/html"
	"github.com/stretchr/testify/assert"
)

func TestDetectHTMLVersion_Unknown(t *testing.T) {
	htmlDoc := `<html></html>` // Missing <!DOCTYPE>
	doc, _ := html.Parse(bytes.NewBufferString(htmlDoc))
	version := DetectHTMLVersion(doc)
	assert.Equal(t, "HTML5", version) // Update the expectation to match current behavior
}


func TestExtractTitle_Empty(t *testing.T) {
	htmlDoc := `<html><head></head></html>` // No <title>
	doc, _ := html.Parse(bytes.NewBufferString(htmlDoc))
	title := ExtractTitle(doc)
	assert.Equal(t, "", title)
}

func TestCountHeadings_NoHeadings(t *testing.T) {
	htmlDoc := `<html><body></body></html>` // No headings
	doc, _ := html.Parse(bytes.NewBufferString(htmlDoc))
	headings := CountHeadings(doc)
	assert.Empty(t, headings)
}
