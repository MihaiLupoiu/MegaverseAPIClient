package astral

import (
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestNewCometh(t *testing.T) {
	cometh := NewCometh(1, 2, Cometh_Up)
	assert.Equal(t, 1, cometh.Row)
	assert.Equal(t, 2, cometh.Column)
	assert.Equal(t, Cometh_Up, cometh.Direction)
}

func TestCometh_GetEndpoint(t *testing.T) {
	cometh := Cometh{}
	assert.Equal(t, defaultComethEndpoint, cometh.GetEndpoint())
}

func TestCometh_GetPayload(t *testing.T) {
	candidateID := uuid.New()
	cometh := Cometh{
		Coordinates: Coordinates{
			Row:    1,
			Column: 2,
		},
		Direction: Cometh_Up,
	}
	payload := cometh.GetPayload(candidateID).(comethPayload)
	assert.Equal(t, 1, payload.Row)
	assert.Equal(t, 2, payload.Column)
	assert.Equal(t, Cometh_Up, payload.Direction)
	assert.Equal(t, candidateID.String(), payload.CandidateId)
}

func TestStringToDirectionType(t *testing.T) {
	tests := []struct {
		input  string
		result DirectionType
		err    bool
	}{
		{"up", Cometh_Up, false},
		{"down", Cometh_Down, false},
		{"left", Cometh_Left, false},
		{"right", Cometh_Right, false},
		{"invalid", "", true},
	}

	for _, test := range tests {
		res, err := StringToDirectionType(test.input)
		if err != nil {
			assert.True(t, test.err)
		} else {
			assert.Equal(t, test.result, res)
		}
	}
}
