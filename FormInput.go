package bs

import (
	"github.com/gouniverse/hb"
)

func FormInput() *hb.Tag {
	return hb.NewInput().Class("form-control")
}
