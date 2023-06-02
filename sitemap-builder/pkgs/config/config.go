package config

import "flag"

type Config struct {
	Depth int
	Url   string
}

func GetConfig() *Config {
	dep := flag.Int("depth", 3, "set depth of the crawler")
	url := flag.String("url", "http://example.com", "set url of the website you want to make a sitemap of")

	flag.Parse()

	return &Config{
		Depth: *dep,
		Url:   *url,
	}
}
