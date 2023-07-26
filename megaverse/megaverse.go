package megaverse

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/cenkalti/backoff/v4"
	"github.com/google/uuid"
	"golang.org/x/time/rate"
)

const (
	// TODO: [improvement] API should have version. e.g: "https://challenge.crossmint.io/v1/api/"
	defaultBaseURL   = "https://challenge.crossmint.io/"
	defaultTimeout   = time.Duration(3) * time.Second
	defaultRateLimit = 20 // Maximum number of requests per second

)

type service struct {
	client *Client
}

// A Client manages communication with the Megaverse API.
type Client struct {
	client  *http.Client
	limiter *rate.Limiter

	BaseURL     *url.URL
	CandidateID uuid.UUID

	common service // Reuse a single struct instead of allocating one for each service on the heap.

	// Services used for talking to different parts of the Megaverse API.
	Astral *AstralService
}

// NewClient returns a new Megaverse API client. If a nil httpClient is
// provided, a new http.Client will be used.
func NewClient(candidateID string, httpClient *http.Client) (*Client, error) {
	if httpClient == nil {
		httpClient = &http.Client{}
	}
	limiter := rate.NewLimiter(rate.Limit(defaultRateLimit), 1)

	// TODO: Add ability to change the baseURL
	baseURL, err := url.Parse(defaultBaseURL)
	if err != nil {
		return nil, err
	}

	candID, err := uuid.Parse(candidateID)
	if err != nil {
		return nil, err
	}

	c := &Client{
		client:      httpClient,
		limiter:     limiter,
		BaseURL:     baseURL,
		CandidateID: candID,
	}

	c.common.client = c
	c.Astral = (*AstralService)(&c.common)

	return c, nil
}

func (c *Client) newRequest(method, path string, payload interface{}) (*http.Request, error) {
	rel := &url.URL{Path: path}
	u := c.BaseURL.ResolveReference(rel)

	var body []byte
	if payload != nil {
		var err error
		body, err = json.Marshal(payload)
		if err != nil {
			return nil, err
		}
	}

	req, err := http.NewRequest(method, u.String(), bytes.NewReader(body))
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", "application/json; charset=utf-8")
	return req, nil
}

func (c *Client) doRequest(req *http.Request) (*http.Response, error) {
	var resp *http.Response
	var err error

	err = c.limiter.Wait(context.Background())
	if err != nil {
		return nil, err
	}

	reqOperation := func() error {
		if resp != nil {
			resp.Body.Close()
		}

		resp, err = c.client.Do(req)
		if err != nil || resp.StatusCode >= http.StatusInternalServerError {
			return fmt.Errorf("request failed")
		}
		return nil
	}

	backoffConfig := backoff.NewExponentialBackOff()
	backoffConfig.MaxElapsedTime = defaultTimeout

	retryErr := backoff.Retry(reqOperation, backoffConfig)
	if retryErr != nil {
		return nil, retryErr
	}

	return resp, nil
}
