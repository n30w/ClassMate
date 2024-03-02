package main

import "testing"

func TestJsonBuilder(t *testing.T) {

	t.Run("single value map", func(t *testing.T) {
		data := map[string]any{
			"i": "1",
		}

		want := []byte{'{', '"', 'i', '"', ':', '"', '1', '"', '}', '\n'}

		var got []byte
		got, _ = jsonBuilder(&data)

		// Compare each index of the byte slice.
		for i := range len(got) {
			if want[i] != got[i] {
				t.Errorf("Incorrect JSON Output: got: %s, want: %s", got, want)
				return
			}
		}
	})

	t.Run("five value map", func(t *testing.T) {
		data := map[string]any{
			"i": "1",
			"y": "2",
			"x": "3",
			"b": "4",
			"a": "5",
		}

		// The Go Standard Library sorts a map in alphabetical order
		// when marshaling a JSON for deterministic purposes.
		want := []byte{
			'{',
			'"', 'a', '"', ':', '"', '5', '"',
			',',
			'"', 'b', '"', ':', '"', '4', '"',
			',',
			'"', 'i', '"', ':', '"', '1', '"',
			',',
			'"', 'x', '"', ':', '"', '3', '"',
			',',
			'"', 'y', '"', ':', '"', '2', '"',
			'}',
			'\n',
		}

		var got []byte

		got, _ = jsonBuilder(&data)

		// Compare each index of the byte slice.
		for i := range len(got) {
			if want[i] != got[i] {
				t.Errorf("Incorrect JSON Output: got: %s, want: %s", got, want)
				return
			}
		}
	})

	t.Run("empty map", func(t *testing.T) {
		data := map[string]any{}

		// The Go Standard Library sorts a map in alphabetical order
		// when marshaling a JSON for deterministic purposes.
		want := []byte{'{', '}', '\n'}

		var got []byte

		got, _ = jsonBuilder(&data)

		// Compare each index of the byte slice.
		for i := range len(got) {
			if want[i] != got[i] {
				t.Errorf("Incorrect JSON Output: got: %s, want: %s", got, want)
				return
			}
		}
	})
}
