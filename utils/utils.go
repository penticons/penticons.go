// Package utils implements some utility functions
package utils

import (
	"crypto/sha1"
	"fmt"
)

// Variables representing Tile colors, background color
// Padding between tiles and size of tile and penticon
var (
	Colors       [5]string
	Background   string
	Padding      int
	TileSize     int
	PenticonSize int
)

// Default values of the colors and dimension variables
func init() {
	Colors[0] = "#eeeeee"
	Colors[1] = "#d6e685"
	Colors[2] = "#8cc665"
	Colors[3] = "#44a340"
	Colors[4] = "#1e6823"

	Background = "#fefefe"

	Padding = 5
	TileSize = 30
	PenticonSize = (5 * TileSize) + 30
}

// Hash returns SHA-1 encrypted value for a string
func Hash(s string) string {
	h := sha1.New()
	h.Write([]byte(s))

	hash := h.Sum(nil)

	return fmt.Sprintf("%x", hash)
}

// Map returns corresponding value of a number from a range to different range
func Map(value int64) int {
	return int((4 * value) / 15)
}
