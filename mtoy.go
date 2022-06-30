package mtoy

import (
	"math/rand"
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
	answers []string
	seed    int64
}

// New returns a reference to an mtoy instance
func New(seed int64) *mtoy {
	return &mtoy{
		answers: answers,
		seed:    seed,
	}
}

// RevealAnswer returns a ramdom string from mtoy posible list of answers.
func (m *mtoy) RevealAnswer() string {
	rand.Seed(m.seed)
	return m.answers[rand.Intn(len(m.answers))]
}
