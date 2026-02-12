package api

import (
	"fmt"

	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
)

func getUsageString[T h.HasID, R any, A APIResource, L APIResourceList](i handlerInput[T, R, A, L]) string {
	listUsage := getUsagePath("", i.endpoint)
	usageStrings := []string{listUsage}

	for _, s := range i.usage {
		usageString := getUsagePath(s, i.endpoint)
		usageStrings = append(usageStrings, usageString)
	}

	usage := h.FormatStringSlice(usageStrings)

	if i.subsections != nil {
		sectionUsage := getUsagePath("/id/{subsection}", i.endpoint)
		sectionStr := h.GetMapKeyStr(i.subsections)
		usage = fmt.Sprintf("%s, '%s'. supported subsections: %s.", usage, sectionUsage, sectionStr)
	}

	return usage
}

func getUsagePath(usage, endpoint string) string {
	return fmt.Sprintf("/api/%s%s", endpoint, usage)
}
