package notion

import "fmt"

// APIError represents a Notion API response
// https://developers.notion.com/reference/errors
type APIError struct {
	Code    string `json:"code"`
	Message string `json:"message"`
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

	return nil
}