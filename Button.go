package bs

import (
	"github.com/gouniverse/hb"
)

func Button() *hb.Tag {
	return hb.NewButton().Class("btn")
}
