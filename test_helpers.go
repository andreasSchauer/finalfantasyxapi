package main

import "testing"


func testResponseChecks(t *testing.T, testCfg *Config, testCase string, checks []testCheck, lenMap map[string]int) {
	for _, c := range checks {
		if len(c.expected) == 0 {
			continue
		}
		if !containsAllResources(testCfg, c.got, c.expected) {
		 	t.Fatalf("%s: %s doesn't contain all wanted resources of %v", testCase, c.name, c.expected)
		}

		expLen, ok := lenMap[c.name]
		if !ok {
			continue
		}

		if !hasExpectedLength(c.got, expLen) {
			t.Fatalf("%s: %s expected length: %d, got: %d", testCase, c.name, expLen, len(c.got))
		}
	}
}


func toIface[T HasAPIResource](in []T) []HasAPIResource {
	out := make([]HasAPIResource, len(in))
	for i, v := range in {
		out[i] = v
	}
	return out
}


func hasExpectedLength[T HasAPIResource](resources []T, expected int) bool {
	return len(resources) == expected
}


func resourcesMatch[T HasAPIResource](cfg *Config, resource T, expectedPath string) bool {
	resourceURL := resource.getAPIResource().getURL()
	expectedURL := cfg.completeURL(expectedPath)
	
	return resourceURL == expectedURL
}


func resourcePtrsMatch[T HasAPIResource](cfg *Config, resourcePtr *T, expectedPathPtr *string) bool {
	if resourcePtr == nil {
		return expectedPathPtr == nil
	}

	resource := *resourcePtr
	expectedPath := *expectedPathPtr

	resourceURL := resource.getAPIResource().getURL()
	expectedURL := cfg.completeURL(expectedPath)

	return resourceURL == expectedURL
}


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


func derefResourcePtr[T IsAPIResource](resourcePtr *T) T {
	if resourcePtr == nil {
		var zeroType T
		return zeroType
	}
	return *resourcePtr
}