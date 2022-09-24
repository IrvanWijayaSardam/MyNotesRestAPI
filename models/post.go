package models

// Post type details
type Post struct {
	ID          int64  `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Date        string `json:"date"`
	UserID      string `json:"userid"`
	// created_at time.Time `json:"created_at"`
	// updated_at time.Time `json:"updated_at"`
}
