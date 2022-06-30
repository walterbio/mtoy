package main

import (
	"bytes"
	"os"
	"os/exec"
	"testing"
	"time"

	"github.com/walterbio/mtoy"
)

func TestRevealAnswer(t *testing.T) {
	want := mtoy.New(time.Now().Unix()).RevealAnswer() + "\n"
	var out bytes.Buffer
	revealAnswer(&out)
	got := out.String()
	if got != want {
		t.Errorf("got: %s, want: %s", got, want)
	}
}

func TestMainFailWhenNoArgs(t *testing.T) {
	cmd := exec.Command(os.Args[0], "-test.run=TestMainFailWhenNoArgs")
	t.Log(len(os.Args))
	t.Log(os.Args)
	err := cmd.Run()
	e, ok := err.(*exec.ExitError)
	if ok {
		if code := e.ExitCode(); code != 1 {
			t.Errorf("main should exit with ExitCode == 1, got: %d", code)
		}

	}
	if !ok {
		t.Error("main should fail when no arguments")
	}
}
