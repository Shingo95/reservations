package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/alexedwards/scs/v2"
	"github.com/elorenzotti/bookings/pkg/config"
	"github.com/elorenzotti/bookings/pkg/handlers"
	"github.com/elorenzotti/bookings/pkg/render"
)

const portNumber = ":8080"
var app config.AppConfig
var session *scs.SessionManager


// main is the main application
func main() {


	// Change to true when in Production
	app.Inproduction = false

	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.Inproduction

	app.Session = session

	tc, err := render.CreateTemplateCache()
	if err != nil{
		log.Fatal(err)
	}

	app.TemplateCache = tc
	app.UseCache = false
	
	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	render.NewTemplates(&app)

	// http.HandleFunc("/", handlers.Repo.Home)
	// http.HandleFunc("/about", handlers.Repo.About)

	fmt.Println(fmt.Sprintf("Starting application on port %s", portNumber))
	// _ = http.ListenAndServe(portNumber, nil)

	srv := &http.Server {
		Addr: portNumber,
		Handler: routes(&app),
	}
	err = srv.ListenAndServe()
	
	log.Println(err)
	log.Fatal(err)

}
