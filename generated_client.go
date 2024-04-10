//go:build exclude

package smssendapi

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
)



// GetUserInfoRequest represents the request body for the GetUserInfo endpoint.
type GetUserInfoRequest struct {
	// Define fields based on your API's request model
}

// GetUserInfoResponse represents the response body for the GetUserInfo endpoint.
type GetUserInfoResponse struct {
	// Define fields based on your API's response model
}



// UpdateUserInfoRequest represents the request body for the UpdateUserInfo endpoint.
type UpdateUserInfoRequest struct {
	// Define fields based on your API's request model
}

// UpdateUserInfoResponse represents the response body for the UpdateUserInfo endpoint.
type UpdateUserInfoResponse struct {
	// Define fields based on your API's response model
}





// GetUserInfo makes a request to the GetUserInfo endpoint.
func GetUserInfo(req GetUserInfoRequest) (GetUserInfoResponse, error) {
	// Marshal request body to JSON
	reqBody, err := json.Marshal(req)
	if err != nil {
		return GetUserInfoResponse{}, err
	}

	// Create HTTP request
	url := c.baseURL + "/user/{id}"
	httpReq, err := http.NewRequest("GET", url, bytes.NewBuffer(reqBody))
	if err != nil {
		return GetUserInfoResponse{}, err
	}

	// Set headers, authentication, etc. if needed
	// httpReq.Header.Set("Content-Type", "application/json")
	// Add authentication headers, etc.

	// Send HTTP request
	resp, err := c.httpClient.Do(httpReq)
	if err != nil {
		return GetUserInfoResponse{}, err
	}
	defer resp.Body.Close()

	// Read response body
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return GetUserInfoResponse{}, err
	}

	// Unmarshal response body
	var response GetUserInfoResponse
	err = json.Unmarshal(respBody, &response)
	if err != nil {
		return GetUserInfoResponse{}, err
	}

	return response, nil
}


// UpdateUserInfo makes a request to the UpdateUserInfo endpoint.
func UpdateUserInfo(req UpdateUserInfoRequest) (UpdateUserInfoResponse, error) {
	// Marshal request body to JSON
	reqBody, err := json.Marshal(req)
	if err != nil {
		return UpdateUserInfoResponse{}, err
	}

	// Create HTTP request
	url := c.baseURL + "/user/{id}"
	httpReq, err := http.NewRequest("PUT", url, bytes.NewBuffer(reqBody))
	if err != nil {
		return UpdateUserInfoResponse{}, err
	}

	// Set headers, authentication, etc. if needed
	// httpReq.Header.Set("Content-Type", "application/json")
	// Add authentication headers, etc.

	// Send HTTP request
	resp, err := c.httpClient.Do(httpReq)
	if err != nil {
		return UpdateUserInfoResponse{}, err
	}
	defer resp.Body.Close()

	// Read response body
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return UpdateUserInfoResponse{}, err
	}

	// Unmarshal response body
	var response UpdateUserInfoResponse
	err = json.Unmarshal(respBody, &response)
	if err != nil {
		return UpdateUserInfoResponse{}, err
	}

	return response, nil
}
