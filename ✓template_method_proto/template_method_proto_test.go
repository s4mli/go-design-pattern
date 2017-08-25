package template_method_proto

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCharDisplay(t *testing.T) {
	expect := "<<< A >< A >< A >< A >< A >>>"
	display := NewCharDisplay('A')
	assert.Equal(t, expect, display.Display())
}

func TestCharDisplayWithoutClose(t *testing.T) {
	expect := "<<< B >< B >< B >< B >< B >"
	display := NewCharDisplayWithoutClose('B')
	assert.Equal(t, expect, display.Display())
}
