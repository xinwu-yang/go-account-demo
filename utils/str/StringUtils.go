package str

func IsEmpty(text string) bool {
	if &text == nil || text == "" {
		return false
	}
	return true
}
