package uid

import (
	"fmt"
	"testing"
)

func TestUID(t *testing.T) {
	for i := 0; i < 100; i++ {
		uid := New()
		info := Parse(uid)
		if info == nil {
			t.Fatal("bad uid", uid)
		}
		fmt.Println(uid, info)
	}

	if Parse("1234567") != nil {
		t.Error("bad parsed")
	}
	if result := Parse("V-B8xTRe2V2Zj2m$"); result != nil {
		t.Error("bad parsed:", result)
	}
}
