package gowebflow

import (
	"reflect"
	"testing"
)

func TestNewClient(t *testing.T) {
	tests := []struct {
		name    string
		secret  string
		wantNil    bool
		wantErr bool
	}{
		{
			name:    "empty secret",
			secret:  "",
			wantNil:    true,
			wantErr: true,
		},
		{
			name:    "provided secret",
			secret:  "1",
			wantNil:   false,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewClient(tt.secret)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewClient() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got == nil, tt.wantNil) {
				t.Errorf("NewClient() got = %v, want nil = %v", got, tt.wantNil)
			}
		})
	}
}