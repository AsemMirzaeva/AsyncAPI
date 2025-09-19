package apiserver

import (
	"encoding/json"
	"errors"
	"net/http"
)

type SignupReequest struct {
	Email string `json:"email"`
	Password string `json:"password"`
}

func (r SignupReequest) Validate() error{
	if r.Email == "" {
		return errors.New("email is required")
	}
	if r.Password == "" {
		return errors.New("password is required")
	}
	return  nil
}

func (s *ApiServer) signupHandler(w http.ResponseWriter, r *http.Request) {
	var req SignupReequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	if err := req.Validate(); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}