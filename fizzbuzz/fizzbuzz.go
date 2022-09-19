package fizzbazz

import (
	"errors"
	"strconv"
)

func fizzbazz(n int) (string, error) {
	if n <= 0 {
		return "", errors.New("input must be greater than 0")
	}
	if n%15 == 0 {
		return "FizzBuzz", nil
	}
	if n%5 == 0 {
		return "Buzz", nil
	}
	if n%3 == 0 {
		return "Fizz", nil
	}
	return strconv.Itoa(n), nil
}
