package helpers

import (
	"fmt"
	"strings"
)

func GetTypeName[T any]() string {
	var zeroType T
	typeString := fmt.Sprintf("%T", zeroType)
	typeOnly := strings.Split(typeString, ".")

	return typeOnly[len(typeOnly)-1]
}
