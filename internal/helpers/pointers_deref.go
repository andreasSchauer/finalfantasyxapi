package helpers

func DerefStringPtr(s *string) string {
	if s == nil {
		return ""
	}

	return *s
}