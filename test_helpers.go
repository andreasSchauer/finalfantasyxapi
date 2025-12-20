package main

import (
	"fmt"
	"testing"
)


func getTestName(name, requestURL string, caseNum int) string {
	return fmt.Sprintf("%s: %d, requestURL: %s", name, caseNum, requestURL)
}


// can also convert into a testing function
func containsAllResources[T HasAPIResource](cfg *Config, resources []T, expectedPaths []string) bool {
	resourceMap := getResourceMap(resources)

	for _, path := range expectedPaths {
		url := cfg.completeURL(path)
		_, ok := resourceMap[url]
		if !ok {
			return false
		}
	}

	return true
}

func getResourceAmountMap[T ResourceAmount](items []T) map[string]int32 {
	amountMap := make(map[string]int32)

	for _, item := range items {
		key := item.GetName()
		amountMap[key] = item.GetVal()
	}

	return amountMap
}

func hasExpectedLength[T HasAPIResource](resources []T, expected int) bool {
	return len(resources) == expected
}


func testResourceAmount[T ResourceAmount](t *testing.T, testName, resName string, resSlice []T, expAmounts map[string]int32) {
	gotResAmount := getResourceAmountMap(resSlice)
	for key, exp := range expAmounts {
		got, ok := gotResAmount[key]
		if !ok {
			t.Fatalf("%s: %s doesn't contain resource %s", testName, resName, key)
		}
		testInt32(t, testName, key, exp, got)
	}
}

func testResourceMatch[T HasAPIResource](t *testing.T, cfg *Config, testName, fieldName, expPath string, gotRes T) {
	gotURL := gotRes.getAPIResource().getURL()
	expURL := cfg.completeURL(expPath)

	if expURL != gotURL {
		t.Fatalf("%s: expected %s %s, got %s", testName, fieldName, expURL, gotURL)
	}
}

func testResourcePtrMatch[T HasAPIResource](t *testing.T, cfg *Config, testName, fieldName string, expPathPtr *string, gotResPtr *T) {
	if expPathPtr == nil {
		if gotResPtr == nil {
			return
		}
		res := *gotResPtr
		gotURL := res.getAPIResource().getURL()
		t.Fatalf("%s: expected nil for %s, but got %s", testName, fieldName, gotURL)
	}

	gotRes := *gotResPtr
	expPath := *expPathPtr

	testResourceMatch(t, cfg, testName, fieldName, expPath, gotRes)
}

func testInt32(t *testing.T, testName, fieldName string, expInt, gotInt int32) {
	if expInt != gotInt {
		t.Fatalf("%s: expected %s %d, got %d", testName, fieldName, expInt, gotInt)
	}
}

func testFloat32(t *testing.T, testName, fieldName string, expFloat, gotFloat float32) {
	if expFloat != gotFloat {
		t.Fatalf("%s: expected %s %.2f, got %.2f", testName, fieldName, expFloat, gotFloat)
	}
}

func testFloat32Ptr(t *testing.T, testName, fieldName string, expFloatPtr, gotFloatPtr *float32) {
	if expFloatPtr == nil {
		if gotFloatPtr == nil {
			return
		}
		t.Fatalf("%s: expected nil for %s, but got %.2f", testName, fieldName, *gotFloatPtr)
	}

	testFloat32(t, testName, fieldName, *expFloatPtr, *gotFloatPtr)
}

func testString(t *testing.T, testName, fieldName, expStr, gotStr string) {
	if expStr != "" && expStr != gotStr {
		t.Fatalf("%s: expected %s %s, got %s", testName, fieldName, expStr, gotStr)
	}
}



func testPaginationURLs(t *testing.T, cfg *Config, testName, fieldName string, gotURLPtr, expPathPtr *string) {
	if expPathPtr == nil {
		if gotURLPtr == nil {
			return
		}
		gotURL := *gotURLPtr
		t.Fatalf("%s: expected nil for %s, but got %s", testName, fieldName, gotURL)
	}

	gotURL := *gotURLPtr
	expPath := *expPathPtr
	expURL := cfg.completeURL(expPath)

	testString(t, testName, fieldName, expURL, gotURL)
}
