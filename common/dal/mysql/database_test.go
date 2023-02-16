package mysql

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDatabaseOp(t *testing.T) {
	InitDB()
	err := CreateLike(123, 321, 111)
	assert.NoError(t, err)
	id, err := FindIDinLike(321, 111)
	assert.NoError(t, err)
	del := DelLike(id)
	assert.Nil(t, del)
}
