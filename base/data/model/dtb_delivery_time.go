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

// DTBDeliveryTime is an object representing the database table.
type DTBDeliveryTime struct {
	ID                uint      `boil:"id" json:"id" toml:"id" yaml:"id"`
	DeliveryID        null.Uint `boil:"delivery_id" json:"delivery_id,omitempty" toml:"delivery_id" yaml:"delivery_id,omitempty"`
	DeliveryTime      string    `boil:"delivery_time" json:"delivery_time" toml:"delivery_time" yaml:"delivery_time"`
	SortNo            uint16    `boil:"sort_no" json:"sort_no" toml:"sort_no" yaml:"sort_no"`
	Visible           bool      `boil:"visible" json:"visible" toml:"visible" yaml:"visible"`
	CreateDate        time.Time `boil:"create_date" json:"create_date" toml:"create_date" yaml:"create_date"`
	UpdateDate        time.Time `boil:"update_date" json:"update_date" toml:"update_date" yaml:"update_date"`
	DiscriminatorType string    `boil:"discriminator_type" json:"discriminator_type" toml:"discriminator_type" yaml:"discriminator_type"`

	R *dtbDeliveryTimeR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L dtbDeliveryTimeL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var DTBDeliveryTimeColumns = struct {
	ID                string
	DeliveryID        string
	DeliveryTime      string
	SortNo            string
	Visible           string
	CreateDate        string
	UpdateDate        string
	DiscriminatorType string
}{
	ID:                "id",
	DeliveryID:        "delivery_id",
	DeliveryTime:      "delivery_time",
	SortNo:            "sort_no",
	Visible:           "visible",
	CreateDate:        "create_date",
	UpdateDate:        "update_date",
	DiscriminatorType: "discriminator_type",
}

// Generated where

var DTBDeliveryTimeWhere = struct {
	ID                whereHelperuint
	DeliveryID        whereHelpernull_Uint
	DeliveryTime      whereHelperstring
	SortNo            whereHelperuint16
	Visible           whereHelperbool
	CreateDate        whereHelpertime_Time
	UpdateDate        whereHelpertime_Time
	DiscriminatorType whereHelperstring
}{
	ID:                whereHelperuint{field: `id`},
	DeliveryID:        whereHelpernull_Uint{field: `delivery_id`},
	DeliveryTime:      whereHelperstring{field: `delivery_time`},
	SortNo:            whereHelperuint16{field: `sort_no`},
	Visible:           whereHelperbool{field: `visible`},
	CreateDate:        whereHelpertime_Time{field: `create_date`},
	UpdateDate:        whereHelpertime_Time{field: `update_date`},
	DiscriminatorType: whereHelperstring{field: `discriminator_type`},
}

// DTBDeliveryTimeRels is where relationship names are stored.
var DTBDeliveryTimeRels = struct {
	Delivery string
}{
	Delivery: "Delivery",
}

// dtbDeliveryTimeR is where relationships are stored.
type dtbDeliveryTimeR struct {
	Delivery *DTBDelivery
}

// NewStruct creates a new relationship struct
func (*dtbDeliveryTimeR) NewStruct() *dtbDeliveryTimeR {
	return &dtbDeliveryTimeR{}
}

// dtbDeliveryTimeL is where Load methods for each relationship are stored.
type dtbDeliveryTimeL struct{}

var (
	dtbDeliveryTimeColumns               = []string{"id", "delivery_id", "delivery_time", "sort_no", "visible", "create_date", "update_date", "discriminator_type"}
	dtbDeliveryTimeColumnsWithoutDefault = []string{"delivery_id", "delivery_time", "sort_no", "create_date", "update_date", "discriminator_type"}
	dtbDeliveryTimeColumnsWithDefault    = []string{"id", "visible"}
	dtbDeliveryTimePrimaryKeyColumns     = []string{"id"}
)

type (
	// DTBDeliveryTimeSlice is an alias for a slice of pointers to DTBDeliveryTime.
	// This should generally be used opposed to []DTBDeliveryTime.
	DTBDeliveryTimeSlice []*DTBDeliveryTime
	// DTBDeliveryTimeHook is the signature for custom DTBDeliveryTime hook methods
	DTBDeliveryTimeHook func(context.Context, boil.ContextExecutor, *DTBDeliveryTime) error

	dtbDeliveryTimeQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	dtbDeliveryTimeType                 = reflect.TypeOf(&DTBDeliveryTime{})
	dtbDeliveryTimeMapping              = queries.MakeStructMapping(dtbDeliveryTimeType)
	dtbDeliveryTimePrimaryKeyMapping, _ = queries.BindMapping(dtbDeliveryTimeType, dtbDeliveryTimeMapping, dtbDeliveryTimePrimaryKeyColumns)
	dtbDeliveryTimeInsertCacheMut       sync.RWMutex
	dtbDeliveryTimeInsertCache          = make(map[string]insertCache)
	dtbDeliveryTimeUpdateCacheMut       sync.RWMutex
	dtbDeliveryTimeUpdateCache          = make(map[string]updateCache)
	dtbDeliveryTimeUpsertCacheMut       sync.RWMutex
	dtbDeliveryTimeUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var dtbDeliveryTimeBeforeInsertHooks []DTBDeliveryTimeHook
var dtbDeliveryTimeBeforeUpdateHooks []DTBDeliveryTimeHook
var dtbDeliveryTimeBeforeDeleteHooks []DTBDeliveryTimeHook
var dtbDeliveryTimeBeforeUpsertHooks []DTBDeliveryTimeHook

var dtbDeliveryTimeAfterInsertHooks []DTBDeliveryTimeHook
var dtbDeliveryTimeAfterSelectHooks []DTBDeliveryTimeHook
var dtbDeliveryTimeAfterUpdateHooks []DTBDeliveryTimeHook
var dtbDeliveryTimeAfterDeleteHooks []DTBDeliveryTimeHook
var dtbDeliveryTimeAfterUpsertHooks []DTBDeliveryTimeHook

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *DTBDeliveryTime) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range dtbDeliveryTimeBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *DTBDeliveryTime) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range dtbDeliveryTimeBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *DTBDeliveryTime) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range dtbDeliveryTimeBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *DTBDeliveryTime) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range dtbDeliveryTimeBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *DTBDeliveryTime) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range dtbDeliveryTimeAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterSelectHooks executes all "after Select" hooks.
func (o *DTBDeliveryTime) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range dtbDeliveryTimeAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *DTBDeliveryTime) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range dtbDeliveryTimeAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *DTBDeliveryTime) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range dtbDeliveryTimeAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *DTBDeliveryTime) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range dtbDeliveryTimeAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddDTBDeliveryTimeHook registers your hook function for all future operations.
func AddDTBDeliveryTimeHook(hookPoint boil.HookPoint, dtbDeliveryTimeHook DTBDeliveryTimeHook) {
	switch hookPoint {
	case boil.BeforeInsertHook:
		dtbDeliveryTimeBeforeInsertHooks = append(dtbDeliveryTimeBeforeInsertHooks, dtbDeliveryTimeHook)
	case boil.BeforeUpdateHook:
		dtbDeliveryTimeBeforeUpdateHooks = append(dtbDeliveryTimeBeforeUpdateHooks, dtbDeliveryTimeHook)
	case boil.BeforeDeleteHook:
		dtbDeliveryTimeBeforeDeleteHooks = append(dtbDeliveryTimeBeforeDeleteHooks, dtbDeliveryTimeHook)
	case boil.BeforeUpsertHook:
		dtbDeliveryTimeBeforeUpsertHooks = append(dtbDeliveryTimeBeforeUpsertHooks, dtbDeliveryTimeHook)
	case boil.AfterInsertHook:
		dtbDeliveryTimeAfterInsertHooks = append(dtbDeliveryTimeAfterInsertHooks, dtbDeliveryTimeHook)
	case boil.AfterSelectHook:
		dtbDeliveryTimeAfterSelectHooks = append(dtbDeliveryTimeAfterSelectHooks, dtbDeliveryTimeHook)
	case boil.AfterUpdateHook:
		dtbDeliveryTimeAfterUpdateHooks = append(dtbDeliveryTimeAfterUpdateHooks, dtbDeliveryTimeHook)
	case boil.AfterDeleteHook:
		dtbDeliveryTimeAfterDeleteHooks = append(dtbDeliveryTimeAfterDeleteHooks, dtbDeliveryTimeHook)
	case boil.AfterUpsertHook:
		dtbDeliveryTimeAfterUpsertHooks = append(dtbDeliveryTimeAfterUpsertHooks, dtbDeliveryTimeHook)
	}
}

// One returns a single dtbDeliveryTime record from the query.
func (q dtbDeliveryTimeQuery) One(ctx context.Context, exec boil.ContextExecutor) (*DTBDeliveryTime, error) {
	o := &DTBDeliveryTime{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "model: failed to execute a one query for dtb_delivery_time")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all DTBDeliveryTime records from the query.
func (q dtbDeliveryTimeQuery) All(ctx context.Context, exec boil.ContextExecutor) (DTBDeliveryTimeSlice, error) {
	var o []*DTBDeliveryTime

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "model: failed to assign all query results to DTBDeliveryTime slice")
	}

	if len(dtbDeliveryTimeAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all DTBDeliveryTime records in the query.
func (q dtbDeliveryTimeQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "model: failed to count dtb_delivery_time rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q dtbDeliveryTimeQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "model: failed to check if dtb_delivery_time exists")
	}

	return count > 0, nil
}

// Delivery pointed to by the foreign key.
func (o *DTBDeliveryTime) Delivery(mods ...qm.QueryMod) dtbDeliveryQuery {
	queryMods := []qm.QueryMod{
		qm.Where("id=?", o.DeliveryID),
	}

	queryMods = append(queryMods, mods...)

	query := DTBDeliveries(queryMods...)
	queries.SetFrom(query.Query, "`dtb_delivery`")

	return query
}

// LoadDelivery allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (dtbDeliveryTimeL) LoadDelivery(ctx context.Context, e boil.ContextExecutor, singular bool, maybeDTBDeliveryTime interface{}, mods queries.Applicator) error {
	var slice []*DTBDeliveryTime
	var object *DTBDeliveryTime

	if singular {
		object = maybeDTBDeliveryTime.(*DTBDeliveryTime)
	} else {
		slice = *maybeDTBDeliveryTime.(*[]*DTBDeliveryTime)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &dtbDeliveryTimeR{}
		}
		if !queries.IsNil(object.DeliveryID) {
			args = append(args, object.DeliveryID)
		}

	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &dtbDeliveryTimeR{}
			}

			for _, a := range args {
				if queries.Equal(a, obj.DeliveryID) {
					continue Outer
				}
			}

			if !queries.IsNil(obj.DeliveryID) {
				args = append(args, obj.DeliveryID)
			}

		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(qm.From(`dtb_delivery`), qm.WhereIn(`id in ?`, args...))
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load DTBDelivery")
	}

	var resultSlice []*DTBDelivery
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice DTBDelivery")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results of eager load for dtb_delivery")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for dtb_delivery")
	}

	if len(dtbDeliveryTimeAfterSelectHooks) != 0 {
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
		object.R.Delivery = foreign
		if foreign.R == nil {
			foreign.R = &dtbDeliveryR{}
		}
		foreign.R.DeliveryDTBDeliveryTimes = append(foreign.R.DeliveryDTBDeliveryTimes, object)
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if queries.Equal(local.DeliveryID, foreign.ID) {
				local.R.Delivery = foreign
				if foreign.R == nil {
					foreign.R = &dtbDeliveryR{}
				}
				foreign.R.DeliveryDTBDeliveryTimes = append(foreign.R.DeliveryDTBDeliveryTimes, local)
				break
			}
		}
	}

	return nil
}

// SetDelivery of the dtbDeliveryTime to the related item.
// Sets o.R.Delivery to related.
// Adds o to related.R.DeliveryDTBDeliveryTimes.
func (o *DTBDeliveryTime) SetDelivery(ctx context.Context, exec boil.ContextExecutor, insert bool, related *DTBDelivery) error {
	var err error
	if insert {
		if err = related.Insert(ctx, exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE `dtb_delivery_time` SET %s WHERE %s",
		strmangle.SetParamNames("`", "`", 0, []string{"delivery_id"}),
		strmangle.WhereClause("`", "`", 0, dtbDeliveryTimePrimaryKeyColumns),
	)
	values := []interface{}{related.ID, o.ID}

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, updateQuery)
		fmt.Fprintln(boil.DebugWriter, values)
	}

	if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	queries.Assign(&o.DeliveryID, related.ID)
	if o.R == nil {
		o.R = &dtbDeliveryTimeR{
			Delivery: related,
		}
	} else {
		o.R.Delivery = related
	}

	if related.R == nil {
		related.R = &dtbDeliveryR{
			DeliveryDTBDeliveryTimes: DTBDeliveryTimeSlice{o},
		}
	} else {
		related.R.DeliveryDTBDeliveryTimes = append(related.R.DeliveryDTBDeliveryTimes, o)
	}

	return nil
}

// RemoveDelivery relationship.
// Sets o.R.Delivery to nil.
// Removes o from all passed in related items' relationships struct (Optional).
func (o *DTBDeliveryTime) RemoveDelivery(ctx context.Context, exec boil.ContextExecutor, related *DTBDelivery) error {
	var err error

	queries.SetScanner(&o.DeliveryID, nil)
	if _, err = o.Update(ctx, exec, boil.Whitelist("delivery_id")); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	o.R.Delivery = nil
	if related == nil || related.R == nil {
		return nil
	}

	for i, ri := range related.R.DeliveryDTBDeliveryTimes {
		if queries.Equal(o.DeliveryID, ri.DeliveryID) {
			continue
		}

		ln := len(related.R.DeliveryDTBDeliveryTimes)
		if ln > 1 && i < ln-1 {
			related.R.DeliveryDTBDeliveryTimes[i] = related.R.DeliveryDTBDeliveryTimes[ln-1]
		}
		related.R.DeliveryDTBDeliveryTimes = related.R.DeliveryDTBDeliveryTimes[:ln-1]
		break
	}
	return nil
}

// DTBDeliveryTimes retrieves all the records using an executor.
func DTBDeliveryTimes(mods ...qm.QueryMod) dtbDeliveryTimeQuery {
	mods = append(mods, qm.From("`dtb_delivery_time`"))
	return dtbDeliveryTimeQuery{NewQuery(mods...)}
}

// FindDTBDeliveryTime retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindDTBDeliveryTime(ctx context.Context, exec boil.ContextExecutor, iD uint, selectCols ...string) (*DTBDeliveryTime, error) {
	dtbDeliveryTimeObj := &DTBDeliveryTime{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from `dtb_delivery_time` where `id`=?", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, dtbDeliveryTimeObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "model: unable to select from dtb_delivery_time")
	}

	return dtbDeliveryTimeObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *DTBDeliveryTime) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("model: no dtb_delivery_time provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(dtbDeliveryTimeColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	dtbDeliveryTimeInsertCacheMut.RLock()
	cache, cached := dtbDeliveryTimeInsertCache[key]
	dtbDeliveryTimeInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			dtbDeliveryTimeColumns,
			dtbDeliveryTimeColumnsWithDefault,
			dtbDeliveryTimeColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(dtbDeliveryTimeType, dtbDeliveryTimeMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(dtbDeliveryTimeType, dtbDeliveryTimeMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO `dtb_delivery_time` (`%s`) %%sVALUES (%s)%%s", strings.Join(wl, "`,`"), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO `dtb_delivery_time` () VALUES ()%s%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			cache.retQuery = fmt.Sprintf("SELECT `%s` FROM `dtb_delivery_time` WHERE %s", strings.Join(returnColumns, "`,`"), strmangle.WhereClause("`", "`", 0, dtbDeliveryTimePrimaryKeyColumns))
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
		return errors.Wrap(err, "model: unable to insert into dtb_delivery_time")
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
	if lastID != 0 && len(cache.retMapping) == 1 && cache.retMapping[0] == dtbDeliveryTimeMapping["ID"] {
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
		return errors.Wrap(err, "model: unable to populate default values for dtb_delivery_time")
	}

CacheNoHooks:
	if !cached {
		dtbDeliveryTimeInsertCacheMut.Lock()
		dtbDeliveryTimeInsertCache[key] = cache
		dtbDeliveryTimeInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the DTBDeliveryTime.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *DTBDeliveryTime) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	dtbDeliveryTimeUpdateCacheMut.RLock()
	cache, cached := dtbDeliveryTimeUpdateCache[key]
	dtbDeliveryTimeUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			dtbDeliveryTimeColumns,
			dtbDeliveryTimePrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("model: unable to update dtb_delivery_time, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE `dtb_delivery_time` SET %s WHERE %s",
			strmangle.SetParamNames("`", "`", 0, wl),
			strmangle.WhereClause("`", "`", 0, dtbDeliveryTimePrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(dtbDeliveryTimeType, dtbDeliveryTimeMapping, append(wl, dtbDeliveryTimePrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "model: unable to update dtb_delivery_time row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "model: failed to get rows affected by update for dtb_delivery_time")
	}

	if !cached {
		dtbDeliveryTimeUpdateCacheMut.Lock()
		dtbDeliveryTimeUpdateCache[key] = cache
		dtbDeliveryTimeUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q dtbDeliveryTimeQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "model: unable to update all for dtb_delivery_time")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "model: unable to retrieve rows affected for dtb_delivery_time")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o DTBDeliveryTimeSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), dtbDeliveryTimePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE `dtb_delivery_time` SET %s WHERE %s",
		strmangle.SetParamNames("`", "`", 0, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, dtbDeliveryTimePrimaryKeyColumns, len(o)))

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "model: unable to update all in dtbDeliveryTime slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "model: unable to retrieve rows affected all in update all dtbDeliveryTime")
	}
	return rowsAff, nil
}

var mySQLDTBDeliveryTimeUniqueColumns = []string{
	"id",
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *DTBDeliveryTime) Upsert(ctx context.Context, exec boil.ContextExecutor, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("model: no dtb_delivery_time provided for upsert")
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(dtbDeliveryTimeColumnsWithDefault, o)
	nzUniques := queries.NonZeroDefaultSet(mySQLDTBDeliveryTimeUniqueColumns, o)

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

	dtbDeliveryTimeUpsertCacheMut.RLock()
	cache, cached := dtbDeliveryTimeUpsertCache[key]
	dtbDeliveryTimeUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			dtbDeliveryTimeColumns,
			dtbDeliveryTimeColumnsWithDefault,
			dtbDeliveryTimeColumnsWithoutDefault,
			nzDefaults,
		)
		update := updateColumns.UpdateColumnSet(
			dtbDeliveryTimeColumns,
			dtbDeliveryTimePrimaryKeyColumns,
		)

		if len(update) == 0 {
			return errors.New("model: unable to upsert dtb_delivery_time, could not build update column list")
		}

		ret = strmangle.SetComplement(ret, nzUniques)
		cache.query = buildUpsertQueryMySQL(dialect, "dtb_delivery_time", update, insert)
		cache.retQuery = fmt.Sprintf(
			"SELECT %s FROM `dtb_delivery_time` WHERE %s",
			strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, ret), ","),
			strmangle.WhereClause("`", "`", 0, nzUniques),
		)

		cache.valueMapping, err = queries.BindMapping(dtbDeliveryTimeType, dtbDeliveryTimeMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(dtbDeliveryTimeType, dtbDeliveryTimeMapping, ret)
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
		return errors.Wrap(err, "model: unable to upsert for dtb_delivery_time")
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
	if lastID != 0 && len(cache.retMapping) == 1 && cache.retMapping[0] == dtbDeliveryTimeMapping["id"] {
		goto CacheNoHooks
	}

	uniqueMap, err = queries.BindMapping(dtbDeliveryTimeType, dtbDeliveryTimeMapping, nzUniques)
	if err != nil {
		return errors.Wrap(err, "model: unable to retrieve unique values for dtb_delivery_time")
	}
	nzUniqueCols = queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), uniqueMap)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.retQuery)
		fmt.Fprintln(boil.DebugWriter, nzUniqueCols...)
	}

	err = exec.QueryRowContext(ctx, cache.retQuery, nzUniqueCols...).Scan(returns...)
	if err != nil {
		return errors.Wrap(err, "model: unable to populate default values for dtb_delivery_time")
	}

CacheNoHooks:
	if !cached {
		dtbDeliveryTimeUpsertCacheMut.Lock()
		dtbDeliveryTimeUpsertCache[key] = cache
		dtbDeliveryTimeUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single DTBDeliveryTime record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *DTBDeliveryTime) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("model: no DTBDeliveryTime provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), dtbDeliveryTimePrimaryKeyMapping)
	sql := "DELETE FROM `dtb_delivery_time` WHERE `id`=?"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "model: unable to delete from dtb_delivery_time")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "model: failed to get rows affected by delete for dtb_delivery_time")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q dtbDeliveryTimeQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("model: no dtbDeliveryTimeQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "model: unable to delete all from dtb_delivery_time")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "model: failed to get rows affected by deleteall for dtb_delivery_time")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o DTBDeliveryTimeSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("model: no DTBDeliveryTime slice provided for delete all")
	}

	if len(o) == 0 {
		return 0, nil
	}

	if len(dtbDeliveryTimeBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), dtbDeliveryTimePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM `dtb_delivery_time` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, dtbDeliveryTimePrimaryKeyColumns, len(o))

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args)
	}

	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "model: unable to delete all from dtbDeliveryTime slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "model: failed to get rows affected by deleteall for dtb_delivery_time")
	}

	if len(dtbDeliveryTimeAfterDeleteHooks) != 0 {
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
func (o *DTBDeliveryTime) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindDTBDeliveryTime(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *DTBDeliveryTimeSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := DTBDeliveryTimeSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), dtbDeliveryTimePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT `dtb_delivery_time`.* FROM `dtb_delivery_time` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, dtbDeliveryTimePrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "model: unable to reload all in DTBDeliveryTimeSlice")
	}

	*o = slice

	return nil
}

// DTBDeliveryTimeExists checks if the DTBDeliveryTime row exists.
func DTBDeliveryTimeExists(ctx context.Context, exec boil.ContextExecutor, iD uint) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from `dtb_delivery_time` where `id`=? limit 1)"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, iD)
	}

	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "model: unable to check if dtb_delivery_time exists")
	}

	return exists, nil
}
