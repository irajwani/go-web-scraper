package webscraper

import (
	"github.com/irajwani/gowebscraper/internal/config"
	"github.com/irajwani/gowebscraper/internal/fetcher"
	"github.com/irajwani/gowebscraper/internal/parser"
)

type Options struct {
	// expose only what you want users to control
	UserAgent string
	// add more fields as needed
}

func Extract(url string, opts *Options) (string, error) {
	// map Options → internal config
	cfg := &config.Config{
		URL: url,
		// fill other fields from opts if needed
	}

	content, err := fetcher.FetchAndExtractContent(cfg.URL)
	if err != nil {
		return "", err
	}

	text, err := parser.ExtractTextFromHTML(content, cfg)
	if err != nil {
		return "", err
	}

	return text, nil
}
