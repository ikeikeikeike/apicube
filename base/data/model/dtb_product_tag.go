// Code generated by SQLBoiler (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package model

import (
	"context"
	"database/sql"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/pkg/errors"
	"github.com/volatiletech/null"
	"github.com/volatiletech/sqlboiler/boil"
	"github.com/volatiletech/sqlboiler/queries"
	"github.com/volatiletech/sqlboiler/queries/qm"
	"github.com/volatiletech/sqlboiler/queries/qmhelper"
	"github.com/volatiletech/sqlboiler/strmangle"
)

// DTBProductTag is an object representing the database table.
type DTBProductTag struct {
	ID                uint      `boil:"id" json:"id" toml:"id" yaml:"id"`
	ProductID         null.Uint `boil:"product_id" json:"product_id,omitempty" toml:"product_id" yaml:"product_id,omitempty"`
	TagID             null.Uint `boil:"tag_id" json:"tag_id,omitempty" toml:"tag_id" yaml:"tag_id,omitempty"`
	CreatorID         null.Uint `boil:"creator_id" json:"creator_id,omitempty" toml:"creator_id" yaml:"creator_id,omitempty"`
	CreateDate        time.Time `boil:"create_date" json:"create_date" toml:"create_date" yaml:"create_date"`
	DiscriminatorType string    `boil:"discriminator_type" json:"discriminator_type" toml:"discriminator_type" yaml:"discriminator_type"`

	R *dtbProductTagR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L dtbProductTagL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var DTBProductTagColumns = struct {
	ID                string
	ProductID         string
	TagID             string
	CreatorID         string
	CreateDate        string
	DiscriminatorType string
}{
	ID:                "id",
	ProductID:         "product_id",
	TagID:             "tag_id",
	CreatorID:         "creator_id",
	CreateDate:        "create_date",
	DiscriminatorType: "discriminator_type",
}

// Generated where

var DTBProductTagWhere = struct {
	ID                whereHelperuint
	ProductID         whereHelpernull_Uint
	TagID             whereHelpernull_Uint
	CreatorID         whereHelpernull_Uint
	CreateDate        whereHelpertime_Time
	DiscriminatorType whereHelperstring
}{
	ID:                whereHelperuint{field: `id`},
	ProductID:         whereHelpernull_Uint{field: `product_id`},
	TagID:             whereHelpernull_Uint{field: `tag_id`},
	CreatorID:         whereHelpernull_Uint{field: `creator_id`},
	CreateDate:        whereHelpertime_Time{field: `create_date`},
	DiscriminatorType: whereHelperstring{field: `discriminator_type`},
}

// DTBProductTagRels is where relationship names are stored.
var DTBProductTagRels = struct {
	Product string
	Tag     string
}{
	Product: "Product",
	Tag:     "Tag",
}

// dtbProductTagR is where relationships are stored.
type dtbProductTagR struct {
	Product *DTBProduct
	Tag     *DTBTag
}

// NewStruct creates a new relationship struct
func (*dtbProductTagR) NewStruct() *dtbProductTagR {
	return &dtbProductTagR{}
}

// dtbProductTagL is where Load methods for each relationship are stored.
type dtbProductTagL struct{}

var (
	dtbProductTagColumns               = []string{"id", "product_id", "tag_id", "creator_id", "create_date", "discriminator_type"}
	dtbProductTagColumnsWithoutDefault = []string{"product_id", "tag_id", "creator_id", "create_date", "discriminator_type"}
	dtbProductTagColumnsWithDefault    = []string{"id"}
	dtbProductTagPrimaryKeyColumns     = []string{"id"}
)

type (
	// DTBProductTagSlice is an alias for a slice of pointers to DTBProductTag.
	// This should generally be used opposed to []DTBProductTag.
	DTBProductTagSlice []*DTBProductTag
	// DTBProductTagHook is the signature for custom DTBProductTag hook methods
	DTBProductTagHook func(context.Context, boil.ContextExecutor, *DTBProductTag) error

	dtbProductTagQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	dtbProductTagType                 = reflect.TypeOf(&DTBProductTag{})
	dtbProductTagMapping              = queries.MakeStructMapping(dtbProductTagType)
	dtbProductTagPrimaryKeyMapping, _ = queries.BindMapping(dtbProductTagType, dtbProductTagMapping, dtbProductTagPrimaryKeyColumns)
	dtbProductTagInsertCacheMut       sync.RWMutex
	dtbProductTagInsertCache          = make(map[string]insertCache)
	dtbProductTagUpdateCacheMut       sync.RWMutex
	dtbProductTagUpdateCache          = make(map[string]updateCache)
	dtbProductTagUpsertCacheMut       sync.RWMutex
	dtbProductTagUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var dtbProductTagBeforeInsertHooks []DTBProductTagHook
var dtbProductTagBeforeUpdateHooks []DTBProductTagHook
var dtbProductTagBeforeDeleteHooks []DTBProductTagHook
var dtbProductTagBeforeUpsertHooks []DTBProductTagHook

var dtbProductTagAfterInsertHooks []DTBProductTagHook
var dtbProductTagAfterSelectHooks []DTBProductTagHook
var dtbProductTagAfterUpdateHooks []DTBProductTagHook
var dtbProductTagAfterDeleteHooks []DTBProductTagHook
var dtbProductTagAfterUpsertHooks []DTBProductTagHook

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *DTBProductTag) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range dtbProductTagBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *DTBProductTag) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range dtbProductTagBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *DTBProductTag) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range dtbProductTagBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *DTBProductTag) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range dtbProductTagBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *DTBProductTag) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range dtbProductTagAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterSelectHooks executes all "after Select" hooks.
func (o *DTBProductTag) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range dtbProductTagAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *DTBProductTag) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range dtbProductTagAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *DTBProductTag) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range dtbProductTagAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *DTBProductTag) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range dtbProductTagAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddDTBProductTagHook registers your hook function for all future operations.
func AddDTBProductTagHook(hookPoint boil.HookPoint, dtbProductTagHook DTBProductTagHook) {
	switch hookPoint {
	case boil.BeforeInsertHook:
		dtbProductTagBeforeInsertHooks = append(dtbProductTagBeforeInsertHooks, dtbProductTagHook)
	case boil.BeforeUpdateHook:
		dtbProductTagBeforeUpdateHooks = append(dtbProductTagBeforeUpdateHooks, dtbProductTagHook)
	case boil.BeforeDeleteHook:
		dtbProductTagBeforeDeleteHooks = append(dtbProductTagBeforeDeleteHooks, dtbProductTagHook)
	case boil.BeforeUpsertHook:
		dtbProductTagBeforeUpsertHooks = append(dtbProductTagBeforeUpsertHooks, dtbProductTagHook)
	case boil.AfterInsertHook:
		dtbProductTagAfterInsertHooks = append(dtbProductTagAfterInsertHooks, dtbProductTagHook)
	case boil.AfterSelectHook:
		dtbProductTagAfterSelectHooks = append(dtbProductTagAfterSelectHooks, dtbProductTagHook)
	case boil.AfterUpdateHook:
		dtbProductTagAfterUpdateHooks = append(dtbProductTagAfterUpdateHooks, dtbProductTagHook)
	case boil.AfterDeleteHook:
		dtbProductTagAfterDeleteHooks = append(dtbProductTagAfterDeleteHooks, dtbProductTagHook)
	case boil.AfterUpsertHook:
		dtbProductTagAfterUpsertHooks = append(dtbProductTagAfterUpsertHooks, dtbProductTagHook)
	}
}

// One returns a single dtbProductTag record from the query.
func (q dtbProductTagQuery) One(ctx context.Context, exec boil.ContextExecutor) (*DTBProductTag, error) {
	o := &DTBProductTag{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "model: failed to execute a one query for dtb_product_tag")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all DTBProductTag records from the query.
func (q dtbProductTagQuery) All(ctx context.Context, exec boil.ContextExecutor) (DTBProductTagSlice, error) {
	var o []*DTBProductTag

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "model: failed to assign all query results to DTBProductTag slice")
	}

	if len(dtbProductTagAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all DTBProductTag records in the query.
func (q dtbProductTagQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "model: failed to count dtb_product_tag rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q dtbProductTagQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "model: failed to check if dtb_product_tag exists")
	}

	return count > 0, nil
}

// Product pointed to by the foreign key.
func (o *DTBProductTag) Product(mods ...qm.QueryMod) dtbProductQuery {
	queryMods := []qm.QueryMod{
		qm.Where("id=?", o.ProductID),
	}

	queryMods = append(queryMods, mods...)

	query := DTBProducts(queryMods...)
	queries.SetFrom(query.Query, "`dtb_product`")

	return query
}

// Tag pointed to by the foreign key.
func (o *DTBProductTag) Tag(mods ...qm.QueryMod) dtbTagQuery {
	queryMods := []qm.QueryMod{
		qm.Where("id=?", o.TagID),
	}

	queryMods = append(queryMods, mods...)

	query := DTBTags(queryMods...)
	queries.SetFrom(query.Query, "`dtb_tag`")

	return query
}

// LoadProduct allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (dtbProductTagL) LoadProduct(ctx context.Context, e boil.ContextExecutor, singular bool, maybeDTBProductTag interface{}, mods queries.Applicator) error {
	var slice []*DTBProductTag
	var object *DTBProductTag

	if singular {
		object = maybeDTBProductTag.(*DTBProductTag)
	} else {
		slice = *maybeDTBProductTag.(*[]*DTBProductTag)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &dtbProductTagR{}
		}
		if !queries.IsNil(object.ProductID) {
			args = append(args, object.ProductID)
		}

	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &dtbProductTagR{}
			}

			for _, a := range args {
				if queries.Equal(a, obj.ProductID) {
					continue Outer
				}
			}

			if !queries.IsNil(obj.ProductID) {
				args = append(args, obj.ProductID)
			}

		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(qm.From(`dtb_product`), qm.WhereIn(`id in ?`, args...))
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load DTBProduct")
	}

	var resultSlice []*DTBProduct
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice DTBProduct")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results of eager load for dtb_product")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for dtb_product")
	}

	if len(dtbProductTagAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(ctx, e); err != nil {
				return err
			}
		}
	}

	if len(resultSlice) == 0 {
		return nil
	}

	if singular {
		foreign := resultSlice[0]
		object.R.Product = foreign
		if foreign.R == nil {
			foreign.R = &dtbProductR{}
		}
		foreign.R.ProductDTBProductTags = append(foreign.R.ProductDTBProductTags, object)
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if queries.Equal(local.ProductID, foreign.ID) {
				local.R.Product = foreign
				if foreign.R == nil {
					foreign.R = &dtbProductR{}
				}
				foreign.R.ProductDTBProductTags = append(foreign.R.ProductDTBProductTags, local)
				break
			}
		}
	}

	return nil
}

// LoadTag allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (dtbProductTagL) LoadTag(ctx context.Context, e boil.ContextExecutor, singular bool, maybeDTBProductTag interface{}, mods queries.Applicator) error {
	var slice []*DTBProductTag
	var object *DTBProductTag

	if singular {
		object = maybeDTBProductTag.(*DTBProductTag)
	} else {
		slice = *maybeDTBProductTag.(*[]*DTBProductTag)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &dtbProductTagR{}
		}
		if !queries.IsNil(object.TagID) {
			args = append(args, object.TagID)
		}

	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &dtbProductTagR{}
			}

			for _, a := range args {
				if queries.Equal(a, obj.TagID) {
					continue Outer
				}
			}

			if !queries.IsNil(obj.TagID) {
				args = append(args, obj.TagID)
			}

		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(qm.From(`dtb_tag`), qm.WhereIn(`id in ?`, args...))
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load DTBTag")
	}

	var resultSlice []*DTBTag
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice DTBTag")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results of eager load for dtb_tag")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for dtb_tag")
	}

	if len(dtbProductTagAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(ctx, e); err != nil {
				return err
			}
		}
	}

	if len(resultSlice) == 0 {
		return nil
	}

	if singular {
		foreign := resultSlice[0]
		object.R.Tag = foreign
		if foreign.R == nil {
			foreign.R = &dtbTagR{}
		}
		foreign.R.TagDTBProductTags = append(foreign.R.TagDTBProductTags, object)
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if queries.Equal(local.TagID, foreign.ID) {
				local.R.Tag = foreign
				if foreign.R == nil {
					foreign.R = &dtbTagR{}
				}
				foreign.R.TagDTBProductTags = append(foreign.R.TagDTBProductTags, local)
				break
			}
		}
	}

	return nil
}

// SetProduct of the dtbProductTag to the related item.
// Sets o.R.Product to related.
// Adds o to related.R.ProductDTBProductTags.
func (o *DTBProductTag) SetProduct(ctx context.Context, exec boil.ContextExecutor, insert bool, related *DTBProduct) error {
	var err error
	if insert {
		if err = related.Insert(ctx, exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE `dtb_product_tag` SET %s WHERE %s",
		strmangle.SetParamNames("`", "`", 0, []string{"product_id"}),
		strmangle.WhereClause("`", "`", 0, dtbProductTagPrimaryKeyColumns),
	)
	values := []interface{}{related.ID, o.ID}

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, updateQuery)
		fmt.Fprintln(boil.DebugWriter, values)
	}

	if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	queries.Assign(&o.ProductID, related.ID)
	if o.R == nil {
		o.R = &dtbProductTagR{
			Product: related,
		}
	} else {
		o.R.Product = related
	}

	if related.R == nil {
		related.R = &dtbProductR{
			ProductDTBProductTags: DTBProductTagSlice{o},
		}
	} else {
		related.R.ProductDTBProductTags = append(related.R.ProductDTBProductTags, o)
	}

	return nil
}

// RemoveProduct relationship.
// Sets o.R.Product to nil.
// Removes o from all passed in related items' relationships struct (Optional).
func (o *DTBProductTag) RemoveProduct(ctx context.Context, exec boil.ContextExecutor, related *DTBProduct) error {
	var err error

	queries.SetScanner(&o.ProductID, nil)
	if _, err = o.Update(ctx, exec, boil.Whitelist("product_id")); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	o.R.Product = nil
	if related == nil || related.R == nil {
		return nil
	}

	for i, ri := range related.R.ProductDTBProductTags {
		if queries.Equal(o.ProductID, ri.ProductID) {
			continue
		}

		ln := len(related.R.ProductDTBProductTags)
		if ln > 1 && i < ln-1 {
			related.R.ProductDTBProductTags[i] = related.R.ProductDTBProductTags[ln-1]
		}
		related.R.ProductDTBProductTags = related.R.ProductDTBProductTags[:ln-1]
		break
	}
	return nil
}

// SetTag of the dtbProductTag to the related item.
// Sets o.R.Tag to related.
// Adds o to related.R.TagDTBProductTags.
func (o *DTBProductTag) SetTag(ctx context.Context, exec boil.ContextExecutor, insert bool, related *DTBTag) error {
	var err error
	if insert {
		if err = related.Insert(ctx, exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE `dtb_product_tag` SET %s WHERE %s",
		strmangle.SetParamNames("`", "`", 0, []string{"tag_id"}),
		strmangle.WhereClause("`", "`", 0, dtbProductTagPrimaryKeyColumns),
	)
	values := []interface{}{related.ID, o.ID}

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, updateQuery)
		fmt.Fprintln(boil.DebugWriter, values)
	}

	if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	queries.Assign(&o.TagID, related.ID)
	if o.R == nil {
		o.R = &dtbProductTagR{
			Tag: related,
		}
	} else {
		o.R.Tag = related
	}

	if related.R == nil {
		related.R = &dtbTagR{
			TagDTBProductTags: DTBProductTagSlice{o},
		}
	} else {
		related.R.TagDTBProductTags = append(related.R.TagDTBProductTags, o)
	}

	return nil
}

// RemoveTag relationship.
// Sets o.R.Tag to nil.
// Removes o from all passed in related items' relationships struct (Optional).
func (o *DTBProductTag) RemoveTag(ctx context.Context, exec boil.ContextExecutor, related *DTBTag) error {
	var err error

	queries.SetScanner(&o.TagID, nil)
	if _, err = o.Update(ctx, exec, boil.Whitelist("tag_id")); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	o.R.Tag = nil
	if related == nil || related.R == nil {
		return nil
	}

	for i, ri := range related.R.TagDTBProductTags {
		if queries.Equal(o.TagID, ri.TagID) {
			continue
		}

		ln := len(related.R.TagDTBProductTags)
		if ln > 1 && i < ln-1 {
			related.R.TagDTBProductTags[i] = related.R.TagDTBProductTags[ln-1]
		}
		related.R.TagDTBProductTags = related.R.TagDTBProductTags[:ln-1]
		break
	}
	return nil
}

// DTBProductTags retrieves all the records using an executor.
func DTBProductTags(mods ...qm.QueryMod) dtbProductTagQuery {
	mods = append(mods, qm.From("`dtb_product_tag`"))
	return dtbProductTagQuery{NewQuery(mods...)}
}

// FindDTBProductTag retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindDTBProductTag(ctx context.Context, exec boil.ContextExecutor, iD uint, selectCols ...string) (*DTBProductTag, error) {
	dtbProductTagObj := &DTBProductTag{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from `dtb_product_tag` where `id`=?", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, dtbProductTagObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "model: unable to select from dtb_product_tag")
	}

	return dtbProductTagObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *DTBProductTag) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("model: no dtb_product_tag provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(dtbProductTagColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	dtbProductTagInsertCacheMut.RLock()
	cache, cached := dtbProductTagInsertCache[key]
	dtbProductTagInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			dtbProductTagColumns,
			dtbProductTagColumnsWithDefault,
			dtbProductTagColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(dtbProductTagType, dtbProductTagMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(dtbProductTagType, dtbProductTagMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO `dtb_product_tag` (`%s`) %%sVALUES (%s)%%s", strings.Join(wl, "`,`"), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO `dtb_product_tag` () VALUES ()%s%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			cache.retQuery = fmt.Sprintf("SELECT `%s` FROM `dtb_product_tag` WHERE %s", strings.Join(returnColumns, "`,`"), strmangle.WhereClause("`", "`", 0, dtbProductTagPrimaryKeyColumns))
		}

		cache.query = fmt.Sprintf(cache.query, queryOutput, queryReturning)
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.query)
		fmt.Fprintln(boil.DebugWriter, vals)
	}

	result, err := exec.ExecContext(ctx, cache.query, vals...)

	if err != nil {
		return errors.Wrap(err, "model: unable to insert into dtb_product_tag")
	}

	var lastID int64
	var identifierCols []interface{}

	if len(cache.retMapping) == 0 {
		goto CacheNoHooks
	}

	lastID, err = result.LastInsertId()
	if err != nil {
		return ErrSyncFail
	}

	o.ID = uint(lastID)
	if lastID != 0 && len(cache.retMapping) == 1 && cache.retMapping[0] == dtbProductTagMapping["ID"] {
		goto CacheNoHooks
	}

	identifierCols = []interface{}{
		o.ID,
	}

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.retQuery)
		fmt.Fprintln(boil.DebugWriter, identifierCols...)
	}

	err = exec.QueryRowContext(ctx, cache.retQuery, identifierCols...).Scan(queries.PtrsFromMapping(value, cache.retMapping)...)
	if err != nil {
		return errors.Wrap(err, "model: unable to populate default values for dtb_product_tag")
	}

CacheNoHooks:
	if !cached {
		dtbProductTagInsertCacheMut.Lock()
		dtbProductTagInsertCache[key] = cache
		dtbProductTagInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the DTBProductTag.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *DTBProductTag) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	dtbProductTagUpdateCacheMut.RLock()
	cache, cached := dtbProductTagUpdateCache[key]
	dtbProductTagUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			dtbProductTagColumns,
			dtbProductTagPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("model: unable to update dtb_product_tag, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE `dtb_product_tag` SET %s WHERE %s",
			strmangle.SetParamNames("`", "`", 0, wl),
			strmangle.WhereClause("`", "`", 0, dtbProductTagPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(dtbProductTagType, dtbProductTagMapping, append(wl, dtbProductTagPrimaryKeyColumns...))
		if err != nil {
			return 0, err
		}
	}

	values := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), cache.valueMapping)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.query)
		fmt.Fprintln(boil.DebugWriter, values)
	}

	var result sql.Result
	result, err = exec.ExecContext(ctx, cache.query, values...)
	if err != nil {
		return 0, errors.Wrap(err, "model: unable to update dtb_product_tag row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "model: failed to get rows affected by update for dtb_product_tag")
	}

	if !cached {
		dtbProductTagUpdateCacheMut.Lock()
		dtbProductTagUpdateCache[key] = cache
		dtbProductTagUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q dtbProductTagQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "model: unable to update all for dtb_product_tag")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "model: unable to retrieve rows affected for dtb_product_tag")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o DTBProductTagSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	ln := int64(len(o))
	if ln == 0 {
		return 0, nil
	}

	if len(cols) == 0 {
		return 0, errors.New("model: update all requires at least one column argument")
	}

	colNames := make([]string, len(cols))
	args := make([]interface{}, len(cols))

	i := 0
	for name, value := range cols {
		colNames[i] = name
		args[i] = value
		i++
	}

	// Append all of the primary key values for each column
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), dtbProductTagPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE `dtb_product_tag` SET %s WHERE %s",
		strmangle.SetParamNames("`", "`", 0, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, dtbProductTagPrimaryKeyColumns, len(o)))

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "model: unable to update all in dtbProductTag slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "model: unable to retrieve rows affected all in update all dtbProductTag")
	}
	return rowsAff, nil
}

var mySQLDTBProductTagUniqueColumns = []string{
	"id",
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *DTBProductTag) Upsert(ctx context.Context, exec boil.ContextExecutor, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("model: no dtb_product_tag provided for upsert")
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(dtbProductTagColumnsWithDefault, o)
	nzUniques := queries.NonZeroDefaultSet(mySQLDTBProductTagUniqueColumns, o)

	if len(nzUniques) == 0 {
		return errors.New("cannot upsert with a table that cannot conflict on a unique column")
	}

	// Build cache key in-line uglily - mysql vs psql problems
	buf := strmangle.GetBuffer()
	buf.WriteString(strconv.Itoa(updateColumns.Kind))
	for _, c := range updateColumns.Cols {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	buf.WriteString(strconv.Itoa(insertColumns.Kind))
	for _, c := range insertColumns.Cols {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	for _, c := range nzDefaults {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	for _, c := range nzUniques {
		buf.WriteString(c)
	}
	key := buf.String()
	strmangle.PutBuffer(buf)

	dtbProductTagUpsertCacheMut.RLock()
	cache, cached := dtbProductTagUpsertCache[key]
	dtbProductTagUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			dtbProductTagColumns,
			dtbProductTagColumnsWithDefault,
			dtbProductTagColumnsWithoutDefault,
			nzDefaults,
		)
		update := updateColumns.UpdateColumnSet(
			dtbProductTagColumns,
			dtbProductTagPrimaryKeyColumns,
		)

		if len(update) == 0 {
			return errors.New("model: unable to upsert dtb_product_tag, could not build update column list")
		}

		ret = strmangle.SetComplement(ret, nzUniques)
		cache.query = buildUpsertQueryMySQL(dialect, "dtb_product_tag", update, insert)
		cache.retQuery = fmt.Sprintf(
			"SELECT %s FROM `dtb_product_tag` WHERE %s",
			strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, ret), ","),
			strmangle.WhereClause("`", "`", 0, nzUniques),
		)

		cache.valueMapping, err = queries.BindMapping(dtbProductTagType, dtbProductTagMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(dtbProductTagType, dtbProductTagMapping, ret)
			if err != nil {
				return err
			}
		}
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)
	var returns []interface{}
	if len(cache.retMapping) != 0 {
		returns = queries.PtrsFromMapping(value, cache.retMapping)
	}

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.query)
		fmt.Fprintln(boil.DebugWriter, vals)
	}

	result, err := exec.ExecContext(ctx, cache.query, vals...)

	if err != nil {
		return errors.Wrap(err, "model: unable to upsert for dtb_product_tag")
	}

	var lastID int64
	var uniqueMap []uint64
	var nzUniqueCols []interface{}

	if len(cache.retMapping) == 0 {
		goto CacheNoHooks
	}

	lastID, err = result.LastInsertId()
	if err != nil {
		return ErrSyncFail
	}

	o.ID = uint(lastID)
	if lastID != 0 && len(cache.retMapping) == 1 && cache.retMapping[0] == dtbProductTagMapping["id"] {
		goto CacheNoHooks
	}

	uniqueMap, err = queries.BindMapping(dtbProductTagType, dtbProductTagMapping, nzUniques)
	if err != nil {
		return errors.Wrap(err, "model: unable to retrieve unique values for dtb_product_tag")
	}
	nzUniqueCols = queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), uniqueMap)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.retQuery)
		fmt.Fprintln(boil.DebugWriter, nzUniqueCols...)
	}

	err = exec.QueryRowContext(ctx, cache.retQuery, nzUniqueCols...).Scan(returns...)
	if err != nil {
		return errors.Wrap(err, "model: unable to populate default values for dtb_product_tag")
	}

CacheNoHooks:
	if !cached {
		dtbProductTagUpsertCacheMut.Lock()
		dtbProductTagUpsertCache[key] = cache
		dtbProductTagUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single DTBProductTag record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *DTBProductTag) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("model: no DTBProductTag provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), dtbProductTagPrimaryKeyMapping)
	sql := "DELETE FROM `dtb_product_tag` WHERE `id`=?"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "model: unable to delete from dtb_product_tag")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "model: failed to get rows affected by delete for dtb_product_tag")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q dtbProductTagQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("model: no dtbProductTagQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "model: unable to delete all from dtb_product_tag")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "model: failed to get rows affected by deleteall for dtb_product_tag")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o DTBProductTagSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("model: no DTBProductTag slice provided for delete all")
	}

	if len(o) == 0 {
		return 0, nil
	}

	if len(dtbProductTagBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), dtbProductTagPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM `dtb_product_tag` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, dtbProductTagPrimaryKeyColumns, len(o))

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args)
	}

	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "model: unable to delete all from dtbProductTag slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "model: failed to get rows affected by deleteall for dtb_product_tag")
	}

	if len(dtbProductTagAfterDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	return rowsAff, nil
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *DTBProductTag) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindDTBProductTag(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *DTBProductTagSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := DTBProductTagSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), dtbProductTagPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT `dtb_product_tag`.* FROM `dtb_product_tag` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, dtbProductTagPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "model: unable to reload all in DTBProductTagSlice")
	}

	*o = slice

	return nil
}

// DTBProductTagExists checks if the DTBProductTag row exists.
func DTBProductTagExists(ctx context.Context, exec boil.ContextExecutor, iD uint) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from `dtb_product_tag` where `id`=? limit 1)"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, iD)
	}

	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "model: unable to check if dtb_product_tag exists")
	}

	return exists, nil
}
