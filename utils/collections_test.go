package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRemoveIndexFromArr(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5}

	result := RemoveIndexFromArr(arr, 2)

	expected := []int{1, 2, 4, 5}
	assert.Equal(t, expected, result)
}

func TestInsertAtIndex(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5}

	result := InsertAtIndex(arr, 2, 10)

	expected := []int{1, 2, 10, 3, 4, 5}
	assert.Equal(t, expected, result)
}
