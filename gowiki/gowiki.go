package main

/*Creating a data structure with:
  load and save methods
  Using the net/http package to build web applications
  Using the html/template package to process HTML templates
  Using the regexp package to validate user input
  Using closures */

import (
	"html/template"
	"io/ioutil"
	"net/http"
	"regexp"
)



/*The function template.Must is a convenience wrapper that panics when passed a non-nil error value. Otherwise, it returns the *Template unaltered. A panic is appropriate here; if the templates can't be
loaded the only sensible thing to do is exit the program.
*/

/*For security, create a validPath that prevents users from creating arbitrary path to be read/write. Validate the title with a regular expression.
 */
var validPath = regexp.MustCompile("^/(edit|view|save)/([a-zA-Z0-9]+)$")

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
	filename := "./data/" + p.Title + ".txt"
	return ioutil.WriteFile(filename, p.Body, 0600)
}

/*create a method to load pages called loadPage, that constructs the file name from the title parameter, reads the file's contents into a new variable body, and returns a pointer to a Page literal constructed with the proper title and body values.If the page is not found it should return empty body and an error*/

func loadPage(title string) (*Page, error) {
	filename := "./data/" + title + ".txt"
	body, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return &Page{Title: title, Body: body}, nil
}


//Using Template Caching. templates are stored in tmpl/
var template_path = ".tmpl/"
var templates = template.Must(template.ParseFiles(
	"./tmpl/edit.html", "./tmpl/view.html"))
	
//Create a function to render templates
func renderTemplate(w http.ResponseWriter, tmpl string, p *Page) {
	//ExecuteTemplate applies the template associated with the receiver,
	//in this case a template, that has the name value of tmpl
	//to the specified data object p and writes the output to w
	//if an error occurs, execution stops
	err := templates.ExecuteTemplate(w, tmpl+".html", p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

}

//Catching the error condition in each handler introduces a lot of repeated code.
//Let's wrap each of the handlers in a function that does this validation and error checking.The wrapper function will take a function of the viewHandler type, and return a function of type http.HandlerFunc
//(suitable to be passed to the function http.HandleFunc):

func makeHandler(fn func(w http.ResponseWriter, r *http.Request, title string)) http.HandlerFunc {
	//return a function that uses the validPath expression
	//to validate path and extract the page title
	return func(w http.ResponseWriter, r *http.Request) {
		m := validPath.FindStringSubmatch(r.URL.Path)
		if m == nil {
			http.NotFound(w, r)
			return
		}
		//it returns fn(w, r, m[2]). It is called a closure,
		//because it uses variables defined outside of it
		fn(w, r, m[2]) // m[2] is for submatch group 2: title.
	}
}

//Modify the handler functions to match the argument of the makeHandler
//wrapper function. Add a title argument of type string.
//The goal is to DRY the code that catches the error condition,
//which is repeated in every handler. Since we have a makeHandler func
//it is safe to remove the error catching condition from the handlers.
//There is no need for a getTitle function.
//Create a function viewHandler of type http.HandlerFunc with 3 arguments
//title string in addition to the two arguments: http.ResponseWriter and a pointer to http.Request
//viewHandler will allow users to view wiki pages. It will handle URL prefixed with "/view/"

func viewHandler(w http.ResponseWriter, r *http.Request, title string) {
	p, err := loadPage(title)
	if err != nil {
		http.Redirect(w, r, "/edit/"+title, http.StatusFound)
	}
	renderTemplate(w, "view", p)
}

//make the same modification as applied to viewHandler.
func editHandler(w http.ResponseWriter, r *http.Request, title string) {
	p, err := loadPage(title)
	if err != nil {
		p = &Page{Title: title}
	}
	renderTemplate(w, "edit", p)
}

func saveHandler(w http.ResponseWriter, r *http.Request, title string) {
	body := r.FormValue("body")
	//The value returned by FormValue is of type string.
	// We must convert that value to []byte
	//before it will fit into the Page struct.
	// We use []byte(body) to perform the conversion.
	p := &Page{Title: title, Body: []byte(body)}
	err := p.save()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/view/"+title, http.StatusFound)
}

/*TODO
Add a handler to make the web root redirect to /view/FrontPage.
*/
func webrootHandler(w http.ResponseWriter, r *http.Request) {
	//title := "FrontPage"
	http.Redirect(w, r, "/view/FrontPage", http.StatusFound)
}

/*

*/

//wrap the handler functions with makeHandler in main, before they are registered with the http package:
func main() {
	http.HandleFunc("/view/", makeHandler(viewHandler))
	http.HandleFunc("/edit/", makeHandler(editHandler))
	http.HandleFunc("/save/", makeHandler(saveHandler))
	http.HandleFunc("/", webrootHandler)

	http.ListenAndServe(":8080", nil)
}
