package megaverse

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/MihaiLupoiu/MegaverseAPIClient/megaverse/astral"
)

type AstralService service

func (as *AstralService) Generate(
	ctx context.Context,
	astralObject astral.AstralObject,
) error {
	endpoint := astralObject.GetEndpoint()
	payload := astralObject.GetPayload(as.client.CandidateID)

	return as.executeRequest(ctx, http.MethodPost, endpoint, payload, nil)
}

func (as *AstralService) Delete(
	ctx context.Context,
	astralObject astral.AstralObject,
) error {
	endpoint := astralObject.GetEndpoint()
	payload := astralObject.GetPayload(as.client.CandidateID)

	return as.executeRequest(ctx, http.MethodDelete, endpoint, payload, nil)
}

func (as *AstralService) GetGoalMap(ctx context.Context) (*astral.GoalMap, error) {
	endpoint := "/api/map/" + as.client.CandidateID.String() + "/goal"

	var goalMap astral.GoalMap
	err := as.executeRequest(ctx, http.MethodGet, endpoint, nil, &goalMap)
	if err != nil {
		return nil, err
	}
	return &goalMap, nil
}

func (as *AstralService) GetMap(ctx context.Context) (*astral.CurrentMap, error) {
	endpoint := "/api/map/" + as.client.CandidateID.String() + "/"

	var currentMap astral.CurrentMap
	err := as.executeRequest(ctx, http.MethodGet, endpoint, nil, &currentMap)
	if err != nil {
		return nil, err
	}
	return &currentMap, nil
}

func (as *AstralService) executeRequest(
	ctx context.Context,
	method, endpoint string,
	payload interface{},
	target interface{},
) error {
	// Prepare the request using the newRequest method.
	req, err := as.client.newRequest(method, endpoint, payload)
	if err != nil {
		return err
	}

	// Execute the request using the doRequest method.
	resp, err := as.client.doRequest(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Check if the response status code is OK (200).
	if resp.StatusCode != http.StatusOK {
		b, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("unexpected status code: %d, msg: %s", resp.StatusCode, b)
	}

	if target != nil {
		// Parse the JSON response into the target struct.
		err = json.NewDecoder(resp.Body).Decode(target)
		if err != nil {
			return err
		}
	}

	return nil
}
