package chapter20

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCoinGame(t *testing.T) {
	winner := GameWinner(1, "you")
	assert.Equal(t, "them", winner)
	winner = GameWinner(2, "you")
	assert.Equal(t, "you", winner)
	winner = GameWinner(3, "you")
	assert.Equal(t, "you", winner)
	winner = GameWinner(4, "you")
	assert.Equal(t, "them", winner)
	winner = GameWinner(5, "you")
	assert.Equal(t, "you", winner)
	winner = GameWinner(6, "you")
	assert.Equal(t, "you", winner)
	winner = GameWinner(7, "you")
	assert.Equal(t, "them", winner)
}
