package util

import (
	"strings"

	"github.com/google/uuid"
)

func First[T any](defaultArg T, args []T) T {
	if len(args) > 0 {
		defaultArg = args[0]
	}
	return defaultArg
}

func UUID(n ...int) string {
	len := First(16, n)
	uuidStr := uuid.New().String()
	uuidStr = strings.ReplaceAll(uuidStr, "-", "")
	return uuidStr[0:len]
}
