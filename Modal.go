package bs

import (
	"github.com/gouniverse/hb"
)

func Modal() *hb.Tag {
	return hb.NewDiv().Class("modal")
}

func ModalDialog() *hb.Tag {
	return hb.NewDiv().Class("modal-dialog")
}

func ModalContent() *hb.Tag {
	return hb.NewDiv().Class("modal-content")
}

func ModalHeader() *hb.Tag {
	return hb.NewDiv().Class("modal-header")
}

func ModalBody() *hb.Tag {
	return hb.NewDiv().Class("modal-body")
}

func ModalFooter() *hb.Tag {
	return hb.NewDiv().Class("modal-footer")
}
