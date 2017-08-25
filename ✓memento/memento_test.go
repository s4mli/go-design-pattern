package memento

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMemento(t *testing.T) {
	game := &Game{100}
	m := game.CreateMemento()
	game.Money = 0
	game.RestoreMemento(m)
	assert.Equal(t, game.Money, 100)
}
