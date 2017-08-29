// Package bork will output the string "bork".
package bork

import (
	"errors"
	"fmt"
	"log"
)

// Bork will output "bork" to stdout. Parameter is just used to decide whether to return an error or not.
//
// Provide "cat" to get a bork.ErrUnborkable error.
// Provide "cate" to get a native error, properly formatted with the package name in front.
func Bork(out string) error {
	if out == "cat" {
		return &ErrUnborkable
	}

	// When we get a non-bork error, we'll add the "bork" prefix (package name) here, since this method is exported.
	err := simulateExternalPackage(out)
	if err != nil {
		return fmt.Errorf("bork: %v", err)
	}

	// Output bork.
	log.Println("bork")

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
