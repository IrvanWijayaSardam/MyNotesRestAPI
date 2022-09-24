package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/IrvanWijayaSardam/GOData/driver"
	ph "github.com/IrvanWijayaSardam/GOData/handler/http"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func main() {
	dbName := "MyNotes"
	dbPass := "root"
	dbHost := "localhost"
	dbPort := "3308"

	connection, err := driver.KoneksiSQL(dbHost, dbPort, "root", dbPass, dbName)
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}

	r := chi.NewRouter()
	r.Use(middleware.Recoverer)
	r.Use(middleware.Logger)

	pHandler := ph.NewPostHandler(connection)
	uHandler := ph.NewUserHandler(connection)
	r.Route("/", func(rt chi.Router) {
		rt.Mount("/notes", postRouter(pHandler))
		rt.Mount("/user", userRouter(uHandler))
	})

	fmt.Println("Server Listen at : 8006")
	http.ListenAndServe(":8006", r)
}

func postRouter(pHandler *ph.Post) http.Handler {
	r := chi.NewRouter()
	r.Get("/", pHandler.Fetch)
	r.Get("/{id:[0-9]+}", pHandler.GetByID)
	r.Post("/", pHandler.Create)
	r.Put("/{id:[0-9]+}", pHandler.Update)
	r.Delete("/{id:[0-9]+}", pHandler.Delete)

	return r
}

func userRouter(uHandler *ph.PostUser) http.Handler {
	r := chi.NewRouter()
	r.Get("/", uHandler.FetchUser)
	r.Get("/{id:[0-9]+}", uHandler.GetUserByID)
	r.Post("/", uHandler.CreateUser)
	r.Put("/{id:[0-9]+}", uHandler.UpdateUser)
	r.Delete("/{id:[0-9]+}", uHandler.DeleteUser)

	return r
}
