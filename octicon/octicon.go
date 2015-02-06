package octicon

import (
 //"fmt"
 "github.com/pravj/octicons.go/utils"
)

type Octicon struct {
 Hash string
}

func New(arg string) *Octicon {
 return &Octicon{Hash: utils.Hash(arg)}
}
