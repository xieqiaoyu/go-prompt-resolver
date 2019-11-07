package main

import (
	prompt "github.com/c-bata/go-prompt"
)

//Completer a completer can return complete suggest with string input
type Completer interface {
	Complete(...string) []prompt.Suggest
}

//Resolver resolver can decide  next step of a command
type Resolver interface {
	Resolve(...string)
}
