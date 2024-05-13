package models

type Credential interface {
	// Valid validates whether a certain type of credential
	// is taken or fits to a defined parameters for what type
	// of values are allowed.
	Valid() error
	String() string
}
