package main

import "testing"

func TestGetItemType(t *testing.T) {
	tests := []struct {
		name string
		arg  string
		want int
	}{
		{
			"Should return 16 for 'vJrwpWtwJgWrhcsFMMfFFhFp'",
			"vJrwpWtwJgWrhcsFMMfFFhFp",
			16,
		},
		{
			"Should return 38 for 'jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL'",
			"jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL",
			38,
		},
		{
			"Should return 42 for 'PmmdzqPrVvPwwTWBwg'",
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

func TestGetGroupItemType(t *testing.T) {
	tests := []struct {
		name string
		arg  []string
		want int
	}{
		{
			"Duplicate r should return 18",
			[]string{
				"vJrwpWtwJgWrhcsFMMfFFhFp",
				"jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL",
				"PmmdzqPrVvPwwTWBwg",
			},
			18,
		},
		{
			"Duplicate Z should return 52",
			[]string{
				"wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn",
				"ttgJtRGJQctTZtZT",
				"CrZsJsPPZsGzwwsLwLmpwMDw",
			},
			52,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetGroupItemType(tt.arg); got != tt.want {
				t.Errorf("GetGroupItemType() = %v, want %v", got, tt.want)
			}
		})
	}
}
