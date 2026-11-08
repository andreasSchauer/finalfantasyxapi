package errs

import (
	"fmt"


	ts "github.com/andreasSchauer/finalfantasyxapi/internal/typestrings"
)

func FieldErrorMsg(fieldName ts.FieldName, msg string) string {
	return fmt.Sprintf("field '%s': %s", fieldName, msg)
}

func GetFieldRangeErrStart(val int32, fieldName ts.FieldName, intType string, isArray bool) string {
	errStart := fmt.Sprintf("%s '%d'", intType, val)

	if !isArray {
		errStart += fmt.Sprintf(" of field '%s'", fieldName)
	}

	errStart += " is out of range"

	return errStart
}
