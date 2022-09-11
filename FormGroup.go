package bs

import (
	"github.com/gouniverse/hb"
)

func FormGroup() *hb.Tag {
	return hb.NewDiv().Class("form-group mt-3")
}
