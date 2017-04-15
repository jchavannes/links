package main

import (
	"github.com/jchavannes/jgo/web"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
)

var (
	indexRoute = web.Route{
		Pattern: "/",
		Handler: func(r *web.Response) {
			r.Helper["Links"] = getLinks()
			r.Render()
		},
	}
)

func main() {
	server := web.Server{
		Port: 8251,
		StaticFilesDir: "web",
		TemplatesDir: "web",
		Routes: []web.Route{
			indexRoute,
		},
	}
	server.Run()
}

type Link struct {
	Name string
	Url  string
	Date string
}

func getLinks() []Link {
	data, err := ioutil.ReadFile("links.yaml")
	if err != nil {
		log.Fatal(err)
	}
	links := struct{
		Links []Link
	}{}
	err = yaml.Unmarshal(data, &links)
	if err != nil {
		log.Fatal(err)
	}
	return links.Links
}
