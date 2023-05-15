package domain

import (
	"fmt"
	"time"
)

type RangeError struct {
	Name string
	Min  int32
	Max  int32
}

func (e *RangeError) Error() string {
	return fmt.Sprintf("Range Error: %s must be between %d and %d.\n", e.Name, e.Min, e.Max)
}

type AlreadyExistsError struct {
	Class string
	Name  string
}

func (e *AlreadyExistsError) Error() string {
	return fmt.Sprintf("Already Exists Error: %s %s already exists.\n", e.Class, e.Name)
}

type DoesNotExistsError struct {
	Class string
	Name  string
}

func (e *DoesNotExistsError) Error() string {
	return fmt.Sprintf("Does Not Exists Error: %s %s does not exists.\n", e.Class, e.Name)
}

type ExpiredError struct {
	Name        string
	ExpiredDate time.Time
}

func (e *ExpiredError) Error() string {
	return fmt.Sprintf("Expired Error: The date %s of %s was expired.\n", e.ExpiredDate, e.Name)
}
