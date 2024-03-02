package gofacades

import "fmt"

const (
	ALPHABETS = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	NUMBERS   = "1234567890"
)

var (
	ALPHA_NUM = fmt.Sprintf("%s%s", ALPHABETS, NUMBERS)
)
