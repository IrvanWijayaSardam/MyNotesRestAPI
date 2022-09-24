package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/IrvanWijayaSardam/GOData/driver"
	models "github.com/IrvanWijayaSardam/GOData/models"
	repository "github.com/IrvanWijayaSardam/GOData/repository"
	post "github.com/IrvanWijayaSardam/GOData/repository/post"
	postuser "github.com/IrvanWijayaSardam/GOData/repository/user"
	"github.com/go-chi/chi"
)

type HandlerClient interface {
	Fetch(w http.ResponseWriter, r *http.Request)
	GetByID(w http.ResponseWriter, r *http.Request)
	Create(w http.ResponseWriter, r *http.Request)
	Update(w http.ResponseWriter, r *http.Request)
	Delete(w http.ResponseWriter, r *http.Request)
}

func NewPostHandler(db *driver.DB) *Post {
	return &Post{
		repo: post.NewSQLPostRepo(db.SQL),
	}
}

func NewUserHandler(db *driver.DB) *PostUser {
	return &PostUser{
		repo: postuser.NewSQLUserRepo(db.SQL),
	}
}

type Post struct {
	repo repository.PostRepo
}

type PostUser struct {
	repo repository.UserRepo
}

func (p *Post) Fetch(w http.ResponseWriter, r *http.Request) {
	payload, _ := p.repo.Fetch(r.Context(), 5)

	respondwithJSON(w, http.StatusOK, payload)
}

func (p *PostUser) FetchUser(w http.ResponseWriter, r *http.Request) {
	payload, _ := p.repo.Fetch(r.Context(), 5)

	respondwithJSON(w, http.StatusOK, payload)
}

// Create a new notes
func (p *Post) Create(w http.ResponseWriter, r *http.Request) {
	post := models.Post{}
	json.NewDecoder(r.Body).Decode(&post)

	newID, err := p.repo.Create(r.Context(), &post)
	fmt.Println(newID)
	if err != nil {
		fmt.Println(err)
		respondWithError(w, http.StatusInternalServerError, "Server Error")
		return
	}

	respondwithJSON(w, http.StatusCreated, map[string]string{"message": "Successfully Created"})
}

// Create a new User
func (p *PostUser) CreateUser(w http.ResponseWriter, r *http.Request) {
	postUser := models.User{}
	json.NewDecoder(r.Body).Decode(&postUser)

	newID, err := p.repo.Create(r.Context(), &postUser)
	fmt.Println(newID)
	if err != nil {
		fmt.Println(err)
		respondWithError(w, http.StatusInternalServerError, "Server Error")
		return
	}

	respondwithJSON(w, http.StatusCreated, map[string]string{"message": "Successfully Created"})
}

// Update a post by id
func (p *Post) Update(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))
	data := models.Post{ID: int64(id)}
	json.NewDecoder(r.Body).Decode(&data)
	payload, err := p.repo.Update(r.Context(), &data)

	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
	}

	respondwithJSON(w, http.StatusOK, payload)
}

// Update user by id
func (p *PostUser) UpdateUser(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))
	data := models.User{UserID: int64(id)}
	json.NewDecoder(r.Body).Decode(&data)
	payload, err := p.repo.Update(r.Context(), &data)

	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
	}

	respondwithJSON(w, http.StatusOK, payload)
}

// GetByID returns a post details
func (p *Post) GetByID(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))
	payload, err := p.repo.GetByID(r.Context(), int64(id))

	if err != nil {
		respondWithError(w, http.StatusNoContent, "Content not found")
	}

	respondwithJSON(w, http.StatusOK, payload)
}

// GetByID returns a post details
func (p *PostUser) GetUserByID(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))
	payload, err := p.repo.GetByID(r.Context(), int64(id))

	if err != nil {
		respondWithError(w, http.StatusNoContent, "Content not found")
	}

	respondwithJSON(w, http.StatusOK, payload)
}

// Delete a post
func (p *Post) Delete(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))
	_, err := p.repo.Delete(r.Context(), int64(id))

	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Server Error")
	}

	respondwithJSON(w, http.StatusMovedPermanently, map[string]string{"message": "Delete Successfully"})
}

// Delete a post
func (p *PostUser) DeleteUser(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))
	_, err := p.repo.Delete(r.Context(), int64(id))

	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Server Error")
	}

	respondwithJSON(w, http.StatusMovedPermanently, map[string]string{"message": "Delete Successfully"})
}

// respondwithJSON write json response format
func respondwithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

// respondwithError return error message
func respondWithError(w http.ResponseWriter, code int, msg string) {
	respondwithJSON(w, code, map[string]string{"message": msg})
}
