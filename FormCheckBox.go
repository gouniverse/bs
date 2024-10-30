package bs

import (
	"github.com/gouniverse/hb"
)

func FormCheckBox() hb.TagInterface {
	return hb.Input().Class("form-check-input").Type(hb.TYPE_CHECKBOX)
}
