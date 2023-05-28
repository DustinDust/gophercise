package html_test

import (
	"htmllinkparser/pkgs/html"
	"testing"
)

func TestParserInvalidPath(t *testing.T) {
	path := "Invalid file path"
	_, err := html.ParseHtml(path)
	if err == nil {
		t.Fatalf("Must return error")
	}
}
