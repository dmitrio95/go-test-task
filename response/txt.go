package response

import "fmt"

type simpleFormatter struct{}

func (simpleFormatter) FormatResponse(data ResponseData) (string, error) {
	return fmt.Sprint(data), nil
}

func newSimpleFormatter() ResponseFormatter {
	return &simpleFormatter{}
}
