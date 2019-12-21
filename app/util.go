package app

import (
	"bufio"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

func normalizationURL(queryURL string) (string, error) {
	u, err := url.Parse(queryURL)
	if err != nil {
		return "", fmt.Errorf("failed parse URL: %w", err)
	}
	return fmt.Sprintf("%s:%s%s", u.Scheme, u.Host, u.Path), nil
}

// len(<title>) == 7
var tagWord = 7

func getTitleFromURL(queryURL string) (string, error) {
	queryRes, err := http.Get(queryURL)
	if err != nil {
		return "", fmt.Errorf("failed get URL title: %w", err)
	}
	scanner := bufio.NewScanner(queryRes.Body)
	for scanner.Scan() {
		line := scanner.Text()
		if i := strings.Index(line, "<title>"); i >= 0 {
			if j := strings.Index(line, "</title>"); j >= 0 {
				return line[i+tagWord : j], nil
			}
		}
	}
	return "", errors.New("failed search title")
}
