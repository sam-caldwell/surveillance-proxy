package ubiquity

import (
	"encoding/json"
	"github.com/google/uuid"
	"net/http"
)

type Handler func(w http.ResponseWriter, r *http.Request)

// WebhookHandlerFactory - Return a webhook handler function
func WebhookHandlerFactory(authToken, jiraUser, jiraToken, jiraBaseURL, jiraProject *string) Handler {
	// handler - handle webhook events.
	return func(w http.ResponseWriter, r *http.Request) {

		//Authenticate using a bearer token
		auth := r.Header.Get("Authorization")
		if auth != "Bearer "+(*authToken) {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		if r.Method != http.MethodPost {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		var event WebhookEvent
		if err := json.NewDecoder(r.Body).Decode(&event); err != nil {
			http.Error(w, "Bad request", http.StatusBadRequest)
			return
		}

		eventID := uuid.New().String()
		// asynchronously handle the event.
		go handleEvent(eventID, event, jiraUser, jiraToken, jiraBaseURL, jiraProject)
		w.WriteHeader(http.StatusAccepted)
	}
}
