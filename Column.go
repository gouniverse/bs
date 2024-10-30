package bs

import (
	"github.com/gouniverse/hb"
	"github.com/gouniverse/utils"
)

func Column(width int) hb.TagInterface {
	return hb.Div().Class("col-sm-" + utils.ToString(width))
}
