package util

import "strconv"

func StringToUint(str string) (result uint, err error) {
	result64, err := strconv.ParseUint(str, 10, 64)
	return uint(result64), err
}
