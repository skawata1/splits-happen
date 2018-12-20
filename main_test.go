package main

import (
	"testing"
)

func Test(t *testing.T) {
	tests := []struct {
		name string
		in []string
		out int
	}{
		{"perfect", []string{"X","X","X","X","X","X","X","X","X","X","X","X"}, 300},
		{"9gutters", []string{"9","-","9","-","9","-","9","-","9","-","9","-","9","-","9","-","9","-","9","-"}, 90},
		{"5spares", []string{"5","/","5","/","5","/","5","/","5","/","5","/","5","/","5","/","5","/","5","/","5"}, 150},
		{"mixedbag", []string{"X","7","/","9","-","X","-","8","8","/","-","6","X","X","X","8","1"}, 167},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T){
			result:= calculateBowlingScore(test.in)

			if result != test.out {
				t.Errorf("got %v, wanted %v", result, test.out)
			}
		})

	}
}