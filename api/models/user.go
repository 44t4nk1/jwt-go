package models

//User ...
type User struct {
	ID       uint64 `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

//Response ...
type Response struct {
	Error   bool   `json:"error"`
	Message string `json:"message"`
}
