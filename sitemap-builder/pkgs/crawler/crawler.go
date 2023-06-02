package crawler

import (
	"encoding/xml"
	"net/http"
	"os"
	"strings"

	"golang.org/x/net/html"
	"gophercise.dd.com/htmllinkparser/pkgs/links"
	"gophercise.dd.com/sitemap-builder/pkgs/queue"
)

type URL struct {
	XMLName xml.Name `xml:"url"`
	Loc     string   `xml:"loc"`
}

type Sitemap struct {
	XMLName xml.Name `xml:"urlset"`
	XMLNS   string   `xml:"xmlns,attr"`
	Urls    []URL
}

func CrawlAllLinkFromNodes(url string, depth int) ([]URL, error) {
	urlMap := make([]URL, 0)
	node, err := GetHtml(url)
	if err != nil {
		return nil, err
	}

	firstPageLinks := make([]links.Link, 0)
	visited := make(map[string]interface{})
	links.ProccessNodes(node, &firstPageLinks)
	qu := &queue.Queue[string]{
		Size: 100000,
	}
	for _, link := range firstPageLinks {
		dep := 0
		href := FormatUrl(url, link.Href)
		if err := qu.Enqueue(href); err != nil {
			return nil, err
		}
		visited[href] = true
		for !qu.IsEmpty() {
			if dep >= depth {
				break
			}
			href, err := qu.Dequeue()
			if err != nil {
				return nil, err
			}
			node, err := GetHtml(href)
			if err != nil {
				return nil, err
			}
			linkPage := make([]links.Link, 0)
			links.ProccessNodes(node, &linkPage)
			for _, link := range linkPage {
				nref := FormatUrl(href, link.Href)
				if _, ok := visited[nref]; !ok {
					urlMap = append(urlMap, URL{
						Loc: nref,
					})
					visited[nref] = true
					qu.Enqueue(nref)
				}
			}
			dep++
		}
	}

	return urlMap, nil
}

func FormatUrl(baseUrl string, url string) string {
	if strings.HasPrefix(url, "/") {
		return baseUrl + url
	}
	if !strings.HasPrefix(url, "http") {
		return baseUrl + "/" + url
	}
	return url
}

func GetHtml(url string) (*html.Node, error) {
	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	return html.Parse(res.Body)
}

func MakeSitemapXML(url []URL, filepath string) error {
	sitemap := Sitemap{
		XMLNS: "http://www.sitemaps.org/schemas/sitemap/0.9",
		Urls:  url,
	}
	output, err := xml.MarshalIndent(sitemap, "", "  ")
	if err != nil {
		return err
	}
	file, err := os.Create(filepath)
	if err != nil {
		return err
	}
	file.Write([]byte(xml.Header))
	if _, err = file.Write(output); err != nil {
		return err
	}
	return nil
}
