package newwords

import (
	"os"
	"testing"
)

func TestWordwriter(t *testing.T) {
	if _, err := WordWriter(); os.IsNotExist(err) {
		t.Error("Failure during assigning", err)
	}
}
