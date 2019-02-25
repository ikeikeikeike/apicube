package loades

import (
	"context"

	"github.com/ikeikeikeike/apicube/base/data/repo"
	"github.com/ikeikeikeike/apicube/base/domain/usecase"
	"github.com/ikeikeikeike/apicube/base/util"
	"github.com/ikeikeikeike/apicube/base/util/dlm"
	"github.com/ikeikeikeike/apicube/base/util/logger"
	"github.com/spf13/cast"
)

type (
	loadProducts struct {
		Env             util.Environment `inject:""`
		DLM             *dlm.DLM         `inject:""`
		RepoDTBProducts repo.DTBProducts `inject:""`
		UsecaseProducts usecase.Products `inject:""`
	}
)

func (ld *loadProducts) load() {
	// Distributed lock
	mux := ld.DLM.Mutex("job.loadProducts", expires)
	if err := mux.Lock(); err != nil {
		logger.Warnf("a loadProducts job is locked by DLM: %s", err)
		return
	}
	defer mux.Unlock()

	// Prevent to lose data when a server is going to restart
	//
	// XXX: no need prevent data lost
	//
	//  end := make(chan struct{})
	//  defer func() {
	//  	end <- struct{}{}
	//  }()
	//  graceful.PostHook(func() {
	//  	select {
	//  	case <-time.After(expires):
	//  		logger.Println("[ERROR] loadProducts job lock timeout")
	//  		break
	//  	case <-end:
	//  		break
	//  	}
	//  })

	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, expires)
	defer cancel()

	products, err := ld.RepoDTBProducts.ListPublic(ctx)
	if err != nil {
		logger.Errorf("loades job ListPublic: %s", err.Error())
		return
	}

	for _, m := range products {
		if err := ld.UsecaseProducts.ESUpsert(ctx, m); err != nil {
			s := "loades job upsert to products: id=%d err=%s"
			logger.Warnf(s, m.ID, err)
		}
	}

	products, err = ld.RepoDTBProducts.ListPrivate(ctx)
	if err != nil {
		logger.Errorf("loades job ListPrivate: %s", err.Error())
		return
	}

	for _, m := range products {
		if err := ld.UsecaseProducts.ESDelete(ctx, cast.ToInt(m.ID)); err != nil {
			s := "loades job delete to products: id=%d err=%s"
			logger.Warnf(s, m.ID, err)
		}
	}
}
