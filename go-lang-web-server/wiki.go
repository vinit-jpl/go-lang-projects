package main

import (
	"errors"
	"html/template"
	"log"
	"net/http"
	"os"
	"regexp"
)

var templates = template.Must(template.ParseFiles("edit.html", "view.html"))
var validPath = regexp.MustCompile("^/(edit|save|view)/([a-zA-Z0-9]+)$")

// declaring a struct to hold wiki page data
type WikiPage struct {
	Title string
	Body  []byte
}

// method with a receiver p of type pointer to WikiPage.
// This means the method is called on a *WikiPage type, like page.Save().
func (p *WikiPage) Save() error {
	filename := p.Title + ".txt"
	return os.WriteFile(filename, p.Body, 0600)
}

func getTitle(w http.ResponseWriter, r *http.Request) (string, error) {
	m := validPath.FindStringSubmatch(r.URL.Path)

	if m == nil {
		http.NotFound(w, r)
		return "", errors.New("Invalid page title")
	}
	return m[2], nil
}

// function to load a wiki page from a file using the filename based on the title.
func loadWikiPage(title string) (*WikiPage, error) {
	filename := title + ".txt"
	body, err := os.ReadFile(filename)

	if err != nil {
		return nil, err
	}
	return &WikiPage{Title: title, Body: body}, nil
}

func renderTemplate(w http.ResponseWriter, tmpl string, p *WikiPage) {
	err := templates.ExecuteTemplate(w, tmpl+".html", p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

}

func viewHandler(w http.ResponseWriter, r *http.Request) {
	title, err := getTitle(w, r)

	if err != nil {
		return
	}

	p, err := loadWikiPage(title)
	if err != nil {
		http.Redirect(w, r, "/edit/"+title, http.StatusFound)
		return
	}
	// t, _ := template.ParseFiles("view.html")
	// t.Execute(w, p)
	renderTemplate(w, "view", p)
}

func editHandler(w http.ResponseWriter, r *http.Request) {
	title, err := getTitle(w, r)

	if err != nil {
		return
	}

	p, err := loadWikiPage(title)
	if err != nil {
		p = &WikiPage{Title: title}
	}
	// t, _ := template.ParseFiles("edit.html")
	// t.Execute(w, p)
	renderTemplate(w, "edit", p)
}

func saveHandler(w http.ResponseWriter, r *http.Request) {
	title, err := getTitle(w, r)

	if err != nil {
		return
	}

	body := r.FormValue("body")
	p := &WikiPage{Title: title, Body: []byte(body)}
	err = p.Save()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/view/"+title, http.StatusFound)
}

func main() {
	// p1 := &WikiPage{Title: "TestPage", Body: []byte("This is a test page.")}
	// p1.Save()
	// p2, _ := loadWikiPage("TestPage")
	// fmt.Println(string(p2.Body)) // p2.Body is a byte slice, so we convert it to string for printing

	http.HandleFunc("/view/", viewHandler)
	http.HandleFunc("/edit/", editHandler)
	http.HandleFunc("/save/", saveHandler)

	log.Println("Server started on port: 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
