package tests

import "testing"



func Test_PositivePow(t *testing.T) {
    tests := []struct{
        name string
        number int
        degree uint
        want int
    }{
        {"1", 10, 2, 100},
        {"2", 2, 10, 1024},
        {"3", 3, 4, 81},
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            got := PositivePow(tt.number, tt.degree)
            if tt.want != got {
                t.Errorf("got %d, want %d", got, tt.want)
            }

        })
    }

}
