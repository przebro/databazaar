package store

import (
	"testing"
)

func TestOptions(t *testing.T) {

	conStrings := []struct {
		addr string
		err  error
	}{

		{addr: "driver;127.0.0.1:2323/path?option=abc", err: nil},
		{addr: ";127.0.0.1:2323/path?option=abc", err: errOptMissingDriver},
		{addr: "driver;127.0.0.1:ABCD//path/?option=abc", err: errOptInvalidPort},
		{addr: "driver;//path/?option=abc", err: errOptInvalidHost},
		{addr: "driver;127.0.0.1:2323/path/?option=", err: erroOptInvalidOption},
		{addr: "driver;127.0.0.1:2323/path/?=option&", err: erroOptInvalidOption},
		{addr: "driver;127.0.0.1:2323/path/?=&", err: erroOptInvalidOption},
		{addr: "driver;127.0.0.1:2323/path/?&&&", err: erroOptInvalidOption},
		{addr: "driver;127.0.0.1:2323/path/?opt=abc&opt=", err: erroOptInvalidOption},
		//{addr: "driver;127.0.0.1:2323", err: nil}, :TODO
	}

	for x := range conStrings {

		_, err := BuildOptions(conStrings[x].addr)
		if err != err {
			t.Error(err)
		}
	}

}
