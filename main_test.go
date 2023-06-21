package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSumFound_1(t *testing.T) {
	c := assert.New(t)
	output := sum([]int{1, 0, 5, 20, -4, 12, 16, 7}, 12)
	c.Len(output, 3)
}

func TestSumFound_2(t *testing.T) {
	c := assert.New(t)
	output := sum([]int{1, 0, 5, 20, -4, 12, 16, 7}, 0)
	c.Len(output, 1)
}

func TestSumNofFound(t *testing.T) {
	c := assert.New(t)
	output := sum([]int{1, 0, 5, 20, -4, 12, 16, 7}, 99)
	c.Len(output, 0)
}
