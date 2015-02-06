package octicon

import (
 "strconv"
 //"fmt"
 "github.com/pravj/octicons.go/utils"
 "github.com/pravj/octicons.go/svg"
)

type Octicon struct {
 Hash string
 Svg *svg.SVG
 Tile [25]uint8
 TileColor [25]string
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
 o.mark_tiles()
 o.fill_tiles()
 o.reflect_tiles()

 o.place_tiles()
}

func (o *Octicon) mark_tiles() {
 i := 0

 for i < 15 {
  o.Tile[i] = o.Hash[i]>>uint(8-((i%8)+1))&1
  i = i + 1
 }
}

func (o *Octicon) fill_tiles() {
 i := 0

 for i < 15 {
  if o.Tile[i] == 1 {
   integer, _ := strconv.ParseInt(string(o.Hash[i]), 16, 0)
   o.TileColor[i] = utils.Colors[int(utils.Map(integer))]
  } else {
   o.TileColor[i] = utils.Colors[0]
  }
  i = i + 1
 }
}

func (o *Octicon) reflect_tiles() {
 i := 15

 for i < 20 {
  o.TileColor[i] = o.TileColor[i - 10]
  i = i + 1
 }

 for i < 25 {
  o.TileColor[i] = o.TileColor[i - 20]
  i = i + 1
 }
}

func (o *Octicon) place_tiles() {
 padding := 5
 i := 0

 for i < 25 {
  args := make(map[string]string)
  args["fill"] = o.TileColor[i]

  o.Svg.Rect(padding + (35*(i/5)), padding + (35*(i%5)), 30, 30, args)
  i = i + 1
 }
}
