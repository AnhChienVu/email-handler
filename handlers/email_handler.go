package handlers

import (
	"email-handler/services"
	"encoding/json"
	"fmt"
	"net/http"
)

func SendEmailHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	// json:"..." tags are used to map the struct fields to JSON keys during encoding and decoding the request body.
	var request struct {
		Provider string `json:"provider"`
		To       string `json:"to"`
		Subject  string `json:"subject"`
		Body     string `json:"body"`
	}

	// Try to read a JSON object from the request body into the struct.
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&request); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	emailService, err := services.EmailServiceFactory(request.Provider)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error: %v", err), http.StatusBadRequest)
		return
	}

	err = emailService.SendEmail(request.To, request.Subject, request.Body)
	if err != nil {
		http.Error(w, "Failed to send email", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Email sent successfully"})
}
