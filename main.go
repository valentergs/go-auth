package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type person struct {
	Nome string
}

func main() {
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

	http.HandleFunc("/encode", foo)
	http.HandleFunc("/decode", bar)
	http.ListenAndServe(":8080", nil)
}

func foo(w http.ResponseWriter, r *http.Request) {
	p1 := person{
		Nome: "Rodrigo",
	}

	err := json.NewEncoder(w).Encode(p1)
	if err != nil {
		log.Println("Encoded bad data")
	}
}

func bar(w http.ResponseWriter, r *http.Request) {
	var p1 person
	err := json.NewDecoder(r.Body).Decode(&p1)
	if err != nil {
		log.Println("Decoded bad data")
	}

	log.Println("Pessoa:", p1)
}
