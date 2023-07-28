package megaverse_test

import (
	"context"
	"net/http"
	"testing"

	"github.com/MihaiLupoiu/MegaverseAPIClient/megaverse"
	"github.com/MihaiLupoiu/MegaverseAPIClient/megaverse/astral"
	"github.com/stretchr/testify/assert"
)

// Replace with your package path

func TestNewClient(t *testing.T) {
	// Mock the HTTP client to return a response with HTTP status 200 and the expected JSON.
	mockHTTPClient := &http.Client{
		Transport: roundTripFunc(func(req *http.Request) *http.Response {
			resp := &http.Response{
				StatusCode: http.StatusOK,
			}
			return resp
		}),
	}

	// Create the client with the mock HTTP client.
	client, err := megaverse.NewClient("67f01a7f-64e2-4e40-b781-04113f1af7c5", mockHTTPClient)
	if err != nil {
		t.Fatalf("failed to create client: %v", err)
	}

	// Perform the API request.
	ctx := context.Background()
	astralObject := astral.NewPolyanet(0, 0)
	err = client.Astral.Generate(ctx, astralObject)
	assert.NoError(t, err)
}

// roundTripFunc is a custom HTTP transport that uses a function to handle the round-trip.
type roundTripFunc func(*http.Request) *http.Response

func (fn roundTripFunc) RoundTrip(req *http.Request) (*http.Response, error) {
	return fn(req), nil
}

func TestNewClientInvalidCondaditaID(t *testing.T) {
	// Create the client with the mock HTTP client.
	_, err := megaverse.NewClient("04113f1af7c5", nil)
	assert.Error(t, err)
}
