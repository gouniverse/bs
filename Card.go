package bs

import (
	"github.com/gouniverse/hb"
)

func Card() *hb.Tag {
	return hb.NewUL().Class("card")
}
