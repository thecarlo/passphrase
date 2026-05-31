package main

import (
	"strconv"
)

func run() {
	wordCount := wf.Config.GetInt("WORD_COUNT", 3)

	args := wf.Args()
	if len(args) > 0 {
		if n, err := strconv.Atoi(args[0]); err == nil {
			if n < 2 {
				n = 2
			} else if n > 10 {
				n = 10
			}
			wordCount = n
		}
	}

	phrases, err := suggestions(wordCount)
	if err != nil {
		wf.FatalError(err)
		return
	}

	for _, phrase := range phrases {
		wf.NewItem(phrase).
			Subtitle("Press Enter to copy to clipboard").
			Arg(phrase).
			Valid(true)
	}

	wf.SendFeedback()
}
