package tests

import (
	"math/rand"
	"reflect"
	"testing"
	"testing/quick"
)

func TestGenerateRandomStringWithConfig(t *testing.T) {
	t.Parallel()
	f := func(n int) bool {
		if n < 0 || n > 1000 {
			return true
		}
		result := GenerateRandomString(n)
		return len(result) == n
	}

	config := &quick.Config{
		MaxCount: 1000,
		Values: func(v []reflect.Value, r *rand.Rand) {
			v[0] = reflect.ValueOf(r.Intn(1000))
		},
	}

	if err := quick.Check(f, config); err != nil {
		t.Errorf("Test failed: %v", err)
	}
}
