package file

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetSubDirectories(t *testing.T) {
	dir := "./config"
	subDirs, err := GetSubDirectories(dir)
	assert.Nil(t, err)
	assert.Equal(t, []string{"test_data/subdir1", "test_data/subdir2"}, subDirs)

}
