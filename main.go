package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"fmt"
	"io"
	"net/http"
)

func main() {
	http.HandleFunc("/", foo)
	http.HandleFunc("/submit", bar)
	http.ListenAndServe(":8080", nil)
}

func foo(w http.ResponseWriter, r *http.Request) {

	c, err := r.Cookie("session")
	if err != nil {
		c = &http.Cookie{}
	}

	html := `
	<!DOCTYPE html>
<html lang="en">
  <head>
    <meta charset="UTF-8" />
    <meta name="viewport" content="width=device-width, initial-scale=1.0" />
    <title>HAC Example</title>
  </head>
  <body>
    <h1>HMAC Example</h1>
    <form action="/submit" method="post">
      <input type="email" name="email" />
	  <input type="submit" />
	  <p>Cookie: ` + c.Value + `</p>
    </form>
  </body>
</html>
	`
	io.WriteString(w, html)
}

func bar(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	email := r.FormValue("email")
	if email == "" {
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}

	code := getCode(email)

	// hash + "stored value"
	c := http.Cookie{
		Name:  "session",
		Value: code + "|" + email,
	}

	http.SetCookie(w, &c)
	http.Redirect(w, r, "/", 303)
}

func getCode(msg string) string {
	h := hmac.New(sha256.New, []byte("somekey"))
	h.Write([]byte(msg))
	return fmt.Sprintf("%x", h.Sum(nil))
}
