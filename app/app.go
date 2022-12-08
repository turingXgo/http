package app

import (
	"encoding/json"
	"net/http"

	"github.com/turingXgo/http/model"
)

type Repository interface {
	GetUser(name string) (*model.User, error)
	InsertUser(*model.User) error
}

type Application struct {
	repository Repository
}

func (a Application) CreateUser(w http.ResponseWriter, r *http.Request) {
	var data model.User
	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err := a.repository.InsertUser(&data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write([]byte("Success"))
}

type GetUser struct {
	Name string `json:"name"`
}

func (a Application) GetUserByName(w http.ResponseWriter, r *http.Request) {
	var request GetUser
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	user, err := a.repository.GetUser(request.Name)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(user)
}

func NewApplication(repository Repository) *Application {
	return &Application{
		repository: repository,
	}
}
