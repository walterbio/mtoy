package mtoy

import (
	"math/rand"
	"reflect"
	"testing"
	"time"
)

var (
	seed = time.Now().Unix()
)

func TestNew(t *testing.T) {
	want := &mtoy{
		answers: answers,
		seed:    seed,
	}

	got := New(seed)

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got: %v, want: %v", got, want)
	}
}

func TestRevealAnswer(t *testing.T) {
	rand.Seed(seed)
	want := answers[rand.Intn(len(answers))]
	got := New(seed).RevealAnswer()

	if want != got {
		t.Errorf("got: %s, want: %s", got, want)
	}
}
