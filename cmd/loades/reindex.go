package main

import (
	"context"
	"database/sql"
	"strings"
	"sync"

	"github.com/pkg/errors"

	"github.com/ikeikeikeike/apicube/base/data/es"
	"github.com/ikeikeikeike/apicube/base/data/model"
	"github.com/ikeikeikeike/apicube/base/data/repo"
	"github.com/ikeikeikeike/apicube/base/domain/usecase"
	"github.com/ikeikeikeike/apicube/base/util"
	"github.com/ikeikeikeike/apicube/base/util/logger"
)

type (
	// Reindex defines ...
	Reindex interface {
		Do(ctx context.Context, name, index string) (*es.Result, error)
	}

	reindex struct {
		Env             *util.Env        `inject:""`
		DB              *sql.DB          `inject:""`
		Cmd             es.Cmd           `inject:""`
		RepoDTBProducts repo.DTBProducts `inject:""`
		UsecaseProducts usecase.Products `inject:""`
	}
)

func (ri *reindex) Do(ctx context.Context, name, index string) (*es.Result, error) {
	rch := make(chan *es.Result, 1)

	go func() {
		res, err := ri.reIndex(ctx, name, index)
		rch <- &es.Result{Res: res, Err: err}
	}()

	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	case cr := <-rch:
		return cr, cr.Err
	}
}

// reIndex makes latest index which is zero-downtime re-index processing between older and newer
//
func (ri *reindex) reIndex(ctx context.Context, name, index string) (*es.Result, error) {
	// First, create new index if cluster doesn't have that
	// However returns the 404 status that alias appears does not exists.
	//
	if _, err := ri.Cmd.Aliases(ctx, name); err != nil {
		newname := es.MakeIndexName(name)

		if cr, err := ri.Cmd.CreateIndex(ctx, newname, index); err != nil {
			return cr, errors.Wrap(err, "couldnt put index within first")
		}
		if cr, err := ri.Cmd.PutAlias(ctx, newname, name); err != nil {
			return cr, errors.Wrap(err, "couldnt put alias within first")
		}
	}

	// Second, make sure currently index name
	//
	cr, err := ri.Cmd.Aliases(ctx, name)
	if err != nil {
		return cr, errors.Wrap(err, "alias aint yet within second")
	}
	// This indexes appears older, newer name.
	//
	newname, oldname := es.MakeIndexName(name), cr.Indices(name)[0]

	// Third, create brand new index before import process.
	//
	if cr, err := ri.Cmd.CreateIndex(ctx, newname, index); err != nil {
		return cr, errors.Wrap(err, "couldnt put alias within third")
	}

	// Fourth, import data from database
	//
	if err := ri.dataMigrate(ctx, newname); err != nil {
		return cr, errors.Wrap(err, "couldnt import data to es within Fourth")
	}

	// Finally, replace older alias to newer and then remvoe old index
	//
	if cr, err := ri.Cmd.UpdateAliases(ctx, name, oldname, newname); err != nil {
		return cr, errors.Wrap(err, "couldnt change aliases within Finally")
	}
	if cr, err := ri.Cmd.DeleteIndex(ctx, oldname); err != nil {
		return cr, errors.Wrap(err, "couldnt delete index within Finally")
	}

	// XXX: store data again to alias name
	if err := ri.dataMigrate(ctx, name); err != nil {
		return cr, errors.Wrap(err, "couldnt import data to es within Fourth Again")
	}

	// That's all
	return cr, nil
}

func (ri *reindex) dataMigrate(ctx context.Context, name string) error {
	switch {
	default:
		return errors.New("data migration is not working")
	case strings.HasPrefix(name, es.ProductsName):
		return ri.migrateBy(ctx, name, ri.upsertProducts)
	}
}

func (ri *reindex) migrateBy(ctx context.Context, name string, upsert func(ctx context.Context, m *model.DTBProduct) error) error {
	products, err := ri.RepoDTBProducts.All(ctx)
	if err != nil {
		return errors.Wrap(err, "couldnt get products from database")
	}

	limit := make(chan struct{}, 5) // concurrency
	wg := sync.WaitGroup{}

	for _, product := range products {
		wg.Add(1)

		go func(m *model.DTBProduct) {
			limit <- struct{}{}
			defer wg.Done()

			if err := upsert(ctx, m); err != nil {
				logger.Println("[ERROR] failed to migrate data :", err)
			}

			<-limit
		}(product)
	}

	wg.Wait()

	return nil
}

func (ri *reindex) upsertProducts(ctx context.Context, m *model.DTBProduct) error {

	if err := ri.UsecaseProducts.ESUpsert(m); err != nil {
		s := "[WARN] failed to upsert to products:"
		logger.Println(s, m.ID, err)
	}

	return nil
}
