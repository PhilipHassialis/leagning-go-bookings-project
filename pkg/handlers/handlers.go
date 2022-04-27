package handlers

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/PhilipHassialis/leagning-go-bookings-project/pkg/config"
	"github.com/PhilipHassialis/leagning-go-bookings-project/pkg/models"
	"github.com/PhilipHassialis/leagning-go-bookings-project/pkg/render"
)

// Repo is the Repository for the handlers
var Repo *Repository

// Repository is the repository for the handlers
type Repository struct {
	App *config.AppConfig
}

// NewRepo creates a new Repository
func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{App: a}
}

// NewHandlers sets the repository for the handlers
func NewHandlers(r *Repository) {
	Repo = r
}

// Home is the home page handler
func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	// fmt.Fprintf(w, "This is the home page!")

	remoteIP := r.RemoteAddr
	m.App.Session.Put(r.Context(), "remoteIP", remoteIP)

	render.RenderTemplate(w, "home.page.tmpl", &models.TemplateData{})
}

// About is the about page handler
func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	// perform some logic
	stringMap := make(map[string]string)
	stringMap["test"] = "Hello again"

	remoteIP := m.App.Session.GetString(r.Context(), "remoteIP")
	stringMap["remoteIP"] = remoteIP

	render.RenderTemplate(w, "about.page.tmpl", &models.TemplateData{
		StringMap: stringMap,
	})
}

func (m *Repository) Divide(w http.ResponseWriter, r *http.Request) {
	var x float32 = 100
	var y float32 = 10
	f, err := divideValues(x, y)
	if err != nil {
		fmt.Fprintf(w, "cannot divide by zero")
	} else {
		fmt.Fprintf(w, "%.2f divided by %.2f is %.2f", x, y, f)
	}

}

func divideValues(x, y float32) (float32, error) {
	if y == 0 {
		return 0, errors.New("cannot divide by zero")
	}
	return x / y, nil
}
