package auth_test

import (
	"github.com/bootdotdev/learn-cicd-starter/internal/auth"
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		headers http.Header
		want    string
		wantErr bool
	}{
		{
			name: "AuthorizationHeaderAPIKeyValid",
			headers: func() http.Header {
				h := http.Header{}
				h.Set("Authorization", "ApiKey tes")
				return h
			}(),
			want:    "tes",
			wantErr: false,
		},
		{
			name: "NoAuthorizationHeaderAPIKey",
			headers: func() http.Header {
				h := http.Header{}
				return h
			}(),
			want:    "",
			wantErr: true,
		},
		{
			name: "AuthorizationHeaderAPIKeyInvalidSplitTooShort",
			headers: func() http.Header {
				h := http.Header{}
				h.Set("Authorization", "ApiKey")
				return h
			}(),
			want:    "",
			wantErr: true,
		},
		{
			name: "AuthorizationHeaderAPIKeyInvalidStartsWithNoApiKey",
			headers: func() http.Header {
				h := http.Header{}
				h.Set("Authorization", "tes tes tes")
				return h
			}(),
			want:    "",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, gotErr := auth.GetAPIKey(tt.headers)
			if gotErr != nil {
				if !tt.wantErr {
					t.Errorf("GetAPIKey() failed: %v", gotErr)
				}
				return
			}
			if tt.wantErr {
				t.Fatal("GetAPIKey() succeeded unexpectedly")
			}
			if got != tt.want {
				t.Errorf("GetAPIKey() = %v, want %v", got, tt.want)
			}
		})
	}
}
