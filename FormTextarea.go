package bs

import (
	"github.com/gouniverse/hb"
)

func FormTextArea() *hb.Tag {
	return hb.NewTextArea().Class("form-control")
}
