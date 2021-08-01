package helper

import "testing"

func TestInitialize(t *testing.T) {
	tests := []struct {
		name    string
		wantErr bool
	}{
		{
			name:    "nil-test",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := Initialize(); (err != nil) != tt.wantErr {
				t.Errorf("Initialize() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
