package handler

import (
	"FileStore-Server/service"
	"io/ioutil"
	"net/http"
)

func Sginup(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		data, err := ioutil.ReadFile("./static/html/signup.html")
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.Write(data)
		return
		//http.Redirect(w, r, "/static/html/signup.html", http.StatusFound)
		//return
	}
	err := r.ParseForm()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Sign up failed"))
	}
	username := r.Form.Get("username")
	password := r.Form.Get("password")

	if len(username) < 4 || len(password) < 5 {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Username or Password format error"))
		return
	}
	err = service.SignUp(username, password)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	w.Write([]byte("sign up success"))
}
