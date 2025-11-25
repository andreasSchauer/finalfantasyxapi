package helpers

import (
	"fmt"
	"strings"
)

func GetErr(s string, err error, msgs ...string) error {
	if len(msgs) > 0 {
		msg := strings.Join(msgs, ": ")
		return fmt.Errorf("%s: %s: %v", s, msg, err)
	}

	return fmt.Errorf("%s: %v", s, err)
}

func JoinSubjects(subjects ...string) string {
	return strings.Join(subjects, ": ")
}