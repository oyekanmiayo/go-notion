package version1

import "fmt"

// APIError represents a Notion API response
// https://developers.notion.com/reference/errors
// Object is always "error"
type APIError struct {
	Object  string `json:"object,omitempty"`
	Status  int32  `json:"status,omitempty"`
	Code    string `json:"code,omitempty"`
	Message string `json:"message,omitempty"`
}

func (e APIError) Error() string {
	return fmt.Sprintf("code: %v, message: %v", e.Code, e.Message)
}

// relevantError returns any http-related error if it exists
// if http-related errors don't exist, it returns apiError if it exists
// else it returns nil
func relevantError(httpError error, apiError APIError) error {
	if httpError != nil {
		return httpError
	}

	if (APIError{}) == apiError {
		return nil
	}

	return apiError
}
