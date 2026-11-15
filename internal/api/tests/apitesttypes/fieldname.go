package apitesttypes

import "fmt"

func childFn(fn, childName string) string {
	return fmt.Sprintf("%s - %s", fn, childName)
}
