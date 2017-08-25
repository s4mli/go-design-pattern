package template_method

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCharDisplay(t *testing.T) {
	expect := "<<< A >< A >< A >< A >< A >>>"
	display := NewCharDisplay('A')
	// using self as the para of Display(printer)
	assert.Equal(t, expect, display.Display(display))
}

func TestStringDisplay(t *testing.T) {
	expect := "+-----+\n< | ABC |\n >< | ABC |\n >< | ABC |\n >< | ABC |\n >< | ABC |\n >+-----+\n"
	display := NewStringDisplay("ABC")
	// using self as the para of Display(printer)
	assert.Equal(t, expect, display.Display(display))
}
