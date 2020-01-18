package resolver

import (
	prompt "github.com/c-bata/go-prompt"
	"strings"
)

func GetExecutorWithSpaceSplit(r Resolver) prompt.Executor {
	return func(in string) {
		in = strings.TrimSpace(in)
		blocks := strings.Split(in, " ")
		r.Resolve(blocks...)
	}
}

func GetCompleterCompleterWithSpaceSplit(c Completer) prompt.Completer {
	return func(in prompt.Document) []prompt.Suggest {
		input := in.TextBeforeCursor()
		if input == "" {
			return []prompt.Suggest{}
		}
		blocks := strings.Split(input, " ")
		return c.Complete(blocks...)
	}
}
