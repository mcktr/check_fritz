package fritzutil

import "testing"

func TestContains(t *testing.T) {
	s := []string{
		"Hello",
		"we",
		"just",
		"test",
		"here",
	}

	f := Contains(s, "Hello")

	if f != true {
		t.Errorf("Contains was incorrect, want: true, got %t", f)
	}

	f = Contains(s, "CatDog")

	if f != false {
		t.Errorf("Contains was incorrect, want: false, got %t", f)
	}

	f = Contains(s, "hello")

	if f != false {
		t.Errorf("Contains was incorrect, want: false, got %t", f)
	}
}
