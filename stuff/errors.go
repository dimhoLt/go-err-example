package stuff

// Error is the type that is returned on unexpected states in the stuff package.
type Error struct {
	s string
}

// Error returns the error message.
func (e Error) Error() string {
	return e.s
}

var (
	// ErrUnknownType is returned if the provided parameter is of unknown type.
	ErrUnknownType = Error{"stuff: unknown type"}

	// ErrUnhandledType is returned if the provided parameter is recognized, but not implemented.
	ErrUnhandledType = Error{"stuff: unhandled type"}
)
