package observer

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestObserver(t *testing.T) {
	r := NewRandomNumberGenerator()
	o1 := &DigitObserver{}
	o2 := &DigitObserver{}
	r.AddObserver(o1)
	r.AddObserver(o2)

	results := r.Execute()
	assert.Equal(t, 2, len(results))

	for _, result := range results {
		assert.Equal(t, true, result < 100)
		assert.Equal(t, true, result >= 0)
	}
}
