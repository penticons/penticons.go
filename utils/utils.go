package utils

import (
 "fmt"
 "crypto/sha1"
)

var Colors [5]string
var Background string

func init() {
 Colors[0] = "#eeeeee"
 Colors[1] = "#d6e685"
 Colors[2] = "#8cc665"
 Colors[3] = "#44a340"
 Colors[4] = "#1e6823"

 Background = "#fefefe"
}

func Hash(s string) string {
 h := sha1.New()
 h.Write([]byte(s))

 hash := h.Sum(nil)

 return fmt.Sprintf("%x", hash)
}
