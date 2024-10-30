package bs

import (
	"github.com/gouniverse/hb"
)

func NavTabs() *hb.Tag {
	return hb.UL().Class("nav nav-tabs")
}
