package minus_test

import (
	"testing"

	"github.com/runnerdave/testing-example/src/minus"
)

func TestMinus(t *testing.T) {
	total := minus.Minus(2, 1)
	if total != 1 {
		t.Errorf("Should have totalled 1 but instead was %d", total)
	}
}
