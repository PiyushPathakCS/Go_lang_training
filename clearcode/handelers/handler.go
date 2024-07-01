package handle

import (
	"clear_code/usecase"
	"encoding/json"
	"net/http"
)

type val struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Password string `json:"password"`
}

type Handleit interface {
	Sign(w http.ResponseWriter, r *http.Request)
}

type handler struct {
	user usecase.Servce
}

func NewHandler(user usecase.Servce) *handler {
	return &handler{
		user: user,
	}
}

// type u_case interface {
// }

func (handler *handler) Sign(w http.ResponseWriter, r *http.Request) {
	var d usecase.Val
	w.Header().Set("Content-Type", "application/json")
	var v val
	json.NewDecoder(r.Body).Decode(&v)

	d.Id = v.Id
	d.Name = v.Name
	d.Password = v.Password

	err := handler.user.Signup(&d)

	if err != nil {
		json.NewEncoder(w).Encode("Error in Signup")
	} else {
		json.NewEncoder(w).Encode("Success")
	}

}

func Signin(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	

}
