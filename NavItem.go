package bs

import (
	"github.com/gouniverse/hb"
)

func NavItem() *hb.Tag {
	return hb.NewLI().Class("nav-item")
}
