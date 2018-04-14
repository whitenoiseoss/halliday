package errors

type ErrorType int

type HallidayError struct {
	Type   ErrorType
	Detail string
}

func (he *HallidayError) Error() string {
	return he.Detail
}
