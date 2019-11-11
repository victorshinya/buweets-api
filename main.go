package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	_ "github.com/joho/godotenv/autoload"
	"github.com/victorshinya/buweets-api/handler"
)

func main() {
	router := mux.NewRouter()
	s := router.PathPrefix("/api").Subrouter()
	s.HandleFunc("/get-emotion", handler.GetEmotion).Methods(http.MethodGet)

	port := os.Getenv("PORT")
	fmt.Printf(`Server is up and running at port %s`, port)
	log.Fatal(http.ListenAndServe(":"+port, s))
}
