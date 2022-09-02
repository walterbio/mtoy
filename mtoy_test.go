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

func TestRevealAnswerCustomAnswers(t *testing.T) {
	t.Parallel()

	answers := []string{
		"Como yo lo veo, sí",
		"Concéntrate y pregunta otra vez",
		"Es cierto",
		"Es decididamente así",
		"Las señales apuntan a que sí",
		"Más probable",
		"Mejor no decirte ahora",
		"Mi respuesta es no",
		"Mis fuentes dicen que no",
		"Muy dudoso",
		"no cuentes con eso",
		"No se puede predecir ahora",
		"Perspectiva no tan buena",
		"Perspectivas buena",
		"Pregunta de nuevo más tarde",
		"Puedes confiar en ello",
		"Respuesta confusa, intenta otra vez",
		"Sí",
		"Sí definitivamente",
		"Sin duda",
	}

	m := mtoy.New(42, answers...)
	want := answers[5]
	got := m.RevealAnswer()
	if want != got {
		t.Errorf("got: %s, want: %s", got, want)
	}

	m = mtoy.New(1, answers...)
	want2 := answers[1]
	got2 := m.RevealAnswer()
	if want2 != got2 {
		t.Errorf("got: %s, want: %s", got2, want2)
	}
}
