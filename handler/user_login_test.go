package handler

import (
	"context"
	"testing"

	"github.com/cloudwego/hertz/pkg/app"
)

func TestUserLogin(t *testing.T) {
	type args struct {
		ctx context.Context
		c   *app.RequestContext
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			UserLogin(tt.args.ctx, tt.args.c)
		})
	}
}
