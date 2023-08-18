package model

type Phone struct {
	Id int `json:"id"`
	PhoneDto
}

type PhoneDto struct {
	Phone       string `json:"phone"`
	IsFax       bool   `json:"isFax"`
	Description string `json:"description"`
	UserId      int    `json:"userId"`
}

type PhoneDeleteDto struct {
	Id     int `json:"id"`
	UserId int `json:"userId"`
}
