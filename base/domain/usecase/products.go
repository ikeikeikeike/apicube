package usecase

import (
	"context"

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
		ESUpsert(context.Context, *model.DTBProduct) error
		ESDelete(context.Context, int) error
		Similars(context.Context, *pb.ListSimilarsRequest) (*pb.ListSimilarsResponse, error)
	}

	products struct {
		Env        util.Environment `inject:""`
		Trans      trans.Products   `inject:""`
		ESProducts es.Products      `inject:""`
	}
)

// ESUpsert upserts document
//
func (p *products) ESUpsert(ctx context.Context, m *model.DTBProduct) error {
	data, err := p.Trans.DBToES(m)
	if err != nil {
		return errors.Wrap(err, "products esupsert translate")
	}

	if err := p.ESProducts.Upsert(ctx, data); err != nil {
		return errors.Wrap(err, "products esupsert upsert")
	}

	return nil
}

// ESDelete deletes
//
func (p *products) ESDelete(ctx context.Context, id int) error {
	if err := p.ESProducts.Delete(ctx, id); err != nil {
		return errors.Wrap(err, "products esdelete")
	}

	return nil
}

// Similars gets similar products
//
func (p *products) Similars(ctx context.Context, req *pb.ListSimilarsRequest) (*pb.ListSimilarsResponse, error) {
	uids, err := p.ESProducts.Similars(ctx, req.Name)
	if err != nil {
		return nil, errors.Wrap(err, "esproducts similarls")
	}
	items, err := p.Trans.IDsToDBs(ctx, uids...)
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
