package gofacades

import (
	"math/rand"
	"strconv"
	"strings"
)

func GetFileExtension(mime string) string {
	split := strings.Split(mime, "/")
	if len(split) > 1 {
		return split[len(split)-1]
	}

	return ""
}

func GenerateRandomString(n uint, lowercase bool) string {
	return getRandomChars(ALPHABETS, n)
}

func GenerateRandomNumbers(n uint) int64 {
	num, _ := strconv.Atoi(getRandomChars(NUMBERS, n))

	return int64(num)
}

func getRandomChars(charset string, n uint) string {
	if n == 0 || charset == "" || len(charset) == 0 {
		return ""
	}

	b := make([]byte, n)
	for i := range b {
		b[i] = charset[rand.Intn(len(charset))]
	}
	return string(b)
}
