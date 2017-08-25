package bridge

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBridge(t *testing.T) {

	dh := NewDefaultDisplayWithImp(2, &StringDisplayImp{})
	//DisplayHelper{&StringDisplayImp{}, 2}
	expect := "+-+\n|2|\n+-+\n"
	assert.Equal(t, expect, dh.Display())

	dh = NewDefaultDisplayWithImp(2, &IntegerDisplayImp{})
	//DisplayHelper{&IntegerDisplayImp{}, 2}
	expect = "+--+\n|2|\n+--+\n"
	assert.Equal(t, expect, dh.Display())

	cd := NewCountDisplayWithImp(1, &StringDisplayImp{})
	//CountDisplay{&DisplayHelper{&StringDisplayImp{}, 1}}
	expect = "+-+\n|1|\n+-+\n"
	assert.Equal(t, expect, cd.Display())

	expect = "+-+\n|1|\n|1|\n+-+\n"
	assert.Equal(t, expect, cd.MultiDisplay(2))

	cd = NewCountDisplayWithImp(1, &IntegerDisplayImp{})
	//CountDisplay{&DisplayHelper{&IntegerDisplayImp{}, 1}}
	expect = "+-+\n|1|\n+-+\n"
	assert.Equal(t, expect, cd.Display())

	expect = "+-+\n|1|\n|1|\n+-+\n"
	assert.Equal(t, expect, cd.MultiDisplay(2))
}
