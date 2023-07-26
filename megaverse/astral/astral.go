package astral

import "github.com/google/uuid"

type Coordinates struct {
	Row    int `json:"row"`
	Column int `json:"column"`
}

type AstralObject interface {
	GetEndpoint() string
	GetPayload(candidateId uuid.UUID) interface{}
}
