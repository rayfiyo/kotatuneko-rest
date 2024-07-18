package entity

type User struct {
	Id       string `json:"id"`       // user_id (UUID)
	Name     string `json:"name"`     // user_name (and display_name)
	Password string `json:"password"` // user_password (hashed)
	Score    int64  `json:"score"`    // game_score
}
