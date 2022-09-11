package bs

import (
	"github.com/gouniverse/hb"
)

func NavTabs() *hb.Tag {
	return hb.NewLI().Class("nav nav-tabs")
}
