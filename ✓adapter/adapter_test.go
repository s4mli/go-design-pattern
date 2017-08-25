package adapter

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAdapterByEmbedded(t *testing.T) {
	decorator := NewEmbeddedDecorateBanner("A")
	assert.Equal(t, "* A *", decorator.Decorate())
}

func TestAdapterByDelegate(t *testing.T) {
	decorator := NewCompositionDecorateBanner("A")
	assert.Equal(t, "* A *", decorator.Decorate())
}
