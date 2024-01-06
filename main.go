package main

import (
	"html/template"
	"log"
	"net/http"
	"strings"
)

var navigation = Navigation{
	Routes: []Route{
		{Title: "Home", Href: "/", Target: SELF.String(), Relation: NONE.String()},
		{Title: "Form", Href: "/form", Target: SELF.String(), Relation: NONE.String()},
	},
}

var templ *template.Template
var err error
var fileHttp = http.NewServeMux()

func indexHandler(w http.ResponseWriter, r *http.Request) {
	if strings.Contains(r.URL.Path, ".") {
		fileHttp.ServeHTTP(w, r)
		return
	}

	if !strings.EqualFold(r.URL.Path, "/") {
		http.Error(w, "404 | Not found", http.StatusNotFound)
		return
	}

	page := Page{
		Title:       "Website",
		Description: "Website Description",
		Index:       "index",
		Follow:      "follow",
		Navigation:  navigation,
		Body: []Block[any]{
			{Type: "hero", Content: Hero{
				Headline: "Hero",
			}},
			{Type: "text", Content: Text{
				Headline: "Text",
				Bodytext: "Lorem ipsum...",
			}},
		},
	}

	templ.ExecuteTemplate(w, "index.html", page)
}

func formHandler(w http.ResponseWriter, r *http.Request) {
	page := Page{
		Title:       "Form",
		Description: "Form Description",
		Index:       "index",
		Follow:      "follow",
		Navigation:  navigation,
		Body: []Block[any]{
			{Type: "hero", Content: Text{
				Headline: "Hero",
				Bodytext: "Lorem ipsum...",
			}},
			{Type: "form", Content: Form{
				Headline: "Form",
			}},
		},
	}

	templ.ExecuteTemplate(w, "index.html", page)
}

func sendHandler(w http.ResponseWriter, r *http.Request) {
	firstname := r.FormValue("firstname")
	lastname := r.FormValue("lastname")
	email := r.FormValue("email")

	if firstname == "" || lastname == "" || email == "" {
		http.Redirect(w, r, "/form", http.StatusNoContent)
		return
	}

	page := Page{
		Title:       "Send",
		Description: "Send Description",
		Index:       "noindex",
		Follow:      "nofollow",
		Navigation:  navigation,
		Body: []Block[any]{
			{Type: "hero", Content: Hero{
				Headline: "Thank you",
				Bodytext: "Lorem ipsum...",
			}},
			{Type: "confirm", Content: Confirmation{
				Headline:  "Submitted data",
				Bodytext:  "Lorem ipsum...",
				Firstname: firstname,
				Lastname:  lastname,
				Email:     email,
			}},
		},
	}

	templ.ExecuteTemplate(w, "index.html", page)
}

func main() {
	templ, err = template.ParseGlob("templates/*.html")

	if err != nil {
		log.Fatal(err)
	}

	fileHttp.Handle("/", http.FileServer(http.Dir("./public")))
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/form", formHandler)
	http.HandleFunc("/send", sendHandler)

	servErr := http.ListenAndServe(":80", nil)

	if servErr != nil {
		log.Fatal(servErr)
	}
}
