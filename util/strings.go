package util

import (
	"strconv"
	"strings"
)

func check1stdigit(numStr, firstdigit string) bool {
	if strings.LastIndex(numStr, firstdigit) != len(numStr)-1 {
		return true
	}
	return false
}

// TODO: 1, 2, 3の時のsuffixがずれる（例：1nd, 3stなど）
func convNumNth(num int) string {
	var nth string
	numStr := strconv.Itoa(num)
	if numStr == "11" || numStr == "12" || numStr == "13" {
		nth = numStr + "th"
	} else if check1stdigit(numStr, "1") {
		nth = numStr + "st"
	} else if check1stdigit(numStr, "2") {
		nth = numStr + "nd"
	} else if check1stdigit(numStr, "3") {
		nth = numStr + "rd"
	} else {
		nth = numStr + "th"
	}
	return nth
}
