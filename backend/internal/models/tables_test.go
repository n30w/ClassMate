package models

import "testing"

func TestTable_String(t *testing.T) {
	tests := []struct {
		table    Table
		expected string
	}{
		{USERS, "users"},
		{COURSES, "courses"},
		{MEDIA, "media"},
		{PROJECTS, "projects"},
		{MESSAGES, "messages"},
		{ASSIGNMENTS, "assignments"},
		{SUBMISSIONS, "submissions"},
		{USER_COURSES, "user_courses"},
		{COURSE_MESSAGES, "course_messages"},
		{COURSE_TEACHERS, "course_teachers"},
		{COURSE_ROSTER, "course_roster"},
		{COURSE_ASSIGNMENTS, "course_assignments"},
		{ASSIGNMENT_SUBMISSIONS, "assignment_submissions"},
		{MESSAGE_MEDIA, "message_media"},
	}

	for _, tc := range tests {
		t.Run(
			tc.expected, func(t *testing.T) {
				result := tc.table.String()
				if result != tc.expected {
					t.Errorf("Expected %s, got %s", tc.expected, result)
				}
			},
		)
	}

	// Test the default case
	t.Run(
		"invalid", func(t *testing.T) {
			result := Table(-1).String()
			if result != "invalid table" {
				t.Errorf("Expected 'invalid table', got '%s'", result)
			}
		},
	)
}
