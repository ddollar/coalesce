package coalesce

import "reflect"

func Any[T any](values ...T) T {
	var zero T

	for _, value := range values {
		if !reflect.DeepEqual(value, zero) {
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
