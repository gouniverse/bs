package bs

import (
	"github.com/gouniverse/hb"
)

func ButtonLink() hb.TagInterface {
	return hb.Hyperlink().Class("btn")
}
