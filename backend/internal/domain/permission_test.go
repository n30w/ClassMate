package domain

import (
	"fmt"
	"testing"
)

func TestPermission_String(t *testing.T) {
	t.Run(
		"all false", func(t *testing.T) {
			p := permission{}
			p.read = false
			p.write = false
			p.update = false
			p.delete = false

			want := "----"

			got := fmt.Sprintf("%s", p)

			if want != got {
				t.Errorf("got %s, want %s", got, want)
			}
		},
	)

	t.Run(
		"all true", func(t *testing.T) {
			p := permission{}
			p.read = true
			p.write = true
			p.update = true
			p.delete = true

			want := "----"

			got := fmt.Sprintf("%s", p)

			if want != got {
				t.Errorf("got %s, want %s", got, want)
			}
		},
	)

	t.Run(
		"one true", func(t *testing.T) {
			p := permission{}
			p.read = true
			p.write = false
			p.update = false
			p.delete = false

			want := "r---"

			got := fmt.Sprintf("%s", p)

			if want != got {
				t.Errorf("got %s, want %s", got, want)
			}
		},
	)
}

func TestPermission_fromString(t *testing.T) {
	t.Run(
		"all false", func(t *testing.T) {
			got := fromString("----")
			want := permission{}

			if got != want {
				t.Errorf("got %v, want %v", got, want)
			}
		},
	)

	t.Run(
		"all true", func(t *testing.T) {
			got := fromString("rwud")
			want := permission{read: true, write: true, update: true,
				delete: true}

			if got != want {
				t.Errorf("got %v, want %v", got, want)
			}
		},
	)

	t.Run(
		"one true", func(t *testing.T) {
			got := fromString("r---")
			want := permission{read: true}

			if got != want {
				t.Errorf("got %v, want %v", got, want)
			}
		},
	)
}
