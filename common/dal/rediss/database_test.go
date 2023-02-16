package rediss

import (
	"context"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func init() {
	initRedis()
	fmt.Println("init")
}

func TestSetToken(t *testing.T) {
	var err error
	ctx := context.Background()
	err = SetToken(ctx, "hhh", "123456")
	assert.NoError(t, err)
	err = SetToken(ctx, "bbb", "217014")
	assert.NoError(t, err)

}

func TestGetTokenByName(t *testing.T) {
	tests := []struct {
		name    string
		args    string
		want    string
		wantErr bool
	}{
		// cases
		{
			name:    "right case",
			args:    "hhh",
			want:    "123456",
			wantErr: false,
		},
		{
			name:    "wrong case",
			args:    "bb",
			want:    "",
			wantErr: true,
		},
	}

	ctx := context.Background()

	for _, tt := range tests {
		got, err := GetTokenByName(ctx, tt.args)
		assert.Equal(t, tt.want, got)
		assert.Equal(t, tt.wantErr, err != nil)
	}
}
