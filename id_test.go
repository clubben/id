package id

import (
	"strings"
	"testing"
)

func TestIdType(t *testing.T) {
	type test struct {
		input TypeId
		want  string
	}

	tables := []test{
		{
			input: User,
			want:  "u",
		},
		{
			input: Company,
			want:  "c",
		},
	}

	for _, tc := range tables {
		id := New(tc.input)

		got := strings.Split(id, "_")

		if got[0] != tc.want {
			t.Fatalf("expected: %v, got: %v", tc.want, got)
		}
	}
}
