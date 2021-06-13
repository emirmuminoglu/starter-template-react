package user

// Model is the model of user
type Model struct {
	ID           int    `json:"id" bson:"_id"`
	Username     string `json:"username" bson:"username"`
	PasswordHash string `json:"passwordHash" bson:"password_hash"`
}
