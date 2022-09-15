package bs

import (
	"github.com/gouniverse/hb"
)

func CardTitle() *hb.Tag {
	return hb.NewHeading5().Class("card-title")
}
