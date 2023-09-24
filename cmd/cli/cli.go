package main

import (
	CustomCommand "skeleton/app/command"
	_ "skeleton/internal/bootstrap"
	"skeleton/internal/command"
)

func main() {
	command.New().AddCommand(
		&CustomCommand.FooCommand{},
	).Execute()
}
