package trans

import (
	"context"
	"time"

	"github.com/gogo/protobuf/types"
	"github.com/pkg/errors"
	"github.com/spf13/cast"
	"github.com/volatiletech/sqlboiler/queries/qm"

	"github.com/ikeikeikeike/apicube/base/data/es"
	"github.com/ikeikeikeike/apicube/base/data/model"
	"github.com/ikeikeikeike/apicube/base/data/repo"
	"github.com/ikeikeikeike/apicube/base/util"
	pb "github.com/ikeikeikeike/apicube/rpc/pb/apicube/product"
)

type (
	// Products manifests ...
	Products interface {
		DBToES(*model.DTBProduct) (*es.ProductsSchema, error)
		DBToPB(*model.DTBProduct) (*pb.Product, error)
		IDsToDBs(...uint) ([]*model.DTBProduct, error)
	}

	products struct {
		Env             util.Environment `inject:""`
		RepoDTBProducts repo.DTBProducts `inject:""`
	}
)

// DBToES translates to es format
func (t *products) DBToES(m *model.DTBProduct) (*es.ProductsSchema, error) {
	msg := &es.ProductsSchema{
		NameRuby:     m.Name,
		NameAnything: m.Name,
	}

	if err := util.Merge(msg, m); err != nil {
		return nil, errors.Wrap(err, "trans merge product")
	}

	return msg, nil
}

// DBToES translates to pb format
func (t *products) DBToPB(m *model.DTBProduct) (*pb.Product, error) {
	msg := &pb.Product{
		ID:                &types.Int64Value{Value: cast.ToInt64(m.ID)}, // XXX: Must be set types.Int64Value to primaryID
		Name:              m.Name,
		DescriptionDetail: &types.StringValue{Value: m.DescriptionDetail.String},
	}

	return msg, nil
}

// IDsToDBs translates to records with keep sort order
func (t *products) IDsToDBs(uids ...uint) ([]*model.DTBProduct, error) {
	ctx, cancel := context.WithTimeout(context.TODO(), time.Minute)
	defer cancel()

	args := make([]interface{}, len(uids))
	for i, id := range uids {
		args[i] = id
	}
	items, err := t.RepoDTBProducts.ListBy(ctx, []qm.QueryMod{qm.WhereIn("id in ?", args...)})
	if err != nil {
		return nil, errors.Wrap(err, "trans IDsToDBs")
	}

	// keep sort order
	keep := map[uint]*model.DTBProduct{}
	for _, item := range items {
		keep[item.ID] = item
	}

	r := make([]*model.DTBProduct, 0)
	for _, id := range uids {
		if v, ok := keep[id]; ok {
			r = append(r, v)
		}
	}

	return r, nil
}
