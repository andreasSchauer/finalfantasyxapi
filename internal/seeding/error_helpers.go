package seeding

import "fmt"

func getErr(obj, err error) error {
	return fmt.Errorf("%v: %v", obj, err)
}

func getDbErr(obj, err error, msg string) error {
	return fmt.Errorf("%v: %s: %v", obj, msg, err)
}

func getStrErr(s string, err error, msg string) error {
	return fmt.Errorf("%s %s: %v", s, msg, err)
}
