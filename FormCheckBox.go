package bs

import (
	"github.com/gouniverse/hb"
)

func FormCheckbox() *hb.Tag {
	return hb.NewInput().Class("form-check-input").Type(hb.TYPE_CHECKBOX)
}
