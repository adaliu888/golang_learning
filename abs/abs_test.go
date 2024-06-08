package abs

import (
	"testing"
)

func TestAbs(t *testing.T) {
	// Test abs function with -1, 0, 1
	if abs(-1) != 1 {
		t.Errorf("Abs(-1) = %d, want 1", abs(-1))
	}
	if abs(1) != 1 {
		t.Errorf("Abs(1) = %d, want 1", abs(1))
	}
	if abs(0) != 0 {
		t.Errorf("Abs(0) = %d, want 0", abs(0))
	}
}

func TestAbs1(t *testing.T) {
	t.Run("positive", func(t *testing.T) {
		if abs(-1) != 1 {
			t.Errorf("Abs(-1) = %d, want 1", abs(-1))
		}
		if abs(1) != 1 {
			t.Errorf("Abs(1) = %d, want 1", abs(1))
		}
		if abs(0) != 0 {
			t.Errorf("Abs(0) = %d, want 0", abs(0))
		}
	})
	t.Run("negative", func(t *testing.T) {
		if abs(-1) != 1 {
			t.Errorf("Abs(-1) = %d, want 1", abs(-1))
		}
		if abs(1) != 1 {
			t.Errorf("Abs(1) = %d, want 1", abs(1))
		}
		if abs(0) != 0 {
			t.Errorf("Abs(0) = %d, want 0", abs(0))
		}
	})
}

func TestSkip(t *testing.T) {
	t.Skip("skipping the remaining the structure")
}

func TestClean(t *testing.T) {
	t.Cleanup(func() {
		t.Log("cleaning up")
	})
	t.Log("running test")
}
