package bs

import "github.com/gouniverse/hb"

func FormLabel(label string) *hb.Tag {
	return hb.Div().Class("form-label").HTML(label)
}
