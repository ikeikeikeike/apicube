package trans

import (
	"github.com/ikeikeikeike/apicube/base/data/es"
	"github.com/ikeikeikeike/apicube/base/data/model"
	"github.com/ikeikeikeike/apicube/base/util"
	"github.com/pkg/errors"
)

type (
	// Products manifests ...
	Products interface {
		Translate(*model.DTBProduct) (*es.ProductsSchema, error)
	}

	products struct {
		Env util.Environment `inject:""`
	}
)

// Translate to message
func (t *products) Translate(m *model.DTBProduct) (*es.ProductsSchema, error) {
	msg := &es.ProductsSchema{
		NameRuby:     m.Name,
		NameAnything: m.Name,
	}

	if err := util.Merge(msg, m); err != nil {
		return nil, errors.Wrap(err, "merge product")
	}

	return msg, nil
}
