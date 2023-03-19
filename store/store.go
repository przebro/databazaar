package store

import (
	"context"
	"fmt"

	"github.com/przebro/databazaar/collection"
)

type initializerFunc func(ConnectionOptions) (DataStore, error)

var storeInitializer map[string]initializerFunc = map[string]initializerFunc{}

func RegisterStoreFactory(name string, initfunc initializerFunc) {

	storeInitializer[name] = initfunc
}

// DataStore - common interface that represents a database
type DataStore interface {
	CreateCollection(ctx context.Context, name string) (collection.DataCollection, error)
	Collection(context.Context, string) (collection.DataCollection, error)
	Status(context.Context) (string, error)
	CollectionExists(ctx context.Context, name string) bool
	Close(ctx context.Context)
}

// NewStore - Creates a new store
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
