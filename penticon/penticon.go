// Package penticon deals with generating a penticon object
package penticon

import (
	"strconv"
	"github.com/pravj/penticons.go/svg"
	"github.com/pravj/penticons.go/utils"
)

// Penticon struct that contains Hash value, Tile representation and SVG object for penticon
type Penticon struct {
	Hash      string
	Svg       *svg.SVG
	Tile      [25]uint8
	TileColor [25]string
}

// New returns a new instance of Penticon struct
func New(arg string) *Penticon {
	return &Penticon{Hash: utils.Hash(arg), Svg: new(svg.SVG)}
}

// Svg_string returns string representation of Penticon SVG object
func (p *Penticon) Svg_string() string {
	p.generate_background()
	p.generate_foreground()

	return p.Svg.String()
}

// generate_background initialize the process by drawing a background
// according to provided dimension
func (p *Penticon) generate_background() {
        size := utils.PenticonSize

	p.Svg.Set_width(size)
	p.Svg.Set_height(size)

	args := make(map[string]string)
	args["fill"] = utils.Background

	p.Svg.Rect(0, 0, size, size, args)
}

// generate_foreground generated the penticon pattern
// by calling required methods one-by-one
func (p *Penticon) generate_foreground() {
	p.mark_tiles()
	p.fill_tiles()
	p.reflect_tiles()

	p.place_tiles()
}

// mark_tiles decides which tile to color or which one to leave default
func (p *Penticon) mark_tiles() {
	i := 0

	for i < 15 {
		p.Tile[i] = p.Hash[i] >> uint(8-((i%8)+1)) & 1
		i = i + 1
	}
}

// fill_tiles colors all the suitable tiles
func (p *Penticon) fill_tiles() {
	i := 0

	for i < 15 {
		if p.Tile[i] == 1 {
			integer, _ := strconv.ParseInt(string(p.Hash[i]), 16, 0)
			p.TileColor[i] = utils.Colors[int(utils.Map(integer))]
		} else {
			p.TileColor[i] = utils.Colors[0]
		}
		i = i + 1
	}
}

// reflect_tiles reflects left two columns of tiles to the right-most two columns
func (p *Penticon) reflect_tiles() {
	i := 15

	for i < 25 {
		p.TileColor[i] = p.TileColor[i-(10*(i/10))]
		i = i + 1
	}
}

// place_tiles places all the tiles to the place where they should be
func (p *Penticon) place_tiles() {
	padding := utils.Padding
	size := utils.TileSize

	i := 0
	for i < 25 {
		args := make(map[string]string)
		args["fill"] = p.TileColor[i]

		p.Svg.Rect(padding+((size+padding)*(i/5)), padding+((size+padding)*(i%5)), size, size, args)
		i = i + 1
	}
}
