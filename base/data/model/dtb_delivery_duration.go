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

// DTBDeliveryDuration is an object representing the database table.
type DTBDeliveryDuration struct {
	ID                uint        `boil:"id" json:"id" toml:"id" yaml:"id"`
	Name              null.String `boil:"name" json:"name,omitempty" toml:"name" yaml:"name,omitempty"`
	Duration          uint16      `boil:"duration" json:"duration" toml:"duration" yaml:"duration"`
	SortNo            uint        `boil:"sort_no" json:"sort_no" toml:"sort_no" yaml:"sort_no"`
	DiscriminatorType string      `boil:"discriminator_type" json:"discriminator_type" toml:"discriminator_type" yaml:"discriminator_type"`

	R *dtbDeliveryDurationR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L dtbDeliveryDurationL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var DTBDeliveryDurationColumns = struct {
	ID                string
	Name              string
	Duration          string
	SortNo            string
	DiscriminatorType string
}{
	ID:                "id",
	Name:              "name",
	Duration:          "duration",
	SortNo:            "sort_no",
	DiscriminatorType: "discriminator_type",
}

// Generated where

var DTBDeliveryDurationWhere = struct {
	ID                whereHelperuint
	Name              whereHelpernull_String
	Duration          whereHelperuint16
	SortNo            whereHelperuint
	DiscriminatorType whereHelperstring
}{
	ID:                whereHelperuint{field: `id`},
	Name:              whereHelpernull_String{field: `name`},
	Duration:          whereHelperuint16{field: `duration`},
	SortNo:            whereHelperuint{field: `sort_no`},
	DiscriminatorType: whereHelperstring{field: `discriminator_type`},
}

// DTBDeliveryDurationRels is where relationship names are stored.
var DTBDeliveryDurationRels = struct {
}{}

// dtbDeliveryDurationR is where relationships are stored.
type dtbDeliveryDurationR struct {
}

// NewStruct creates a new relationship struct
func (*dtbDeliveryDurationR) NewStruct() *dtbDeliveryDurationR {
	return &dtbDeliveryDurationR{}
}

// dtbDeliveryDurationL is where Load methods for each relationship are stored.
type dtbDeliveryDurationL struct{}

var (
	dtbDeliveryDurationColumns               = []string{"id", "name", "duration", "sort_no", "discriminator_type"}
	dtbDeliveryDurationColumnsWithoutDefault = []string{"name", "sort_no", "discriminator_type"}
	dtbDeliveryDurationColumnsWithDefault    = []string{"id", "duration"}
	dtbDeliveryDurationPrimaryKeyColumns     = []string{"id"}
)

type (
	// DTBDeliveryDurationSlice is an alias for a slice of pointers to DTBDeliveryDuration.
	// This should generally be used opposed to []DTBDeliveryDuration.
	DTBDeliveryDurationSlice []*DTBDeliveryDuration
	// DTBDeliveryDurationHook is the signature for custom DTBDeliveryDuration hook methods
	DTBDeliveryDurationHook func(context.Context, boil.ContextExecutor, *DTBDeliveryDuration) error

	dtbDeliveryDurationQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	dtbDeliveryDurationType                 = reflect.TypeOf(&DTBDeliveryDuration{})
	dtbDeliveryDurationMapping              = queries.MakeStructMapping(dtbDeliveryDurationType)
	dtbDeliveryDurationPrimaryKeyMapping, _ = queries.BindMapping(dtbDeliveryDurationType, dtbDeliveryDurationMapping, dtbDeliveryDurationPrimaryKeyColumns)
	dtbDeliveryDurationInsertCacheMut       sync.RWMutex
	dtbDeliveryDurationInsertCache          = make(map[string]insertCache)
	dtbDeliveryDurationUpdateCacheMut       sync.RWMutex
	dtbDeliveryDurationUpdateCache          = make(map[string]updateCache)
	dtbDeliveryDurationUpsertCacheMut       sync.RWMutex
	dtbDeliveryDurationUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var dtbDeliveryDurationBeforeInsertHooks []DTBDeliveryDurationHook
var dtbDeliveryDurationBeforeUpdateHooks []DTBDeliveryDurationHook
var dtbDeliveryDurationBeforeDeleteHooks []DTBDeliveryDurationHook
var dtbDeliveryDurationBeforeUpsertHooks []DTBDeliveryDurationHook

var dtbDeliveryDurationAfterInsertHooks []DTBDeliveryDurationHook
var dtbDeliveryDurationAfterSelectHooks []DTBDeliveryDurationHook
var dtbDeliveryDurationAfterUpdateHooks []DTBDeliveryDurationHook
var dtbDeliveryDurationAfterDeleteHooks []DTBDeliveryDurationHook
var dtbDeliveryDurationAfterUpsertHooks []DTBDeliveryDurationHook

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *DTBDeliveryDuration) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range dtbDeliveryDurationBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *DTBDeliveryDuration) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range dtbDeliveryDurationBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *DTBDeliveryDuration) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range dtbDeliveryDurationBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *DTBDeliveryDuration) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range dtbDeliveryDurationBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *DTBDeliveryDuration) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range dtbDeliveryDurationAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterSelectHooks executes all "after Select" hooks.
func (o *DTBDeliveryDuration) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range dtbDeliveryDurationAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *DTBDeliveryDuration) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range dtbDeliveryDurationAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *DTBDeliveryDuration) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range dtbDeliveryDurationAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *DTBDeliveryDuration) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range dtbDeliveryDurationAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddDTBDeliveryDurationHook registers your hook function for all future operations.
func AddDTBDeliveryDurationHook(hookPoint boil.HookPoint, dtbDeliveryDurationHook DTBDeliveryDurationHook) {
	switch hookPoint {
	case boil.BeforeInsertHook:
		dtbDeliveryDurationBeforeInsertHooks = append(dtbDeliveryDurationBeforeInsertHooks, dtbDeliveryDurationHook)
	case boil.BeforeUpdateHook:
		dtbDeliveryDurationBeforeUpdateHooks = append(dtbDeliveryDurationBeforeUpdateHooks, dtbDeliveryDurationHook)
	case boil.BeforeDeleteHook:
		dtbDeliveryDurationBeforeDeleteHooks = append(dtbDeliveryDurationBeforeDeleteHooks, dtbDeliveryDurationHook)
	case boil.BeforeUpsertHook:
		dtbDeliveryDurationBeforeUpsertHooks = append(dtbDeliveryDurationBeforeUpsertHooks, dtbDeliveryDurationHook)
	case boil.AfterInsertHook:
		dtbDeliveryDurationAfterInsertHooks = append(dtbDeliveryDurationAfterInsertHooks, dtbDeliveryDurationHook)
	case boil.AfterSelectHook:
		dtbDeliveryDurationAfterSelectHooks = append(dtbDeliveryDurationAfterSelectHooks, dtbDeliveryDurationHook)
	case boil.AfterUpdateHook:
		dtbDeliveryDurationAfterUpdateHooks = append(dtbDeliveryDurationAfterUpdateHooks, dtbDeliveryDurationHook)
	case boil.AfterDeleteHook:
		dtbDeliveryDurationAfterDeleteHooks = append(dtbDeliveryDurationAfterDeleteHooks, dtbDeliveryDurationHook)
	case boil.AfterUpsertHook:
		dtbDeliveryDurationAfterUpsertHooks = append(dtbDeliveryDurationAfterUpsertHooks, dtbDeliveryDurationHook)
	}
}

// One returns a single dtbDeliveryDuration record from the query.
func (q dtbDeliveryDurationQuery) One(ctx context.Context, exec boil.ContextExecutor) (*DTBDeliveryDuration, error) {
	o := &DTBDeliveryDuration{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "model: failed to execute a one query for dtb_delivery_duration")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all DTBDeliveryDuration records from the query.
func (q dtbDeliveryDurationQuery) All(ctx context.Context, exec boil.ContextExecutor) (DTBDeliveryDurationSlice, error) {
	var o []*DTBDeliveryDuration

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "model: failed to assign all query results to DTBDeliveryDuration slice")
	}

	if len(dtbDeliveryDurationAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all DTBDeliveryDuration records in the query.
func (q dtbDeliveryDurationQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "model: failed to count dtb_delivery_duration rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q dtbDeliveryDurationQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "model: failed to check if dtb_delivery_duration exists")
	}

	return count > 0, nil
}

// DTBDeliveryDurations retrieves all the records using an executor.
func DTBDeliveryDurations(mods ...qm.QueryMod) dtbDeliveryDurationQuery {
	mods = append(mods, qm.From("`dtb_delivery_duration`"))
	return dtbDeliveryDurationQuery{NewQuery(mods...)}
}

// FindDTBDeliveryDuration retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindDTBDeliveryDuration(ctx context.Context, exec boil.ContextExecutor, iD uint, selectCols ...string) (*DTBDeliveryDuration, error) {
	dtbDeliveryDurationObj := &DTBDeliveryDuration{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from `dtb_delivery_duration` where `id`=?", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, dtbDeliveryDurationObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "model: unable to select from dtb_delivery_duration")
	}

	return dtbDeliveryDurationObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *DTBDeliveryDuration) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("model: no dtb_delivery_duration provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(dtbDeliveryDurationColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	dtbDeliveryDurationInsertCacheMut.RLock()
	cache, cached := dtbDeliveryDurationInsertCache[key]
	dtbDeliveryDurationInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			dtbDeliveryDurationColumns,
			dtbDeliveryDurationColumnsWithDefault,
			dtbDeliveryDurationColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(dtbDeliveryDurationType, dtbDeliveryDurationMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(dtbDeliveryDurationType, dtbDeliveryDurationMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO `dtb_delivery_duration` (`%s`) %%sVALUES (%s)%%s", strings.Join(wl, "`,`"), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO `dtb_delivery_duration` () VALUES ()%s%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			cache.retQuery = fmt.Sprintf("SELECT `%s` FROM `dtb_delivery_duration` WHERE %s", strings.Join(returnColumns, "`,`"), strmangle.WhereClause("`", "`", 0, dtbDeliveryDurationPrimaryKeyColumns))
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
		return errors.Wrap(err, "model: unable to insert into dtb_delivery_duration")
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
	if lastID != 0 && len(cache.retMapping) == 1 && cache.retMapping[0] == dtbDeliveryDurationMapping["ID"] {
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
		return errors.Wrap(err, "model: unable to populate default values for dtb_delivery_duration")
	}

CacheNoHooks:
	if !cached {
		dtbDeliveryDurationInsertCacheMut.Lock()
		dtbDeliveryDurationInsertCache[key] = cache
		dtbDeliveryDurationInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the DTBDeliveryDuration.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *DTBDeliveryDuration) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	dtbDeliveryDurationUpdateCacheMut.RLock()
	cache, cached := dtbDeliveryDurationUpdateCache[key]
	dtbDeliveryDurationUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			dtbDeliveryDurationColumns,
			dtbDeliveryDurationPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("model: unable to update dtb_delivery_duration, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE `dtb_delivery_duration` SET %s WHERE %s",
			strmangle.SetParamNames("`", "`", 0, wl),
			strmangle.WhereClause("`", "`", 0, dtbDeliveryDurationPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(dtbDeliveryDurationType, dtbDeliveryDurationMapping, append(wl, dtbDeliveryDurationPrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "model: unable to update dtb_delivery_duration row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "model: failed to get rows affected by update for dtb_delivery_duration")
	}

	if !cached {
		dtbDeliveryDurationUpdateCacheMut.Lock()
		dtbDeliveryDurationUpdateCache[key] = cache
		dtbDeliveryDurationUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q dtbDeliveryDurationQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "model: unable to update all for dtb_delivery_duration")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "model: unable to retrieve rows affected for dtb_delivery_duration")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o DTBDeliveryDurationSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), dtbDeliveryDurationPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE `dtb_delivery_duration` SET %s WHERE %s",
		strmangle.SetParamNames("`", "`", 0, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, dtbDeliveryDurationPrimaryKeyColumns, len(o)))

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "model: unable to update all in dtbDeliveryDuration slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "model: unable to retrieve rows affected all in update all dtbDeliveryDuration")
	}
	return rowsAff, nil
}

var mySQLDTBDeliveryDurationUniqueColumns = []string{
	"id",
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *DTBDeliveryDuration) Upsert(ctx context.Context, exec boil.ContextExecutor, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("model: no dtb_delivery_duration provided for upsert")
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(dtbDeliveryDurationColumnsWithDefault, o)
	nzUniques := queries.NonZeroDefaultSet(mySQLDTBDeliveryDurationUniqueColumns, o)

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

	dtbDeliveryDurationUpsertCacheMut.RLock()
	cache, cached := dtbDeliveryDurationUpsertCache[key]
	dtbDeliveryDurationUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			dtbDeliveryDurationColumns,
			dtbDeliveryDurationColumnsWithDefault,
			dtbDeliveryDurationColumnsWithoutDefault,
			nzDefaults,
		)
		update := updateColumns.UpdateColumnSet(
			dtbDeliveryDurationColumns,
			dtbDeliveryDurationPrimaryKeyColumns,
		)

		if len(update) == 0 {
			return errors.New("model: unable to upsert dtb_delivery_duration, could not build update column list")
		}

		ret = strmangle.SetComplement(ret, nzUniques)
		cache.query = buildUpsertQueryMySQL(dialect, "dtb_delivery_duration", update, insert)
		cache.retQuery = fmt.Sprintf(
			"SELECT %s FROM `dtb_delivery_duration` WHERE %s",
			strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, ret), ","),
			strmangle.WhereClause("`", "`", 0, nzUniques),
		)

		cache.valueMapping, err = queries.BindMapping(dtbDeliveryDurationType, dtbDeliveryDurationMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(dtbDeliveryDurationType, dtbDeliveryDurationMapping, ret)
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
		return errors.Wrap(err, "model: unable to upsert for dtb_delivery_duration")
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
	if lastID != 0 && len(cache.retMapping) == 1 && cache.retMapping[0] == dtbDeliveryDurationMapping["id"] {
		goto CacheNoHooks
	}

	uniqueMap, err = queries.BindMapping(dtbDeliveryDurationType, dtbDeliveryDurationMapping, nzUniques)
	if err != nil {
		return errors.Wrap(err, "model: unable to retrieve unique values for dtb_delivery_duration")
	}
	nzUniqueCols = queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), uniqueMap)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.retQuery)
		fmt.Fprintln(boil.DebugWriter, nzUniqueCols...)
	}

	err = exec.QueryRowContext(ctx, cache.retQuery, nzUniqueCols...).Scan(returns...)
	if err != nil {
		return errors.Wrap(err, "model: unable to populate default values for dtb_delivery_duration")
	}

CacheNoHooks:
	if !cached {
		dtbDeliveryDurationUpsertCacheMut.Lock()
		dtbDeliveryDurationUpsertCache[key] = cache
		dtbDeliveryDurationUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single DTBDeliveryDuration record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *DTBDeliveryDuration) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("model: no DTBDeliveryDuration provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), dtbDeliveryDurationPrimaryKeyMapping)
	sql := "DELETE FROM `dtb_delivery_duration` WHERE `id`=?"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "model: unable to delete from dtb_delivery_duration")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "model: failed to get rows affected by delete for dtb_delivery_duration")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q dtbDeliveryDurationQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("model: no dtbDeliveryDurationQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "model: unable to delete all from dtb_delivery_duration")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "model: failed to get rows affected by deleteall for dtb_delivery_duration")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o DTBDeliveryDurationSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("model: no DTBDeliveryDuration slice provided for delete all")
	}

	if len(o) == 0 {
		return 0, nil
	}

	if len(dtbDeliveryDurationBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), dtbDeliveryDurationPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM `dtb_delivery_duration` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, dtbDeliveryDurationPrimaryKeyColumns, len(o))

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args)
	}

	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "model: unable to delete all from dtbDeliveryDuration slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "model: failed to get rows affected by deleteall for dtb_delivery_duration")
	}

	if len(dtbDeliveryDurationAfterDeleteHooks) != 0 {
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
func (o *DTBDeliveryDuration) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindDTBDeliveryDuration(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *DTBDeliveryDurationSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := DTBDeliveryDurationSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), dtbDeliveryDurationPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT `dtb_delivery_duration`.* FROM `dtb_delivery_duration` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, dtbDeliveryDurationPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "model: unable to reload all in DTBDeliveryDurationSlice")
	}

	*o = slice

	return nil
}

// DTBDeliveryDurationExists checks if the DTBDeliveryDuration row exists.
func DTBDeliveryDurationExists(ctx context.Context, exec boil.ContextExecutor, iD uint) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from `dtb_delivery_duration` where `id`=? limit 1)"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, iD)
	}

	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "model: unable to check if dtb_delivery_duration exists")
	}

	return exists, nil
}
