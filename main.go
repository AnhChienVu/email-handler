package main

import (
	"email-handler/handlers"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Initialize the email handler with environment variables
	fmt.Println("SENDGRID_API_KEY:", os.Getenv("SENDGRID_API_KEY")) // Example usage of env variable

	http.HandleFunc("/email/send", handlers.SendEmailHandler)
	// http.HandleFunc("emails", handlers.GetEmails)

	fmt.Println("Server is running on port 8080")
	http.ListenAndServe(":8080", nil)
}
