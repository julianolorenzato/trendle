package fail

import (
	"fmt"
	"time"
)

type ExpiredError struct {
	Name        string
	ExpiredDate time.Time
}

func (e *ExpiredError) Error() string {
	return fmt.Sprintf("Expired Error: The date %s of %s was expired.\n", e.ExpiredDate, e.Name)
}
