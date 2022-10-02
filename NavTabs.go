package bs

import (
	"github.com/gouniverse/hb"
)

func NavTabs() *hb.Tag {
	return hb.NewUL().Class("nav nav-tabs")
}
