package bs

import (
	"github.com/gouniverse/hb"
)

func FormSelect() *hb.Tag {
	return hb.Select().Class("form-select")
}
