package bs

import (
	"github.com/gouniverse/hb"
)

func FormInputPlaintext() *hb.Tag {
	return hb.NewInput().Class("form-control-plaintext")
}
