package handlers

import (
	"net/http"

	"github.com/Lien6410/learning-go/web-config/pkg/config"
	"github.com/Lien6410/learning-go/web-config/pkg/models"
	"github.com/Lien6410/learning-go/web-config/pkg/render"
)

// Repo used by the handlers
var Repo *Repository

// Repository is the repository type
type Repository struct {
	App *config.AppConfig
}

// NewRepository creates a new repository
func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

// NewHandlers sets the repository for the handlers
func NewHandlers(r *Repository) {
	Repo = r
}

// Home is the home page handler
func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "home.page.html", &models.TemplateData{})
}

// About is the about page handler
func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	// perfrom some logic
	stringMap := map[string]string{}
	stringMap["test"] = "Hello, again."

	// send the data to the template
	render.RenderTemplate(w, "about.page.html", &models.TemplateData{
		StringMap: stringMap,
	})
}
