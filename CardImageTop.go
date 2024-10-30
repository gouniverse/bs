package bs

import (
	"github.com/gouniverse/hb"
)

func CardImageTop(src string) *hb.Tag {
	return hb.Image(src).Class("card-image-top")
}
