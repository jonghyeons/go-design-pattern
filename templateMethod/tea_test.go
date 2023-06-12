package templateMethod

import "testing"

func TestTea(t *testing.T) {
	tea := &Tea{}
	tea.PrepareRecipe()

	coffee := &Coffee{}
	coffee.PrepareRecipe()
}
