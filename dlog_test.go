package dlog

import (
	"bytes"
	"testing"
	"time"
)

func TestGenerate(t *testing.T) {
	var b bytes.Buffer

	time, _ := time.Parse(time.RFC3339, "2013-03-10T02:55:45Z")

	log := New(&b, map[string]interface{}{
		"t": time,
		"a": 1,
		"b": "some words",
	})
	log.SortOrder([]string{"t", "b"})

	log.Output(map[string]interface{}{
		"c": "there",
	})

	e := "t=2013-03-10T02:55:45Z b=\"some words\" a=1 c=there\n"
	if b.String() != e {
		t.Errorf("expected %q was %q", e, b.String())
	}
}
