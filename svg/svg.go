package svg

import (
 "fmt"
)

type SVG struct {
 svg_string string
 width, height int
}

func (s *SVG) Set_width(w int) {
 s.width = w
}

func (s *SVG) Set_height(h int) {
 s.height = h
}

func (s *SVG) header() string {
 return fmt.Sprintf("<svg xmlns='http://www.w3.org/2000/svg' width='%v' height='%v'>", s.width, s.height)
}

func (s *SVG) footer() string {
 return "</svg>"
}

func (s *SVG) String() string {
 return s.header() + s.svg_string + s.footer()
}

func (s *SVG) Rect(x, y, w, h int, args map[string]string) {
 rect_str := fmt.Sprintf("<rect x='%v' y='%v' width='%v' height='%v' %s />", x, y, w, h, s.Write_args(args))
 s.svg_string += rect_str
}

func (s *SVG) Write_args(args map[string]string) string {
 str := ""

 for k, v := range args {
  str += fmt.Sprintf("%s='%s' ", k, v)
 }

 return str
}
