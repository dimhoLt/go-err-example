// Package stuff implements doing stuff.
package stuff

// RecognizeParam will attempt to see if it recognizes the type of the provided parameter.
func RecognizeParam(p interface{}) error {
	switch p.(type) {
	case int, string:
		return nil
	case float64, float32:
		return &ErrUnhandledType
	}

	return &ErrUnknownType
}
