package bs

import (
	"github.com/gouniverse/hb"
)

func Alert() *hb.Tag {
	return hb.NewDiv().Class("alert")
}
