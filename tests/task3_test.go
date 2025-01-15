package tests

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_MultipleEachSlcieElements(t *testing.T) {
	tests := []struct {
		name       string
		slice      []int
		multiplier int
		want       []int
	}{
		{"1", []int{1, 2, 3, 4, 5}, 5, []int{5, 10, 15, 20, 25}},
		{"2", []int{0, 5, 10, 15}, -1, []int{0, -5, -10, -15}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := MultipleEachSliceElements(tt.slice, tt.multiplier)
			assert.Equal(t, tt.want, got)
		})
	}
}
