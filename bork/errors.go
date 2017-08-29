package bork

// Error is the type that is returned on unexpected states in the bork package.
type Error struct {
	s string
}

// Error returns the error message.
func (e Error) Error() string {
	return e.s
}

var (
	// ErrUnborkable is returned if the parameter provided is not borkable.
	ErrUnborkable = Error{"bork: unborkable"}
)
