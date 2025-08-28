package auth

import (
	"net/http"
	"testing"
)

func TestGetApiKet(t *testing.T) {
	req, _ := http.NewRequest("GET", "", nil)
	token := "Bearer hopuga"
	req.Header.Set("Authorization", token)

	req2, _ := http.NewRequest("GET", "", nil)

	req3, _ := http.NewRequest("GET", "", nil)
	token2 := "ApiKey bruh"
	req3.Header.Set("Authorization", token2)

	tests := []struct {
		name    string
		token   string
		header  http.Header
		wantErr bool
	}{
		{
			name:    "Wrong header token",
			token:   "hopuga",
			header:  req.Header,
			wantErr: true,
		},
		{
			name:    "No header",
			token:   "",
			header:  req2.Header,
			wantErr: true,
		},
		{
			name:    "Correct header",
			token:   "bruh",
			header:  req3.Header,
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ret, err := GetAPIKey(tt.header)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetAPIKey() error = %v, wantErr %v", err, tt.wantErr)
			}
			if (err == nil) && ret != tt.token {
				t.Errorf("GetAPIKey() wrong token, token %v, return %v", token, ret)
			}
		})
	}
}
