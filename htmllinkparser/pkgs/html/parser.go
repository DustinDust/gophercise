package html

import (
	"bytes"
	"os"

	"golang.org/x/net/html"
)

func ParseHtml(path string) (*html.Node, error) {
	htmlBytes, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	r := bytes.NewReader(htmlBytes)
	root, err := html.Parse(r)
	if err != nil {
		return nil, err
	}
	return root, nil
}
