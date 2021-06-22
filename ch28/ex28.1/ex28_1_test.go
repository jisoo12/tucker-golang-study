package main

import "testing"
import "github.com/stretchr/testify/assert"

func TestSquare1(t *testing.T) {
	// rst := square(9)
	// if rst != 81 {
	// 	t.Errorf("square(9) should be 81 but square(9) returns %d\n", rst)
	// }

	assert := assert.New(t)
	assert.Equal(81, square(9), "square(9) should be 81")
}

func TestSquare2(t *testing.T) {
	// rst := square(3)
	// if rst != 9 {
	// 	t.Errorf("square(3) should be 9 but square(3) returns %d\n", rst)
	// }

	assert := assert.New(t)
	assert.Equal(9, square(3), "square(3) should be 9")
}