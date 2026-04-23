package webscraper

import (
	"github.com/irajwani/go-web-scraper/internal/config"
	"github.com/irajwani/go-web-scraper/internal/fetcher"
	"github.com/irajwani/go-web-scraper/internal/parser"
)

func Extract(url string) (string, error) {

	args := make([]string, 0)
	args[0] = url
	cfg, err := config.ParseFlags(args)

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
