package id

import (
	"strings"
	"testing"
)

func TestIdType(t *testing.T) {
	type test struct {
		name  string
		input TypeId
		want  string
	}

	tables := []test{
		{
			name:  "User",
			input: User,
			want:  "u",
		},
		{
			name:  "Company",
			input: Company,
			want:  "c",
		},
	}

	for _, tc := range tables {
		t.Run(tc.name, func(t *testing.T) {
			id := New(tc.input)

			got := strings.Split(id, "_")

			if got[0] != tc.want {
				t.Fatalf("expected: %v, got: %v", tc.want, got)
			}
		})
	}

	t.Run("None", func(t *testing.T) {
		id := New(None)
		if len(id) != 27 {
			t.Error("Length of ksuid should be 27")
		}

		if GetType(id) != None {
			t.Error("Type should be None")
		}

	})
}
