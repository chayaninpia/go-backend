package index

import _ "github.com/go-playground/validator/v10"

type IndexI struct {
	Name     string   `json:"name" validate:"required"`
	LastName []string `json:"lastName" validate:"required"`
}

type IndexO struct {
	Name     string `json:"name"`
	LastName string `json:"lastName"`
}
