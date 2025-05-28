package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

var (
	ErrDivisionByZero       = errors.New("division by zero")
	ErrUnsupportedOperation = errors.New("operation not supported")
	ErrOperandNotInt        = errors.New("operand not int")
	ErrNotEnoughArgs        = errors.New("not enough args")
)

func calculate(exp []string) (int, error) {
	if len(exp) != 3 {
		return 0, ErrNotEnoughArgs
	}

	i1, err := strconv.Atoi(exp[0])
	if err != nil {
		return 0, ErrOperandNotInt
	}
	i2, err := strconv.Atoi(exp[2])
	if err != nil {
		return 0, ErrOperandNotInt
	}

	switch exp[1] {
	case "+":
		return i1 + i2, nil
	case "-":
		return i1 - i2, nil
	case "*":
		return i1 * i2, nil
	case "/":
		if i2 == 0 {
			return 0, ErrDivisionByZero
		}
		return i1 / i2, nil
	default:
		return 0, ErrUnsupportedOperation
	}

}

func fileLen(filename string) (int, error) {
	content, err := os.ReadFile(filename)
	return len(content), err
}

func prefixer(prefix string) func(string) string {
	return func(name string) string {
		return fmt.Sprintf("%s %s", prefix, name)
	}
}

func main() {

}
