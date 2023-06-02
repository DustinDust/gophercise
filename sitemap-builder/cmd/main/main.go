package main

import (
	"log"

	"gophercise.dd.com/sitemap-builder/pkgs/config"
	"gophercise.dd.com/sitemap-builder/pkgs/crawler"
)

func main() {
	config := config.GetConfig()
	urlMap, err := crawler.CrawlAllLinkFromNodes(config.Url, config.Depth)
	if err != nil {
		log.Fatal(err)
	}
	crawler.MakeSitemapXML(urlMap, "./sitemap.xml")
}
