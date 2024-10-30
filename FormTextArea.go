package bs

import (
	"github.com/gouniverse/hb"
)

func FormTextArea() hb.TagInterface {
	return hb.TextArea().Class("form-control")
}
