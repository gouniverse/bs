package bs

import (
	"github.com/gouniverse/hb"
)

func FormTextArea() *hb.Tag {
	return hb.TextArea().Class("form-control")
}
