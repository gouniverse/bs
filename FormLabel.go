package bs

import "github.com/gouniverse/hb"

func FormLabel(label string) hb.TagInterface {
	return hb.Div().Class("form-label").HTML(label)
}
