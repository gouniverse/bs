package bs

import (
	"github.com/gouniverse/hb"
)

func Modal() hb.TagInterface {
	return hb.Div().Class("modal")
}

func ModalDialog() hb.TagInterface {
	return hb.Div().Class("modal-dialog")
}

func ModalContent() hb.TagInterface {
	return hb.Div().Class("modal-content")
}

func ModalHeader() hb.TagInterface {
	return hb.Div().Class("modal-header")
}

func ModalBody() hb.TagInterface {
	return hb.Div().Class("modal-body")
}

func ModalFooter() hb.TagInterface {
	return hb.Div().Class("modal-footer")
}
