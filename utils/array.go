package utils

func IsExistedInt(element int, array []int) bool {
	for _, e := range array {
		if element == e {
			return true
		}
	}
	return false
}

func IsExistedString(element string, array []string) bool {
	for _, e := range array {
		if element == e {
			return true
		}
	}
	return false
}
