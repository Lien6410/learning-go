package render

import (
	"fmt"
	"html/template"
	"net/http"
)

func init() {
	// v1, _ := filepath.Abs("./web-hw")
	// fmt.Println(v1)
}

// RenderTemplate renders template using html/template
func RenderTemplate(w http.ResponseWriter, tmpl string) {
	parssedTemplate, _ := template.ParseFiles("./web-hw/templates/"+tmpl, "./web-hw/templates/base.layout.html")

	err := parssedTemplate.Execute(w, nil)
	if err != nil {
		fmt.Println("errors parsing template: ", err)
		return
	}
}
