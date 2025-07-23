package scannerutil

func StringDifference(a, b string) float64 {
	// In a real implementation, we would use a more sophisticated algorithm
	// to calculate the difference between two strings.
	if a == b {
		return 0.0
	}
	return 1.0
}

func TruncateBody(body string) string {
	max := 4096
	if len(body) > max {
		return body[:max] + "... [truncated]"
	}
	return body
}
