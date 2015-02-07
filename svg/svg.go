// Package svg helps in working with SVG
package svg

import (
	"fmt"
)

// SVG structure that contains necessary elements like
// string representation, width and height
type SVG struct {
	svg_string    string
	width, height int
}

// Set_width sets a width for the SVG pattern
func (s *SVG) Set_width(w int) {
	s.width = w
}

// Set_height sets a height for the SVG pattern
func (s *SVG) Set_height(h int) {
	s.height = h
}

// header returns fix-header-string for a SVG, changes according to dimension of pattern
func (s *SVG) header() string {
	return fmt.Sprintf("<svg xmlns='http://www.w3.org/2000/svg' width='%v' height='%v'>", s.width, s.height)
}

// footer returns fix-footer-string for a SVG pattern
func (s *SVG) footer() string {
	return "</svg>"
}

// String returns string representation of the SVG pattern
func (s *SVG) String() string {
	return s.header() + s.svg_string + s.footer()
}

// Rect generates a Rectangle element SVG string and add it to the SVG pattern object
func (s *SVG) Rect(x, y, w, h int, args map[string]string) {
	rect_str := fmt.Sprintf("<rect x='%v' y='%v' width='%v' height='%v' %s />", x, y, w, h, s.Write_args(args))
	s.svg_string += rect_str
}

// Write_args adds additional attributes to any SVG element
func (s *SVG) Write_args(args map[string]string) string {
	str := ""

	for k, v := range args {
		str += fmt.Sprintf("%s='%s' ", k, v)
	}

	return str
}
