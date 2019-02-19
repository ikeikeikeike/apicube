package usecase

import (
	"github.com/pkg/errors"

	"github.com/ikeikeikeike/apicube/base/data/es"
	"github.com/ikeikeikeike/apicube/base/data/model"
	"github.com/ikeikeikeike/apicube/base/domain/trans"
	"github.com/ikeikeikeike/apicube/base/util"
	"github.com/ikeikeikeike/apicube/base/util/logger"
	pb "github.com/ikeikeikeike/apicube/rpc/protocol/apicube/product"
)

type (
	// Products manifests ...
	Products interface {
		ESUpsert(*model.DTBProduct) error
		Similars(*pb.ListSimilarsRequest) (*pb.ListSimilarsResponse, error)
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
	data, err := p.Trans.DBToES(m)
	if err != nil {
		return errors.Wrap(err, "esupsert translate")
	}

	if err := p.ESProducts.Upsert(data); err != nil {
		return errors.Wrap(err, "esupsert upsert")
	}

	return nil
}

// Similars gets similar products
//
func (p *products) Similars(req *pb.ListSimilarsRequest) (*pb.ListSimilarsResponse, error) {
	items, err := p.ESProducts.Similars(req.Name)
	if err != nil {
		return nil, errors.Wrap(err, "esproducts similarls")
	}

	products := make([]*pb.Product, len(items))
	for i, item := range items {
		product, err := p.Trans.ESToPB(item)
		if err != nil {
			logger.Warnf("esupsert translate: %s", err)
			continue
		}

		products[i] = product
	}

	return &pb.ListSimilarsResponse{Products: products}, nil
}
