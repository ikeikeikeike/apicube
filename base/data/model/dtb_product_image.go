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

// DTBProductImage is an object representing the database table.
type DTBProductImage struct {
	ID                uint      `boil:"id" json:"id" toml:"id" yaml:"id"`
	ProductID         null.Uint `boil:"product_id" json:"product_id,omitempty" toml:"product_id" yaml:"product_id,omitempty"`
	CreatorID         null.Uint `boil:"creator_id" json:"creator_id,omitempty" toml:"creator_id" yaml:"creator_id,omitempty"`
	FileName          string    `boil:"file_name" json:"file_name" toml:"file_name" yaml:"file_name"`
	SortNo            uint16    `boil:"sort_no" json:"sort_no" toml:"sort_no" yaml:"sort_no"`
	CreateDate        time.Time `boil:"create_date" json:"create_date" toml:"create_date" yaml:"create_date"`
	DiscriminatorType string    `boil:"discriminator_type" json:"discriminator_type" toml:"discriminator_type" yaml:"discriminator_type"`

	R *dtbProductImageR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L dtbProductImageL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var DTBProductImageColumns = struct {
	ID                string
	ProductID         string
	CreatorID         string
	FileName          string
	SortNo            string
	CreateDate        string
	DiscriminatorType string
}{
	ID:                "id",
	ProductID:         "product_id",
	CreatorID:         "creator_id",
	FileName:          "file_name",
	SortNo:            "sort_no",
	CreateDate:        "create_date",
	DiscriminatorType: "discriminator_type",
}

// Generated where

var DTBProductImageWhere = struct {
	ID                whereHelperuint
	ProductID         whereHelpernull_Uint
	CreatorID         whereHelpernull_Uint
	FileName          whereHelperstring
	SortNo            whereHelperuint16
	CreateDate        whereHelpertime_Time
	DiscriminatorType whereHelperstring
}{
	ID:                whereHelperuint{field: `id`},
	ProductID:         whereHelpernull_Uint{field: `product_id`},
	CreatorID:         whereHelpernull_Uint{field: `creator_id`},
	FileName:          whereHelperstring{field: `file_name`},
	SortNo:            whereHelperuint16{field: `sort_no`},
	CreateDate:        whereHelpertime_Time{field: `create_date`},
	DiscriminatorType: whereHelperstring{field: `discriminator_type`},
}

// DTBProductImageRels is where relationship names are stored.
var DTBProductImageRels = struct {
	Product string
}{
	Product: "Product",
}

// dtbProductImageR is where relationships are stored.
type dtbProductImageR struct {
	Product *DTBProduct
}

// NewStruct creates a new relationship struct
func (*dtbProductImageR) NewStruct() *dtbProductImageR {
	return &dtbProductImageR{}
}

// dtbProductImageL is where Load methods for each relationship are stored.
type dtbProductImageL struct{}

var (
	dtbProductImageColumns               = []string{"id", "product_id", "creator_id", "file_name", "sort_no", "create_date", "discriminator_type"}
	dtbProductImageColumnsWithoutDefault = []string{"product_id", "creator_id", "file_name", "sort_no", "create_date", "discriminator_type"}
	dtbProductImageColumnsWithDefault    = []string{"id"}
	dtbProductImagePrimaryKeyColumns     = []string{"id"}
)

type (
	// DTBProductImageSlice is an alias for a slice of pointers to DTBProductImage.
	// This should generally be used opposed to []DTBProductImage.
	DTBProductImageSlice []*DTBProductImage
	// DTBProductImageHook is the signature for custom DTBProductImage hook methods
	DTBProductImageHook func(context.Context, boil.ContextExecutor, *DTBProductImage) error

	dtbProductImageQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	dtbProductImageType                 = reflect.TypeOf(&DTBProductImage{})
	dtbProductImageMapping              = queries.MakeStructMapping(dtbProductImageType)
	dtbProductImagePrimaryKeyMapping, _ = queries.BindMapping(dtbProductImageType, dtbProductImageMapping, dtbProductImagePrimaryKeyColumns)
	dtbProductImageInsertCacheMut       sync.RWMutex
	dtbProductImageInsertCache          = make(map[string]insertCache)
	dtbProductImageUpdateCacheMut       sync.RWMutex
	dtbProductImageUpdateCache          = make(map[string]updateCache)
	dtbProductImageUpsertCacheMut       sync.RWMutex
	dtbProductImageUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var dtbProductImageBeforeInsertHooks []DTBProductImageHook
var dtbProductImageBeforeUpdateHooks []DTBProductImageHook
var dtbProductImageBeforeDeleteHooks []DTBProductImageHook
var dtbProductImageBeforeUpsertHooks []DTBProductImageHook

var dtbProductImageAfterInsertHooks []DTBProductImageHook
var dtbProductImageAfterSelectHooks []DTBProductImageHook
var dtbProductImageAfterUpdateHooks []DTBProductImageHook
var dtbProductImageAfterDeleteHooks []DTBProductImageHook
var dtbProductImageAfterUpsertHooks []DTBProductImageHook

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *DTBProductImage) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range dtbProductImageBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *DTBProductImage) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range dtbProductImageBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *DTBProductImage) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range dtbProductImageBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *DTBProductImage) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range dtbProductImageBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *DTBProductImage) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range dtbProductImageAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterSelectHooks executes all "after Select" hooks.
func (o *DTBProductImage) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range dtbProductImageAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *DTBProductImage) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range dtbProductImageAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *DTBProductImage) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range dtbProductImageAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *DTBProductImage) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range dtbProductImageAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddDTBProductImageHook registers your hook function for all future operations.
func AddDTBProductImageHook(hookPoint boil.HookPoint, dtbProductImageHook DTBProductImageHook) {
	switch hookPoint {
	case boil.BeforeInsertHook:
		dtbProductImageBeforeInsertHooks = append(dtbProductImageBeforeInsertHooks, dtbProductImageHook)
	case boil.BeforeUpdateHook:
		dtbProductImageBeforeUpdateHooks = append(dtbProductImageBeforeUpdateHooks, dtbProductImageHook)
	case boil.BeforeDeleteHook:
		dtbProductImageBeforeDeleteHooks = append(dtbProductImageBeforeDeleteHooks, dtbProductImageHook)
	case boil.BeforeUpsertHook:
		dtbProductImageBeforeUpsertHooks = append(dtbProductImageBeforeUpsertHooks, dtbProductImageHook)
	case boil.AfterInsertHook:
		dtbProductImageAfterInsertHooks = append(dtbProductImageAfterInsertHooks, dtbProductImageHook)
	case boil.AfterSelectHook:
		dtbProductImageAfterSelectHooks = append(dtbProductImageAfterSelectHooks, dtbProductImageHook)
	case boil.AfterUpdateHook:
		dtbProductImageAfterUpdateHooks = append(dtbProductImageAfterUpdateHooks, dtbProductImageHook)
	case boil.AfterDeleteHook:
		dtbProductImageAfterDeleteHooks = append(dtbProductImageAfterDeleteHooks, dtbProductImageHook)
	case boil.AfterUpsertHook:
		dtbProductImageAfterUpsertHooks = append(dtbProductImageAfterUpsertHooks, dtbProductImageHook)
	}
}

// One returns a single dtbProductImage record from the query.
func (q dtbProductImageQuery) One(ctx context.Context, exec boil.ContextExecutor) (*DTBProductImage, error) {
	o := &DTBProductImage{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "model: failed to execute a one query for dtb_product_image")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all DTBProductImage records from the query.
func (q dtbProductImageQuery) All(ctx context.Context, exec boil.ContextExecutor) (DTBProductImageSlice, error) {
	var o []*DTBProductImage

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "model: failed to assign all query results to DTBProductImage slice")
	}

	if len(dtbProductImageAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all DTBProductImage records in the query.
func (q dtbProductImageQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "model: failed to count dtb_product_image rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q dtbProductImageQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "model: failed to check if dtb_product_image exists")
	}

	return count > 0, nil
}

// Product pointed to by the foreign key.
func (o *DTBProductImage) Product(mods ...qm.QueryMod) dtbProductQuery {
	queryMods := []qm.QueryMod{
		qm.Where("id=?", o.ProductID),
	}

	queryMods = append(queryMods, mods...)

	query := DTBProducts(queryMods...)
	queries.SetFrom(query.Query, "`dtb_product`")

	return query
}

// LoadProduct allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (dtbProductImageL) LoadProduct(ctx context.Context, e boil.ContextExecutor, singular bool, maybeDTBProductImage interface{}, mods queries.Applicator) error {
	var slice []*DTBProductImage
	var object *DTBProductImage

	if singular {
		object = maybeDTBProductImage.(*DTBProductImage)
	} else {
		slice = *maybeDTBProductImage.(*[]*DTBProductImage)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &dtbProductImageR{}
		}
		if !queries.IsNil(object.ProductID) {
			args = append(args, object.ProductID)
		}

	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &dtbProductImageR{}
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

	if len(dtbProductImageAfterSelectHooks) != 0 {
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
		foreign.R.ProductDTBProductImages = append(foreign.R.ProductDTBProductImages, object)
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if queries.Equal(local.ProductID, foreign.ID) {
				local.R.Product = foreign
				if foreign.R == nil {
					foreign.R = &dtbProductR{}
				}
				foreign.R.ProductDTBProductImages = append(foreign.R.ProductDTBProductImages, local)
				break
			}
		}
	}

	return nil
}

// SetProduct of the dtbProductImage to the related item.
// Sets o.R.Product to related.
// Adds o to related.R.ProductDTBProductImages.
func (o *DTBProductImage) SetProduct(ctx context.Context, exec boil.ContextExecutor, insert bool, related *DTBProduct) error {
	var err error
	if insert {
		if err = related.Insert(ctx, exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE `dtb_product_image` SET %s WHERE %s",
		strmangle.SetParamNames("`", "`", 0, []string{"product_id"}),
		strmangle.WhereClause("`", "`", 0, dtbProductImagePrimaryKeyColumns),
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
		o.R = &dtbProductImageR{
			Product: related,
		}
	} else {
		o.R.Product = related
	}

	if related.R == nil {
		related.R = &dtbProductR{
			ProductDTBProductImages: DTBProductImageSlice{o},
		}
	} else {
		related.R.ProductDTBProductImages = append(related.R.ProductDTBProductImages, o)
	}

	return nil
}

// RemoveProduct relationship.
// Sets o.R.Product to nil.
// Removes o from all passed in related items' relationships struct (Optional).
func (o *DTBProductImage) RemoveProduct(ctx context.Context, exec boil.ContextExecutor, related *DTBProduct) error {
	var err error

	queries.SetScanner(&o.ProductID, nil)
	if _, err = o.Update(ctx, exec, boil.Whitelist("product_id")); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	o.R.Product = nil
	if related == nil || related.R == nil {
		return nil
	}

	for i, ri := range related.R.ProductDTBProductImages {
		if queries.Equal(o.ProductID, ri.ProductID) {
			continue
		}

		ln := len(related.R.ProductDTBProductImages)
		if ln > 1 && i < ln-1 {
			related.R.ProductDTBProductImages[i] = related.R.ProductDTBProductImages[ln-1]
		}
		related.R.ProductDTBProductImages = related.R.ProductDTBProductImages[:ln-1]
		break
	}
	return nil
}

// DTBProductImages retrieves all the records using an executor.
func DTBProductImages(mods ...qm.QueryMod) dtbProductImageQuery {
	mods = append(mods, qm.From("`dtb_product_image`"))
	return dtbProductImageQuery{NewQuery(mods...)}
}

// FindDTBProductImage retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindDTBProductImage(ctx context.Context, exec boil.ContextExecutor, iD uint, selectCols ...string) (*DTBProductImage, error) {
	dtbProductImageObj := &DTBProductImage{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from `dtb_product_image` where `id`=?", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, dtbProductImageObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "model: unable to select from dtb_product_image")
	}

	return dtbProductImageObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *DTBProductImage) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("model: no dtb_product_image provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(dtbProductImageColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	dtbProductImageInsertCacheMut.RLock()
	cache, cached := dtbProductImageInsertCache[key]
	dtbProductImageInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			dtbProductImageColumns,
			dtbProductImageColumnsWithDefault,
			dtbProductImageColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(dtbProductImageType, dtbProductImageMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(dtbProductImageType, dtbProductImageMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO `dtb_product_image` (`%s`) %%sVALUES (%s)%%s", strings.Join(wl, "`,`"), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO `dtb_product_image` () VALUES ()%s%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			cache.retQuery = fmt.Sprintf("SELECT `%s` FROM `dtb_product_image` WHERE %s", strings.Join(returnColumns, "`,`"), strmangle.WhereClause("`", "`", 0, dtbProductImagePrimaryKeyColumns))
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
		return errors.Wrap(err, "model: unable to insert into dtb_product_image")
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
	if lastID != 0 && len(cache.retMapping) == 1 && cache.retMapping[0] == dtbProductImageMapping["ID"] {
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
		return errors.Wrap(err, "model: unable to populate default values for dtb_product_image")
	}

CacheNoHooks:
	if !cached {
		dtbProductImageInsertCacheMut.Lock()
		dtbProductImageInsertCache[key] = cache
		dtbProductImageInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the DTBProductImage.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *DTBProductImage) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	dtbProductImageUpdateCacheMut.RLock()
	cache, cached := dtbProductImageUpdateCache[key]
	dtbProductImageUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			dtbProductImageColumns,
			dtbProductImagePrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("model: unable to update dtb_product_image, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE `dtb_product_image` SET %s WHERE %s",
			strmangle.SetParamNames("`", "`", 0, wl),
			strmangle.WhereClause("`", "`", 0, dtbProductImagePrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(dtbProductImageType, dtbProductImageMapping, append(wl, dtbProductImagePrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "model: unable to update dtb_product_image row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "model: failed to get rows affected by update for dtb_product_image")
	}

	if !cached {
		dtbProductImageUpdateCacheMut.Lock()
		dtbProductImageUpdateCache[key] = cache
		dtbProductImageUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q dtbProductImageQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "model: unable to update all for dtb_product_image")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "model: unable to retrieve rows affected for dtb_product_image")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o DTBProductImageSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), dtbProductImagePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE `dtb_product_image` SET %s WHERE %s",
		strmangle.SetParamNames("`", "`", 0, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, dtbProductImagePrimaryKeyColumns, len(o)))

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "model: unable to update all in dtbProductImage slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "model: unable to retrieve rows affected all in update all dtbProductImage")
	}
	return rowsAff, nil
}

var mySQLDTBProductImageUniqueColumns = []string{
	"id",
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *DTBProductImage) Upsert(ctx context.Context, exec boil.ContextExecutor, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("model: no dtb_product_image provided for upsert")
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(dtbProductImageColumnsWithDefault, o)
	nzUniques := queries.NonZeroDefaultSet(mySQLDTBProductImageUniqueColumns, o)

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

	dtbProductImageUpsertCacheMut.RLock()
	cache, cached := dtbProductImageUpsertCache[key]
	dtbProductImageUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			dtbProductImageColumns,
			dtbProductImageColumnsWithDefault,
			dtbProductImageColumnsWithoutDefault,
			nzDefaults,
		)
		update := updateColumns.UpdateColumnSet(
			dtbProductImageColumns,
			dtbProductImagePrimaryKeyColumns,
		)

		if len(update) == 0 {
			return errors.New("model: unable to upsert dtb_product_image, could not build update column list")
		}

		ret = strmangle.SetComplement(ret, nzUniques)
		cache.query = buildUpsertQueryMySQL(dialect, "dtb_product_image", update, insert)
		cache.retQuery = fmt.Sprintf(
			"SELECT %s FROM `dtb_product_image` WHERE %s",
			strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, ret), ","),
			strmangle.WhereClause("`", "`", 0, nzUniques),
		)

		cache.valueMapping, err = queries.BindMapping(dtbProductImageType, dtbProductImageMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(dtbProductImageType, dtbProductImageMapping, ret)
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
		return errors.Wrap(err, "model: unable to upsert for dtb_product_image")
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
	if lastID != 0 && len(cache.retMapping) == 1 && cache.retMapping[0] == dtbProductImageMapping["id"] {
		goto CacheNoHooks
	}

	uniqueMap, err = queries.BindMapping(dtbProductImageType, dtbProductImageMapping, nzUniques)
	if err != nil {
		return errors.Wrap(err, "model: unable to retrieve unique values for dtb_product_image")
	}
	nzUniqueCols = queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), uniqueMap)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.retQuery)
		fmt.Fprintln(boil.DebugWriter, nzUniqueCols...)
	}

	err = exec.QueryRowContext(ctx, cache.retQuery, nzUniqueCols...).Scan(returns...)
	if err != nil {
		return errors.Wrap(err, "model: unable to populate default values for dtb_product_image")
	}

CacheNoHooks:
	if !cached {
		dtbProductImageUpsertCacheMut.Lock()
		dtbProductImageUpsertCache[key] = cache
		dtbProductImageUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single DTBProductImage record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *DTBProductImage) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("model: no DTBProductImage provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), dtbProductImagePrimaryKeyMapping)
	sql := "DELETE FROM `dtb_product_image` WHERE `id`=?"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "model: unable to delete from dtb_product_image")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "model: failed to get rows affected by delete for dtb_product_image")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q dtbProductImageQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("model: no dtbProductImageQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "model: unable to delete all from dtb_product_image")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "model: failed to get rows affected by deleteall for dtb_product_image")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o DTBProductImageSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("model: no DTBProductImage slice provided for delete all")
	}

	if len(o) == 0 {
		return 0, nil
	}

	if len(dtbProductImageBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), dtbProductImagePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM `dtb_product_image` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, dtbProductImagePrimaryKeyColumns, len(o))

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args)
	}

	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "model: unable to delete all from dtbProductImage slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "model: failed to get rows affected by deleteall for dtb_product_image")
	}

	if len(dtbProductImageAfterDeleteHooks) != 0 {
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
func (o *DTBProductImage) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindDTBProductImage(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *DTBProductImageSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := DTBProductImageSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), dtbProductImagePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT `dtb_product_image`.* FROM `dtb_product_image` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, dtbProductImagePrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "model: unable to reload all in DTBProductImageSlice")
	}

	*o = slice

	return nil
}

// DTBProductImageExists checks if the DTBProductImage row exists.
func DTBProductImageExists(ctx context.Context, exec boil.ContextExecutor, iD uint) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from `dtb_product_image` where `id`=? limit 1)"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, iD)
	}

	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "model: unable to check if dtb_product_image exists")
	}

	return exists, nil
}
