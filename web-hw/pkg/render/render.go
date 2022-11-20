package render

import (
	"fmt"
	"html/template"
	"net/http"
)

// RenderTemplate renders template using html/template
func RenderTemplate(w http.ResponseWriter, tmpl string) {
	parssedTemplate, _ := template.ParseFiles("./templates/"+tmpl, "./templates/base.layout.html")
	err := parssedTemplate.Execute(w, nil)
	if err != nil {
		fmt.Println("errors parsing template: ", err)
		return
	}
}
