package mtoy

import (
	"math/rand"
	"time"
)

var (
	answers = []string{
		"Ask again later.",
		"As I see it, yes.",
		"Better not tell you now.",
		"Cannot predict now.",
		"Concentrate and ask again.",
		"Don't count on it.",
		"It is certain.",
		"It is decidedly so.",
		"Most likely.",
		"My reply is no.",
		"My sources say no.",
		"Outlook good.",
		"Outlook not so good.",
		"Reply hazy, try again.",
		"Signs point to yes.",
		"Very doubtful. ",
		"Without a doubt.",
		"Yes.",
		"Yes definitely.",
		"You may rely on it.",
	}
)

type mtoy struct {
	seed int64
}

// to be called by the tests
func New(seed int64) mtoy {
	return mtoy{
		seed: seed,
	}
}

func (m mtoy) RevealAnswer() string {
	rand.Seed(m.seed)
	return answers[rand.Intn(len(answers))]
}

// Default behaviour
func RevealAnswer() string {
	m := New(time.Now().Unix())
	return m.RevealAnswer()
}
