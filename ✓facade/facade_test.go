package facade

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFacade(t *testing.T) {
	maker := PageMaker{}
	result := maker.MakeWelcomePage("a@a.com")
	expect := "# Welcome to a's page!"

	assert.Equal(t, expect, result)
}
