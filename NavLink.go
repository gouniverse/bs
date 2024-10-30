package bs

import (
	"github.com/gouniverse/hb"
)

func NavLink() hb.TagInterface {
	return hb.Hyperlink().Class("nav-link")
}
