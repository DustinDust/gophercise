package main

import (
	"fmt"
	"log"

	html "htmllinkparser/pkgs/html"
	"htmllinkparser/pkgs/links"
)

func main() {
	root, err := html.ParseHtml("test.html")
	if err != nil {
		log.Fatalf("Error parsing html %v", err)
	}
	linksList := make([]links.Link, 0)
	links.ProccessNodes(root, &linksList)
	for _, link := range linksList {
		fmt.Println(link.Stringer())
	}
}
