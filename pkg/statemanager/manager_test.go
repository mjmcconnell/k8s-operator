package statemanager

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStateManagerSetGet(t *testing.T) {
	// arrange
	smA := GetStateManager()
	smB := GetStateManager()

	// act
	smA.Set("Foo", "foo")
	smB.Set("Bar", "bar")

	// assert
	assert.Nil(t, smA.Get("FooBar"))
	assert.Equal(t, smA.Get("Bar"), "bar")
	assert.Equal(t, smB.Get("Foo"), "foo")
	assert.Equal(t, smA, smB)
}

func TestStateManagerClear(t *testing.T) {
	// arrange
	sm := GetStateManager()

	// act
	sm.Set("Foo", "foo")
	sm.Clear()

	// assert
	assert.Nil(t, sm.Get("Foo"))
}
