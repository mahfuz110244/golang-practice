package main

import (
	"regexp"
	"testing"
)

func TestHelloWorld(t *testing.T) {
	name := "World"
	want := regexp.MustCompile(`Hello, World!`)
	msg, err := HelloWorld(name)
	if !want.MatchString(msg) || err != nil {
		t.Errorf("HelloWorld(%q) = %q, %v", name, msg, err)
	}
}
