// Package bork will output the string "bork".
package bork

import (
	"errors"
	"fmt"
)

// Bork will output "bork" to stdout, along with the provided parameter, as formatted JSON.
func Bork(out string) error {
	if out == "cat" {
		return &ErrUnborkable
	}

	err := simulateExternalPackage(out)
	if err != nil {
		// When we get a non-bork error, we'll add the "bork" prefix (package name) here, since this method is exported.
		return fmt.Errorf("bork: %v", err)
	}

	// Return explicit nil at function end.
	return nil
}

// A helper function to simulate returning a non-bork error.
func simulateExternalPackage(out string) error {
	if out == "cate" {
		return errors.New("native error")
	}

	return nil
}
