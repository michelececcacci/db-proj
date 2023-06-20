package util

// holds a message in case of success and an error. Components
// can choose what to display. currently used in error.go
type OptionalError struct {
	Message string
	Err     error
}

type UpdateUsername struct {
	Username *string
}
