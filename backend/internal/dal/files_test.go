package dal

import "testing"

func TestExcelStore_Get(t *testing.T) {
	es := NewExcelStore()
	p := "../../resources/"
	fileName := "grade-offline-template.xlsx"

	want := [][]string{
		{
			"Name", "Net ID", "Grade", "Feedback", "Submission ID",
		},
	}

	got, err := es.Get(p + fileName)
	if err != nil {
		t.Errorf("%+v", err)
	}

	if len(got) != len(want) {
		t.Errorf("length got %d, want %d", len(got), len(want))
	}

	if len(got[0]) != len(want[0]) {
		t.Errorf("length got %d, want %d", len(got[0]), len(want[0]))
	}

	for i := range want {
		for j := range want[i] {
			if got[i][j] != want[i][j] {
				t.Errorf("got %s, want %s", got[i][j], want[i][j])
			}
		}
	}
}
