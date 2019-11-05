package main

import (
	"fmt"
	"net/http"
	"os"

	_ "github.com/joho/godotenv/autoload"
	"github.com/victorshinya/buweets-api/handler"
)

func main() {
	http.HandleFunc("/api/get-emotion", handler.GetEmotion)

	port := os.Getenv("PORT")
	fmt.Printf(`Server is up and running at port %s`, port)
	http.ListenAndServe(":"+port, nil)
}
