package unpack

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestUnpack(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{input: "a4bc2d5e", expected: "aaaabccddddde"},
		{input: "abccd", expected: "abccd"},
		{input: "", expected: ""},
		{input: "aaa0b", expected: "aab"},
		{input: `qwe\4\5`, expected: `qwe45`},
		{input: `qwe\45`, expected: `qwe44444`},
		{input: `qwe\\5`, expected: `qwe\\\\\`},
		{input: `qwe\\\3`, expected: `qwe\3`},
		{input: "d\n5abc", expected: "d\n\n\n\n\nabc"},
		{input: "-3qwe+4", expected: "---qwe++++"},
		{input: `\\4\\5\\6`, expected: `\\\\\\\\\\\\\\\`},
		{input: `\\4\\5\6`, expected: `\\\\\\\\\6`},
		{input: `\6abc`, expected: `6abc`},
		{input: `+\6abc`, expected: `+6abc`},
		{input: `aaa0\0`, expected: `aa0`},
		{input: `aaa0\\0`, expected: `aa`},
		{input: `\-6abcd`, expected: `------abcd`},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.input, func(t *testing.T) {
			result, err := Unpack(tc.input)
			require.NoError(t, err)
			require.Equal(t, tc.expected, result)
		})
	}
}

func TestUnpackInvalidString(t *testing.T) {
	tests := []struct {
		input         string
		expectedError string
	}{
		{input: "3abc", expectedError: "first element is number"},
		{input: "45", expectedError: "first element is number"},
		{input: "aaa10b", expectedError: "element is number"},
		{input: `\\666`, expectedError: "element is number"},
		{input: `6\6abv`, expectedError: "first element is number"},
		{input: `aaa00`, expectedError: "element is number"},
	}
	for _, tc := range tests {
		tc := tc
		t.Run(tc.input, func(t *testing.T) {
			_, err := Unpack(tc.input)
			require.EqualErrorf(t, err, tc.expectedError, "error message %v", err)
		})
	}
}
