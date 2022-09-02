// mtoy implements Magic Ball 8 Game
package mtoy

import (
	"math/rand"
	"time"
)

type mtoy struct {
	seed    int64
	answers []string
}

// New returns a mtoy instance that can be used to customize the seed receives a
// seed and an optional list of answers to be used instead of the default list
// of answers.
func New(seed int64, answers ...string) mtoy {
	// TODO: is there another way to do this?
	// to support custom answers and parallel tests I can't use a reference to
	// default_answers variable
	// NOTE: Copy copies the minimun of len(dst) and len(src), if you create a
	// slice with cap: 0, no elements will be copied.
	local_answers := make([]string, len(default_answers))
	copy(local_answers, default_answers)

	if len(answers) > 0 {
		local_answers = answers
	}

	return mtoy{
		seed:    seed,
		answers: local_answers,
	}
}

// RevealAnswer returns a randomized answer based on Magic 8 Ball game
func (m mtoy) RevealAnswer() string {
	rand.Seed(m.seed)
	return m.answers[rand.Intn(len(m.answers))]
}

// RevealAnswer returns a randomized answer based on Magic 8 Ball game
func RevealAnswer() string {
	m := New(time.Now().Unix())
	return m.RevealAnswer()
}
