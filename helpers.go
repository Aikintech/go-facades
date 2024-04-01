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
	num := rand.Int63()

	numStr, _ := strconv.Atoi(strconv.FormatInt(num, 10)[0:n])

	return int64(numStr)
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
