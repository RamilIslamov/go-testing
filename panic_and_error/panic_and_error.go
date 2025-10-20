package panic_and_error

import "errors"

func ValidateName(name string) error {
	var err error = nil
	if name == "" {
		err = errors.New("name is empty")
	}
	return err
}

func MustAt[T any](i int, xs []T) T {
	if i < 0 || i >= len(xs) {
		panic("index out of range")
	}
	return xs[i]
}
