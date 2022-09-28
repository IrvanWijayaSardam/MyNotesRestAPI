package main

import (
	"fmt"
	"net/http"
	"os"
	"time"

	jwt "github.com/dgrijalva/jwt-go"

	"github.com/IrvanWijayaSardam/GOData/driver"
	ph "github.com/IrvanWijayaSardam/GOData/handler/http"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

var mySigningKey = []byte("mysupersecretkey")

func generateJWT() (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)

	claims["authorized"] = true
	claims["user"] = "Irvan Wijaya"
	claims["exp"] = time.Now().Add(time.Minute * 30).Unix()

	tokenString, err := token.SignedString(mySigningKey)

	if err != nil {
		fmt.Errorf("Something went wrong : %s", err.Error())
		return "", err
	}

	return tokenString, nil

}

func main() {
	dbName := "MyNotes"
	dbPass := "root"
	dbHost := "localhost"
	dbPort := "3308"

	tokenString, err := generateJWT()
	if err != nil {
		fmt.Println("Error generating token")
	}
	fmt.Println(tokenString)

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
		rt.Mount("/auth", authRouter(uHandler))
		rt.Mount("/jwt", isAuthorised(getJWTRouter))
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

func authRouter(uHandler *ph.PostUser) http.Handler {
	r := chi.NewRouter()
	r.Get("/{email:}", uHandler.GetUserByEmail)
	return r
}

func getJWTRouter(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Ini rahasia")
}

func isAuthorised(endpoint func(http.ResponseWriter, *http.Request)) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Header["token"] != nil {
			token, err := jwt.Parse(r.Header["token"][0], func(token *jwt.Token) (interface{}, error) {
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, fmt.Errorf("There was an error")
				}
				return mySigningKey, nil
			})

			if err != nil {
				fmt.Fprintf(w, err.Error())
			}

			if token.Valid {
				endpoint(w, r)
			}
		} else {
			fmt.Fprintf(w, "Not Authorized")
			fmt.Fprintf(w, "Header", r.Header)
		}
	})
}
