package controller

import (
	"fmt"
	"net/http"

	"github.com/pt-arvind/gocleanarchitecture/domain"
)

// RegisterHandler represents the services required for this controller.
type RegisterHandler struct {
	UserService domain.UserInteractor
	ViewService domain.ViewCase
}

// Index displays the register screen.
func (h *RegisterHandler) Index(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		h.Store(w, r)
		return
	}

	h.ViewService.SetTemplate("register/index")
	h.ViewService.Render(w, r)
}

// Store adds a user to the database.
func (h *RegisterHandler) Store(w http.ResponseWriter, r *http.Request) {
	// Don't continue if required fields are missing.
	for _, v := range []string{"firstname", "lastname", "email", "password"} {
		if len(r.FormValue(v)) == 0 {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprint(w, `<html>One or more required fields are missing. `+
				`Click <a href="/register">here</a> to try again.</html>`)
			return
		}
	}

	// Build the user from the form values.
	u := new(domain.User)
	u.FirstName = r.FormValue("firstname")
	u.LastName = r.FormValue("lastname")
	u.Email = r.FormValue("email")
	u.Password = r.FormValue("password")

	// Add the user to the database.
	err := h.UserService.CreateUser(u)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, err)
		return
	}

	w.WriteHeader(http.StatusCreated)
	fmt.Fprint(w, `<html>User created. `+
		`Click <a href="/">here</a> to login.</html>`)
}
