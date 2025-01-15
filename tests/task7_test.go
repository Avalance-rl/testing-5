package tests

import (
	"testing"
	"time"
)

func Test_GetGreetingMessage(t *testing.T) {
	tests := []struct {
		name          string
		mockTime      time.Time
		expectedGreet string
	}{
		{
			name:          "Morning",
			mockTime:      time.Date(2025, 1, 15, 9, 0, 0, 0, time.UTC),
			expectedGreet: "Good morning!",
		},
		{
			name:          "Afternoon",
			mockTime:      time.Date(2025, 1, 15, 15, 0, 0, 0, time.UTC),
			expectedGreet: "Good afternoon!",
		},
		{
			name:          "Evening",
			mockTime:      time.Date(2025, 1, 15, 20, 0, 0, 0, time.UTC),
			expectedGreet: "Good evening!",
		},
	}

	t.Parallel()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := GetGreetingMessage(tt.mockTime)
			if got != tt.expectedGreet {
				t.Errorf("expected %q, got %q", tt.expectedGreet, got)
			}
		})
	}
}
