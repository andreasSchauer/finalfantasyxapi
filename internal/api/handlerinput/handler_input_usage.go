package handlerinput

import (
	"fmt"

	h "github.com/andreasSchauer/finalfantasyxapi/internal/general/helpers"
	ts "github.com/andreasSchauer/finalfantasyxapi/internal/general/typestrings"
)

func (i Main[T, R, A]) CompileUsage() string {
	listUsage := getUsagePath("", i.Endpoint)
	usageStrings := []string{listUsage}

	for _, s := range i.Usage {
		usageString := getUsagePath(s, i.Endpoint)
		usageStrings = append(usageStrings, usageString)
	}

	usage := h.FormatStringSlice(usageStrings)

	if i.Subsections != nil {
		sectionUsage := getUsagePath("/id/{subsection}", i.Endpoint)
		sectionStr := FormatSectionNames(i.Subsections)
		usage = fmt.Sprintf("%s, '%s'. supported subsections: %s", usage, sectionUsage, sectionStr)
	}

	return usage
}

func getUsagePath(usage string, endpoint ts.EndpointName) string {
	return fmt.Sprintf("/api/%s%s", endpoint, usage)
}
