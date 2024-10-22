package greetings

import (
	"regexp"
	"testing"
)

// Checks for valid return
func TestHelloName(t *testing.T) {
	name := "tom"
	want := regexp.MustCompile(`\b` + name + `\b`)
	msg, err := Hello("tom")
	if !want.MatchString(msg) || err != nil {
		t.Fatalf(`hello("tom") = %q, %v, want match for %#q, nil`, msg, err, want)
	}
}

// Call with empty string and check for error
func TestHelloEmpty(t *testing.T) {
	msg, err := Hello("")
	if msg != "" || err == nil {
		t.Fatalf(`Hello("") = %q, %v, want "", error`, msg, err)
	}
}
