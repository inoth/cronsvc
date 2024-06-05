package util

import (
	"strings"

	"github.com/google/uuid"
)

func UUID(n ...int) string {
	len := First(16, n)
	uuidStr := uuid.New().String()
	uuidStr = strings.ReplaceAll(uuidStr, "-", "")
	return uuidStr[0:len]
}
