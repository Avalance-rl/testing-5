package tests

import (
	"errors"
	"testing"
)

func TestCheckPositiveNumber(t *testing.T) {
	tests := []struct {
		name    string
		input   int
		wantErr error
	}{
		{
			name:    "Positive number",
			input:   10,
			wantErr: nil},
		{
			name:    "Negative number",
			input:   -5,
			wantErr: ErrNegativeNumber,
		},
	}

	t.Parallel()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := CheckPositiveNumber(tt.input)

			if !errors.Is(err, tt.wantErr) {
				t.Errorf("expected error %v, got %v", tt.wantErr, err)
			}
		})
	}
}
