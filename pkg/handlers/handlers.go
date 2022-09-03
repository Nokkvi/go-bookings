package handlers

import (
	"net/http"

	"github.com/nokkvi/go-bookings/pkg/config"
	"github.com/nokkvi/go-bookings/pkg/models"
	"github.com/nokkvi/go-bookings/pkg/render"
)

// Repo is the repository used by the handlers
var Repo *Repository


// Repository is the repository type
type Repository struct {
	App *config.AppConfig
}

// NewRepo creates a new Repository
func NewRepo(a *config.AppConfig) *Repository {
	return &Repository {
		App: a,
	}
}

// NewHandlers sets the repository for the handlers
func NewHandlers(r *Repository) {
	Repo = r
}

// Home is the home page handler
func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	m.App.Session.Put(r.Context(), "remote_ip", r.RemoteAddr)

	render.RenderTemplate(w, "home.page.tmpl", &models.TemplateData{})
}

// About is the about page handler
func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	// perform some logic
	stringMap := make(map[string]string)
	stringMap["test"] = "Hello, matey!"

	stringMap["remote_ip"] = m.App.Session.GetString(r.Context(), "remote_ip")

	// send data to template
	render.RenderTemplate(w, "about.page.tmpl", &models.TemplateData{
		StringMap: stringMap,
	})
}

