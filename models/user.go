package models

// Post type details
type User struct {
	UserID   int64  `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Profile  string `json:"profile"`
	Jk       string `json:"jk"`
	// created_at time.Time `json:"created_at"`
	// updated_at time.Time `json:"updated_at"`
}
