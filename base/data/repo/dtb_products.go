package repo

import (
	"context"
	"database/sql"

	"github.com/volatiletech/sqlboiler/queries/qm"

	"github.com/ikeikeikeike/apicube/base/data/model"
	"github.com/ikeikeikeike/apicube/base/util"
)

type (
	// DTBProducts manifests api interface
	DTBProducts interface {
		// Find ...
		Find(context.Context, uint) (*model.DTBProduct, error)
		FindBy(context.Context, []qm.QueryMod, ...string) (*model.DTBProduct, error)
		FindPreload(context.Context, uint, ...string) (*model.DTBProduct, error)
		// List ...
		ListBy(ctx context.Context, where []qm.QueryMod, loads ...string) (model.DTBProductSlice, error)
		ListPublic(ctx context.Context, loads ...string) (model.DTBProductSlice, error)
		ListPrivate(ctx context.Context, loads ...string) (model.DTBProductSlice, error)
		// All returns sort ordered records
		All(context.Context) (model.DTBProductSlice, error)
		AllPreload(context.Context, ...string) (model.DTBProductSlice, error)
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

// ListPublic is retriver
func (r *dtbProducts) ListPublic(ctx context.Context, loads ...string) (model.DTBProductSlice, error) {
	where := []qm.QueryMod{ProductPublic, DescOrder}
	mods, err := preloadBy(where, loads...)
	if err != nil {
		return nil, err
	}

	return model.DTBProducts(mods...).All(ctx, r.DB)
}

// ListPrivate is retriver
func (r *dtbProducts) ListPrivate(ctx context.Context, loads ...string) (model.DTBProductSlice, error) {
	where := []qm.QueryMod{ProductPrivate, DescOrder}
	mods, err := preloadBy(where, loads...)
	if err != nil {
		return nil, err
	}

	return model.DTBProducts(mods...).All(ctx, r.DB)
}

// ListBy is retriver with eager preloading
func (r *dtbProducts) ListBy(ctx context.Context, where []qm.QueryMod, loads ...string) (model.DTBProductSlice, error) {
	mods, err := preloadBy(where, loads...)
	if err != nil {
		return nil, err
	}

	return model.DTBProducts(append(mods, DescOrder)...).All(ctx, r.DB)
}

// All returns sort ordered records
func (r *dtbProducts) All(ctx context.Context) (model.DTBProductSlice, error) {
	return model.DTBProducts(DescOrder).All(ctx, r.DB)
}

// AllPreload returns sort ordered records
func (r *dtbProducts) AllPreload(ctx context.Context, loads ...string) (model.DTBProductSlice, error) {
	return model.DTBProducts(preloads(loads...)...).All(ctx, r.DB)
}

// Exists determines existence a record
func (r *dtbProducts) Exists(ctx context.Context, id uint) (bool, error) {
	return model.DTBProductExists(ctx, r.DB, id)
}
