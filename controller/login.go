package controller

import (
	"fmt"
	"net/http"

	"github.com/pt-arvind/gocleanarchitecture/domain"
)

// LoginHandler represents the services required for this controller.
type LoginHandler struct {
	UserService domain.UserCase
	ViewService domain.ViewCase
}

// Index displays the logon screen.
func (h *LoginHandler) Index(w http.ResponseWriter, r *http.Request) {
	// Handle 404.
	if r.URL.Path != "/" {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprint(w, "404 Page Not Found")
		return
	}

	if r.Method == "POST" {
		h.Store(w, r)
		return
	}

	h.ViewService.SetTemplate("login/index")
	h.ViewService.Render(w, r)
}

// Store handles the submission of the login information.
func (h *LoginHandler) Store(w http.ResponseWriter, r *http.Request) {
	// Don't continue if required fields are missing.
	for _, v := range []string{"email", "password"} {
		if len(r.FormValue(v)) == 0 {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprint(w, `<html>One or more required fields are missing. `+
				`Click <a href="/">here</a> to try again.</html>`)
			return
		}
	}

	u := new(domain.User)
	u.Email = r.FormValue("email")
	u.Password = r.FormValue("password")

	err := h.UserService.Authenticate(u)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Fprint(w, `<html>Login failed. `+
			`Click <a href="/">here</a> to try again.</html>`)
		return
	}

	fmt.Fprint(w, "<html>Login successful!</html>")
}
