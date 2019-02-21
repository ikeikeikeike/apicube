package usecase

import (
	"github.com/pkg/errors"

	"github.com/ikeikeikeike/apicube/base/data/es"
	"github.com/ikeikeikeike/apicube/base/data/model"
	"github.com/ikeikeikeike/apicube/base/domain/trans"
	"github.com/ikeikeikeike/apicube/base/util"
	"github.com/ikeikeikeike/apicube/base/util/logger"
	pb "github.com/ikeikeikeike/apicube/rpc/pb/apicube/product"
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
	uids, err := p.ESProducts.Similars(req.Name)
	if err != nil {
		return nil, errors.Wrap(err, "esproducts similarls")
	}
	items, err := p.Trans.IDsToDBs(uids...)
	if err != nil {
		return nil, errors.Wrap(err, "esproducts similarls")
	}

	msgs := make([]*pb.Product, len(items))
	for i, item := range items {
		msg, err := p.Trans.DBToPB(item)
		if err != nil {
			logger.Warnf("esupsert translate: %s", err)
			continue
		}

		msgs[i] = msg
	}

	return &pb.ListSimilarsResponse{Products: msgs}, nil
}
