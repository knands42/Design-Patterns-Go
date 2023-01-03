package breaking_the_principle

import "testing"

func Test_UseItRectangle(t *testing.T) {
	rc := Rectangle{2, 3}
	UseIt(&rc)
	if rc.width != 2 {
		t.Error("Expected the width to remain unchanged")
	}
	if rc.height != 10 {
		t.Error("Expected the height to change")
	}
}

func Test_UseItSquare(t *testing.T) {
	sq := Square{Rectangle{5, 5}}
	UseIt(&sq)
	if sq.width != 5 {
		t.Error("Expected the width to remain unchanged")
	}
	if sq.height != 10 {
		t.Error("Expected the height to change")
	}
}
