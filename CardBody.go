package bs

import (
	"github.com/gouniverse/hb"
)

func CardBody() hb.TagInterface {
	return hb.NewDiv().Class("card-body")
}
