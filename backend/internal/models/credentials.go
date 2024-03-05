package models

type credentials interface {
	// valid validates whether a certain type of credential
	// is taken or fits to a defined parameters for what type
	// of values are allowed.
	valid() error
}
