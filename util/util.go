package util

func PointerToString(s *string) string {
	if s != nil {
		return *s
	}
	return ""
}
