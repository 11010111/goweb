package main

import (
	"html/template"
	"log"
	"net/http"
	"strings"

	"github.com/11010111/goweb/types"
)

var routes = []types.Route{
	{Title: "Home", Href: "/", Target: types.TARGET_SELF.String(), Relation: types.REL_NONE.String()},
	{Title: "Form", Href: "/form", Target: types.TARGET_SELF.String(), Relation: types.REL_NONE.String()},
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

	page := types.NewPage()
	page.SetTitle("Website")
	page.SetDescription("Website Description")
	page.SetIndex("index")
	page.SetFollow("follow")
	page.SetRoutes(routes)

	hero := types.NewHero()
	hero.SetHeadline("Hero Block")
	hero.SetBodytext("Lorem ipsum...")

	heroContent := types.NewContent()
	heroContent.SetBlock(hero)
	heroContent.SetContentType("hero")
	page.AppendContent(heroContent)

	text := types.NewText()
	text.SetHeadline("Text Block")
	text.SetBodytext("Lorem ipsum...")

	textContent := types.NewContent()
	textContent.SetContentType("text")
	textContent.SetBlock(text)
	page.AppendContent(textContent)

	templ.ExecuteTemplate(w, "index.html", page)
}

func formHandler(w http.ResponseWriter, r *http.Request) {
	page := types.NewPage()
	page.SetTitle("Website Form")
	page.SetDescription("Website Form Description")
	page.SetIndex("index")
	page.SetFollow("follow")
	page.SetRoutes(routes)

	hero := types.NewHero()
	hero.SetHeadline("Hero Block")
	hero.SetBodytext("Lorem ipsum...")

	heroContent := types.NewContent()
	heroContent.SetBlock(hero)
	heroContent.SetContentType("hero")
	page.AppendContent(heroContent)

	form := types.NewForm()
	form.SetHeadline("Text Block")

	formContent := types.NewContent()
	formContent.SetBlock(form)
	formContent.SetContentType("form")
	page.AppendContent(formContent)

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

	page := types.NewPage()
	page.SetTitle("Website Send Route")
	page.SetDescription("Website Send Route Description")
	page.SetIndex("noindex")
	page.SetFollow("nofollow")
	page.SetRoutes(routes)

	hero := types.NewHero()
	hero.SetHeadline("Hero Block")
	hero.SetBodytext("Lorem ipsum...")

	heroContent := types.NewContent()
	heroContent.SetBlock(hero)
	heroContent.SetContentType("hero")
	page.AppendContent(heroContent)

	confirm := types.NewConfirmation()
	confirm.SetHeadline("Submitted data")
	confirm.SetBodytext("Lorem ipsum...")
	confirm.SetFirstname(firstname)
	confirm.SetLastname(lastname)
	confirm.SetEmail(email)

	confirmContent := types.NewContent()
	confirmContent.SetBlock(confirm)
	confirmContent.SetContentType("confirm")
	page.AppendContent(confirmContent)

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
