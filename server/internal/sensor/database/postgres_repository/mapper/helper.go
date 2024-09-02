package mapper

func stringToPointer(s string) *string {
	if s == "" {
		return nil
	}
	return &s
}

func pointerToString(s *string) string {
	if s == nil {
		return ""
	}
	return *s
}
