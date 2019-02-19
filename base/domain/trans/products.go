package trans

import (
	"github.com/gogo/protobuf/types"
	"github.com/pkg/errors"
	"github.com/spf13/cast"

	"github.com/ikeikeikeike/apicube/base/data/es"
	"github.com/ikeikeikeike/apicube/base/data/model"
	"github.com/ikeikeikeike/apicube/base/util"
	pb "github.com/ikeikeikeike/apicube/rpc/protocol/apicube/product"
)

type (
	// Products manifests ...
	Products interface {
		DBToES(*model.DTBProduct) (*es.ProductsSchema, error)
		ESToPB(*es.ProductsSchema) (*pb.Product, error)
	}

	products struct {
		Env util.Environment `inject:""`
	}
)

// DBToES translates to es format
func (t *products) DBToES(m *model.DTBProduct) (*es.ProductsSchema, error) {
	msg := &es.ProductsSchema{
		NameRuby:     m.Name,
		NameAnything: m.Name,
	}

	if err := util.Merge(msg, m); err != nil {
		return nil, errors.Wrap(err, "merge product")
	}

	return msg, nil
}

// DBToES translates to pb format
func (t *products) ESToPB(ps *es.ProductsSchema) (*pb.Product, error) {
	msg := &pb.Product{
		ID:                &types.Int64Value{Value: cast.ToInt64(ps.ID)}, // XXX: Must be set types.Int64Value to primaryID
		DescriptionDetail: &types.StringValue{Value: ps.DescriptionDetail.String},
	}

	return msg, nil
}
