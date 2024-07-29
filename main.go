package main

import (
	"flag"
	"log"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/geziyor/geziyor"
	"github.com/geziyor/geziyor/client"
	"github.com/geziyor/geziyor/export"
)

var (
	baseURL      string
	searchQuery  string
	outputFile   string
	concurrency  int
)

func init() {
	flag.StringVar(&baseURL, "base-url", "https://www.youtube.com", "Base URL for YouTube")
	flag.StringVar(&searchQuery, "query", "geziyor golang", "Search query for YouTube")
	flag.StringVar(&outputFile, "output", "videos.json", "Output JSON file name")
	flag.IntVar(&concurrency, "concurrency", 1, "Number of concurrent requests")
	flag.Parse()
}

func main() {
	log.Printf("Starting YouTube scraper with query: %s\n", searchQuery)

	geziyor.NewGeziyor(&geziyor.Options{
		StartRequestsFunc: func(g *geziyor.Geziyor) {
			searchURL := baseURL + "/results?search_query=" + strings.ReplaceAll(searchQuery, " ", "+")
			g.GetRendered(searchURL, g.Opt.ParseFunc)
		},
		ParseFunc:                   videosParse,
		Exporters:                   []export.Exporter{&export.JSON{FileName: outputFile}},
		RobotsTxtDisabled:           true,
		ConcurrentRequests:          concurrency,
		ConcurrentRequestsPerDomain: concurrency,
		LogDisabled:                 false,
		UserAgent:                   "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.124 Safari/537.36",
		AllowedDomains:              []string{"www.youtube.com", "youtube.com"},
	}).Start()

	log.Println("Scraping completed. Results saved to", outputFile)
}

func videosParse(g *geziyor.Geziyor, r *client.Response) {
    if r.HTMLDoc == nil {
        log.Println("Error: HTML document is nil")
        return
    }

    r.HTMLDoc.Find("ytd-video-renderer.ytd-item-section-renderer").Each(func(i int, s *goquery.Selection) {
        titleAnchor := s.Find("a#video-title")
        channelAnchor := s.Find("ytd-channel-name#channel-name a[spellcheck='false']")

        metaLine := s.Find("div#metadata-line span")

        relativeDate := metaLine.Last().Text()
        views := strings.TrimSuffix(metaLine.First().Text(), " views")

        channelName := channelAnchor.First().Text()

        g.Exports <- map[string]interface{}{
            "id":            i,
            "title":         titleAnchor.AttrOr("title", ""),
            "url":           baseURL + titleAnchor.AttrOr("href", ""),
            "views":         views,
            "relative_date": relativeDate,
            "channel_name":  channelName,
            "channel_url":   baseURL + channelAnchor.AttrOr("href", ""),
        }
    })
}