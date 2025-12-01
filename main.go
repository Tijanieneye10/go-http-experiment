package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	 "github.com/joho/godotenv"
)

type Author struct {
	Name string
	Age  int
}

func main() {

	godotenv.Load()

	PORT := os.Getenv("PORT")

	fmt.Println(PORT)

	author := []Author{
		{
			Name: "Tijani Usman",
			Age:  25,
		},
		{
			Name: "John Doe",
			Age:  29,
		},
	}

	mux := http.NewServeMux()

	mux.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {

		err := json.NewEncoder(w).Encode(author)

		w.Header().Set("Content-Type", "application/json")
		
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	})

	http.ListenAndServe(fmt.Sprintf(":%s", PORT), mux)
}
