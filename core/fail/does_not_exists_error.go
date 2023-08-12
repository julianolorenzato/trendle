package fail

import "fmt"

type DoesNotExistsError struct {
	Class string
	Name  string
}

func (e *DoesNotExistsError) Error() string {
	return fmt.Sprintf("Does Not Exists Error: %s %s does not exists.\n", e.Class, e.Name)
}
