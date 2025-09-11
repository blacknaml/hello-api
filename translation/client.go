package translation

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"log"
	"net/http"
)

type APIClient struct {
	endpoint string
}

var _ HelloClient = &APIClient{}

// NewHelloClient creates instance of client with a given endpoint.
func NewHelloClient(endpoint string) *APIClient {
	return &APIClient{
		endpoint: endpoint,
	}
}

// Translate will call external client for translation.
func (c *APIClient) Translate(word, language string) (string, error) {
	req := map[string]interface{}{
		"word":     word,
		"language": language,
	}

	b, err := json.Marshal(req)
	if err != nil {
		return "", errors.New("unable to encode msg")
	}

	resp, err := http.Post(c.endpoint, "application/json", bytes.NewBuffer(b))
	if err != nil {
		log.Println(err)
		return "", errors.New("call to api failed")
	}

	if resp.StatusCode == http.StatusNotFound {
		return "", nil
	}

	if resp.StatusCode == http.StatusInternalServerError {
		return "", errors.New("errors in API")
	}

	b, _ = io.ReadAll(resp.Body)

	defer func() {
		if closeErr := resp.Body.Close(); closeErr != nil {
			// Log the error, but don't necessarily return it as it might overshadow
			// a more significant error from the main function logic.
			log.Printf("Error closing response body: %v", closeErr)
		}
	}()

	var m map[string]interface{}
	if err := json.Unmarshal(b, &m); err != nil {
		return "", errors.New("unable to decode response")
	}
	return m["translate"].(string), nil
}
