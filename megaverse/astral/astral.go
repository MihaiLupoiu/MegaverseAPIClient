package astral

import "github.com/google/uuid"

const (
	POLYANET = "POLYANET"
	COMETH   = "COMETH"
	SOLOON   = "SOLOON"
	SPACE    = "SPACE"
)

type Coordinates struct {
	Row    int `json:"row"`
	Column int `json:"column"`
}

type AstralObject interface {
	GetEndpoint() string
	GetPayload(candidateId uuid.UUID) interface{}
}
