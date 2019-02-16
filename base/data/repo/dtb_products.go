package repo

import (
	"context"
	"database/sql"

	"github.com/volatiletech/sqlboiler/queries/qm"

	"github.com/ikeikeikeike/apicube/base/data/model"
	"github.com/ikeikeikeike/apicube/base/util"
)

type (
	// Products manifests ...
	DTBProducts interface {
		Find(context.Context, uint) (*model.DTBProduct, error)
		FindBy(context.Context, []qm.QueryMod, ...string) (*model.DTBProduct, error)
		FindPreload(context.Context, uint, ...string) (*model.DTBProduct, error)
		All(context.Context) (model.DTBProductSlice, error)                   // All returns sort ordered records
		AllPreload(context.Context, ...string) (model.DTBProductSlice, error) // All with relation(child) recorrd returns sort ordered records
		Exists(context.Context, uint) (bool, error)
	}

	dtbProducts struct {
		Env util.Environment `inject:""`
		DB  *sql.DB          `inject:""`
	}
)

// Find is retriver
func (r *dtbProducts) Find(ctx context.Context, id uint) (*model.DTBProduct, error) {
	return model.FindDTBProduct(ctx, r.DB, id)
}

// FindBy is retriver with eager preloading
func (r *dtbProducts) FindBy(ctx context.Context, where []qm.QueryMod, loads ...string) (*model.DTBProduct, error) {
	mods, err := preloadBy(where, loads...)
	if err != nil {
		return nil, err
	}

	return model.DTBProducts(mods...).One(ctx, r.DB)
}

// FindPreload is retriver with eager preloading
func (r *dtbProducts) FindPreload(ctx context.Context, id uint, loads ...string) (*model.DTBProduct, error) {
	return model.DTBProducts(preload(id, loads...)...).One(ctx, r.DB)
}

// All returns sort ordered records
func (r *dtbProducts) All(ctx context.Context) (model.DTBProductSlice, error) {
	return model.DTBProducts(qm.OrderBy("id ASC")).All(ctx, r.DB)
}

// AllPreload returns sort ordered records
func (r *dtbProducts) AllPreload(ctx context.Context, loads ...string) (model.DTBProductSlice, error) {
	return model.DTBProducts(preloads(loads...)...).All(ctx, r.DB)
}

// Exists determines existance a record
func (r *dtbProducts) Exists(ctx context.Context, id uint) (bool, error) {
	return model.DTBProductExists(ctx, r.DB, id)
}
