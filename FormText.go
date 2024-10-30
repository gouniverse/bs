package bs

import (
	"github.com/gouniverse/hb"
)

func FormText(text string) hb.TagInterface {
	return hb.Div().Class("form-text").HTML(text)
}
