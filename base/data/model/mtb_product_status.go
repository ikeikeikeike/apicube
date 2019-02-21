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
	"github.com/volatiletech/sqlboiler/boil"
	"github.com/volatiletech/sqlboiler/queries"
	"github.com/volatiletech/sqlboiler/queries/qm"
	"github.com/volatiletech/sqlboiler/queries/qmhelper"
	"github.com/volatiletech/sqlboiler/strmangle"
)

// MTBProductStatus is an object representing the database table.
type MTBProductStatus struct {
	ID                uint16 `boil:"id" json:"id" toml:"id" yaml:"id"`
	Name              string `boil:"name" json:"name" toml:"name" yaml:"name"`
	SortNo            uint16 `boil:"sort_no" json:"sort_no" toml:"sort_no" yaml:"sort_no"`
	DiscriminatorType string `boil:"discriminator_type" json:"discriminator_type" toml:"discriminator_type" yaml:"discriminator_type"`

	R *mtbProductStatusR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L mtbProductStatusL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var MTBProductStatusColumns = struct {
	ID                string
	Name              string
	SortNo            string
	DiscriminatorType string
}{
	ID:                "id",
	Name:              "name",
	SortNo:            "sort_no",
	DiscriminatorType: "discriminator_type",
}

// Generated where

var MTBProductStatusWhere = struct {
	ID                whereHelperuint16
	Name              whereHelperstring
	SortNo            whereHelperuint16
	DiscriminatorType whereHelperstring
}{
	ID:                whereHelperuint16{field: `id`},
	Name:              whereHelperstring{field: `name`},
	SortNo:            whereHelperuint16{field: `sort_no`},
	DiscriminatorType: whereHelperstring{field: `discriminator_type`},
}

// MTBProductStatusRels is where relationship names are stored.
var MTBProductStatusRels = struct {
	ProductStatusDTBProducts string
}{
	ProductStatusDTBProducts: "ProductStatusDTBProducts",
}

// mtbProductStatusR is where relationships are stored.
type mtbProductStatusR struct {
	ProductStatusDTBProducts DTBProductSlice
}

// NewStruct creates a new relationship struct
func (*mtbProductStatusR) NewStruct() *mtbProductStatusR {
	return &mtbProductStatusR{}
}

// mtbProductStatusL is where Load methods for each relationship are stored.
type mtbProductStatusL struct{}

var (
	mtbProductStatusColumns               = []string{"id", "name", "sort_no", "discriminator_type"}
	mtbProductStatusColumnsWithoutDefault = []string{"id", "name", "sort_no", "discriminator_type"}
	mtbProductStatusColumnsWithDefault    = []string{}
	mtbProductStatusPrimaryKeyColumns     = []string{"id"}
)

type (
	// MTBProductStatusSlice is an alias for a slice of pointers to MTBProductStatus.
	// This should generally be used opposed to []MTBProductStatus.
	MTBProductStatusSlice []*MTBProductStatus
	// MTBProductStatusHook is the signature for custom MTBProductStatus hook methods
	MTBProductStatusHook func(context.Context, boil.ContextExecutor, *MTBProductStatus) error

	mtbProductStatusQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	mtbProductStatusType                 = reflect.TypeOf(&MTBProductStatus{})
	mtbProductStatusMapping              = queries.MakeStructMapping(mtbProductStatusType)
	mtbProductStatusPrimaryKeyMapping, _ = queries.BindMapping(mtbProductStatusType, mtbProductStatusMapping, mtbProductStatusPrimaryKeyColumns)
	mtbProductStatusInsertCacheMut       sync.RWMutex
	mtbProductStatusInsertCache          = make(map[string]insertCache)
	mtbProductStatusUpdateCacheMut       sync.RWMutex
	mtbProductStatusUpdateCache          = make(map[string]updateCache)
	mtbProductStatusUpsertCacheMut       sync.RWMutex
	mtbProductStatusUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var mtbProductStatusBeforeInsertHooks []MTBProductStatusHook
var mtbProductStatusBeforeUpdateHooks []MTBProductStatusHook
var mtbProductStatusBeforeDeleteHooks []MTBProductStatusHook
var mtbProductStatusBeforeUpsertHooks []MTBProductStatusHook

var mtbProductStatusAfterInsertHooks []MTBProductStatusHook
var mtbProductStatusAfterSelectHooks []MTBProductStatusHook
var mtbProductStatusAfterUpdateHooks []MTBProductStatusHook
var mtbProductStatusAfterDeleteHooks []MTBProductStatusHook
var mtbProductStatusAfterUpsertHooks []MTBProductStatusHook

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *MTBProductStatus) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range mtbProductStatusBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *MTBProductStatus) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range mtbProductStatusBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *MTBProductStatus) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range mtbProductStatusBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *MTBProductStatus) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range mtbProductStatusBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *MTBProductStatus) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range mtbProductStatusAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterSelectHooks executes all "after Select" hooks.
func (o *MTBProductStatus) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range mtbProductStatusAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *MTBProductStatus) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range mtbProductStatusAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *MTBProductStatus) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range mtbProductStatusAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *MTBProductStatus) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range mtbProductStatusAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddMTBProductStatusHook registers your hook function for all future operations.
func AddMTBProductStatusHook(hookPoint boil.HookPoint, mtbProductStatusHook MTBProductStatusHook) {
	switch hookPoint {
	case boil.BeforeInsertHook:
		mtbProductStatusBeforeInsertHooks = append(mtbProductStatusBeforeInsertHooks, mtbProductStatusHook)
	case boil.BeforeUpdateHook:
		mtbProductStatusBeforeUpdateHooks = append(mtbProductStatusBeforeUpdateHooks, mtbProductStatusHook)
	case boil.BeforeDeleteHook:
		mtbProductStatusBeforeDeleteHooks = append(mtbProductStatusBeforeDeleteHooks, mtbProductStatusHook)
	case boil.BeforeUpsertHook:
		mtbProductStatusBeforeUpsertHooks = append(mtbProductStatusBeforeUpsertHooks, mtbProductStatusHook)
	case boil.AfterInsertHook:
		mtbProductStatusAfterInsertHooks = append(mtbProductStatusAfterInsertHooks, mtbProductStatusHook)
	case boil.AfterSelectHook:
		mtbProductStatusAfterSelectHooks = append(mtbProductStatusAfterSelectHooks, mtbProductStatusHook)
	case boil.AfterUpdateHook:
		mtbProductStatusAfterUpdateHooks = append(mtbProductStatusAfterUpdateHooks, mtbProductStatusHook)
	case boil.AfterDeleteHook:
		mtbProductStatusAfterDeleteHooks = append(mtbProductStatusAfterDeleteHooks, mtbProductStatusHook)
	case boil.AfterUpsertHook:
		mtbProductStatusAfterUpsertHooks = append(mtbProductStatusAfterUpsertHooks, mtbProductStatusHook)
	}
}

// One returns a single mtbProductStatus record from the query.
func (q mtbProductStatusQuery) One(ctx context.Context, exec boil.ContextExecutor) (*MTBProductStatus, error) {
	o := &MTBProductStatus{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "model: failed to execute a one query for mtb_product_status")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all MTBProductStatus records from the query.
func (q mtbProductStatusQuery) All(ctx context.Context, exec boil.ContextExecutor) (MTBProductStatusSlice, error) {
	var o []*MTBProductStatus

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "model: failed to assign all query results to MTBProductStatus slice")
	}

	if len(mtbProductStatusAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all MTBProductStatus records in the query.
func (q mtbProductStatusQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "model: failed to count mtb_product_status rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q mtbProductStatusQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "model: failed to check if mtb_product_status exists")
	}

	return count > 0, nil
}

// ProductStatusDTBProducts retrieves all the dtb_product's DTBProducts with an executor via product_status_id column.
func (o *MTBProductStatus) ProductStatusDTBProducts(mods ...qm.QueryMod) dtbProductQuery {
	var queryMods []qm.QueryMod
	if len(mods) != 0 {
		queryMods = append(queryMods, mods...)
	}

	queryMods = append(queryMods,
		qm.Where("`dtb_product`.`product_status_id`=?", o.ID),
	)

	query := DTBProducts(queryMods...)
	queries.SetFrom(query.Query, "`dtb_product`")

	if len(queries.GetSelect(query.Query)) == 0 {
		queries.SetSelect(query.Query, []string{"`dtb_product`.*"})
	}

	return query
}

// LoadProductStatusDTBProducts allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for a 1-M or N-M relationship.
func (mtbProductStatusL) LoadProductStatusDTBProducts(ctx context.Context, e boil.ContextExecutor, singular bool, maybeMTBProductStatus interface{}, mods queries.Applicator) error {
	var slice []*MTBProductStatus
	var object *MTBProductStatus

	if singular {
		object = maybeMTBProductStatus.(*MTBProductStatus)
	} else {
		slice = *maybeMTBProductStatus.(*[]*MTBProductStatus)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &mtbProductStatusR{}
		}
		args = append(args, object.ID)
	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &mtbProductStatusR{}
			}

			for _, a := range args {
				if queries.Equal(a, obj.ID) {
					continue Outer
				}
			}

			args = append(args, obj.ID)
		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(qm.From(`dtb_product`), qm.WhereIn(`product_status_id in ?`, args...))
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load dtb_product")
	}

	var resultSlice []*DTBProduct
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice dtb_product")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results in eager load on dtb_product")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for dtb_product")
	}

	if len(dtbProductAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(ctx, e); err != nil {
				return err
			}
		}
	}
	if singular {
		object.R.ProductStatusDTBProducts = resultSlice
		for _, foreign := range resultSlice {
			if foreign.R == nil {
				foreign.R = &dtbProductR{}
			}
			foreign.R.ProductStatus = object
		}
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if queries.Equal(local.ID, foreign.ProductStatusID) {
				local.R.ProductStatusDTBProducts = append(local.R.ProductStatusDTBProducts, foreign)
				if foreign.R == nil {
					foreign.R = &dtbProductR{}
				}
				foreign.R.ProductStatus = local
				break
			}
		}
	}

	return nil
}

// AddProductStatusDTBProducts adds the given related objects to the existing relationships
// of the mtb_product_status, optionally inserting them as new records.
// Appends related to o.R.ProductStatusDTBProducts.
// Sets related.R.ProductStatus appropriately.
func (o *MTBProductStatus) AddProductStatusDTBProducts(ctx context.Context, exec boil.ContextExecutor, insert bool, related ...*DTBProduct) error {
	var err error
	for _, rel := range related {
		if insert {
			queries.Assign(&rel.ProductStatusID, o.ID)
			if err = rel.Insert(ctx, exec, boil.Infer()); err != nil {
				return errors.Wrap(err, "failed to insert into foreign table")
			}
		} else {
			updateQuery := fmt.Sprintf(
				"UPDATE `dtb_product` SET %s WHERE %s",
				strmangle.SetParamNames("`", "`", 0, []string{"product_status_id"}),
				strmangle.WhereClause("`", "`", 0, dtbProductPrimaryKeyColumns),
			)
			values := []interface{}{o.ID, rel.ID}

			if boil.DebugMode {
				fmt.Fprintln(boil.DebugWriter, updateQuery)
				fmt.Fprintln(boil.DebugWriter, values)
			}

			if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
				return errors.Wrap(err, "failed to update foreign table")
			}

			queries.Assign(&rel.ProductStatusID, o.ID)
		}
	}

	if o.R == nil {
		o.R = &mtbProductStatusR{
			ProductStatusDTBProducts: related,
		}
	} else {
		o.R.ProductStatusDTBProducts = append(o.R.ProductStatusDTBProducts, related...)
	}

	for _, rel := range related {
		if rel.R == nil {
			rel.R = &dtbProductR{
				ProductStatus: o,
			}
		} else {
			rel.R.ProductStatus = o
		}
	}
	return nil
}

// SetProductStatusDTBProducts removes all previously related items of the
// mtb_product_status replacing them completely with the passed
// in related items, optionally inserting them as new records.
// Sets o.R.ProductStatus's ProductStatusDTBProducts accordingly.
// Replaces o.R.ProductStatusDTBProducts with related.
// Sets related.R.ProductStatus's ProductStatusDTBProducts accordingly.
func (o *MTBProductStatus) SetProductStatusDTBProducts(ctx context.Context, exec boil.ContextExecutor, insert bool, related ...*DTBProduct) error {
	query := "update `dtb_product` set `product_status_id` = null where `product_status_id` = ?"
	values := []interface{}{o.ID}
	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, query)
		fmt.Fprintln(boil.DebugWriter, values)
	}

	_, err := exec.ExecContext(ctx, query, values...)
	if err != nil {
		return errors.Wrap(err, "failed to remove relationships before set")
	}

	if o.R != nil {
		for _, rel := range o.R.ProductStatusDTBProducts {
			queries.SetScanner(&rel.ProductStatusID, nil)
			if rel.R == nil {
				continue
			}

			rel.R.ProductStatus = nil
		}

		o.R.ProductStatusDTBProducts = nil
	}
	return o.AddProductStatusDTBProducts(ctx, exec, insert, related...)
}

// RemoveProductStatusDTBProducts relationships from objects passed in.
// Removes related items from R.ProductStatusDTBProducts (uses pointer comparison, removal does not keep order)
// Sets related.R.ProductStatus.
func (o *MTBProductStatus) RemoveProductStatusDTBProducts(ctx context.Context, exec boil.ContextExecutor, related ...*DTBProduct) error {
	var err error
	for _, rel := range related {
		queries.SetScanner(&rel.ProductStatusID, nil)
		if rel.R != nil {
			rel.R.ProductStatus = nil
		}
		if _, err = rel.Update(ctx, exec, boil.Whitelist("product_status_id")); err != nil {
			return err
		}
	}
	if o.R == nil {
		return nil
	}

	for _, rel := range related {
		for i, ri := range o.R.ProductStatusDTBProducts {
			if rel != ri {
				continue
			}

			ln := len(o.R.ProductStatusDTBProducts)
			if ln > 1 && i < ln-1 {
				o.R.ProductStatusDTBProducts[i] = o.R.ProductStatusDTBProducts[ln-1]
			}
			o.R.ProductStatusDTBProducts = o.R.ProductStatusDTBProducts[:ln-1]
			break
		}
	}

	return nil
}

// MTBProductStatuses retrieves all the records using an executor.
func MTBProductStatuses(mods ...qm.QueryMod) mtbProductStatusQuery {
	mods = append(mods, qm.From("`mtb_product_status`"))
	return mtbProductStatusQuery{NewQuery(mods...)}
}

// FindMTBProductStatus retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindMTBProductStatus(ctx context.Context, exec boil.ContextExecutor, iD uint16, selectCols ...string) (*MTBProductStatus, error) {
	mtbProductStatusObj := &MTBProductStatus{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from `mtb_product_status` where `id`=?", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, mtbProductStatusObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "model: unable to select from mtb_product_status")
	}

	return mtbProductStatusObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *MTBProductStatus) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("model: no mtb_product_status provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(mtbProductStatusColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	mtbProductStatusInsertCacheMut.RLock()
	cache, cached := mtbProductStatusInsertCache[key]
	mtbProductStatusInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			mtbProductStatusColumns,
			mtbProductStatusColumnsWithDefault,
			mtbProductStatusColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(mtbProductStatusType, mtbProductStatusMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(mtbProductStatusType, mtbProductStatusMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO `mtb_product_status` (`%s`) %%sVALUES (%s)%%s", strings.Join(wl, "`,`"), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO `mtb_product_status` () VALUES ()%s%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			cache.retQuery = fmt.Sprintf("SELECT `%s` FROM `mtb_product_status` WHERE %s", strings.Join(returnColumns, "`,`"), strmangle.WhereClause("`", "`", 0, mtbProductStatusPrimaryKeyColumns))
		}

		cache.query = fmt.Sprintf(cache.query, queryOutput, queryReturning)
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.query)
		fmt.Fprintln(boil.DebugWriter, vals)
	}

	_, err = exec.ExecContext(ctx, cache.query, vals...)

	if err != nil {
		return errors.Wrap(err, "model: unable to insert into mtb_product_status")
	}

	var identifierCols []interface{}

	if len(cache.retMapping) == 0 {
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
		return errors.Wrap(err, "model: unable to populate default values for mtb_product_status")
	}

CacheNoHooks:
	if !cached {
		mtbProductStatusInsertCacheMut.Lock()
		mtbProductStatusInsertCache[key] = cache
		mtbProductStatusInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the MTBProductStatus.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *MTBProductStatus) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	mtbProductStatusUpdateCacheMut.RLock()
	cache, cached := mtbProductStatusUpdateCache[key]
	mtbProductStatusUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			mtbProductStatusColumns,
			mtbProductStatusPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("model: unable to update mtb_product_status, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE `mtb_product_status` SET %s WHERE %s",
			strmangle.SetParamNames("`", "`", 0, wl),
			strmangle.WhereClause("`", "`", 0, mtbProductStatusPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(mtbProductStatusType, mtbProductStatusMapping, append(wl, mtbProductStatusPrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "model: unable to update mtb_product_status row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "model: failed to get rows affected by update for mtb_product_status")
	}

	if !cached {
		mtbProductStatusUpdateCacheMut.Lock()
		mtbProductStatusUpdateCache[key] = cache
		mtbProductStatusUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q mtbProductStatusQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "model: unable to update all for mtb_product_status")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "model: unable to retrieve rows affected for mtb_product_status")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o MTBProductStatusSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), mtbProductStatusPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE `mtb_product_status` SET %s WHERE %s",
		strmangle.SetParamNames("`", "`", 0, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, mtbProductStatusPrimaryKeyColumns, len(o)))

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "model: unable to update all in mtbProductStatus slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "model: unable to retrieve rows affected all in update all mtbProductStatus")
	}
	return rowsAff, nil
}

var mySQLMTBProductStatusUniqueColumns = []string{
	"id",
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *MTBProductStatus) Upsert(ctx context.Context, exec boil.ContextExecutor, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("model: no mtb_product_status provided for upsert")
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(mtbProductStatusColumnsWithDefault, o)
	nzUniques := queries.NonZeroDefaultSet(mySQLMTBProductStatusUniqueColumns, o)

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

	mtbProductStatusUpsertCacheMut.RLock()
	cache, cached := mtbProductStatusUpsertCache[key]
	mtbProductStatusUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			mtbProductStatusColumns,
			mtbProductStatusColumnsWithDefault,
			mtbProductStatusColumnsWithoutDefault,
			nzDefaults,
		)
		update := updateColumns.UpdateColumnSet(
			mtbProductStatusColumns,
			mtbProductStatusPrimaryKeyColumns,
		)

		if len(update) == 0 {
			return errors.New("model: unable to upsert mtb_product_status, could not build update column list")
		}

		ret = strmangle.SetComplement(ret, nzUniques)
		cache.query = buildUpsertQueryMySQL(dialect, "mtb_product_status", update, insert)
		cache.retQuery = fmt.Sprintf(
			"SELECT %s FROM `mtb_product_status` WHERE %s",
			strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, ret), ","),
			strmangle.WhereClause("`", "`", 0, nzUniques),
		)

		cache.valueMapping, err = queries.BindMapping(mtbProductStatusType, mtbProductStatusMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(mtbProductStatusType, mtbProductStatusMapping, ret)
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

	_, err = exec.ExecContext(ctx, cache.query, vals...)

	if err != nil {
		return errors.Wrap(err, "model: unable to upsert for mtb_product_status")
	}

	var uniqueMap []uint64
	var nzUniqueCols []interface{}

	if len(cache.retMapping) == 0 {
		goto CacheNoHooks
	}

	uniqueMap, err = queries.BindMapping(mtbProductStatusType, mtbProductStatusMapping, nzUniques)
	if err != nil {
		return errors.Wrap(err, "model: unable to retrieve unique values for mtb_product_status")
	}
	nzUniqueCols = queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), uniqueMap)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.retQuery)
		fmt.Fprintln(boil.DebugWriter, nzUniqueCols...)
	}

	err = exec.QueryRowContext(ctx, cache.retQuery, nzUniqueCols...).Scan(returns...)
	if err != nil {
		return errors.Wrap(err, "model: unable to populate default values for mtb_product_status")
	}

CacheNoHooks:
	if !cached {
		mtbProductStatusUpsertCacheMut.Lock()
		mtbProductStatusUpsertCache[key] = cache
		mtbProductStatusUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single MTBProductStatus record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *MTBProductStatus) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("model: no MTBProductStatus provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), mtbProductStatusPrimaryKeyMapping)
	sql := "DELETE FROM `mtb_product_status` WHERE `id`=?"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "model: unable to delete from mtb_product_status")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "model: failed to get rows affected by delete for mtb_product_status")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q mtbProductStatusQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("model: no mtbProductStatusQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "model: unable to delete all from mtb_product_status")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "model: failed to get rows affected by deleteall for mtb_product_status")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o MTBProductStatusSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("model: no MTBProductStatus slice provided for delete all")
	}

	if len(o) == 0 {
		return 0, nil
	}

	if len(mtbProductStatusBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), mtbProductStatusPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM `mtb_product_status` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, mtbProductStatusPrimaryKeyColumns, len(o))

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args)
	}

	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "model: unable to delete all from mtbProductStatus slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "model: failed to get rows affected by deleteall for mtb_product_status")
	}

	if len(mtbProductStatusAfterDeleteHooks) != 0 {
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
func (o *MTBProductStatus) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindMTBProductStatus(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *MTBProductStatusSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := MTBProductStatusSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), mtbProductStatusPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT `mtb_product_status`.* FROM `mtb_product_status` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, mtbProductStatusPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "model: unable to reload all in MTBProductStatusSlice")
	}

	*o = slice

	return nil
}

// MTBProductStatusExists checks if the MTBProductStatus row exists.
func MTBProductStatusExists(ctx context.Context, exec boil.ContextExecutor, iD uint16) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from `mtb_product_status` where `id`=? limit 1)"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, iD)
	}

	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "model: unable to check if mtb_product_status exists")
	}

	return exists, nil
}
