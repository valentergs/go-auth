package main

import (
	"crypto/sha256"
	"fmt"
)

func main() {

	// FROM FILE
	// f, err := os.Open("somefile.txt")
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// defer f.Close()

	// h := sha256.New()
	// if _, err := io.Copy(h, f); err != nil {
	// 	log.Fatal(err)
	// }

	// //fmt.Printf("%T\", f)
	// fmt.Printf("%T\n", h.Sum(nil))
	// //fmt.Println(len(h.Sum(nil)))

	// FROM STRING
	h := sha256.New()
	h.Write([]byte("RodrigoValente"))
	fmt.Printf("%x\n", h.Sum(nil))
}
