package main

import (
	"testing"

	"github.com/elorenzotti/bookings/internal/config"
	"github.com/go-chi/chi/v5"
)

func TestRoutes(t *testing.T) {
	var app config.AppConfig
	mux := routes(&app)

	switch v := mux.(type) {
	case *chi.Mux:
		//do Nothing; test passed
	default:
		t.Errorf("Type is not valid= Expected mux receive %T", v)
	}
}
