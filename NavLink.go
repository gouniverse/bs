package bs

import (
	"github.com/gouniverse/hb"
)

func NavLink() *hb.Tag {
	return hb.NewHyperlink().Class("nav-link")
}
