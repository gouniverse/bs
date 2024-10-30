package bs

import (
	"github.com/gouniverse/hb"
)

func FormCheckRadio() hb.TagInterface {
	return hb.Input().Class("form-check-input").Type(hb.TYPE_RADIO)
}
