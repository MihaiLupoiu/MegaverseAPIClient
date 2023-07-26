package megaverse

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/MihaiLupoiu/MegaverseAPIClient/megaverse/astral"
)

type AstralService service

func (as *AstralService) Generate(
	ctx context.Context,
) error {
	return nil
}

func (as *AstralService) Delete(
	ctx context.Context,
) error {
	return nil
}

func (as *AstralService) GetGoalMap(ctx context.Context) (*astral.GoalMap, error) {
	// Define the API endpoint path for getting the goal map.
	endpoint := "/api/map/" + as.client.CandidateID.String() + "/goal"

	// Prepare the request using the newRequest method.
	req, err := as.client.newRequest(http.MethodGet, endpoint, nil)
	if err != nil {
		return nil, err
	}

	// Execute the request using the doRequest method.
	resp, err := as.client.doRequest(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// Handle the response and parse the goal map JSON.
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	var goalMap astral.GoalMap
	err = json.NewDecoder(resp.Body).Decode(&goalMap)
	if err != nil {
		return nil, err
	}

	return &goalMap, nil
}

func (as *AstralService) GetMap(ctx context.Context) (*astral.CurrentMap, error) {
	// Define the API endpoint path for getting the goal map.
	endpoint := "/api/map/" + as.client.CandidateID.String() + "/"

	// Prepare the request using the newRequest method.
	req, err := as.client.newRequest(http.MethodGet, endpoint, nil)
	if err != nil {
		return nil, err
	}

	// Execute the request using the doRequest method.
	resp, err := as.client.doRequest(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// Handle the response and parse the goal map JSON.
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	var currentMap astral.CurrentMap
	err = json.NewDecoder(resp.Body).Decode(&currentMap)
	if err != nil {
		return nil, err
	}

	return &currentMap, nil
}
