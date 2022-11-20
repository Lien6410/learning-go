package render

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

func init() {
	// v1, _ := filepath.Abs("./web-hw")
	// fmt.Println(v1)
}

// RenderTemplate renders template using html/template
func RenderTemplateTest(w http.ResponseWriter, tmpl string) {
	parssedTemplate, _ := template.ParseFiles("./web-cache2/templates/"+tmpl, "./web-cache2/templates/base.layout.html")

	err := parssedTemplate.Execute(w, nil)
	if err != nil {
		fmt.Println("errors parsing template: ", err)
		return
	}
}

var tc = make(map[string]*template.Template)

func RenderTemplate(w http.ResponseWriter, t string) {
	var tmpl *template.Template
	var err error

	// check to see if we already have the template in our cache
	_, inMap := tc[t]
	if !inMap {
		// need to create the template
		log.Println("creating template and adding to cache")
		err = createTemplateCache(t)
		if err != nil {
			log.Println(err)
		}
	} else {
		// we have the template in the cache
		log.Println("using cached template")
	}

	tmpl = tc[t]

	err = tmpl.Execute(w, nil)
	if err != nil {
		log.Println(err)
	}
}

func createTemplateCache(t string) error {
	templates := []string{
		fmt.Sprintf("./web-cache2/templates/%s", t),
		"./web-cache2/templates/base.layout.html",
	}

	// parse the template
	tmpl, err := template.ParseFiles(templates...)
	if err != nil {
		return err
	}

	// add template to cahce (map)
	tc[t] = tmpl

	return nil
}
