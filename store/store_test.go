package store

import (
	"strings"
	"testing"
)

func TestNoInitFunc(t *testing.T) {
	_, err := NewStore("local;127.0.0.1:3232/")
	if err == nil {
		t.Error("unexpected result")
	}
	if !strings.HasPrefix(err.Error(), "storeinit") {
		t.Error("unexpected result")
	}

}
