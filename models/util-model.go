package models

type ResponseAPI[T any] struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    T      `json:"data"`
}

type Token struct {
	Token string `json:"token"`
}
