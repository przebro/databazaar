package collection

import (
	"context"

	"github.com/przebro/databazaar/result"
	"github.com/przebro/databazaar/selector"
)

type BazaarDocument interface {
	ID() string
	Revision() string
}

// DataCollection - Common interface for database operations
type DataCollection interface {
	Create(ctx context.Context, document interface{}) (*result.BazaarResult, error)
	Get(ctx context.Context, id string, result interface{}) error
	Update(ctx context.Context, doc interface{}) error
	Delete(ctx context.Context, id string) error
	CreateMany(ctx context.Context, docs []interface{}) ([]result.BazaarResult, error)
	BulkUpdate(ctx context.Context, docs []interface{}) error
	All(ctx context.Context) (BazaarCursor, error)
	Count(ctx context.Context) (int64, error)
	AsQuerable() (QuerableCollection, error)
	Type() string
}

// QuerableCollection - Collection which allows performing select queries
type QuerableCollection interface {
	DataCollection
	Select(ctx context.Context, s selector.Expr, fld selector.Fields) (BazaarCursor, error)
}
