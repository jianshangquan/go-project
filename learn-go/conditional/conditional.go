package conditional

func IsPassed(score int) bool {
	if score >= 60 {
		return true
	} else {
		return false
	}
}