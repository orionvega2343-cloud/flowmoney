package handlers

import (
	"encoding/json"
	"flowmoney/api/internal/config"
	"flowmoney/api/internal/models"
	"flowmoney/api/internal/service"
	"net/http"
	"strconv"
)

type UserHandler interface {
	CreateUser(w http.ResponseWriter, r *http.Request)
	Login(w http.ResponseWriter, r *http.Request)
	GetUserById(w http.ResponseWriter, r *http.Request)
	UpdateBalance(w http.ResponseWriter, r *http.Request)
}

type UserHandlerImpl struct {
	Us  service.UserService
	Cfg config.Jwt
}

func NewUserHandler(us service.UserService, cfg config.Jwt) *UserHandlerImpl {
	return &UserHandlerImpl{Us: us, Cfg: cfg}
}

func (u *UserHandlerImpl) CreateUser(w http.ResponseWriter, r *http.Request) {
	var userData models.User
	err := json.NewDecoder(r.Body).Decode(&userData)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	user, err := u.Us.CreateUser(userData)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	err = json.NewEncoder(w).Encode(user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

}

func (u *UserHandlerImpl) Login(w http.ResponseWriter, r *http.Request) {
	var ud models.User
	err := json.NewDecoder(r.Body).Decode(&ud)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	user, err := u.Us.Login(ud.Email, ud.Password, u.Cfg.JWTSecret)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(user)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (u *UserHandlerImpl) GetUserById(w http.ResponseWriter, r *http.Request) {
	url := r.PathValue("id")
	parsed, err := strconv.Atoi(url)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	user, err := u.Us.GetUserById(parsed)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	err = json.NewEncoder(w).Encode(user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

}

func (u *UserHandlerImpl) UpdateBalance(w http.ResponseWriter, r *http.Request) {
	var ud models.User
	url := r.PathValue("id")

	parsed, err := strconv.Atoi(url)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = json.NewDecoder(r.Body).Decode(&ud)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	user, err := u.Us.UpdateBalance(parsed, ud.Balance)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	err = json.NewEncoder(w).Encode(user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
