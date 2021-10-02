package sio

import (
	"fmt"
)

type Output = []string

const lineSeparator = "\n"
const arraySeparator = ", "

func Write(output Output) error {
	for _, value := range output {
		fmt.Println(value)
	}

	return nil
}

func boolToStr(v bool) string {
	if v {
		return "yes"
	}

	return "no"
}
