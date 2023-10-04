package main

import (
	"log"
	"net/http"
	"time"

	"github.com/alexedwards/scs/v2"
	"github.com/ausates/bookings/pkg/config"
	"github.com/ausates/bookings/pkg/handlers"
	"github.com/ausates/bookings/pkg/render"
)

var app config.AppConfig
var session *scs.SessionManager

const portNumber = ":8080"

func main() {

	// change value for if prod/local testing
	app.InProduction = false

	session = scs.New()
	// determines how long the cookie is valid:
	session.Lifetime = 24 * time.Hour
	// determines if the cookie should persist after the browser is closed:
	session.Cookie.Persist = true
	// sets the strictness of the cookie
	session.Cookie.SameSite = http.SameSiteLaxMode
	// cookie encrypted t/f (https)
	session.Cookie.Secure = app.InProduction

	app.Session = session

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("cannot create template cache")
	}

	app.TemplateCache = tc
	app.UseCache = false
	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	render.NewTemplates(&app)

	srv := &http.Server{
		Addr:    portNumber,
		Handler: routes(&app),
	}

	_ = srv.ListenAndServe()
	// http.HandleFunc("/", handlers.Repo.Home)
	// http.HandleFunc("/about", handlers.Repo.About)

}
