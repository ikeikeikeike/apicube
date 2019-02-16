package usecase

import (
	"github.com/pkg/errors"

	"github.com/ikeikeikeike/apicube/base/data/es"
	"github.com/ikeikeikeike/apicube/base/data/model"
	"github.com/ikeikeikeike/apicube/base/domain/trans"
	"github.com/ikeikeikeike/apicube/base/util"
)

type (
	// Products manifests ...
	Products interface {
		ESUpsert(*model.DTBProduct) error
	}

	products struct {
		Env        util.Environment `inject:""`
		Trans      trans.Products   `inject:""`
		ESProducts es.Products      `inject:""`
	}
)

// Compare returns pb message
//
func (p *products) ESUpsert(m *model.DTBProduct) error {
	data, err := p.Trans.Translate(m)
	if err != nil {
		return errors.Wrap(err, "esupsert translate")
	}

	if err := p.ESProducts.Upsert(data); err != nil {
		return errors.Wrap(err, "esupsert upsert")
	}

	return nil
}
