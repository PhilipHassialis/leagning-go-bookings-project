package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/PhilipHassialis/leagning-go-bookings-project/pkg/config"
	"github.com/PhilipHassialis/leagning-go-bookings-project/pkg/handlers"
	"github.com/PhilipHassialis/leagning-go-bookings-project/pkg/render"
	"github.com/alexedwards/scs/v2"
)

const PORT_NUMBER = ":8080"

var app config.AppConfig // global visibility in all of main package
var session *scs.SessionManager

// main is the entry point for the program
func main() {

	app.InProduction = false // set it to true for production

	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
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

	// http.HandleFunc("/", handlers.Repo.Home)
	// http.HandleFunc("/about", handlers.Repo.About)
	// http.HandleFunc("/divide", handlers.Repo.Divide)
	fmt.Println("Server is listening on port", PORT_NUMBER)
	// http.ListenAndServe(PORT_NUMBER, nil)

	srv := &http.Server{
		Addr:    PORT_NUMBER,
		Handler: routes(&app),
	}

	err = srv.ListenAndServe()
	log.Fatal(err)
}
