package main

import (
	"testing"
)

func TestClient_Get(t *testing.T) {
	tests := []struct {
		name       string
		endpoint   string
		wantStatus int
		wantErr    bool
	}{
		{"succeeded pattern", "http://success.com", 200, false},
		{"failed pattern(server error)", "http://fail.com", 500, true},
		{"failed pattern(client error)", "http://xxxx.com", 400, true},
		//{"failed pattern(client error)", "http://xxxx.com", 200, true},
	}

	cli := NewClient("0.0.0", "test")
	for _, tt := range tests {
		status, err := cli.Get(tt.endpoint)
		if status != tt.wantStatus || (err != nil) != tt.wantErr {
			t.Fatalf("want status = %d, got status = %d; want error = %t, got error = %t",
				tt.wantStatus, status, tt.wantErr, (err != nil),
			)
		}
	}
}
