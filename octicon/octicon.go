package octicon

import (
 //"fmt"
 "github.com/pravj/octicons.go/utils"
 "github.com/pravj/octicons.go/svg"
)

type Octicon struct {
 Hash string
 Svg *svg.SVG
}

func New(arg string) *Octicon {
 return &Octicon{Hash: utils.Hash(arg), Svg: new(svg.SVG)}
}

func (o *Octicon) Svg_string() string {
 o.generate_background()
 o.generate_foreground()

 return o.Svg.String()
}

func (o *Octicon) generate_background() {
 o.Svg.Set_width(180)
 o.Svg.Set_height(180)

 args := make(map[string]string)
 args["fill"] = utils.Background

 o.Svg.Rect(0, 0, 180, 180, args)
}

func (o *Octicon) generate_foreground() {
 //
}
