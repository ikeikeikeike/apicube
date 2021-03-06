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

// MTBSaleType is an object representing the database table.
type MTBSaleType struct {
	ID                uint16 `boil:"id" json:"id" toml:"id" yaml:"id"`
	Name              string `boil:"name" json:"name" toml:"name" yaml:"name"`
	SortNo            uint16 `boil:"sort_no" json:"sort_no" toml:"sort_no" yaml:"sort_no"`
	DiscriminatorType string `boil:"discriminator_type" json:"discriminator_type" toml:"discriminator_type" yaml:"discriminator_type"`

	R *mtbSaleTypeR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L mtbSaleTypeL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var MTBSaleTypeColumns = struct {
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

var MTBSaleTypeWhere = struct {
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

// MTBSaleTypeRels is where relationship names are stored.
var MTBSaleTypeRels = struct {
	SaleTypeDTBDeliveries string
}{
	SaleTypeDTBDeliveries: "SaleTypeDTBDeliveries",
}

// mtbSaleTypeR is where relationships are stored.
type mtbSaleTypeR struct {
	SaleTypeDTBDeliveries DTBDeliverySlice
}

// NewStruct creates a new relationship struct
func (*mtbSaleTypeR) NewStruct() *mtbSaleTypeR {
	return &mtbSaleTypeR{}
}

// mtbSaleTypeL is where Load methods for each relationship are stored.
type mtbSaleTypeL struct{}

var (
	mtbSaleTypeColumns               = []string{"id", "name", "sort_no", "discriminator_type"}
	mtbSaleTypeColumnsWithoutDefault = []string{"id", "name", "sort_no", "discriminator_type"}
	mtbSaleTypeColumnsWithDefault    = []string{}
	mtbSaleTypePrimaryKeyColumns     = []string{"id"}
)

type (
	// MTBSaleTypeSlice is an alias for a slice of pointers to MTBSaleType.
	// This should generally be used opposed to []MTBSaleType.
	MTBSaleTypeSlice []*MTBSaleType
	// MTBSaleTypeHook is the signature for custom MTBSaleType hook methods
	MTBSaleTypeHook func(context.Context, boil.ContextExecutor, *MTBSaleType) error

	mtbSaleTypeQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	mtbSaleTypeType                 = reflect.TypeOf(&MTBSaleType{})
	mtbSaleTypeMapping              = queries.MakeStructMapping(mtbSaleTypeType)
	mtbSaleTypePrimaryKeyMapping, _ = queries.BindMapping(mtbSaleTypeType, mtbSaleTypeMapping, mtbSaleTypePrimaryKeyColumns)
	mtbSaleTypeInsertCacheMut       sync.RWMutex
	mtbSaleTypeInsertCache          = make(map[string]insertCache)
	mtbSaleTypeUpdateCacheMut       sync.RWMutex
	mtbSaleTypeUpdateCache          = make(map[string]updateCache)
	mtbSaleTypeUpsertCacheMut       sync.RWMutex
	mtbSaleTypeUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var mtbSaleTypeBeforeInsertHooks []MTBSaleTypeHook
var mtbSaleTypeBeforeUpdateHooks []MTBSaleTypeHook
var mtbSaleTypeBeforeDeleteHooks []MTBSaleTypeHook
var mtbSaleTypeBeforeUpsertHooks []MTBSaleTypeHook

var mtbSaleTypeAfterInsertHooks []MTBSaleTypeHook
var mtbSaleTypeAfterSelectHooks []MTBSaleTypeHook
var mtbSaleTypeAfterUpdateHooks []MTBSaleTypeHook
var mtbSaleTypeAfterDeleteHooks []MTBSaleTypeHook
var mtbSaleTypeAfterUpsertHooks []MTBSaleTypeHook

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *MTBSaleType) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range mtbSaleTypeBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *MTBSaleType) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range mtbSaleTypeBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *MTBSaleType) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range mtbSaleTypeBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *MTBSaleType) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range mtbSaleTypeBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *MTBSaleType) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range mtbSaleTypeAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterSelectHooks executes all "after Select" hooks.
func (o *MTBSaleType) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range mtbSaleTypeAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *MTBSaleType) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range mtbSaleTypeAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *MTBSaleType) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range mtbSaleTypeAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *MTBSaleType) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range mtbSaleTypeAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddMTBSaleTypeHook registers your hook function for all future operations.
func AddMTBSaleTypeHook(hookPoint boil.HookPoint, mtbSaleTypeHook MTBSaleTypeHook) {
	switch hookPoint {
	case boil.BeforeInsertHook:
		mtbSaleTypeBeforeInsertHooks = append(mtbSaleTypeBeforeInsertHooks, mtbSaleTypeHook)
	case boil.BeforeUpdateHook:
		mtbSaleTypeBeforeUpdateHooks = append(mtbSaleTypeBeforeUpdateHooks, mtbSaleTypeHook)
	case boil.BeforeDeleteHook:
		mtbSaleTypeBeforeDeleteHooks = append(mtbSaleTypeBeforeDeleteHooks, mtbSaleTypeHook)
	case boil.BeforeUpsertHook:
		mtbSaleTypeBeforeUpsertHooks = append(mtbSaleTypeBeforeUpsertHooks, mtbSaleTypeHook)
	case boil.AfterInsertHook:
		mtbSaleTypeAfterInsertHooks = append(mtbSaleTypeAfterInsertHooks, mtbSaleTypeHook)
	case boil.AfterSelectHook:
		mtbSaleTypeAfterSelectHooks = append(mtbSaleTypeAfterSelectHooks, mtbSaleTypeHook)
	case boil.AfterUpdateHook:
		mtbSaleTypeAfterUpdateHooks = append(mtbSaleTypeAfterUpdateHooks, mtbSaleTypeHook)
	case boil.AfterDeleteHook:
		mtbSaleTypeAfterDeleteHooks = append(mtbSaleTypeAfterDeleteHooks, mtbSaleTypeHook)
	case boil.AfterUpsertHook:
		mtbSaleTypeAfterUpsertHooks = append(mtbSaleTypeAfterUpsertHooks, mtbSaleTypeHook)
	}
}

// One returns a single mtbSaleType record from the query.
func (q mtbSaleTypeQuery) One(ctx context.Context, exec boil.ContextExecutor) (*MTBSaleType, error) {
	o := &MTBSaleType{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "model: failed to execute a one query for mtb_sale_type")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all MTBSaleType records from the query.
func (q mtbSaleTypeQuery) All(ctx context.Context, exec boil.ContextExecutor) (MTBSaleTypeSlice, error) {
	var o []*MTBSaleType

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "model: failed to assign all query results to MTBSaleType slice")
	}

	if len(mtbSaleTypeAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all MTBSaleType records in the query.
func (q mtbSaleTypeQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "model: failed to count mtb_sale_type rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q mtbSaleTypeQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "model: failed to check if mtb_sale_type exists")
	}

	return count > 0, nil
}

// SaleTypeDTBDeliveries retrieves all the dtb_delivery's DTBDeliveries with an executor via sale_type_id column.
func (o *MTBSaleType) SaleTypeDTBDeliveries(mods ...qm.QueryMod) dtbDeliveryQuery {
	var queryMods []qm.QueryMod
	if len(mods) != 0 {
		queryMods = append(queryMods, mods...)
	}

	queryMods = append(queryMods,
		qm.Where("`dtb_delivery`.`sale_type_id`=?", o.ID),
	)

	query := DTBDeliveries(queryMods...)
	queries.SetFrom(query.Query, "`dtb_delivery`")

	if len(queries.GetSelect(query.Query)) == 0 {
		queries.SetSelect(query.Query, []string{"`dtb_delivery`.*"})
	}

	return query
}

// LoadSaleTypeDTBDeliveries allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for a 1-M or N-M relationship.
func (mtbSaleTypeL) LoadSaleTypeDTBDeliveries(ctx context.Context, e boil.ContextExecutor, singular bool, maybeMTBSaleType interface{}, mods queries.Applicator) error {
	var slice []*MTBSaleType
	var object *MTBSaleType

	if singular {
		object = maybeMTBSaleType.(*MTBSaleType)
	} else {
		slice = *maybeMTBSaleType.(*[]*MTBSaleType)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &mtbSaleTypeR{}
		}
		args = append(args, object.ID)
	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &mtbSaleTypeR{}
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

	query := NewQuery(qm.From(`dtb_delivery`), qm.WhereIn(`sale_type_id in ?`, args...))
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load dtb_delivery")
	}

	var resultSlice []*DTBDelivery
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice dtb_delivery")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results in eager load on dtb_delivery")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for dtb_delivery")
	}

	if len(dtbDeliveryAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(ctx, e); err != nil {
				return err
			}
		}
	}
	if singular {
		object.R.SaleTypeDTBDeliveries = resultSlice
		for _, foreign := range resultSlice {
			if foreign.R == nil {
				foreign.R = &dtbDeliveryR{}
			}
			foreign.R.SaleType = object
		}
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if queries.Equal(local.ID, foreign.SaleTypeID) {
				local.R.SaleTypeDTBDeliveries = append(local.R.SaleTypeDTBDeliveries, foreign)
				if foreign.R == nil {
					foreign.R = &dtbDeliveryR{}
				}
				foreign.R.SaleType = local
				break
			}
		}
	}

	return nil
}

// AddSaleTypeDTBDeliveries adds the given related objects to the existing relationships
// of the mtb_sale_type, optionally inserting them as new records.
// Appends related to o.R.SaleTypeDTBDeliveries.
// Sets related.R.SaleType appropriately.
func (o *MTBSaleType) AddSaleTypeDTBDeliveries(ctx context.Context, exec boil.ContextExecutor, insert bool, related ...*DTBDelivery) error {
	var err error
	for _, rel := range related {
		if insert {
			queries.Assign(&rel.SaleTypeID, o.ID)
			if err = rel.Insert(ctx, exec, boil.Infer()); err != nil {
				return errors.Wrap(err, "failed to insert into foreign table")
			}
		} else {
			updateQuery := fmt.Sprintf(
				"UPDATE `dtb_delivery` SET %s WHERE %s",
				strmangle.SetParamNames("`", "`", 0, []string{"sale_type_id"}),
				strmangle.WhereClause("`", "`", 0, dtbDeliveryPrimaryKeyColumns),
			)
			values := []interface{}{o.ID, rel.ID}

			if boil.DebugMode {
				fmt.Fprintln(boil.DebugWriter, updateQuery)
				fmt.Fprintln(boil.DebugWriter, values)
			}

			if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
				return errors.Wrap(err, "failed to update foreign table")
			}

			queries.Assign(&rel.SaleTypeID, o.ID)
		}
	}

	if o.R == nil {
		o.R = &mtbSaleTypeR{
			SaleTypeDTBDeliveries: related,
		}
	} else {
		o.R.SaleTypeDTBDeliveries = append(o.R.SaleTypeDTBDeliveries, related...)
	}

	for _, rel := range related {
		if rel.R == nil {
			rel.R = &dtbDeliveryR{
				SaleType: o,
			}
		} else {
			rel.R.SaleType = o
		}
	}
	return nil
}

// SetSaleTypeDTBDeliveries removes all previously related items of the
// mtb_sale_type replacing them completely with the passed
// in related items, optionally inserting them as new records.
// Sets o.R.SaleType's SaleTypeDTBDeliveries accordingly.
// Replaces o.R.SaleTypeDTBDeliveries with related.
// Sets related.R.SaleType's SaleTypeDTBDeliveries accordingly.
func (o *MTBSaleType) SetSaleTypeDTBDeliveries(ctx context.Context, exec boil.ContextExecutor, insert bool, related ...*DTBDelivery) error {
	query := "update `dtb_delivery` set `sale_type_id` = null where `sale_type_id` = ?"
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
		for _, rel := range o.R.SaleTypeDTBDeliveries {
			queries.SetScanner(&rel.SaleTypeID, nil)
			if rel.R == nil {
				continue
			}

			rel.R.SaleType = nil
		}

		o.R.SaleTypeDTBDeliveries = nil
	}
	return o.AddSaleTypeDTBDeliveries(ctx, exec, insert, related...)
}

// RemoveSaleTypeDTBDeliveries relationships from objects passed in.
// Removes related items from R.SaleTypeDTBDeliveries (uses pointer comparison, removal does not keep order)
// Sets related.R.SaleType.
func (o *MTBSaleType) RemoveSaleTypeDTBDeliveries(ctx context.Context, exec boil.ContextExecutor, related ...*DTBDelivery) error {
	var err error
	for _, rel := range related {
		queries.SetScanner(&rel.SaleTypeID, nil)
		if rel.R != nil {
			rel.R.SaleType = nil
		}
		if _, err = rel.Update(ctx, exec, boil.Whitelist("sale_type_id")); err != nil {
			return err
		}
	}
	if o.R == nil {
		return nil
	}

	for _, rel := range related {
		for i, ri := range o.R.SaleTypeDTBDeliveries {
			if rel != ri {
				continue
			}

			ln := len(o.R.SaleTypeDTBDeliveries)
			if ln > 1 && i < ln-1 {
				o.R.SaleTypeDTBDeliveries[i] = o.R.SaleTypeDTBDeliveries[ln-1]
			}
			o.R.SaleTypeDTBDeliveries = o.R.SaleTypeDTBDeliveries[:ln-1]
			break
		}
	}

	return nil
}

// MTBSaleTypes retrieves all the records using an executor.
func MTBSaleTypes(mods ...qm.QueryMod) mtbSaleTypeQuery {
	mods = append(mods, qm.From("`mtb_sale_type`"))
	return mtbSaleTypeQuery{NewQuery(mods...)}
}

// FindMTBSaleType retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindMTBSaleType(ctx context.Context, exec boil.ContextExecutor, iD uint16, selectCols ...string) (*MTBSaleType, error) {
	mtbSaleTypeObj := &MTBSaleType{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from `mtb_sale_type` where `id`=?", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, mtbSaleTypeObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "model: unable to select from mtb_sale_type")
	}

	return mtbSaleTypeObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *MTBSaleType) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("model: no mtb_sale_type provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(mtbSaleTypeColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	mtbSaleTypeInsertCacheMut.RLock()
	cache, cached := mtbSaleTypeInsertCache[key]
	mtbSaleTypeInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			mtbSaleTypeColumns,
			mtbSaleTypeColumnsWithDefault,
			mtbSaleTypeColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(mtbSaleTypeType, mtbSaleTypeMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(mtbSaleTypeType, mtbSaleTypeMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO `mtb_sale_type` (`%s`) %%sVALUES (%s)%%s", strings.Join(wl, "`,`"), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO `mtb_sale_type` () VALUES ()%s%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			cache.retQuery = fmt.Sprintf("SELECT `%s` FROM `mtb_sale_type` WHERE %s", strings.Join(returnColumns, "`,`"), strmangle.WhereClause("`", "`", 0, mtbSaleTypePrimaryKeyColumns))
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
		return errors.Wrap(err, "model: unable to insert into mtb_sale_type")
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
		return errors.Wrap(err, "model: unable to populate default values for mtb_sale_type")
	}

CacheNoHooks:
	if !cached {
		mtbSaleTypeInsertCacheMut.Lock()
		mtbSaleTypeInsertCache[key] = cache
		mtbSaleTypeInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the MTBSaleType.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *MTBSaleType) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	mtbSaleTypeUpdateCacheMut.RLock()
	cache, cached := mtbSaleTypeUpdateCache[key]
	mtbSaleTypeUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			mtbSaleTypeColumns,
			mtbSaleTypePrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("model: unable to update mtb_sale_type, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE `mtb_sale_type` SET %s WHERE %s",
			strmangle.SetParamNames("`", "`", 0, wl),
			strmangle.WhereClause("`", "`", 0, mtbSaleTypePrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(mtbSaleTypeType, mtbSaleTypeMapping, append(wl, mtbSaleTypePrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "model: unable to update mtb_sale_type row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "model: failed to get rows affected by update for mtb_sale_type")
	}

	if !cached {
		mtbSaleTypeUpdateCacheMut.Lock()
		mtbSaleTypeUpdateCache[key] = cache
		mtbSaleTypeUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q mtbSaleTypeQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "model: unable to update all for mtb_sale_type")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "model: unable to retrieve rows affected for mtb_sale_type")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o MTBSaleTypeSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), mtbSaleTypePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE `mtb_sale_type` SET %s WHERE %s",
		strmangle.SetParamNames("`", "`", 0, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, mtbSaleTypePrimaryKeyColumns, len(o)))

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "model: unable to update all in mtbSaleType slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "model: unable to retrieve rows affected all in update all mtbSaleType")
	}
	return rowsAff, nil
}

var mySQLMTBSaleTypeUniqueColumns = []string{
	"id",
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *MTBSaleType) Upsert(ctx context.Context, exec boil.ContextExecutor, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("model: no mtb_sale_type provided for upsert")
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(mtbSaleTypeColumnsWithDefault, o)
	nzUniques := queries.NonZeroDefaultSet(mySQLMTBSaleTypeUniqueColumns, o)

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

	mtbSaleTypeUpsertCacheMut.RLock()
	cache, cached := mtbSaleTypeUpsertCache[key]
	mtbSaleTypeUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			mtbSaleTypeColumns,
			mtbSaleTypeColumnsWithDefault,
			mtbSaleTypeColumnsWithoutDefault,
			nzDefaults,
		)
		update := updateColumns.UpdateColumnSet(
			mtbSaleTypeColumns,
			mtbSaleTypePrimaryKeyColumns,
		)

		if len(update) == 0 {
			return errors.New("model: unable to upsert mtb_sale_type, could not build update column list")
		}

		ret = strmangle.SetComplement(ret, nzUniques)
		cache.query = buildUpsertQueryMySQL(dialect, "mtb_sale_type", update, insert)
		cache.retQuery = fmt.Sprintf(
			"SELECT %s FROM `mtb_sale_type` WHERE %s",
			strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, ret), ","),
			strmangle.WhereClause("`", "`", 0, nzUniques),
		)

		cache.valueMapping, err = queries.BindMapping(mtbSaleTypeType, mtbSaleTypeMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(mtbSaleTypeType, mtbSaleTypeMapping, ret)
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
		return errors.Wrap(err, "model: unable to upsert for mtb_sale_type")
	}

	var uniqueMap []uint64
	var nzUniqueCols []interface{}

	if len(cache.retMapping) == 0 {
		goto CacheNoHooks
	}

	uniqueMap, err = queries.BindMapping(mtbSaleTypeType, mtbSaleTypeMapping, nzUniques)
	if err != nil {
		return errors.Wrap(err, "model: unable to retrieve unique values for mtb_sale_type")
	}
	nzUniqueCols = queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), uniqueMap)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.retQuery)
		fmt.Fprintln(boil.DebugWriter, nzUniqueCols...)
	}

	err = exec.QueryRowContext(ctx, cache.retQuery, nzUniqueCols...).Scan(returns...)
	if err != nil {
		return errors.Wrap(err, "model: unable to populate default values for mtb_sale_type")
	}

CacheNoHooks:
	if !cached {
		mtbSaleTypeUpsertCacheMut.Lock()
		mtbSaleTypeUpsertCache[key] = cache
		mtbSaleTypeUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single MTBSaleType record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *MTBSaleType) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("model: no MTBSaleType provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), mtbSaleTypePrimaryKeyMapping)
	sql := "DELETE FROM `mtb_sale_type` WHERE `id`=?"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "model: unable to delete from mtb_sale_type")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "model: failed to get rows affected by delete for mtb_sale_type")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q mtbSaleTypeQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("model: no mtbSaleTypeQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "model: unable to delete all from mtb_sale_type")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "model: failed to get rows affected by deleteall for mtb_sale_type")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o MTBSaleTypeSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("model: no MTBSaleType slice provided for delete all")
	}

	if len(o) == 0 {
		return 0, nil
	}

	if len(mtbSaleTypeBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), mtbSaleTypePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM `mtb_sale_type` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, mtbSaleTypePrimaryKeyColumns, len(o))

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args)
	}

	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "model: unable to delete all from mtbSaleType slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "model: failed to get rows affected by deleteall for mtb_sale_type")
	}

	if len(mtbSaleTypeAfterDeleteHooks) != 0 {
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
func (o *MTBSaleType) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindMTBSaleType(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *MTBSaleTypeSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := MTBSaleTypeSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), mtbSaleTypePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT `mtb_sale_type`.* FROM `mtb_sale_type` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, mtbSaleTypePrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "model: unable to reload all in MTBSaleTypeSlice")
	}

	*o = slice

	return nil
}

// MTBSaleTypeExists checks if the MTBSaleType row exists.
func MTBSaleTypeExists(ctx context.Context, exec boil.ContextExecutor, iD uint16) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from `mtb_sale_type` where `id`=? limit 1)"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, iD)
	}

	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "model: unable to check if mtb_sale_type exists")
	}

	return exists, nil
}
