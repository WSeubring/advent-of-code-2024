package day9

import (
	"testing"

	assert "github.com/stretchr/testify/require"
)

func TestDecompress(t *testing.T) {
	diskMap := LoadInputDiskMap("")
	assert.Equal(t, "", diskMap.Decompress())
}

func TestDecompress2(t *testing.T) {
	diskMap := LoadInputDiskMap("1")
	assert.Equal(t, "0", diskMap.Decompress())
}

func TestDecompress3(t *testing.T) {
	diskMap := LoadInputDiskMap("22")
	assert.Equal(t, "00..", diskMap.Decompress())
}

func TestSolveA(t *testing.T) {
	SolveA()
}
