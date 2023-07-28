package astral

import (
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestNewSoloon(t *testing.T) {
	soloon := NewSoloon(3, 4, Soloon_Red)
	assert.Equal(t, 3, soloon.Row)
	assert.Equal(t, 4, soloon.Column)
	assert.Equal(t, Soloon_Red, soloon.Color)
}

func TestSoloon_GetEndpoint(t *testing.T) {
	soloon := Soloon{}
	assert.Equal(t, defaultSoloonEndpoint, soloon.GetEndpoint())
}

func TestSoloon_GetPayload(t *testing.T) {
	candidateID := uuid.New()
	soloon := Soloon{
		Coordinates: Coordinates{
			Row:    3,
			Column: 4,
		},
		Color: Soloon_Blue,
	}
	payload := soloon.GetPayload(candidateID).(soloonPayload)
	assert.Equal(t, 3, payload.Row)
	assert.Equal(t, 4, payload.Column)
	assert.Equal(t, Soloon_Blue, payload.Color)
	assert.Equal(t, candidateID.String(), payload.CandidateId)
}

func TestStringToColorType(t *testing.T) {
	testCases := []struct {
		input    string
		expected ColorType
		err      bool
	}{
		{"blue", Soloon_Blue, false},
		{"red", Soloon_Red, false},
		{"purple", Soloon_Purple, false},
		{"white", Soloon_White, false},
		{"invalid_color", "", true},
	}

	for _, tc := range testCases {
		t.Run(tc.input, func(t *testing.T) {
			colorType, err := StringToColorType(tc.input)
			if tc.err {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tc.expected, colorType)
			}
		})
	}
}
