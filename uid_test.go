package uid

import (
	"fmt"
	"testing"
)

func TestUID(t *testing.T) {
	for i := 0; i < 100; i++ {
		uid := New()
		info, err := Parse(uid)
		if err != nil {
			t.Fatal("bad uid", uid)
		}
		fmt.Println(uid, info)
	}

	if _, err := Parse("1234567"); err == nil {
		t.Error("bad parsed")
	}

	if parsed, err := Parse("V-B8xTRe2V2Zj2m$"); err == nil {
		t.Error("bad parsed:", parsed)
	}
}
