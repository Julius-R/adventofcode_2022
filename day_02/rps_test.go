package main

import "testing"

func TestGetRoundScore(t *testing.T) {
	tests := []struct {
		name string
		arg  string
		want int
	}{
		{
			"Win",
			"A Y",
			8,
		},
		{
			"Lose",
			"B X",
			1,
		},
		{
			"Draw",
			"C Z",
			6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetRoundScore(tt.arg); got != tt.want {
				t.Errorf("GetRoundScore() = %v, want %v", got, tt.want)
			}
		})
	}
}
