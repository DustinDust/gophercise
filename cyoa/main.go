package main

import (
	"cyoa/stories"
	"flag"
	"html/template"
	"log"
	"net/http"
	"os"
)

type Config struct {
	json *string
}

type StoryHandler struct {
	Story    *stories.Story
	Template *template.Template
}

func main() {
	config := LoadConfig()
	storiesHandler := StoryHandler{}
	err := SetupHandler(config, &storiesHandler)
	if err != nil {
		panic(err)
	}
	log.Printf("Listening at :8080")
	http.ListenAndServe(":8080", &storiesHandler)
}

func LoadConfig() Config {
	json := flag.String("json", "story.json", "path to the story json file")
	flag.Parse()
	return Config{
		json: json,
	}
}

func SetupHandler(config Config, storyHandler *StoryHandler) error {
	if storyHandler.Story == nil {
		storyHandler.Story = &stories.Story{}
	}
	jsonContent, err := os.ReadFile(*config.json)
	if err != nil {
		return err
	}
	err = stories.ParseStory(jsonContent, storyHandler.Story)
	if err != nil {
		return err
	}
	storyHandler.Template, err = template.ParseFiles("view/view.html")
	if err != nil {
		return err
	}
	return nil
}

func (s *StoryHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		if r.URL.Path == "/story" {
			arc := r.URL.Query().Get("arc")
			if len(arc) == 0 {
				arc = "intro"
			}
			storyArc, ok := (*s.Story)[arc]
			if !ok {
				w.Write([]byte("Internal Server Error - story arc not found"))
				w.WriteHeader(500)
				return
			}
			err := s.Template.Execute(w, storyArc)
			if err != nil {
				w.Write([]byte("Internal Server Error - execute template failed"))
				w.WriteHeader(500)
				return
			}
		}
	}
}
