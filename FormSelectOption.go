package bs

import (
	"github.com/gouniverse/hb"
)

func FormSelectOption(name string, value string) *hb.Tag {
	return hb.NewOption().Attr("value", name).HTML(value)
}
