package helpers

import (
	"fmt"
	"time"
)

func MeasureTime(description string) func() {
	start := time.Now()

	return func() {
		duration := time.Since(start)
		fmt.Printf("%s took %.3f seconds\n", description, duration.Seconds())
	}
}
