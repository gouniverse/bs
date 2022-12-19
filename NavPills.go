package bs

import (
	"github.com/gouniverse/hb"
)

func NavPills() *hb.Tag {
	return hb.NewUL().Class("nav nav-pills")
}
