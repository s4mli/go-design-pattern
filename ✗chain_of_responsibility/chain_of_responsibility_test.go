package chain_of_responsibility

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestChainOfResponsibility(t *testing.T) {
	a := NewNoSupport("A")
	b := NewLimitSupport("B", 2)
	c := NewLimitSupport("C", 3)
	a.SetNext(c)
	c.SetNext(b)
	//a.SetNext(b)
	//b.SetNext(c)

	result := a.Handle(a, &Trouble{1})
	expect := "Trouble : 1 is resolved by C"
	assert.Equal(t, expect, result)

	result = a.Handle(a, &Trouble{2})
	expect = "Trouble : 2 is resolved by C"

	assert.Equal(t, expect, result)

	result = a.Handle(a, &Trouble{3})
	expect = "Trouble : 3 cannot be resolved"
	assert.Equal(t, expect, result)
}
