package decorator

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDecorator(t *testing.T) {
	d1 := NewStringDisplay("A")
	result := d1.Show(d1)
	assert.Equal(t, "A", result)
	d2 := NewSideBorder(d1, "|")
	result = d2.Show(d2)
	assert.Equal(t, "|A|", result)

}
