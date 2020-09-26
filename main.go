package main

import (
	"fmt"
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

var db = map[string][]byte{}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/register", register)
	http.HandleFunc("/login", login)
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, r *http.Request) {

	html := `
	<!DOCTYPE html>
<html lang="en">
  <head>
    <meta charset="UTF-8" />
    <meta name="viewport" content="width=device-width, initial-scale=1.0" />
    <title>JWT Example</title>
  </head>
  <body>
    <h1>JWT Example</h1>
    <h4>Login page</h4>
    <form action="/login" method="POST">
      <input type="email" name="email" />
      <input type="password" name="password" />
      <input type="submit" />
    </form>

    <p>Not registered yet? <a href="/">Click here</a></p>
  </body>
</html>
	`
	fmt.Fprint(w, html)
}

func register(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		http.Redirect(w, r, "/", 303)
		fmt.Println("Wrong Request method!")
	}

	e := r.FormValue("email")
	p := r.FormValue("password")

	senhaCripto, err := bcrypt.GenerateFromPassword([]byte(p), bcrypt.DefaultCost)
	if err != nil {
		http.Error(w, "Erro de servidor", http.StatusInternalServerError)
	}
	db[p] = senhaCripto
	fmt.Println(string(senhaCripto), e)

	http.Redirect(w, r, "/", 303)
}

func login(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		http.Redirect(w, r, "/", 303)
		fmt.Println("Wrong Request method!")
	}

	e := r.FormValue("email")
	p := r.FormValue("password")

	senhaCripto, err := bcrypt.GenerateFromPassword([]byte(p), bcrypt.DefaultCost)
	if err != nil {
		http.Error(w, "Erro de servidor", http.StatusInternalServerError)
	}
	db[p] = senhaCripto
	fmt.Println(string(senhaCripto), e)

	http.Redirect(w, r, "/", 303)
}
