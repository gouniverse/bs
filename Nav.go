package bs

import (
	"github.com/gouniverse/hb"
)

func Nav() *hb.Tag {
	return hb.NewUL().Class("nav")
}
