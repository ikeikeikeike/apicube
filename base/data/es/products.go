package es

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/olivere/elastic"
	"github.com/pkg/errors"
	"github.com/spf13/cast"
	"github.com/volatiletech/null"

	"github.com/ikeikeikeike/apicube/base/util"
)

type (
	// Products manifests ...
	Products interface {
		Upsert(context.Context, *ProductsSchema) error
		Similars(context.Context, string) ([]uint, error)
	}

	products struct {
		Env util.Environment `inject:""`
		ES  *elastic.Client  `inject:""`
		Cmd Cmd              `inject:""`
	}

	// ProductsSchema defines elasticsearch's put json.
	ProductsSchema struct {
		ID                uint        `json:"id"`
		ProductStatusID   null.Uint16 `json:"product_status_id"`
		Name              string      `json:"name"`
		NameRuby          string      `json:"name_ruby"`
		NameAnything      string      `json:"name_anything"`
		DescriptionDetail null.String `json:"description_detail,omitempty"`
		SearchWord        null.String `json:"search_word,omitempty"`
		CreateDate        time.Time   `json:"create_date"`
		UpdateDate        time.Time   `json:"update_date"`
	}
)

func (e *products) Upsert(ctx context.Context, data *ProductsSchema) error {
	doc, err := json.Marshal(data)
	if err != nil {
		return err
	}

	ctx, cancel := context.WithTimeout(ctx, 1*time.Minute)
	defer cancel()

	_, err = e.Cmd.PostDocument(ctx, ProductsName, cast.ToInt(data.ID), string(doc))
	if err != nil {
		return err
	}

	return nil
}

// Similars searches similar products
func (e *products) Similars(ctx context.Context, wd string) ([]uint, error) {
	name := ProductsName

	s := fmt.Sprintf(`{
      "_source": ["_id"],
      "min_score": 0.1,
      "query": {
        "bool": {
          "should": [
            { "terms": { "name": ["%s"] }},
            { "terms": { "name_ruby": ["%s"] }},
            { "terms": { "name_anything" : ["%s"]}},
            { "terms": { "description_detail": ["%s"] }},
            { "terms": { "search_word": ["%s"] }}
          ],
          "filter": [
            { "terms": { "product_status_id": [1] }}
          ]
        }
      }
    }`, wd, wd, wd, wd, wd)

	search := e.ES.Search().
		Pretty(e.Env.IsDebug()).
		Index(name).Type(name).
		Source(s)

	ctx, cancel := context.WithTimeout(ctx, 1*time.Minute)
	defer cancel()

	rs, err := e.Cmd.Search(ctx, search)
	if err != nil {
		return nil, errors.Wrap(err, "Similars failed")
	}
	if rs.Err != nil {
		return nil, errors.Wrap(rs.Err, "Similars failed")
	}

	results, ok := rs.Values().(*elastic.SearchResult)
	if !ok {
		return nil, errors.New("Similars: cast *SearchResult")
	}

	uids := make([]uint, len(results.Hits.Hits))
	for i, hit := range results.Hits.Hits {
		uids[i] = cast.ToUint(hit.Id)
	}

	return uids, nil
}
