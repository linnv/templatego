package util

import "strings"

func InIntSlice(target int, slice []int) bool {
	for i := 0; i < len(slice); i++ {
		if slice[i] == target {
			return true
		}
	}
	return false

}

func InInt64Slice(target int64, slice []int64) bool {
	for i := 0; i < len(slice); i++ {
		if slice[i] == target {
			return true
		}
	}
	return false

}

func RemoveNewLine(str string) string {
	if strings.Contains(str, "<audio") {
		return str
	}

	if strings.Contains(str, "<speak") {
		return str
	}

	items := strings.Split(str, "\n")
	var ret string
	for i := 0; i < len(items); i++ {
		items[i] = strings.TrimSpace(items[i])
		itemsiLen := len(items[i])
		if itemsiLen > 0 {
			if items[i][itemsiLen-1] == '\r' {
				items[i] = items[i][:itemsiLen-1]
			}
		}
		ret += items[i]
	}
	return ret
}
