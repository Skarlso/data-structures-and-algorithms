package chapter09

import (
	"testing"
)

func TestLinter(t *testing.T) {
	testCases := []struct {
		desc string
		name string
		want string
		data string
	}{
		{
			desc: "everything is fine",
			name: "matching",
			data: "[]",
		},
		{
			desc: "everything is fine",
			name: "matching",
			data: "[asdf]",
		},
		{
			desc: "everything is fine",
			name: "matching",
			data: "(asdf)",
		},
		{
			desc: "everything is fine",
			name: "matching",
			data: "{asdfasdf}",
		},
		{
			desc: "missing closing",
			name: "matching",
			want: "stack was not empty at the end of the string, contained: {",
			data: "{asdfasdf",
		},
		{
			desc: "missing opening",
			name: "matching",
			want: "} didn't have an opening",
			data: "asdfasdf}",
		},
		{
			desc: "not matching",
			name: "matching",
			want: "last opening { did not match encounter closing )",
			data: "{asdfasdf)",
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			linter := NewLinter()
			got := linter.Lint(tC.data)
			if tC.want != "" {
				if got == nil {
					t.Fatalf("expected error but none received")
				}
				if tC.want != got.Error() {
					t.Fatalf("want: %s; got: %s", tC.want, got)
				}
			}
		})
	}
}
