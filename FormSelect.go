package bs

import (
	"github.com/gouniverse/hb"
)

func FormSelect() *hb.Tag {
	return hb.NewSelect().Class("form-select")
}
