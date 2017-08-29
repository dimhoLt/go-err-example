package main

import (
	"log"

	"github.com/dimholt/go-err-example/bork"
	"github.com/dimholt/go-err-example/stuff"
)

var stuffParams []interface{}
var borkParams []string

func main() {
	// First, let's set up some random parameters that will yield different result from the `stuff`-package.
	stuffParams = []interface{}{
		1,
		"abc123",
		10.5,             // Trigger ErrUnhandledType
		[]string{"test"}, // Trigger ErrUnknownType
	}

	// Now, let's do the same for the bork package.
	borkParams = []string{
		"bork",
		"spork",
		"bonk",
		"cat",  // Trigger bork error
		"cate", // Trigger native error
	}

	log.Println("running showcaseErrorHandling")
	log.Println("----------")
	showcaseErrorHandling()

	log.Println("just outputting the errors")
	log.Println("----------")
	justOutputErrors()
}

// showcaseErrorHandling will show how you can handle some errors specifically.
func showcaseErrorHandling() {
	// Note that all errors should be of the `error` type. This would not be possible - or interchangeable - if the
	// functions we call would specify their own types as values.
	var err error

	for _, sp := range stuffParams {
		err = stuff.RecognizeParam(sp)

		// Now, we check if the returned error points to our declared variable.
		switch err {
		case &stuff.ErrUnknownType:
			log.Println("successfully recognized a stuff.ErrUnknownType")
		case &stuff.ErrUnhandledType:
			log.Println("successfully recognized a stuff.ErrUnhandledType")
		case nil:
			log.Println("no error")
		default:
			// All native errors raised in the bork package, and thus converted.
			log.Printf("no specific handler for stuff error '%T'", err)
		}
	}

	for _, bp := range borkParams {
		err = bork.Bork(bp)

		// Now, we check if the returned error points to our declared variable.
		switch err {
		case &bork.ErrUnborkable:
			log.Printf("The error was successfully considered a bork.ErrUnborkable for param '%v'", bp)
		case nil:
			log.Println("no error")
		default:
			log.Printf("no specific handler for bork error '%T'", err)
		}
	}
}

// justOutputErrors will implement how all existing implementations, not caring about which error you get, will work
// exactly as usual.
func justOutputErrors() {
	// Note that all errors should be of the `error` type. This would not be possible - or interchangeable - if the
	// functions we call would specify their own types as values.
	var err error

	for _, sp := range stuffParams {
		err = stuff.RecognizeParam(sp)
		if err != nil {
			log.Println(err)
		}
	}

	for _, bp := range borkParams {
		err = bork.Bork(bp)
		if err != nil {
			log.Println(err)
		}
	}
}
