package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
)

type Config struct {
	yaml *string
}

func getConfig() Config {
	yamlFilePath := flag.String("yaml", "./url.yaml", "Yaml file path")
	flag.Parse()
	return Config{
		yaml: yamlFilePath,
	}
}

func main() {
	mux := defaultMux()
	config := getConfig()

	// Build the MapHandler using the mux as the fallback
	pathsToUrls := map[string]string{
		"/urlshort-godoc": "https://godoc.org/github.com/gophercises/urlshort",
		"/yaml-godoc":     "https://godoc.org/gopkg.in/yaml.v2",
	}
	mapHandler := MapHandler(pathsToUrls, mux)

	var yaml []byte
	yaml, err := os.ReadFile(*config.yaml)
	if err != nil {
		log.Printf("Error opening file %v", err)
		yaml = []byte(`
- path: /urlshort
  url: https://github.com/gophercises/urlshort
- path: /urlshort-final
  url: https://github.com/gophercises/urlshort/tree/solution
`)
	}
	// Build the YAMLHandler using the mapHandler as the
	// fallback
	yamlHandler, err := YAMLHandler([]byte(yaml), mapHandler)
	if err != nil {
		panic(err)
	}
	fmt.Println("Starting the server on :8080")
	http.ListenAndServe(":8080", yamlHandler)
}

func defaultMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", hello)
	return mux
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, world!")
}
