package store

import (
	"context"
	"errors"
	"fmt"

	"github.com/przebro/databazaar/collection"
)

type initializerFunc func(ConnectionOptions) (DataStore, error)

var storeInitializer map[string]initializerFunc = map[string]initializerFunc{}

const (
	argUsername = "username"
	argPassword = "password"
	argDatabase = "database"
	argAuth     = "auth"
	argCert     = "cert"
	fileScheme  = "file"
)

var (
	errHostEmpty   = errors.New("hostname cannot be empty")
	errInvalidPort = errors.New("invalid port number")
	errBuildOpt    = errors.New("failed to build connection options")

	connArgs = []string{argUsername, argPassword, argDatabase, argCert, argAuth}
)

func RegisterStoreFactory(name string, initfunc initializerFunc) {

	storeInitializer[name] = initfunc
}

type DataStore interface {
	CreateCollection(ctx context.Context, name string) (collection.DataCollection, error)
	Collection(context.Context, string) (collection.DataCollection, error)
	Status(context.Context) (string, error)
	Close(ctx context.Context)
}

func NewStore(connection string) (DataStore, error) {

	var initStoreFunc initializerFunc
	var exists bool

	opt, err := BuildOptions(connection)
	if err != nil {
		return nil, err
	}

	if initStoreFunc, exists = storeInitializer[opt.Scheme]; !exists {
		return nil, fmt.Errorf("storeinit unable to find:%s initalizer", opt.Scheme)
	}

	return initStoreFunc(opt)
}
