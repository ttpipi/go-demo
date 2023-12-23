package main

import (
	"fmt"

	"github.com/pkg/errors"
)

func main() {
	fmt.Printf("err: %+v", c())
}

func c() error {
	return b()
}

func b() error {
	return a()
}

func a() error {
	fmt.Errorf("xxx")
	return errors.Wrap(errors.New("xxx"), "test")
}
