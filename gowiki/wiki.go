package main

import (
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
)

type Page struct {
	Title string
	Body  []byte
}

//The character * in the following functions, represents a pointer

//This function returns an error
func (p *Page) save() error {
	filename := p.Title + ".txt"
	//The number 0600 indicates that the file should be created with
	//read-write permission for the current user only
	return ioutil.WriteFile(filename, p.Body, 0600)
}

//This function returns a Page struct and an error, in case that the file doesn't exists
func loadPage(title string) (*Page, error) {
	filename := title + ".txt"
	body, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	//The & is used to create a new Page in Page struct
	return &Page{Title: title, Body: body}, nil
}

//Function to show the page
//Uses the view.html template instead of hard-coded HTML
func viewHandler(w http.ResponseWriter, r *http.Request, title string) {
	//Let's handle the error of non-existing page
	p, err := loadPage(title)
	if err != nil {
		//The http.Redirect adds an status of 302 and a Location header to the HTTP response
		http.Redirect(w, r, "/edit/"+title, http.StatusFound)
		return
	}
	//Call the renderTemplate func
	renderTemplate(w, "view", p)
}

//Function to edit the page. Creates a new one if it doesn't exist
//Uses the edit.html template instead of hard-coded HTML
func editHandler(w http.ResponseWriter, r *http.Request, title string) {
	p, err := loadPage(title)
	if err != nil {
		//Create a new one of type Page
		p = &Page{Title: title}
	}
	//Call the renderTemplate func
	renderTemplate(w, "edit", p)
}

//Function to save the edited page. Uses the save.html template instead of hard-coded HTMl
func saveHandler(w http.ResponseWriter, r *http.Request, title string) {
	//We reffer to the FormValue which name is "body", and its type is a string
	body := r.FormValue("body")
	p := &Page{Title: title, Body: []byte(body)}
	//Let's handle the possible error saving the new file
	err := p.save()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/view/"+title, http.StatusFound)
}

//Global variable named "templates" that initialize it with ParseFiles method
var templates = template.Must(template.ParseFiles("edit.html", "view.html"))

//Let's create a renderTemplate func to render the edit,save and view html files
func renderTemplate(w http.ResponseWriter, tmpl string, p *Page) {
	//Let's handle the error executing the template files
	err := templates.ExecuteTemplate(w, tmpl+".html", p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

//Global variable to store our validation expression
var validPath = regexp.MustCompile("^/(edit|save|view)/([a-zA-Z0-9]+)$")

//Wrapper function that takes a function of the Handler type and return function of type http.HandlerFunc
func makeHandler(fn func(http.ResponseWriter, *http.Request, string)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//Here we will extract the page title from the Request,
		//and call the provided handler 'fn'
		m := validPath.FindStringSubmatch(r.URL.Path)
		if m == nil {
			http.NotFound(w, r)
			return
		}
		fn(w, r, m[2])
	}
}

//Let's create a main that handle what we've written
func main() {
	//We'll handle any requests under the path /view/
	http.HandleFunc("/view/", makeHandler(viewHandler))
	//We will handle the ability to edit the page and to save the changes
	http.HandleFunc("/edit/", makeHandler(editHandler))
	http.HandleFunc("/save/", makeHandler(saveHandler))

	log.Fatal(http.ListenAndServe(":8080", nil))
}
