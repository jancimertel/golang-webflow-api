package gowebflow

import (
	"github.com/stretchr/testify/assert"
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

func TestNewClientOptionsPageSize(t *testing.T) {
	tests := []struct {
		name    string
		option ClientOption
		wantPageSize uint
	}{
		{
			name:         "empty option",
			option:       nil,
			wantPageSize: pageSize,
		},
		{
			name:         "provided option",
			option:       WithPageSize(1),
			wantPageSize: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cl, err := NewClient("random", tt.option)
			assert.Nil(t, err)
			assert.Equal(t, tt.wantPageSize, cl.pageSize)
		})
	}
}