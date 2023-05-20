package coalesce

func Any[T comparable](values ...T) T {
	var zero T

	for _, value := range values {
		if value != zero {
			return value
		}
	}

	return zero
}

func String(ss ...string) string {
	for _, s := range ss {
		if s != "" {
			return s
		}
	}

	return ""
}
