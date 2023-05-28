package links

import (
	"fmt"
	"strings"

	"golang.org/x/net/html"
)

type Link struct {
	Href string
	Text string
}

func (l Link) Stringer() string {
	return fmt.Sprintf("href: %s, text: %v", l.Href, l.Text)
}

func ProccessNodes(root *html.Node, links *[]Link) {
	for node := root.FirstChild; node != nil; node = node.NextSibling {
		if node.Type == html.ElementNode && node.Data == "a" {
			for _, attr := range node.Attr {
				if attr.Key == "href" {
					textVal := ""
					getAllChildText(node, &textVal)
					*links = append(*links, Link{
						Href: attr.Val,
						Text: textVal,
					})
				}
			}
		} else {
			ProccessNodes(node, links)
		}
	}
}

func getAllChildText(node *html.Node, res *string) {
	for c := node.FirstChild; c != nil; c = c.NextSibling {
		if c.Type == html.TextNode {
			*res = fmt.Sprintf("%s %s", *res, FormatString(c.Data))
			*res = FormatString(*res)
		}
		getAllChildText(c, res)
	}
}

func FormatString(s string) string {
	s = strings.ReplaceAll(s, "\n", "")
	s = strings.TrimSpace(s)
	return s
}
