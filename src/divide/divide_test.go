package divide_test

import (
	"testing"

	"github.com/runnerdave/testing-example/src/divide"
)

func TestDivide(t *testing.T) {
	total := divide.Divide(6, 3)
	if total != 2 {
		t.Errorf("Did not divide properly, should have been 2 but it returned%d", total)
	}
}

// uncomment to achieve 100% coverage
// func TestDivideByZero(t *testing.T) {
// 	total := divide.Divide(6, 0)
// 	t.Logf("total:%d", total)
// }
