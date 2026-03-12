package helpers

import (
    "cmp"
	"fmt"
    "slices"
    "strings"
)


func NameToString(name string, version *int32, spec *string) string {
	var verStr string
	var specStr string

	if version != nil {
		intVer := int(*version)
		verStr = fmt.Sprintf(" - %d", intVer)
	}

	if spec != nil {
		specStr = fmt.Sprintf(" (%s)", *spec)
	}

	return name + verStr + specStr
}

func NameAmountString(name string, version *int32, spec *string, amount int32) string {
	nameStr := NameToString(name, version, spec)

	if amount == 1 {
		return nameStr
	}

	return fmt.Sprintf("%s x%d", nameStr, amount)
}

func DerefVerStr(version *int32) string {
	if version == nil {
		return "null"
	}

	return fmt.Sprintf("%d", *version)
}


func GetMapKeyStr[T any](itemMap map[string]T) string {
	keys := []string{}

	for key := range itemMap {
		keyFormatted := fmt.Sprintf("'%s'", key)
		keys = append(keys, keyFormatted)
	}

	slices.SortStableFunc(keys, func(a, b string) int {
		return cmp.Compare(a, b)
	})

	return strings.Join(keys, ", ")
}


func GetNameWithSpaces(name, separator string) string {
	nameLower := strings.ToLower(name)
	return strings.ReplaceAll(nameLower, separator, " ")
}

func GetNameWithDashes(name, separator string) string {
	nameLower := strings.ToLower(name)
	return strings.ReplaceAll(nameLower, separator, "-")
}

func StringSliceToListString(s []string) string {
	return strings.Join(s, ", ")
}


func FormatStringSlice(items []string) string {
	formattedVals := []string{}

	for _, s := range items {
		formatted := fmt.Sprintf("'%s'", s)
		formattedVals = append(formattedVals, formatted)
	}

	return strings.Join(formattedVals, ", ")
}


func FormatIntSlice(IDs []int32) string {
	formattedIDs := []string{}

	for _, id := range IDs {
		formatted := fmt.Sprintf("'%d'", id)
		formattedIDs = append(formattedIDs, formatted)
	}

	return strings.Join(formattedIDs, ", ")
}
