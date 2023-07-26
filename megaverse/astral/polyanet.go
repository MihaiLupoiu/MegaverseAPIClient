package astral

import "github.com/google/uuid"

const defaultPolyanetEndpoint = "/api/polyanets"

type Polyanet struct {
	Coordinates
}

func NewPolyanet(row, column int) *Polyanet {
	return &Polyanet{
		Coordinates: Coordinates{
			Row:    row,
			Column: column,
		},
	}
}

func (p Polyanet) GetEndpoint() string {
	return defaultPolyanetEndpoint
}

type polyanetPayload struct {
	Row         int    `json:"row"`
	Column      int    `json:"column"`
	CandidateId string `json:"candidateId"`
}

func (p Polyanet) GetPayload(candidateId uuid.UUID) interface{} {
	return polyanetPayload{
		Row:         p.Row,
		Column:      p.Column,
		CandidateId: candidateId.String(),
	}
}
