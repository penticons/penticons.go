package utils

import (
	"crypto/sha1"
	"fmt"
)

var Colors [5]string
var Background string
var Padding int
var Size int

func init() {
	Colors[0] = "#eeeeee"
	Colors[1] = "#d6e685"
	Colors[2] = "#8cc665"
	Colors[3] = "#44a340"
	Colors[4] = "#1e6823"

	Background = "#fefefe"

	Padding = 5
	Size = 30
}

func Hash(s string) string {
	h := sha1.New()
	h.Write([]byte(s))

	hash := h.Sum(nil)

	return fmt.Sprintf("%x", hash)
}

func Map(value int64) int {
	return int((4 * value) / 15)
}
