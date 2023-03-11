package store

import (
	"strings"
	"testing"
)

func TestOptions(t *testing.T) {

	conStrings := []struct {
		addr string

		msg string
	}{
		{addr: "driver;127.0.0.1:2323;/path?option=abc", msg: ""},
		{addr: "driver;;/path?option=abc", msg: ""},
		{addr: "driver;127.0.0.1:2323;", msg: ""},
		{addr: "driver;;/path/?option=abc", msg: ""},
		{addr: ";127.0.0.1:2323;/path?option=abc", msg: errOptMissingDriver.Error()},
		{addr: "driver;127.0.0.1:ABCD;/path/?option=abc", msg: errOptInvalidPort.Error()},
		{addr: "driver;127.0.0.1:2323;/path/?option=", msg: "invalid connection string;invalid option"},
		{addr: "driver;127.0.0.1:2323;/path/?=option&", msg: "invalid connection string;invalid option"},
		{addr: "driver;127.0.0.1:2323;/path/?=&", msg: "invalid connection string;invalid option"},
		{addr: "driver;127.0.0.1:2323;/path/?&&&", msg: "invalid connection string;invalid option"},
		{addr: "driver;127.0.0.1:2323;/path/?opt=abc&opt=", msg: "invalid connection string;invalid option"},
		{addr: ";", msg: errOptInvalidFormat.Error()},
		{addr: "", msg: errOptEmptyString.Error()},
		{addr: "abcdef", msg: errOptInvalidFormat.Error()},
	}

	for x := range conStrings {

		_, err := BuildOptions(conStrings[x].addr)

		if err == nil && conStrings[x].msg != "" {
			t.Error("unexpected error:", err, "test:", x+1)
		}
		if err != nil && !strings.HasPrefix(err.Error(), conStrings[x].msg) {
			t.Error("unexpected error:", err, "test:", x+1)
		}
	}
}
