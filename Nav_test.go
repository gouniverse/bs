package bs

import (
	"testing"
)

func TestNav(t *testing.T) {
	nav := Nav().ToHTML()
	if nav != "<ul class=\"nav\"></ul>" {
		t.Log("Expected: ", `<ul class="nav"></ul>`, " but found:", nav)
		t.Fail()
	}
}
