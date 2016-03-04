package main

//Creating a data structure with:
// load and save methods
//    Using the net/http package to build web applications
//    Using the html/template package to process HTML templates
//    Using the regexp package to validate user input
//    Using closures

import (
	"errors"
	"html/template"
	"io/ioutil"
	"net/http"
	"regexp"
)

//Imported net/http to use http library

//Using Template Caching
var templates = template.Must(template.ParseFiles(
	"edit.html", "view.html"))

//The function template.Must is a convenience wrapper
// that panics when passed a non-nil error value.
// Otherwise, it returns the *Template unaltered.
// A panic is appropriate here; if the templates can't be
// loaded the only sensible thing to do is exit the program.

//For security, create a validPath that prevents users from
//creating arbitrary path to be read/write.
//write a function to validate the title with a regular expression.
var validPath = regexp.MustCompile("^/(edit|view|save)/([a-zA-Z0-9]+)$")

//regexp.MustCompile will parse and compile the regular expression,
// and return a regexp.Regexp. This regexp has two capturing groups
// or parenthesized expressions(..) numbered 1 and 2 from left to right.

//Crete a function that uses the validPath expression
// to validate path and extract the page title:
func getTitle(w http.ResponseWriter, r *http.Request) (string, error) {
	//calling validPath.FindStringSubmatch(r.URL.Path) will return
	//a slice identifying the successive submatches of the expression
	//validPath. Submatch 0 is the entire expression, submatch 1 is the
	//group 1 and submatch 2 is the group 2. what we want is our case
	m := validPath.FindStringSubmatch(r.URL.Path)
	if m == nil {
		http.NotFound(w, r)
		return "", errors.New("Invalid Page Title")
	}
	return m[2], nil // returns submatch group 2 which is the title
}

//creating a struct to hold the data structure
type Page struct {
	Title string
	Body  []byte //io libraries expect a []byte instead of string
}

//add a save method that takes as receiver p, a pointer to Page
//takes no parameters
//that will save p.Body to filename.txt.
//filename will be generated from p.Title+".txt"
//filename create with read-write privileges for current user only (rw- or octal 0600)
//the method returns error, for it is the what WriteFile returns
func (p *Page) save() error {
	filename := p.Title + ".txt"
	return ioutil.WriteFile(filename, p.Body, 0600)
}

/*create a method to load pages called loadPage, that constructs the file name from the title parameter, reads the file's contents into a new variable body, and returns a pointer to a Page literal constructed with the proper title and body values.If the page is not found it should return empty body and an error*/

func loadPage(title string) (*Page, error) {
	filename := title + ".txt"
	body, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return &Page{Title: title, Body: body}, nil
}

//Create a function to render templates
func renderTemplate(w http.ResponseWriter, tmpl string, p *Page) {
	//ExecuteTemplate applies the template associated with the receiver,
	//in this case a template that has the name value of tmpl
	//to the specified data object p and writes the output to w
	//if an error occurs, execution stops
	err := templates.ExecuteTemplate(w, tmpl+".html", p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

}

//Create a function viewHandler of type http.HandlerFunc
//http.HandlerFunc take two arguments: http.ResponseWriter and a pointer to http.Request
//viewHandler will allow users to view wiki pages. It will hadle URL prefixed with "/view/"

func viewHandler(w http.ResponseWriter, r *http.Request) {
	title, err := getTitle(w, r)
	if err != nil {
		return
	}
	p, err := loadPage(title)
	if err != nil {
		//redirect with func Redirect(w ResponseWriter, r *Request, urlStr string, code int)
		http.Redirect(w, r, "/edit/"+title, http.StatusFound)
	}
	renderTemplate(w, "view", p)
}

//Create handlers for edit and save to allow for editing and saving wiki pages
func editHandler(w http.ResponseWriter, r *http.Request) {
	title, err := getTitle(w, r)
	if err != nil {
		return
	}
	p, err := loadPage(title)
	if err != nil {
		p = &Page{Title: title}
	}
	renderTemplate(w, "edit", p)
}

//The function saveHandler will handle the submission
// of forms located on the edit pages.
func saveHandler(w http.ResponseWriter, r *http.Request) {
	title, err := getTitle(w, r)
	if err != nil {
		return
	}
	body := r.FormValue("body")
	//The value returned by FormValue is of type string.
	// We must convert that value to []byte
	//before it will fit into the Page struct.
	// We use []byte(body) to perform the conversion.
	p := &Page{Title: title, Body: []byte(body)}
	err = p.save()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/view/"+title, http.StatusFound)
}

//write a main function to initialize http using the viewHandler to handle any requests under the path /view/
func main() {
	http.HandleFunc("/view/", viewHandler)
	http.HandleFunc("/edit/", editHandler)
	http.HandleFunc("/save/", saveHandler)
	http.ListenAndServe(":8080", nil)
}
