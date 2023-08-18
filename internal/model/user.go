package model

type User struct {
	Id int `json:"id"`
	UserDto
}

type UserDto struct {
	Login    string `json:"login"`
	Password string `json:"password"`
	Name     string `json:"name"`
	Age      int    `json:"age"`
}

type UserSignIn struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

type UserPayload struct {
	UserId int    `json:"userId"`
	Login  string `json:"login"`
}
