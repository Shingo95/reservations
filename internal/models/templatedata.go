package models

import "github.com/elorenzotti/bookings/internal/forms"

// TemplateData is the type of datas we are going to pass to the template

type TemplateData struct {
	StringMap map[string]string
	IntMap    map[string]int
	FloatMap  map[string]float32
	Data      map[string]interface{}
	CSRFToken string
	Flash     string
	Warning   string
	Error     string
	Form      *forms.Form
}
