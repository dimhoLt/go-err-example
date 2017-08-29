package main

import (
	"fmt"
	"log"
	"net/http"

	"bitbucket.org/dimholt/go-err-example/bork"
)

// This file is intended as an example for how you can implement this error handling to convert different errors to
// correct status codes in Go.
//
// Since the "stuff" package checks for types, its errors don't make much sense trying to implement here, but the bork
// one will. Provide
//
// To try the "bork"-package, use the URL "http://localhost:8600/bork?seed=[something]"

func main() {
	h := &handler{}
	p := ":8600"
	done := make(chan bool)

	go func() {
		log.Fatal(http.ListenAndServe(p, h))
		done <- true
	}()

	log.Printf("now listening on port %v", p)
	log.Println("try running:")
	log.Printf("curl -i \"http://localhost%s/bork?seed=cat\"", p)
	log.Printf("use seeds 'cat' and 'cate' for errors, anything else for success")
	<-done
}

type handler struct{}

func (h *handler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	switch req.URL.Path[1:] {
	case "bork":
		h.handleBork(w, req)
	default:
		w.WriteHeader(http.StatusNotFound)
	}
}

func (h *handler) handleBork(w http.ResponseWriter, req *http.Request) {
	p := req.URL.Query().Get("seed")
	err := bork.Bork(p)

	switch {
	// Unborkable
	case err == &bork.ErrUnborkable:
		log.Println("unborkable, sending bad request")
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("unborkable seed"))
		return

	// Everything else.
	case err != nil:
		log.Println("unhandled, sending internal server error")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("unknown error returned from bork"))
		return
	}

	w.WriteHeader(http.StatusOK)

	msg := fmt.Sprintf("successfully borked '%v'", p)
	w.Write([]byte(msg))
}
