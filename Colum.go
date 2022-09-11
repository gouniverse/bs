package bs

import (
	"github.com/gouniverse/hb"
	"github.com/gouniverse/utils"
)

func Column(width int) *hb.Tag {
	return hb.NewDiv().Class("col-sm-" + utils.ToString(width))
}
