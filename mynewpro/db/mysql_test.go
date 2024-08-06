package db

import (
	"testing"
	//断言
	"github.com/stretchr/testify/assert"
)

func TestDBInit(t *testing.T) {
	DBIint()
	assert.NotNil(t, DBConnect)
}
