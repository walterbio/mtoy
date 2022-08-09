package mtoy_test

import (
	"testing"

	"github.com/walterbio/mtoy"
)

// How can I maliciusly pass this test?
func TestRevealAnswer(t *testing.T) {
	t.Parallel()

	m := mtoy.New(42)
	want := "Don't count on it."
	got := m.RevealAnswer()
	if want != got {
		t.Errorf("got: %s, want: %s", got, want)
	}

	m = mtoy.New(1)
	want2 := "As I see it, yes."
	got2 := m.RevealAnswer()
	if want2 != got2 {
		t.Errorf("got: %s, want: %s", got2, want2)
	}
}

func TestRevealAnswer2(t *testing.T) {
	t.Parallel()

	m := mtoy.New(42)
	want := "Don't count on it."
	got := m.RevealAnswer()
	if want != got {
		t.Errorf("got: %s, want: %s", got, want)
	}

	m = mtoy.New(1)
	want2 := "As I see it, yes."
	got2 := m.RevealAnswer()
	if want2 != got2 {
		t.Errorf("got: %s, want: %s", got2, want2)
	}
}

// how can we see the state of random generator without
// accessing it?
//
//
