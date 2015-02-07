package penticon

import (
	"strconv"
	//"fmt"
	"github.com/pravj/penticons.go/svg"
	"github.com/pravj/penticons.go/utils"
)

type Penticon struct {
	Hash      string
	Svg       *svg.SVG
	Tile      [25]uint8
	TileColor [25]string
}

func New(arg string) *Penticon {
	return &Penticon{Hash: utils.Hash(arg), Svg: new(svg.SVG)}
}

func (p *Penticon) Svg_string() string {
	p.generate_background()
	p.generate_foreground()

	return p.Svg.String()
}

func (p *Penticon) generate_background() {
	p.Svg.Set_width(180)
	p.Svg.Set_height(180)

	args := make(map[string]string)
	args["fill"] = utils.Background

	p.Svg.Rect(0, 0, 180, 180, args)
}

func (p *Penticon) generate_foreground() {
	p.mark_tiles()
	p.fill_tiles()
	p.reflect_tiles()

	p.place_tiles()
}

func (p *Penticon) mark_tiles() {
	i := 0

	for i < 15 {
		p.Tile[i] = p.Hash[i] >> uint(8-((i%8)+1)) & 1
		i = i + 1
	}
}

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

func (p *Penticon) reflect_tiles() {
	i := 15

	for i < 20 {
		p.TileColor[i] = p.TileColor[i-10]
		i = i + 1
	}

	for i < 25 {
		p.TileColor[i] = p.TileColor[i-20]
		i = i + 1
	}
}

func (p *Penticon) place_tiles() {
	padding := utils.Padding
	size := utils.Size

	i := 0
	for i < 25 {
		args := make(map[string]string)
		args["fill"] = p.TileColor[i]

		p.Svg.Rect(padding+((size+padding)*(i/5)), padding+((size+padding)*(i%5)), size, size, args)
		i = i + 1
	}
}
