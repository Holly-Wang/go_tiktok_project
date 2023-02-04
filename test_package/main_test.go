package testpackage

import (
	"testing"

	"github.com/cloudwego/hertz/cmd/hz/util/logs"
	"github.com/stretchr/testify/assert"
)

func TestAll(t *testing.T) {
	defer logs.Flush()
	assert.Equal(t, 1, 2)
	logs.Infof("test go package")
}
