package main

import (
	"fmt"
	"log"
	"net/http"
	"renovatedummy/maths"
	"renovatedummy/server"

	"github.com/sanity-io/litter"
)

func main() {
	sum := maths.Add(5, 5)
	fmt.Printf("5 + 5: %d\n", sum)

	type Person struct {
		Name   string
		Age    int
		Parent *Person
	}

	litter.Dump(Person{
		Name: "Bob",
		Age:  20,
		Parent: &Person{
			Name: "Jane",
			Age:  50,
		},
	})

	http.HandleFunc("/double", server.DoubleHandler)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("The server is listening on port 8080")
}
