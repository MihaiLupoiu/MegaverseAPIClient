package astral

import "github.com/google/uuid"

const defaultComethEndpoint = "/api/comeths"

type DirectionType string

const (
	Cometh_Up    DirectionType = "up"
	Cometh_Down  DirectionType = "down"
	Cometh_Left  DirectionType = "left"
	Cometh_Right DirectionType = "right"
)

type Cometh struct {
	Coordinates
	Direction DirectionType
}

func NewCometh(row, column int, direction DirectionType) *Cometh {
	return &Cometh{
		Coordinates: Coordinates{
			Row:    row,
			Column: column,
		},
		Direction: direction,
	}
}

func (c Cometh) GetEndpoint() string {
	return defaultComethEndpoint
}

type comethPayload struct {
	Row         int           `json:"row"`
	Column      int           `json:"column"`
	CandidateId string        `json:"candidateId"`
	Direction   DirectionType `json:"direction"`
}

func (c Cometh) GetPayload(candidateId uuid.UUID) interface{} {
	return comethPayload{
		Row:         c.Row,
		Column:      c.Column,
		Direction:   c.Direction,
		CandidateId: candidateId.String(),
	}
}
