package astral

import "github.com/google/uuid"

const defaultSoloonEndpoint = "/api/soloons"

type ColorType string

const (
	Soloon_Blue   ColorType = "blue"
	Soloon_Red    ColorType = "red"
	Soloon_Purple ColorType = "purple"
	Soloon_White  ColorType = "white"
)

type Soloon struct {
	Coordinates
	Color ColorType
}

func NewSoloon(row, column int, color ColorType) *Soloon {
	return &Soloon{
		Coordinates: Coordinates{
			Row:    row,
			Column: column,
		},
		Color: color,
	}
}

func (s Soloon) GetEndpoint() string {
	return defaultSoloonEndpoint
}

type soloonPayload struct {
	Row         int       `json:"row"`
	Column      int       `json:"column"`
	CandidateId string    `json:"candidateId"`
	Color       ColorType `json:"color"`
}

func (s Soloon) GetPayload(candidateId uuid.UUID) interface{} {
	return soloonPayload{
		Row:         s.Row,
		Column:      s.Column,
		Color:       s.Color,
		CandidateId: candidateId.String(),
	}
}
