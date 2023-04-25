package bs

import (
	"github.com/gouniverse/hb"
)

func FormSelectOption(optionKey string, optionName string) *hb.Tag {
	return hb.NewOption().Attr("value", optionKey).HTML(optionName)
}
