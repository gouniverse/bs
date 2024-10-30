package bs

import (
	"github.com/gouniverse/hb"
)

func NavTabs() hb.TagInterface {
	return hb.UL().Class("nav nav-tabs")
}
