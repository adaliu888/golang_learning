package file

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFN(t *testing.T) {
	file := "2024-08-24.log"
	assert.Equal(t, file, FN())

}
