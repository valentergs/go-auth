package main

import (
	"crypto/hmac"
	"crypto/sha512"
	"fmt"
	"log"

	"golang.org/x/crypto/bcrypt"
)

type person struct {
	Nome string
}

var key = []byte{}

func main() {
	for i := 1; i <= 64; i++ {
		key = append(key, byte(i))
	}
	// p1 := person{
	// 	Nome: "Rodrigo",
	// }

	// p2 := person{
	// 	Nome: "Vanessa",
	// }

	// p3 := person{
	// 	Nome: "Gustavo",
	// }

	// p4 := person{
	// 	Nome: "Eduardo",
	// }

	// xp := []person{p1, p2, p3, p4}

	// bs, err := json.Marshal(xp)
	// if err != nil {
	// 	log.Panic(err)
	// }

	// fmt.Println(string(bs))

	// xp2 := []person{}

	// err = json.Unmarshal(bs, &xp2)
	// if err != nil {
	// 	log.Panic(err)
	// }

	// fmt.Println(xp2)

	// http.HandleFunc("/encode", foo)
	// http.HandleFunc("/decode", bar)
	// http.ListenAndServe(":8080", nil)

	pass := "Rodrigo Valente"

	hashedPass, err := hashPassword(pass)
	if err != nil {
		panic(err)
	}

	err = comparePassword(pass, hashedPass)
	if err != nil {
		log.Fatalln("Not allowed...")
	}

	log.Println(string(hashedPass))
	log.Println(("Logged in!"))
	log.Println(signMessage(hashedPass))
}

// func foo(w http.ResponseWriter, r *http.Request) {
// 	p1 := person{
// 		Nome: "Rodrigo",
// 	}

// 	err := json.NewEncoder(w).Encode(p1)
// 	if err != nil {
// 		log.Println("Encoded bad data")
// 	}
// }

// func bar(w http.ResponseWriter, r *http.Request) {
// 	var p1 person
// 	err := json.NewDecoder(r.Body).Decode(&p1)
// 	if err != nil {
// 		log.Println("Decoded bad data")
// 	}

// 	log.Println("Pessoa:", p1)
// }

func hashPassword(password string) ([]byte, error) {
	bs, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, fmt.Errorf("Error while generating bcrypt: %w", err)
	}
	return bs, nil
}

func comparePassword(password string, hashedPass []byte) error {
	err := bcrypt.CompareHashAndPassword(hashedPass, []byte(password))
	if err != nil {
		return fmt.Errorf("Invalid password: %w", err)
	}
	return nil
}

func signMessage(msg []byte) ([]byte, error) {
	h := hmac.New(sha512.New, key)
	_, err := h.Write(msg)
	if err != nil {
		return nil, fmt.Errorf("Error while hashing singMessage: %w", err)
	}

	signature := h.Sum(nil)
	return signature, nil
}

func checkSig(msg, sig []byte) (bool, error) {
	newSig, err := signMessage(msg)
	if err != nil {
		return false, fmt.Errorf("Error in checkSig while getting signature of message: %w", err)

	}
	same := hmac.Equal(newSig, sig)
	return same, nil
}
