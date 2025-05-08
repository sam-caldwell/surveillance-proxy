package ubiquity

type WebhookEvent struct {
	CameraID   string                 `json:"camera_id"`
	EventTime  string                 `json:"event_time"`
	ThumbURL   string                 `json:"thumbnail"`
	EventType  string                 `json:"event_type"`
	Additional map[string]interface{} `json:"-"`
}
