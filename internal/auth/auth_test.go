package auth

import (
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name     string
		headers  http.Header
		want     string
		wantErr  error
		wantErrs []string
	}{
		{
			name: "ApiKey",
			headers: http.Header{
				"Authorization": []string{"ApiKey 8080"},
			},
			want:    "8080",
			wantErr: nil,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got, err := GetAPIKey(tt.headers)
			if err != nil && err.Error() != tt.wantErr.Error() {
				t.Errorf("GetAPIKey() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if got != tt.want {
				t.Errorf("GetAPIKey() = %v, want %v", got, tt.want)
			}

			for _, wantErr := range tt.wantErrs {
				if err == nil && wantErr != "" {
					t.Errorf("GetAPIKey() error = %v, wantErr %v", err, wantErr)
					return
				}
				if err != nil && wantErr != "" && err.Error() != wantErr {
					t.Errorf("GetAPIKey() error = %v, wantErr %v", err, wantErr)
					return
				}
			}
		})
	}
}
