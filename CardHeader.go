package bs

import (
	"github.com/gouniverse/hb"
)

func CardHeader() *hb.Tag {
	return hb.NewDiv().Class("card-header")
}
