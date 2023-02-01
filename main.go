package main

import (
	"fmt"
	"log"
	"net/http"
	"renovatedummy/maths"
	"renovatedummy/server"

	"github.com/rs/cors"
	"github.com/sanity-io/litter"
	"go.uber.org/zap"
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

	logger := zap.NewExample()

	logger.Info("Just an info log",
		zap.String("url", "http://example.com"),
	)

	c := cors.New(cors.Options{
		AllowedOrigins: []string{"http://foo.com"},
	})

	http.HandleFunc("/double", server.DoubleHandler)
	handler := http.HandlerFunc(server.DoubleHandler)
	err := http.ListenAndServe(":8080", c.Handler(handler))
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("The server is listening on port 8080")
}
