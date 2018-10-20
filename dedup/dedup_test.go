package dedup

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestChecker(t *testing.T) {
	var c checker
	f := "foo"
	b := "bar"
	require.False(t, c.IsDup(f), "first insert of value should be False")
	require.True(t, c.IsDup(f), "second insert of same value should be True")
	require.True(t, c.IsDup(f), "third insert of same value should be still be True")
	require.False(t, c.IsDup(b), "first insert of new value should be False")
	require.True(t, c.IsDup(b), "second insert of new value should be True")
}

func TestLeading(t *testing.T) {
	tests := []struct {
		description string
		input       []string
		expected    []string
	}{
		{
			description: "no duplicates, order should be preserved",
			input: []string{
				"vim",
				"ls",
				"cat",
			},
			expected: []string{
				"vim",
				"ls",
				"cat",
			},
		},
		{
			description: "last value duplicate, trailing order should be preserved",
			input: []string{
				"vim",
				"ls",
				"cat",
				"cat",
			},
			expected: []string{
				"vim",
				"ls",
				"cat",
			},
		},
		{
			description: "first value duplicate, trailing order should be preserved",
			input: []string{
				"vim",
				"vim",
				"ls",
				"cat",
			},
			expected: []string{
				"vim",
				"ls",
				"cat",
			},
		},
		{
			description: "middle value duplicate, trailing order should be preserved",
			input: []string{
				"vim",
				"ls",
				"ls",
				"cat",
			},
			expected: []string{
				"vim",
				"ls",
				"cat",
			},
		},
		{
			description: "various duplicates, trailing order should be preserved",
			input: []string{
				"ls",
				"vim",
				"ls",
				"cat",
				"vim",
				"cat",
				"ls",
				"cat",
				"cat",
				"cat",
				"vim",
				"ls",
				"cat",
			},
			expected: []string{
				"vim",
				"ls",
				"cat",
			},
		},
	}

	for _, test := range tests {
		test := test // Capture range variable.
		t.Run(test.description, func(t *testing.T) {
			t.Parallel()
			actual := Leading(test.input)
			assert.Equal(t, test.expected, actual,
				"expected does not match actual")
		})
	}
}
