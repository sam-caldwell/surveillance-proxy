package ubiquity

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/sam-caldwell/surveillance-proxy/v2/jira"
	"io"
	"log"
	"net/http"
	"time"
)

// handleEvent - Handle a webhook event and create a jira ticket.
// ToDo: decouple this from the JIRA ticket so we can also allow other targets (e.g. slack)
func handleEvent(eventID string, event WebhookEvent, jiraUser, jiraToken, jiraBaseURL, jiraProject *string) {
	log.Printf("[EVENT %s] Processing event from camera %s", eventID, event.CameraID)

	// Download image
	resp, err := http.Get(event.ThumbURL)
	if err != nil {
		log.Printf("[EVENT %s] Failed to download thumbnail: %v", eventID, err)
		return
	}
	defer func() { _ = resp.Body.Close() }()

	var imgData []byte
	if imgData, err = io.ReadAll(resp.Body); err != nil {
		log.Printf("[EVENT %s] Failed to read image: %v", eventID, err)
		return
	}

	// Create Jira ticket
	issue := jira.Issue{
		Fields: jira.Fields{
			Project: jira.Project{Key: *jiraProject},
			Summary: fmt.Sprintf("Alert: %s from camera %s", event.EventType, event.CameraID),
			Description: fmt.Sprintf("Event occurred at %s\n"+
				"Event type: %s\n"+
				"Camera: %s",
				event.EventTime, event.EventType, event.CameraID),
			Issuetype: jira.Type{Name: "Task"},
		},
	}

	issueBody, _ := json.Marshal(issue)
	req, _ := http.NewRequest("POST", fmt.Sprintf("%s/rest/api/2/issue", jiraBaseURL), bytes.NewBuffer(issueBody))
	req.SetBasicAuth(*jiraUser, *jiraToken)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{Timeout: 10 * time.Second}
	resp, err = client.Do(req)
	if err != nil || resp.StatusCode >= 300 {
		log.Printf("[EVENT %s] Failed to create Jira issue: %v", eventID, err)
		return
	}
	defer func() { _ = resp.Body.Close() }()

	var res struct {
		Key string `json:"key"`
	}

	if err = json.NewDecoder(resp.Body).Decode(&res); err != nil {
		return
	}
	issueKey := res.Key

	// Upload image
	var b bytes.Buffer
	w := io.MultiWriter(&b)
	_, _ = w.Write([]byte("--boundary\r\n"))
	_, _ = w.Write([]byte("Content-Disposition: form-data; name=\"file\"; filename=\"thumbnail.jpg\"\r\n"))
	_, _ = w.Write([]byte("Content-Type: image/jpeg\r\n\r\n"))
	_, _ = w.Write(imgData)
	_, _ = w.Write([]byte("\r\n--boundary--\r\n"))

	uploadReq, _ := http.NewRequest(
		"POST",
		fmt.Sprintf("%s/rest/api/2/issue/%s/attachments", *jiraBaseURL, issueKey), &b)
	uploadReq.SetBasicAuth(*jiraUser, *jiraToken)
	uploadReq.Header.Set("X-Atlassian-Token", "no-check")
	uploadReq.Header.Set("Content-Type", "multipart/form-data; boundary=boundary")

	if resp, err = client.Do(uploadReq); err != nil || resp.StatusCode >= 300 {
		log.Printf("[EVENT %s] Failed to upload image: %v", eventID, err)
		return
	}

	log.Printf("[EVENT %s] Ticket %s created and image uploaded", eventID, issueKey)
}
