package main

import (
	"errors"
	"fmt"
)

func f(arg int) (int, error) {
	if arg == 42 {
		return -1, errors.New("42 error")
	}

	return arg + 3, nil
}

var ErrOutOfTea = fmt.Errorf("no more tea variable")
var ErrPower = fmt.Errorf("cant boil water")

func makeTea(arg int) error {
	if arg == 2 {
		return ErrOutOfTea
	} else if arg == 4 {
		return fmt.Errorf("making tea: %w", ErrPower)
	}
	return nil
}

func main() {
	for _, i := range []int{7, 42} {
		if r, e := f(i); e != nil {
			fmt.Println("f failed: ", e)
		} else {
			fmt.Println("f worked: ", r)
		}
	}

	for i := range 5 {
		if err := makeTea(i); err != nil {
			if errors.Is(err, ErrOutOfTea) {
				fmt.Println("buy new tea!")
			} else if errors.Is(err, ErrPower) {
				fmt.Println("now it is dark")
			} else {
				fmt.Printf("unknown error: %s\n", err)
			}
			continue
		}
		fmt.Println("tea is ready")
	}
}
