package utils

// TODO: Improve it better

func CalculateScore(value float64) (int, string) {
	switch {
	case value < 500:
		return 8, "Bro, are you hacking ?"
	case value < 1000:
		return 5, "Almost there, try again later :)"
	case value < 2000:
		return 4, "Maybe tomorrow"
	case value < 3000:
		return 3, "Yeah, one day maybe"
	case value < 10000:
		return 2, "Did you miss some Geography classes ?"
	case value >= 10000:
		return 1, "Bro... ?"
	default:
		return 0, ""
	}
}
