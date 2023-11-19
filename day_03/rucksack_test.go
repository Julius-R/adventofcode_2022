package main

import "testing"

func TestGetItemType(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		arg  string
		want int
	}{
		{
			"Should return 16",
			"vJrwpWtwJgWrhcsFMMfFFhFp",
			16,
		},
		{
			"Should return 38",
			"jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL",
			38,
		},
		{
			"Should return 42",
			"PmmdzqPrVvPwwTWBwg",
			42,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetItemType(tt.arg); got != tt.want {
				t.Errorf("GetItemType() = %v, want %v", got, tt.want)
			}
		})
	}
}
