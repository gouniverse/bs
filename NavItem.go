package bs

import (
	"github.com/gouniverse/hb"
)

func NavItem() *hb.Tag {
	return hb.NewUL().Class("nav nav-tabs")
}
