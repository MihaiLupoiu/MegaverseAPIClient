package megaverse_test

import (
	"context"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/MihaiLupoiu/MegaverseAPIClient/megaverse"
	"github.com/MihaiLupoiu/MegaverseAPIClient/megaverse/astral"
	"github.com/stretchr/testify/assert"
)

// MockHTTPClient is a mock implementation of http.Client for testing purposes.
type MockHTTPClient struct {
	DoFunc func(req *http.Request) (*http.Response, error)
}

func (c *MockHTTPClient) Do(req *http.Request) (*http.Response, error) {
	return c.DoFunc(req)
}

// CustomHTTPClient is a custom implementation of http.Client that wraps the MockHTTPClient.
type CustomHTTPClient struct {
	mockHTTPClient *MockHTTPClient
}

func (c *CustomHTTPClient) Do(req *http.Request) (*http.Response, error) {
	return c.mockHTTPClient.Do(req)
}

func TestAstralService_Generate(t *testing.T) {
	// Define a mock HTTP server using httptest package.
	mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		respJSON := `{}`
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(respJSON))
	}))
	defer mockServer.Close()

	// Create the mock HTTP client with the mock server.
	mockHTTPClient := mockServer.Client()

	client, err := megaverse.NewClient("67f01a7f-64e2-4e40-b781-04113f1af7c5", mockHTTPClient)
	assert.NoError(t, err)

	// Set the BaseURL of the client to the URL of the mock server.
	client.BaseURL, err = url.Parse(mockServer.URL)
	assert.NoError(t, err)

	as := client.Astral

	// Define an AstralObject for testing.
	astralObject := astral.NewPolyanet(1, 1)

	// Perform the Generate API request.
	err = as.Generate(context.Background(), astralObject)
	assert.NoError(t, err, "Generate should not return an error")
}

func TestAstralService_Delete(t *testing.T) {
	// Define a mock HTTP server using httptest package.
	mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		respJSON := `{}`
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(respJSON))
	}))
	defer mockServer.Close()

	client, err := megaverse.NewClient("67f01a7f-64e2-4e40-b781-04113f1af7c5", mockServer.Client())
	assert.NoError(t, err)

	// Set the BaseURL of the client to the URL of the mock server.
	client.BaseURL, err = url.Parse(mockServer.URL)
	assert.NoError(t, err)

	as := client.Astral

	// Define an AstralObject for testing.
	astralObject := astral.NewPolyanet(1, 1)

	// Perform the Delete API request.
	err = as.Delete(context.Background(), astralObject)
	assert.NoError(t, err, "Delete should not return an error")
}

func TestAstralService_GetGoalMap(t *testing.T) {
	// Define a mock HTTP server using httptest package.
	mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		respJSON := `{
			"goal": [
			  ["POLYANET", "POLYANET", "COMETH"],
			  ["SOLOON", "POLYANET", "COMETH"],
			  ["SOLOON", "SOLOON", "COMETH"]
			]
		  }`
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(respJSON))
	}))
	defer mockServer.Close()

	client, err := megaverse.NewClient("67f01a7f-64e2-4e40-b781-04113f1af7c5", mockServer.Client())
	assert.NoError(t, err)

	// Set the BaseURL of the client to the URL of the mock server.
	client.BaseURL, err = url.Parse(mockServer.URL)
	assert.NoError(t, err)

	as := client.Astral
	// Perform the GetGoalMap API request.
	res, err := as.GetGoalMap(context.Background())
	assert.NoError(t, err, "GetGoalMap should not return an error")
	assert.NotNil(t, res, "GetGoalMap should not be nil")
	assert.Len(t, res.Goal, 3)
}

func TestAstralService_GetMap(t *testing.T) {
	// Define a mock HTTP server using httptest package.
	mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		respJSON := `{
			"map": {
				"_id": "12345",
				"content": [
					[{"type": 1}, {"type": 2}, {"type": 3}],
					[{"type": 4}, {"type": 5}, {"type": 6}],
					[{"type": 7}, {"type": 8}, {"type": 9}]
				],
				"candidateId": "67f01a7f-64e2-4e40-b781-04113f1af7c5",
				"phase": 1,
				"__v": 1
			}
		}`
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(respJSON))
	}))
	defer mockServer.Close()

	client, err := megaverse.NewClient("67f01a7f-64e2-4e40-b781-04113f1af7c5", mockServer.Client())
	assert.NoError(t, err)

	// Set the BaseURL of the client to the URL of the mock server.
	client.BaseURL, err = url.Parse(mockServer.URL)
	assert.NoError(t, err)

	as := client.Astral
	// Perform the GetMap API request.
	res, err := as.GetMap(context.Background())
	assert.NoError(t, err, "GetMap should not return an error")
	assert.NotNil(t, res, "GetMap should not be nil")
	assert.NotNil(t, res.Map, "Map should not be nil")
	assert.Len(t, res.Map.Content, 3)
	for _, row := range res.Map.Content {
		assert.Len(t, row, 3)
	}
}
