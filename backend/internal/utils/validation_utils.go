package utils

import "net/url"

//checks if the given string is a valid URL
func IsValidURL(rawURL string) bool {
	_, err := url.ParseRequestURI(rawURL)
	return err == nil
}
