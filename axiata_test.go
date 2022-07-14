package axiatadigitaltest

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAllValidTestFromExample(t *testing.T) {
	t.Log("--- Melakukan test dari soal")
	assert.Equal(t, 7, FindOddLost([]int{3, 3, 5, 9, -3, 1}))
	assert.Equal(t, 9, FindOddLost([]int{1, 5, 3, 15, 7}))
	assert.Equal(t, 1, FindOddLost([]int{-5, -10, -9}))
	assert.Equal(t, 7, FindOddLost([]int{3, 5, 1}))

}

func TestAllValidTestFromOwn(t *testing.T) {
	t.Log("--- Melakukan test custom")
	assert.Equal(t, 11, FindOddLost([]int{1, 3, 5, 7, 9}))
	assert.Equal(t, 13, FindOddLost([]int{11, 9, 5, 7}))
	assert.Equal(t, 103, FindOddLost([]int{101, 107, 105}))
	assert.Equal(t, 501, FindOddLost([]int{495, 499, 497}))
}
