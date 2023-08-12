package fail

import "fmt"

type AlreadyExistsError struct {
	Class string
	Name  string
}

func (e *AlreadyExistsError) Error() string {
	return fmt.Sprintf("Already Exists Error: %s %s already exists.\n", e.Class, e.Name)
}
