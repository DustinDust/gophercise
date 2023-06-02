package solution_test

import (
	"hr/pkgs/solution"
	"testing"
)

func TestCamelcase(t *testing.T) {
	input := "saveChangesInTheEditor"
	expected := int32(5)
	output := solution.Camelcase(input)
	if output != expected {
		t.FailNow()
	}
}