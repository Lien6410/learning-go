package render

import (
	"bytes"
	"html/template"
	"log"
	"net/http"
	"path/filepath"

	"github.com/Lien6410/learning-go/web-config/pkg/config"
)

func init() {
	// v1, _ := filepath.Abs("./web-hw")
	// fmt.Println(v1)
}

var app *config.AppConfig

// NewTemplate sets the config to the template package
func NewTemplate(a *config.AppConfig) {
	app = a
}

// RenderTemplate renders template using html/template
func RenderTemplate(w http.ResponseWriter, tmpl string) {
	var tc map[string]*template.Template
	if app.UseCache {
		// get the template cache from the app config
		tc = app.TemplateCache
	} else {
		tc, _ = CreateTemplateCache()
	}

	// get request template from cache
	t, ok := tc[tmpl]
	if !ok {
		log.Fatal("template cache not found")
	}

	buf := new(bytes.Buffer)

	_ = t.Execute(buf, nil)

	_, err := buf.WriteTo(w)
	if err != nil {
		log.Println("Error writing template to browser: ", err)
	}

}

func CreateTemplateCache() (map[string]*template.Template, error) {
	// create an empty map[string]*template.Template
	myCache := map[string]*template.Template{}

	// get all of the files named *.page.html from ./web-config/templates/
	pages, err := filepath.Glob("./web-config/templates/*.page.html")
	if err != nil {
		return myCache, err
	}

	// range through all files ending with *.page.html
	for _, page := range pages {
		name := filepath.Base(page)
		ts, err := template.New(name).ParseFiles(page)
		if err != nil {
			return myCache, err
		}

		matches, err := filepath.Glob("./web-config/templates/*.layout.html")
		if err != nil {
			return myCache, err
		}

		if len(matches) > 0 {
			ts, err = ts.ParseGlob("./web-config/templates/*.layout.html")
			if err != nil {
				return myCache, err
			}
		}

		myCache[name] = ts
	}

	return myCache, nil
}
