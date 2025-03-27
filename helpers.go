package hxsqlfiltering

import "strings"

func sprintf(format string, a ...string) string {
	var valueIndex int

	var builder strings.Builder

	builder.Grow(len(format)) // preallocate some capacity

	for i := 0; i < len(format); i++ {
		if i+1 < len(format) && format[i] == '%' && format[i+1] == 's' {
			if valueIndex < len(a) {
				builder.WriteString(a[valueIndex])

				valueIndex++
			}

			i++ // Skip 's' as %s is two characters

			continue
		}

		builder.WriteByte(format[i])
	}

	return builder.String()
}

func ternary[T any](condition bool, value1, value2 T) T {
	if condition {
		return value1
	}

	return value2
}
