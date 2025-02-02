package texts

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestDecodeQuery(t *testing.T) {
	testCases := []struct {
		input  string
		output Query
	}{
		{
			input: "",
			output: Query{
				Params: map[string][]string{},
			},
		},
		{
			input: "/add",
			output: Query{
				Params: map[string][]string{
					"add": nil,
				},
			},
		},
		{
			input: "/single&word=2&word=3&params=4",
			output: Query{
				Params: map[string][]string{
					"single": nil,
					"word":   {"2", "3"},
					"params": {"4"},
				},
			},
		},
		{
			input: "/multi /word=2,3 /params=4",
			output: Query{
				Params: map[string][]string{
					"multi":  nil,
					"word":   {"2", "3"},
					"params": {"4"},
				},
			},
		},
		{
			input: "/multi /word=2,3 with texts /params=4 text2",
			output: Query{
				Texts: []string{"with texts", "text2"},
				Params: map[string][]string{
					"multi":  nil,
					"word":   {"2", "3"},
					"params": {"4"},
				},
			},
		},
	}

	for _, tC := range testCases {
		t.Run(tC.input, func(t *testing.T) {
			res := DecodeQuery(tC.input)
			require.Equal(t, tC.output, res)
		})
	}
}

func TestQuery_Encode(t *testing.T) {
	testCases := []struct {
		input  Query
		output string
	}{
		{
			input: Query{
				Params: map[string][]string{},
			},
			output: "",
		},
		{
			input: Query{
				Params: map[string][]string{
					"add": nil,
				},
			},
			output: "/add",
		},
		{
			input: Query{
				Params: map[string][]string{
					"single": nil,
					"word":   {"2", "3"},
					"params": {"4"},
				},
			},
			output: "/params=4&single&word=2,3",
		},
		{
			input: Query{
				Texts: []string{"with texts", "text2"},
				Params: map[string][]string{
					"empty": {},
				},
			},
			output: "/empty with texts text2",
		},
	}
	for _, tC := range testCases {
		t.Run(tC.output, func(t *testing.T) {
			res := tC.input.Encode()
			require.Equal(t, tC.output, res)
		})
	}
}
