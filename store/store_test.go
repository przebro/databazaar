package store

import (
	"strings"
	"testing"
)

func TestNoInitFunc(t *testing.T) {
	_, err := NewStore("local;127.0.0.1:3232;/")
	if err == nil {
		t.Error("unexpected result")
	}
	if !strings.HasPrefix(err.Error(), "storeinit") {
		t.Error("unexpected result")
	}

	_, err = NewStore(";127.0.0.1:3232;/")
	if err == nil {
		t.Error("unexpected result")
	}

	RegisterStoreFactory("local", func(ConnectionOptions) (DataStore, error) { return nil, nil })

	_, err = NewStore("local;127.0.0.1:3232;/")

	if err != nil {
		t.Error("unexpected result", err)
	}

}
func TestRegisterStoreFunc(t *testing.T) {

	RegisterStoreFactory("test", func(ConnectionOptions) (DataStore, error) { return nil, nil })

	if _, exists := storeInitializer["test"]; !exists {
		t.Error("register store factory func error")
	}

}
