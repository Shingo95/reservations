package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/elorenzotti/bookings/internal/config"
	"github.com/elorenzotti/bookings/internal/models"
	"github.com/elorenzotti/bookings/internal/render"
)

// Repo is the variable for the repository
var Repo *Repository

// Repository is the repository type
type Repository struct {
	App *config.AppConfig
}

// NewRepo creates a new repository

func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

// NewHandlers sets the repository for the handlers
func NewHandlers(r *Repository) {
	Repo = r
}

// Home is the homepage handler
func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	remoteIP := r.RemoteAddr
	m.App.Session.Put(r.Context(), "remote_ip", remoteIP)

	render.RenderTemplate(w, r, "home.page.tmpl", &models.TemplateData{})

}

// About is the about page handler
func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	stringMap := make(map[string]string)
	stringMap["test"] = "Hello again"

	remoteIP := m.App.Session.GetString(r.Context(), "remote_ip")
	stringMap["remote_IP"] = remoteIP

	//Send the data to the template
	render.RenderTemplate(w, r, "about.page.tmpl", &models.TemplateData{
		StringMap: stringMap,
	})

}

// Bucolic is the bucolic-room page handler
func (m *Repository) Bucolic(w http.ResponseWriter, r *http.Request) {
	//Send the data to the template
	render.RenderTemplate(w, r, "bucolic-room.page.tmpl", &models.TemplateData{})

}

// London is the Old London Room page handler
func (m *Repository) London(w http.ResponseWriter, r *http.Request) {
	//Send the data to the template
	render.RenderTemplate(w, r, "old-london.page.tmpl", &models.TemplateData{})

}

// Availability is the search-availibility page handler
func (m *Repository) Availability(w http.ResponseWriter, r *http.Request) {
	//Send the data to the template
	render.RenderTemplate(w, r, "search-availability.page.tmpl", &models.TemplateData{})

}

// Post Availability is the search-availibility page handler
func (m *Repository) PostAvailability(w http.ResponseWriter, r *http.Request) {

	start := r.Form.Get("start")
	end := r.Form.Get("end")
	w.Write([]byte(fmt.Sprintf("Start date: %s --- End date: %s", start, end)))

}

type jsonResponse struct {
	OK      bool   `json:"ok"`
	Message string `json:"message"`
}

// AvailabilityJson handles requests for availability and sends back Json response
func (m *Repository) AvailabilityJson(w http.ResponseWriter, r *http.Request) {
	resp := jsonResponse{
		OK:      true,
		Message: "Available!",
	}

	out, err := json.MarshalIndent(resp, "", "     ")
	if err != nil {
		log.Println(err)
	}

	log.Println(string(out))
	w.Header().Set("Content-Type", "applciation/json")
	w.Write(out)
}

// Contact is the contact page handler
func (m *Repository) Contact(w http.ResponseWriter, r *http.Request) {
	//Send the data to the template
	render.RenderTemplate(w, r, "contact.page.tmpl", &models.TemplateData{})

}

// Reservation is the make reservation page handler
func (m *Repository) Reservation(w http.ResponseWriter, r *http.Request) {
	//Send the data to the template
	render.RenderTemplate(w, r, "make-reservation.page.tmpl", &models.TemplateData{})

}
