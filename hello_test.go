package main

import "testing"

//Beautiful testing
func TestHello(t *testing.T) {
	got := Hello("Joe")
	want := "Hello, Joe"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
