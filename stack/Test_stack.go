package stack

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStack(t *testing.T) {
	s := &Stack{}
	s.push(1)
	s.push(3)
	s.pop()
	s.push(5)
	arr := []int{1, 5}
	res := []int{}
	cur := s.tail
	for cur != nil {
		res = append(res, cur.Val)
		cur = cur.Next
	}
	assert.Same(t, arr, res, "Tidak sama")
} 