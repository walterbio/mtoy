package mtoy

import (
	"testing"
	"time"
)

var (
	seed = time.Now().Unix()
)

func TestRevealAnswer(t *testing.T) {
	mtoy := New(seed)
	got := mtoy.RevealAnswer()

	if !mtoy.isMtoyAnswer(got) {
		t.Errorf("Got: %s, not valid mtoy answer", got)
	}
}

func TestIsMtoyAnswer(t *testing.T) {
	mtoy := New(seed)
	answer := mtoy.RevealAnswer()
	want := true
	got := mtoy.isMtoyAnswer(answer)

	if got != want {
		t.Errorf("want: %t, got: %t, answer: %s\n", want, got, answer)
	}
}

func TestIsNotMtoyAnswer(t *testing.T) {
	mtoy := New(seed)
	answer := "ccccccngfdrjuknlunrgcljdtnrhdifunvfhnvjbrrkn"
	want := false
	got := mtoy.isMtoyAnswer(answer)

	if got != want {
		t.Errorf("want: %t, got: %t, answer: %s\n", want, got, answer)
	}
}
