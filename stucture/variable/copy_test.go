package copy_test

import (
	vv "golang_learning/stucture/variable"
	"testing"
)

func TestQCopy(t *testing.T) {
	vv.QCopy()
	vv.QCopy2()
	t.Log("All tests passed")
}
