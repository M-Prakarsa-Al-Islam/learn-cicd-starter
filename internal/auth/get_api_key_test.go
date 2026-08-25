package auth

import (
	"errors"
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	// 1. Define the test cases table
	tests := []struct {
		name          string
		headers       http.Header
		expectedKey   string
		expectedError error
	}{
		{
			name: "Valid API key",
			headers: http.Header{
				"Authorization": []string{"ApiKey secret12345"},
			},
			expectedKey:   "secret12345",
			expectedError: nil,
		},
		{
			name:          "Missing Authorization header",
			headers:       http.Header{},
			expectedKey:   "",
			expectedError: ErrNoAuthHeaderIncluded,
		},
		{
			name: "Malformed header - missing ApiKey prefix",
			headers: http.Header{
				"Authorization": []string{"Bearer secret12345"},
			},
			expectedKey:   "",
			expectedError: errors.New("malformed authorization header"),
		},
		{
			name: "Malformed header - missing token value",
			headers: http.Header{
				"Authorization": []string{"ApiKey"},
			},
			expectedKey:   "",
			expectedError: errors.New("malformed authorization header"),
		},
	}

	// 2. Loop over each test case
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotKey, gotErr := GetAPIKey(tt.headers)

			// Check key output
			if gotKey != tt.expectedKey {
				t.Errorf("GetAPIKey() key = %v, want %v", gotKey, tt.expectedKey)
			}

			// Check error output
			if tt.expectedError != nil {
				if gotErr == nil || gotErr.Error() != tt.expectedError.Error() {
					t.Errorf("GetAPIKey() error = %v, want %v", gotErr, tt.expectedError)
				}
			} else if gotErr != nil {
				t.Errorf("GetAPIKey() unexpected error = %v", gotErr)
			}
		})
	}
}
