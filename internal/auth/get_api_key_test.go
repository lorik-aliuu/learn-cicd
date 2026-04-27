package auth

import (
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	tests := []struct {
		name        string
		headers     http.Header
		wantKey     string
		wantErr     bool
	}{
		{
			name:    "missing authorization header",
			headers: http.Header{},
			wantErr: true,
		},
		{
			name:    "malformed authorization header",
			headers: http.Header{"Authorization": []string{"Bearer abc123"}},
			wantErr: true,
		},
		{
			name:    "valid api key header",
			headers: http.Header{"Authorization": []string{"ApiKey abc123"}},
			wantKey: "abc123",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotKey, err := GetAPIKey(tt.headers)
			if tt.wantErr {
				if err == nil {
					t.Fatalf("GetAPIKey() error = nil, want error")
				}
				return
			}

			if err != nil {
				t.Fatalf("GetAPIKey() error = %v, want nil", err)
			}

			if gotKey != tt.wantKey {
				t.Fatalf("GetAPIKey() = %q, want %q", gotKey, tt.wantKey)
			}
		})
	}
}
