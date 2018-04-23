package errors

import (
	"fmt"
)

// ErrorType is a uint8 that makes sure we stay clean when passing errors around
// There are 255 possible ErrorTypes
type ErrorType uint8

const (
	ExampleError ErrorType = iota
)

// HallidayError holds an ErrorType as well as details of the error
type HallidayError struct {
	Category string
	Type     ErrorType
	Detail   string
}

func (he *HallidayError) Error() string {
	return fmt.Sprintf("urn:halliday:%s:%d : %s", he.Category, he.Type, he.Detail)
}
