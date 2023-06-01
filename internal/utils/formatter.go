package utils

import (
	"fmt"
	"strings"
)

func Uint64sToString(arr []uint64, delimiter string) string {
	return strings.Trim(strings.ReplaceAll(fmt.Sprint(arr), " ", delimiter), "[]")
}
