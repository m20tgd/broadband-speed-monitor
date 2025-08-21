package http_request

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type ApiMethod string

const (
	GET    ApiMethod = "GET"
	POST   ApiMethod = "POST"
	PUT    ApiMethod = "PUT"
	DELETE ApiMethod = "DELETE"
)

func HttpRequest[T any, U any](path string, method ApiMethod, token string, data T, result *U) error {

	// Marshal the data into JSON and create a request body
	var requestBody io.Reader
	dataJSON, err := json.Marshal(data)
	if err != nil {
		return fmt.Errorf("failed to marshal JSON: %v", err)
	}
	requestBody = bytes.NewBuffer(dataJSON)

	// Create a url request
	req, err := http.NewRequest(string(method), path, requestBody)
	if err != nil {
		return fmt.Errorf("failed to create request: %v", err)
	}

	// Set the request headers
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json, text/plain, */*")
	req.Header.Set("authorization", "Bearer "+token)

	// Send the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("failed to send request: %v", err)
	}
	defer resp.Body.Close()

	// Check the response status code
	if resp.StatusCode != http.StatusOK {
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return fmt.Errorf("unexpected status code: %d. Error in parsing response body, so error details not known", resp.StatusCode)
		}
		return fmt.Errorf("unexpected status code: %d, error message: %s", resp.StatusCode, string(body))
	}

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("failed to read response body: %v", err)
	}

	if len(body) == 0 {
		return nil
	}

	if err := json.Unmarshal(body, &result); err != nil {
		return fmt.Errorf("failed to unmarshal JSON: %v", err)
	}

	return nil
}
