package main

import "testing"

func TestFindOverlap(t *testing.T) {
	tests := []struct {
		name string
		arg  string
		want bool
	}{
		{
			"2-8,3-7",
			"2-8,3-7",
			true,
		},
		{
			"6-6,4-6",
			"6-6,4-6",
			true,
		},
		{
			"2-4,6-8",
			"2-4,6-8",
			false,
		},
		{
			"2-6,4-8",
			"2-6,4-8",
			true,
		},
		{
			"5-7,7-9",
			"5-7,7-9",
			true,
		},
		{
			"5-7,3-9",
			"5-7,3-9",
			true,
		},
		{
			"55-68,68-69",
			"55-68,68-69",
			true,
		},
		{
			"39-95,38-96",
			"39-95,38-96",
			true,
		},
		{
			"76-91,51-91",
			"76-91,51-91",
			true,
		},
		{
			"49-51,29-90",
			"49-51,29-90",
			true,
		},
		{
			"3-92,17-92",
			"3-92,17-92",
			true,
		},
		{
			"11-74,7-12",
			"11-74,7-12",
			true,
		},
		{
			"25-97,25-26",
			"25-97,25-26",
			true,
		},
		{
			"23-23,22-46",
			"23-23,22-46",
			true,
		},
		{
			"49-51,29-90",
			"49-51,29-90",
			true,
		},
		{
			"14-70,11-70",
			"14-70,11-70",
			true,
		},
		{
			"2-95,94-96",
			"2-95,94-96",
			true,
		},
		{
			"19-61,61-61",
			"19-61,61-61",
			true,
		},
		{
			"47-88,47-87",
			"47-88,47-87",
			true,
		},
		{
			"1-51,1-52",
			"1-51,1-52",
			true,
		},
		{
			"12-12,13-13",
			"12-12,13-13",
			false,
		},
		/*
				{
						"",
						"",
						true,
					},

			69-90,68-88



		*/
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FindOverlap(tt.arg); got != tt.want {
				t.Errorf("FindOverlap() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFindOverlap1(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FindOverlap(tt.args.str); got != tt.want {
				t.Errorf("FindOverlap() = %v, want %v", got, tt.want)
			}
		})
	}
}
