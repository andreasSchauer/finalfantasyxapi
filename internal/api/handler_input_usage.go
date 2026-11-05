package api

import (
	"fmt"

	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
	ts "github.com/andreasSchauer/finalfantasyxapi/internal/typestrings"
)

func (i handlerInput[T, R, A]) compileUsage() string {
	listUsage := getUsagePath("", i.Endpoint)
	usageStrings := []string{listUsage}

	for _, s := range i.usage {
		usageString := getUsagePath(s, i.Endpoint)
		usageStrings = append(usageStrings, usageString)
	}

	usage := h.FormatStringSlice(usageStrings)

	if i.subsections != nil {
		sectionUsage := getUsagePath("/id/{subsection}", i.Endpoint)
		sectionStr := formatSectionNames(i.subsections)
		usage = fmt.Sprintf("%s, '%s'. supported subsections: %s", usage, sectionUsage, sectionStr)
	}

	return usage
}

func getUsagePath(usage string, endpoint ts.EndpointName) string {
	return fmt.Sprintf("/api/%s%s", endpoint, usage)
}
