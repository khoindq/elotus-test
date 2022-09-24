package main

import (
	"testing"

	"reflect"

	"github.com/stretchr/testify/assert"
)

func TestGrayCode(t *testing.T) {

	// Test
	assert := assert.New(t)
	assert.True(reflect.DeepEqual(grayCode(2), []int{0, 1, 3, 2}))
	assert.True(reflect.DeepEqual(grayCode(1), []int{0, 1}))

}

func TestFindLength(t *testing.T) {

	// Test
	assert := assert.New(t)
	assert.Equal(findLength([]int{1, 2, 3, 2, 1}, []int{3, 2, 1, 4, 7}), 3)
	assert.Equal(findLength([]int{0, 0, 0, 0, 0}, []int{0, 0, 0, 0, 0}), 5)

}
