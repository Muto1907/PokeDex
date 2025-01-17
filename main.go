package main

import (
	"time"

	"github.com/Muto1907/PokeDex/internal"
	"github.com/Muto1907/PokeDex/repl"
)

func main() {
	config := &internal.Config{
		Next:     nil,
		Previous: nil,
		Client:   internal.NewClient(5*time.Minute, 5*time.Second),
		PokeDex:  make(map[string]internal.Pokemon),
	}
	repl.StartREPL(config)
}
