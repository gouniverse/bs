package bs

import (
	"github.com/gouniverse/hb"
)

func FormSelect() hb.TagInterface {
	return hb.Select().Class("form-select")
}
