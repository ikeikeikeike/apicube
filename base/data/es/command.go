package es

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/ikeikeikeike/apicube/base/util"
	"github.com/olivere/elastic"
)

type (
	// Cmd defines command behavior TODO: fill out after.
	Cmd interface {
		Search(ctx context.Context, search *elastic.SearchService) (*Result, error)
		Bulk(ctx context.Context, bulk *elastic.BulkService) (*Result, error)
		PostDocument(ctx context.Context, name string, id int, doc string) (*Result, error)
		DeleteDocument(ctx context.Context, name string, id int) (*Result, error)
		UpdateByScript(ctx context.Context, name string, id int, script string, params map[string]interface{}) (*Result, error)
		UpsertByScript(ctx context.Context, name string, id int, script string, params, upsert map[string]interface{}) (*Result, error)
		CreateIndex(ctx context.Context, name string, index string) (*Result, error)
		DeleteIndex(ctx context.Context, name string) (*Result, error)
		Aliases(ctx context.Context, name string) (*Result, error)
		PutAlias(ctx context.Context, name, alias string) (*Result, error)
		UpdateAliases(ctx context.Context, name, old, new string) (*Result, error)
	}

	// cmd defines interfaces as elasticsearch api.
	cmd struct {
		Env util.Environment `inject:""`
		ES  *elastic.Client  `inject:""`
	}

	// Result TODO: fill out after.
	Result struct {
		Res interface{} // ES Result Buffer
		Err error
	}
)

// Indices returns values which matches alias name
func (cr *Result) Indices(alias string) []string {
	switch value := cr.Res.(type) {
	case *elastic.AliasesResult:
		return value.IndicesByAlias(alias)
	default:
		return []string{}
	}
}

// JSON returns value as JSON
func (cr *Result) JSON() []byte {
	bytes, _ := json.Marshal(cr.Res)
	return bytes
}

// Values returns significant values which was chosen along with any es result
func (cr *Result) Values() interface{} {
	switch value := cr.Res.(type) {
	default:
		return value
	case *elastic.AliasesResult:
		return value.Indices
	case *elastic.IndicesCreateResult:
		return cr.JSON()
	case *Result:
		return value.Values()
	}
}

// MakeIndexName returns name with suffix
func MakeIndexName(name string) string {
	return fmt.Sprintf("%s_%d", name, time.Now().UnixNano())
}

// RestoreIndexName returns actual name
func RestoreIndexName(name string) string {
	return strings.Split(name, "_")[0]
}

func (c *cmd) do(ctx context.Context, fn func(chan *Result)) (*Result, error) {
	rch := make(chan *Result, 1)

	go fn(rch)

	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	case cr := <-rch:
		return cr, cr.Err
	}
}

func (c *cmd) Search(ctx context.Context, search *elastic.SearchService) (*Result, error) {
	fn := func(rch chan *Result) {
		res, err := search.Pretty(c.Env.IsDebug()).Do(ctx)
		rch <- &Result{Res: res, Err: err}
	}

	return c.do(ctx, fn)
}

func (c *cmd) Bulk(ctx context.Context, bulk *elastic.BulkService) (*Result, error) {
	fn := func(rch chan *Result) {
		res, err := bulk.Do(ctx)
		rch <- &Result{Res: res, Err: err}
	}

	return c.do(ctx, fn)
}

func (c *cmd) PostDocument(ctx context.Context, name string, id int, doc string) (*Result, error) {
	fn := func(rch chan *Result) {
		res, err := c.ES.Index().
			Pretty(c.Env.IsDebug()).
			Index(name).Type(name).Id(strconv.Itoa(id)).BodyString(doc).Do(ctx)
		rch <- &Result{Res: res, Err: err}
	}

	return c.do(ctx, fn)
}

func (c *cmd) UpdateByScript(ctx context.Context, name string, id int, script string, params map[string]interface{}) (*Result, error) {
	fn := func(rch chan *Result) {
		script := elastic.NewScript(script).Params(params).Lang("painless")

		res, err := c.ES.Update().
			Pretty(c.Env.IsDebug()).Index(name).Type(name).Id(strconv.Itoa(id)).
			Script(script).Do(ctx)
		rch <- &Result{Res: res, Err: err}
	}

	return c.do(ctx, fn)
}

func (c *cmd) UpsertByScript(ctx context.Context, name string, id int, script string, params, upsert map[string]interface{}) (*Result, error) {
	fn := func(rch chan *Result) {
		script := elastic.NewScript(script).Params(params).Lang("painless")

		res, err := c.ES.Update().
			Pretty(c.Env.IsDebug()).Index(name).Type(name).Id(strconv.Itoa(id)).
			Script(script).ScriptedUpsert(true).Upsert(upsert).Do(ctx)
		rch <- &Result{Res: res, Err: err}
	}

	return c.do(ctx, fn)
}

func (c *cmd) DeleteDocument(ctx context.Context, name string, id int) (*Result, error) {
	fn := func(rch chan *Result) {
		res, err := c.ES.Delete().
			Pretty(c.Env.IsDebug()).
			Index(name).Type(name).Id(strconv.Itoa(id)).Do(ctx)
		rch <- &Result{Res: res, Err: err}
	}

	return c.do(ctx, fn)
}

func (c *cmd) CreateIndex(ctx context.Context, name string, index string) (*Result, error) {
	fn := func(rch chan *Result) {
		res, err := c.ES.CreateIndex(name).
			Pretty(c.Env.IsDebug()).Body(index).Do(ctx)
		rch <- &Result{Res: res, Err: err}
	}

	return c.do(ctx, fn)
}

func (c *cmd) DeleteIndex(ctx context.Context, name string) (*Result, error) {
	fn := func(rch chan *Result) {
		res, err := c.ES.DeleteIndex(name).
			Pretty(c.Env.IsDebug()).Do(ctx)
		rch <- &Result{Res: res, Err: err}
	}

	return c.do(ctx, fn)
}

func (c *cmd) Aliases(ctx context.Context, name string) (*Result, error) {
	fn := func(rch chan *Result) {
		res, err := c.ES.Aliases().
			Pretty(c.Env.IsDebug()).Index(name).Do(ctx)
		rch <- &Result{Res: res, Err: err}
	}

	return c.do(ctx, fn)
}

func (c *cmd) PutAlias(ctx context.Context, name, alias string) (*Result, error) {
	fn := func(rch chan *Result) {
		res, err := c.ES.Alias().
			Pretty(c.Env.IsDebug()).Add(name, alias).Do(ctx)
		rch <- &Result{Res: res, Err: err}
	}

	return c.do(ctx, fn)
}

func (c *cmd) UpdateAliases(ctx context.Context, name, old, new string) (*Result, error) {
	fn := func(rch chan *Result) {
		res, err := c.ES.Alias().
			Pretty(c.Env.IsDebug()).
			Action(elastic.NewAliasRemoveAction(name).Index(old)).
			Action(elastic.NewAliasAddAction(name).Index(new)).
			Do(ctx)
		rch <- &Result{Res: res, Err: err}
	}

	return c.do(ctx, fn)
}
