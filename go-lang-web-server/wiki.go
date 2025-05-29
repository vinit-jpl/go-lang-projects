package main

import (
	"fmt"
	"os"
)

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

// function to load a wiki page from a file using the filename based on the title.
func loadWikiPage(title string) (*WikiPage, error) {
	filename := title + ".txt"
	body, err := os.ReadFile(filename)

	if err != nil {
		return nil, err
	}
	return &WikiPage{Title: title, Body: body}, nil
}

func main() {
	p1 := &WikiPage{Title: "TestPage", Body: []byte("This is a test page.")}
	p1.Save()
	p2, _ := loadWikiPage("TestPage")
	fmt.Println(string(p2.Body))
}
