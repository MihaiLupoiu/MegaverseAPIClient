package astral

import (
	"fmt"
	"strings"

	"github.com/google/uuid"
)

const defaultComethEndpoint = "/api/comeths"

type DirectionType string

const (
	Cometh_Up    DirectionType = "up"
	Cometh_Down  DirectionType = "down"
	Cometh_Left  DirectionType = "left"
	Cometh_Right DirectionType = "right"
)

var (
	DirectionMap = map[string]DirectionType{
		"up":    Cometh_Up,
		"down":  Cometh_Down,
		"left":  Cometh_Left,
		"right": Cometh_Right,
	}
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

// ParseStringToDirectionType parses a string and returns the corresponding DirectionType.
// It returns an error if the string cannot be parsed.
func StringToDirectionType(str string) (DirectionType, error) {
	c, ok := DirectionMap[strings.ToLower(str)]
	if !ok {
		return c, fmt.Errorf(`cannot parse:[%s] as DirectionType`, str)
	}
	return c, nil
}
