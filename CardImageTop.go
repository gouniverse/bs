package bs

import (
	"github.com/gouniverse/hb"
)

func CardImageTop(src string) hb.TagInterface {
	return hb.Image(src).Class("card-image-top")
}
