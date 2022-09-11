package bs

import (
	"github.com/gouniverse/hb"
)

func FormText(text string) *hb.Tag {
	return hb.NewDiv().Class("form-text").HTML(text)
}
