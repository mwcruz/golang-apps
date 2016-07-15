package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
)

//Follow a Webpage link and copy its content to a local file

//Create a struct to hold the data structure
type Page struct {
	Title string
	Body  []byte //io.util library expects []byte, not strings
}

var validImag = regexp.MustCompile("([A-Za-z0-9-_.]+[jpeg|jpg|gif|bmp|png])$")
var fwdslash = regexp.MustCompile("[/|?|=]")
var validDomain = regexp.MustCompile("(?P<proto>^http(?:s)?://)(?P<domain>[w]*[.]?[[:word:]]+.[a-z]{2,3})([[:graph:]]+)")
var validPdf = regexp.MustCompile("[A-Za-z0-9-_.]+pdf$")

func formatFileName(s string) string {
	return fwdslash.ReplaceAllString(s, "_")
}

func (p *Page) save() error {
	filename := "./data/" + p.Title
	return ioutil.WriteFile(filename, p.Body, 0600)
}

//getPage connects to a URL
//returns the content of the page.
//If the page is not found it should return empty body and an error

func getPage(webURL string) (*Page, error) {
	domain_name := validDomain.FindStringSubmatch(webURL) //return type is []string
	title := formatFileName(domain_name[3])               //domain_name[0] is whole matched expresion
	fmt.Println(title)
	if title == "" {
		title = formatFileName(domain_name[2])
	}
	resp, err := http.Get(webURL)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	return &Page{Title: title, Body: []byte(body)}, nil
}

//TODO
/*func getImages(p *Page) []string {}
 */

/*func getPdf(p *Page) []string {}
 */

func main() {
	webURL := "http://www.navegawireless.com/store/pc/msg.asp?message=83"
	p, err := getPage(webURL)
	if err != nil {
		fmt.Println("Page not Found")
		return
	}

	p.save()
	fmt.Println("Saved File Name is " + p.Title)
}
