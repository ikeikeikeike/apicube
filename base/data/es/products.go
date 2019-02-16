package es

import (
	"context"
	"encoding/json"
	"time"

	"github.com/olivere/elastic"
	"github.com/spf13/cast"
	"github.com/volatiletech/null"

	"github.com/ikeikeikeike/apicube/base/util"
)

type (
	// Products manifests ...
	Products interface {
		Upsert(*ProductsSchema) error
	}

	products struct {
		Env util.Environment `inject:""`
		ES  *elastic.Client  `inject:""`
		Cmd Cmd              `inject:""`
	}

	ProductsSchema struct {
		ID                uint        `json:"id"`
		Name              string      `json:"name"`
		NameRuby          string      `json:"name_ruby"`
		NameAnything      string      `json:"name_anything"`
		DescriptionDetail null.String `json:"description_detail,omitempty"`
		SearchWord        null.String `json:"search_word,omitempty"`
		CreateDate        time.Time   `json:"create_date"`
		UpdateDate        time.Time   `json:"update_date"`
	}
)

func (r *products) Upsert(data *ProductsSchema) error {
	doc, err := json.Marshal(data)
	if err != nil {
		return err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Minute)
	defer cancel()

	_, err = r.Cmd.PostDocument(ctx, ProductsName, cast.ToInt(data.ID), string(doc))
	if err != nil {
		return err
	}

	return nil
}
