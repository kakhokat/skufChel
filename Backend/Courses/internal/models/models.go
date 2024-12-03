package models

type SignUpRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Username string `json:"username"`
	Birthday string `json:"birthday"`
	Photo    []byte `json:"photo"`
}

type User struct {
	Id        int    `json:"userId"`
	Email     string `json:"email"`
	Name      string `json:"name"`
	Birthday  string `json:"birthday"`
	IsCreator bool   `json:"isCreator"`
}

type Error struct {
	Error string `json:"error"`
}

type CheckMessage struct {
	CheckKey string `json:"checkKey"`
	Mail     string `json:"mail"`
}
