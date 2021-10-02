package sio

import (
	"flag"
	"errors"
)

type Input struct {
	Arg1 string
}

func (inpt Input) validate() error {
	if len(inpt.Arg1) == 0 {
		return errors.New("first argument is missing")
	}

	return nil
}

func Read() (Input, error) {
	input := parseFlags()

	if err := input.validate(); err != nil {
		return emptyInput(), err
	}

	return input, nil
}

func emptyInput() Input {
	input := Input{
		Arg1: "",
	}

	return input
}

func parseFlags() Input {
	flag.Parse()

	arg1 := flag.Arg(0)
	input := Input{
		Arg1: arg1,
	}

	return input
}
