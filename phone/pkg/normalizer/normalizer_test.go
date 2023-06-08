package normalizer_test

import (
	"phone/pkg/normalizer"
	"testing"
)

type NormalizeTestCase struct {
	Input    string
	Expected string
}

func TestNormalizer(t *testing.T) {
	testCases := []NormalizeTestCase{
		{
			Input:    "1234567890",
			Expected: "1234567890",
		},
		{
			Input:    "123 456 7891",
			Expected: "1234567891",
		},
		{
			Input:    "(123) 456 7892",
			Expected: "1234567892",
		},
		{
			Input:    "(123) 456-7893",
			Expected: "1234567893",
		},
		{
			Input:    "123-456-7894",
			Expected: "1234567894",
		},
		{
			Input:    "123-456-7890",
			Expected: "1234567890",
		},
		{
			Input:    "(123)456-7892",
			Expected: "1234567892",
		},
	}
	for _, tc := range testCases {
		t.Run(tc.Input, func(t *testing.T) {
			output := normalizer.Nomarlize(tc.Input)
			if output != tc.Expected {
				t.Errorf("Got %s; want %s", output, tc.Expected)
			}
		})
	}
}

func TestNormalizeRegex(t *testing.T) {
	testCases := []NormalizeTestCase{
		{
			Input:    "1234567890",
			Expected: "1234567890",
		},
		{
			Input:    "123 456 7891",
			Expected: "1234567891",
		},
		{
			Input:    "(123) 456 7892",
			Expected: "1234567892",
		},
		{
			Input:    "(123) 456-7893",
			Expected: "1234567893",
		},
		{
			Input:    "123-456-7894",
			Expected: "1234567894",
		},
		{
			Input:    "123-456-7890",
			Expected: "1234567890",
		},
		{
			Input:    "(123)456-7892",
			Expected: "1234567892",
		},
	}
	for _, tc := range testCases {
		t.Run(tc.Input, func(t *testing.T) {
			output := normalizer.NormalizeRegex(tc.Input)
			if output != tc.Expected {
				t.Errorf("Got %s; want %s", output, tc.Expected)
			}
		})
	}
}
