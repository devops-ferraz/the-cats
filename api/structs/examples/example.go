package examples

import "github.com/google/uuid"

type Example struct {
	ID    string `json:id`
	Token string `json:"token"`
	Xpto  string `json:"xpto"`
}

func NewExample() *Example {
	example := Example{
		ID: uuid.New().String(),
	}

	return &example
}
