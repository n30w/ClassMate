package models

// Contains anything related to communications,
// such as discussion posts and user messages.
type Discussion struct {
	name            string
	description     string
	participants    []User
	mediaReferences []Media
	commentThreads  []string
}
