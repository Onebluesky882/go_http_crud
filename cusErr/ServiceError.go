package cusErr

import "fmt"

func Handle(err error, msg string) error {
	if err != nil {
		return fmt.Errorf("%s: %w", msg, err) // wrap original error
	}
	return nil
}
