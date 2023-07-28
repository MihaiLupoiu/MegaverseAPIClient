package astral

import (
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestNewPolyanet(t *testing.T) {
	polyanet := NewPolyanet(1, 2)
	assert.Equal(t, 1, polyanet.Row)
	assert.Equal(t, 2, polyanet.Column)
}

func TestPolyanet_GetEndpoint(t *testing.T) {
	polyanet := Polyanet{}
	assert.Equal(t, defaultPolyanetEndpoint, polyanet.GetEndpoint())
}

func TestPolyanet_GetPayload(t *testing.T) {
	candidateID := uuid.New()
	polyanet := Polyanet{
		Coordinates: Coordinates{
			Row:    1,
			Column: 2,
		},
	}
	payload := polyanet.GetPayload(candidateID).(polyanetPayload)
	assert.Equal(t, 1, payload.Row)
	assert.Equal(t, 2, payload.Column)
	assert.Equal(t, candidateID.String(), payload.CandidateId)
}
