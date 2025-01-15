package tests

import (
	"github.com/stretchr/testify/assert"
	"sync"
	"testing"
)

func Test_ParallelIncrement(t *testing.T) {
	counter := &Counter{}
	want := 1000
	amountIterations := 1_000
	wg := &sync.WaitGroup{}
	wg.Add(amountIterations)
	for i := 0; i < amountIterations; i++ {
		go func() {
			counter.Increment()
			defer wg.Done()
		}()
	}
	wg.Wait()
	assert.Equal(t, want, counter.GetValue())
}
