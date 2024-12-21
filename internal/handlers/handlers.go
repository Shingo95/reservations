package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/elorenzotti/bookings/helpers"
	"github.com/elorenzotti/bookings/internal/config"
	"github.com/elorenzotti/bookings/internal/driver"
	"github.com/elorenzotti/bookings/internal/forms"
	"github.com/elorenzotti/bookings/internal/models"
	"github.com/elorenzotti/bookings/internal/render"
	"github.com/elorenzotti/bookings/internal/repository"
	"github.com/elorenzotti/bookings/internal/repository/dbrepo"
)

// Repo is the variable for the repository
var Repo *Repository

// Repository is the repository type
type Repository struct {
	App *config.AppConfig
	DB  repository.DatabaseRepo
}

// NewRepo creates a new repository

func NewRepo(a *config.AppConfig, db *driver.DB) *Repository {
	return &Repository{
		App: a,
		DB:  dbrepo.NewPostgresRepo(db.SQL, a),
	}
}

// NewHandlers sets the repository for the handlers
func NewHandlers(r *Repository) {
	Repo = r
}

// Home is the homepage handler
func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {

	render.RenderTemplate(w, r, "home.page.tmpl", &models.TemplateData{})

}

// About is the about page handler
func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	//Send the data to the template
	render.RenderTemplate(w, r, "about.page.tmpl", &models.TemplateData{})

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
		helpers.ServerError(w, err)
		return
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

	var emptyReservation models.Reservation
	data := make(map[string]interface{})
	data["reservation"] = emptyReservation

	//Send the data to the template
	render.RenderTemplate(w, r, "make-reservation.page.tmpl", &models.TemplateData{
		Form: forms.New(nil),
		Data: data,
	})

}

// PostReservation handles the posting of a reservation form
func (m *Repository) PostReservation(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()

	if err != nil {
		helpers.ServerError(w, err)
		return
	}
	reservation := models.Reservation{
		FirstName: r.Form.Get("first-name"),
		LastName:  r.Form.Get("last-name"),
		Email:     r.Form.Get("email"),
		Phone:     r.Form.Get("phone"),
	}
	form := forms.New(r.PostForm)

	form.Required("first-name", "last-name", "email")
	form.MinLength("first-name", 3)
	form.IsEmail("email")

	if !form.Valid() {
		data := make(map[string]interface{})
		data["reservation"] = reservation
		render.RenderTemplate(w, r, "make-reservation.page.tmpl", &models.TemplateData{
			Form: form,
			Data: data,
		})
		return
	}
	m.App.Session.Put(r.Context(), "reservation", reservation)
	http.Redirect(w, r, "/reservation-summary", http.StatusSeeOther)
}

func (m *Repository) ReservationSummary(w http.ResponseWriter, r *http.Request) {
	reservation, ok := m.App.Session.Get(r.Context(), "reservation").(models.Reservation)
	if !ok {
		m.App.ErrorLog.Println("Cannot get item from session")
		m.App.Session.Put(r.Context(), "error", "Can't get information for this session")
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}

	m.App.Session.Remove(r.Context(), "reservation")
	data := make(map[string]interface{})
	data["reservation"] = reservation

	render.RenderTemplate(w, r, "reservation-summary.page.tmpl", &models.TemplateData{
		Data: data,
	})
}
