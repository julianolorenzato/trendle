package fail

import "fmt"

type RangeError struct {
	Name string
	Min  int32
	Max  int32
}

func (e *RangeError) Error() string {
	return fmt.Sprintf("Range Error: %s must be between %d and %d.\n", e.Name, e.Min, e.Max)
}
