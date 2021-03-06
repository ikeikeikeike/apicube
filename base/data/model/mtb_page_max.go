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

// MTBPageMax is an object representing the database table.
type MTBPageMax struct {
	ID                uint16 `boil:"id" json:"id" toml:"id" yaml:"id"`
	Name              string `boil:"name" json:"name" toml:"name" yaml:"name"`
	SortNo            uint16 `boil:"sort_no" json:"sort_no" toml:"sort_no" yaml:"sort_no"`
	DiscriminatorType string `boil:"discriminator_type" json:"discriminator_type" toml:"discriminator_type" yaml:"discriminator_type"`

	R *mtbPageMaxR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L mtbPageMaxL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var MTBPageMaxColumns = struct {
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

var MTBPageMaxWhere = struct {
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

// MTBPageMaxRels is where relationship names are stored.
var MTBPageMaxRels = struct {
}{}

// mtbPageMaxR is where relationships are stored.
type mtbPageMaxR struct {
}

// NewStruct creates a new relationship struct
func (*mtbPageMaxR) NewStruct() *mtbPageMaxR {
	return &mtbPageMaxR{}
}

// mtbPageMaxL is where Load methods for each relationship are stored.
type mtbPageMaxL struct{}

var (
	mtbPageMaxColumns               = []string{"id", "name", "sort_no", "discriminator_type"}
	mtbPageMaxColumnsWithoutDefault = []string{"id", "name", "sort_no", "discriminator_type"}
	mtbPageMaxColumnsWithDefault    = []string{}
	mtbPageMaxPrimaryKeyColumns     = []string{"id"}
)

type (
	// MTBPageMaxSlice is an alias for a slice of pointers to MTBPageMax.
	// This should generally be used opposed to []MTBPageMax.
	MTBPageMaxSlice []*MTBPageMax
	// MTBPageMaxHook is the signature for custom MTBPageMax hook methods
	MTBPageMaxHook func(context.Context, boil.ContextExecutor, *MTBPageMax) error

	mtbPageMaxQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	mtbPageMaxType                 = reflect.TypeOf(&MTBPageMax{})
	mtbPageMaxMapping              = queries.MakeStructMapping(mtbPageMaxType)
	mtbPageMaxPrimaryKeyMapping, _ = queries.BindMapping(mtbPageMaxType, mtbPageMaxMapping, mtbPageMaxPrimaryKeyColumns)
	mtbPageMaxInsertCacheMut       sync.RWMutex
	mtbPageMaxInsertCache          = make(map[string]insertCache)
	mtbPageMaxUpdateCacheMut       sync.RWMutex
	mtbPageMaxUpdateCache          = make(map[string]updateCache)
	mtbPageMaxUpsertCacheMut       sync.RWMutex
	mtbPageMaxUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var mtbPageMaxBeforeInsertHooks []MTBPageMaxHook
var mtbPageMaxBeforeUpdateHooks []MTBPageMaxHook
var mtbPageMaxBeforeDeleteHooks []MTBPageMaxHook
var mtbPageMaxBeforeUpsertHooks []MTBPageMaxHook

var mtbPageMaxAfterInsertHooks []MTBPageMaxHook
var mtbPageMaxAfterSelectHooks []MTBPageMaxHook
var mtbPageMaxAfterUpdateHooks []MTBPageMaxHook
var mtbPageMaxAfterDeleteHooks []MTBPageMaxHook
var mtbPageMaxAfterUpsertHooks []MTBPageMaxHook

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *MTBPageMax) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range mtbPageMaxBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *MTBPageMax) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range mtbPageMaxBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *MTBPageMax) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range mtbPageMaxBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *MTBPageMax) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range mtbPageMaxBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *MTBPageMax) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range mtbPageMaxAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterSelectHooks executes all "after Select" hooks.
func (o *MTBPageMax) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range mtbPageMaxAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *MTBPageMax) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range mtbPageMaxAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *MTBPageMax) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range mtbPageMaxAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *MTBPageMax) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range mtbPageMaxAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddMTBPageMaxHook registers your hook function for all future operations.
func AddMTBPageMaxHook(hookPoint boil.HookPoint, mtbPageMaxHook MTBPageMaxHook) {
	switch hookPoint {
	case boil.BeforeInsertHook:
		mtbPageMaxBeforeInsertHooks = append(mtbPageMaxBeforeInsertHooks, mtbPageMaxHook)
	case boil.BeforeUpdateHook:
		mtbPageMaxBeforeUpdateHooks = append(mtbPageMaxBeforeUpdateHooks, mtbPageMaxHook)
	case boil.BeforeDeleteHook:
		mtbPageMaxBeforeDeleteHooks = append(mtbPageMaxBeforeDeleteHooks, mtbPageMaxHook)
	case boil.BeforeUpsertHook:
		mtbPageMaxBeforeUpsertHooks = append(mtbPageMaxBeforeUpsertHooks, mtbPageMaxHook)
	case boil.AfterInsertHook:
		mtbPageMaxAfterInsertHooks = append(mtbPageMaxAfterInsertHooks, mtbPageMaxHook)
	case boil.AfterSelectHook:
		mtbPageMaxAfterSelectHooks = append(mtbPageMaxAfterSelectHooks, mtbPageMaxHook)
	case boil.AfterUpdateHook:
		mtbPageMaxAfterUpdateHooks = append(mtbPageMaxAfterUpdateHooks, mtbPageMaxHook)
	case boil.AfterDeleteHook:
		mtbPageMaxAfterDeleteHooks = append(mtbPageMaxAfterDeleteHooks, mtbPageMaxHook)
	case boil.AfterUpsertHook:
		mtbPageMaxAfterUpsertHooks = append(mtbPageMaxAfterUpsertHooks, mtbPageMaxHook)
	}
}

// One returns a single mtbPageMax record from the query.
func (q mtbPageMaxQuery) One(ctx context.Context, exec boil.ContextExecutor) (*MTBPageMax, error) {
	o := &MTBPageMax{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "model: failed to execute a one query for mtb_page_max")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all MTBPageMax records from the query.
func (q mtbPageMaxQuery) All(ctx context.Context, exec boil.ContextExecutor) (MTBPageMaxSlice, error) {
	var o []*MTBPageMax

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "model: failed to assign all query results to MTBPageMax slice")
	}

	if len(mtbPageMaxAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all MTBPageMax records in the query.
func (q mtbPageMaxQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "model: failed to count mtb_page_max rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q mtbPageMaxQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "model: failed to check if mtb_page_max exists")
	}

	return count > 0, nil
}

// MTBPageMaxes retrieves all the records using an executor.
func MTBPageMaxes(mods ...qm.QueryMod) mtbPageMaxQuery {
	mods = append(mods, qm.From("`mtb_page_max`"))
	return mtbPageMaxQuery{NewQuery(mods...)}
}

// FindMTBPageMax retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindMTBPageMax(ctx context.Context, exec boil.ContextExecutor, iD uint16, selectCols ...string) (*MTBPageMax, error) {
	mtbPageMaxObj := &MTBPageMax{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from `mtb_page_max` where `id`=?", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, mtbPageMaxObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "model: unable to select from mtb_page_max")
	}

	return mtbPageMaxObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *MTBPageMax) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("model: no mtb_page_max provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(mtbPageMaxColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	mtbPageMaxInsertCacheMut.RLock()
	cache, cached := mtbPageMaxInsertCache[key]
	mtbPageMaxInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			mtbPageMaxColumns,
			mtbPageMaxColumnsWithDefault,
			mtbPageMaxColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(mtbPageMaxType, mtbPageMaxMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(mtbPageMaxType, mtbPageMaxMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO `mtb_page_max` (`%s`) %%sVALUES (%s)%%s", strings.Join(wl, "`,`"), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO `mtb_page_max` () VALUES ()%s%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			cache.retQuery = fmt.Sprintf("SELECT `%s` FROM `mtb_page_max` WHERE %s", strings.Join(returnColumns, "`,`"), strmangle.WhereClause("`", "`", 0, mtbPageMaxPrimaryKeyColumns))
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
		return errors.Wrap(err, "model: unable to insert into mtb_page_max")
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
		return errors.Wrap(err, "model: unable to populate default values for mtb_page_max")
	}

CacheNoHooks:
	if !cached {
		mtbPageMaxInsertCacheMut.Lock()
		mtbPageMaxInsertCache[key] = cache
		mtbPageMaxInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the MTBPageMax.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *MTBPageMax) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	mtbPageMaxUpdateCacheMut.RLock()
	cache, cached := mtbPageMaxUpdateCache[key]
	mtbPageMaxUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			mtbPageMaxColumns,
			mtbPageMaxPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("model: unable to update mtb_page_max, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE `mtb_page_max` SET %s WHERE %s",
			strmangle.SetParamNames("`", "`", 0, wl),
			strmangle.WhereClause("`", "`", 0, mtbPageMaxPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(mtbPageMaxType, mtbPageMaxMapping, append(wl, mtbPageMaxPrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "model: unable to update mtb_page_max row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "model: failed to get rows affected by update for mtb_page_max")
	}

	if !cached {
		mtbPageMaxUpdateCacheMut.Lock()
		mtbPageMaxUpdateCache[key] = cache
		mtbPageMaxUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q mtbPageMaxQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "model: unable to update all for mtb_page_max")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "model: unable to retrieve rows affected for mtb_page_max")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o MTBPageMaxSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), mtbPageMaxPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE `mtb_page_max` SET %s WHERE %s",
		strmangle.SetParamNames("`", "`", 0, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, mtbPageMaxPrimaryKeyColumns, len(o)))

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "model: unable to update all in mtbPageMax slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "model: unable to retrieve rows affected all in update all mtbPageMax")
	}
	return rowsAff, nil
}

var mySQLMTBPageMaxUniqueColumns = []string{
	"id",
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *MTBPageMax) Upsert(ctx context.Context, exec boil.ContextExecutor, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("model: no mtb_page_max provided for upsert")
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(mtbPageMaxColumnsWithDefault, o)
	nzUniques := queries.NonZeroDefaultSet(mySQLMTBPageMaxUniqueColumns, o)

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

	mtbPageMaxUpsertCacheMut.RLock()
	cache, cached := mtbPageMaxUpsertCache[key]
	mtbPageMaxUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			mtbPageMaxColumns,
			mtbPageMaxColumnsWithDefault,
			mtbPageMaxColumnsWithoutDefault,
			nzDefaults,
		)
		update := updateColumns.UpdateColumnSet(
			mtbPageMaxColumns,
			mtbPageMaxPrimaryKeyColumns,
		)

		if len(update) == 0 {
			return errors.New("model: unable to upsert mtb_page_max, could not build update column list")
		}

		ret = strmangle.SetComplement(ret, nzUniques)
		cache.query = buildUpsertQueryMySQL(dialect, "mtb_page_max", update, insert)
		cache.retQuery = fmt.Sprintf(
			"SELECT %s FROM `mtb_page_max` WHERE %s",
			strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, ret), ","),
			strmangle.WhereClause("`", "`", 0, nzUniques),
		)

		cache.valueMapping, err = queries.BindMapping(mtbPageMaxType, mtbPageMaxMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(mtbPageMaxType, mtbPageMaxMapping, ret)
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
		return errors.Wrap(err, "model: unable to upsert for mtb_page_max")
	}

	var uniqueMap []uint64
	var nzUniqueCols []interface{}

	if len(cache.retMapping) == 0 {
		goto CacheNoHooks
	}

	uniqueMap, err = queries.BindMapping(mtbPageMaxType, mtbPageMaxMapping, nzUniques)
	if err != nil {
		return errors.Wrap(err, "model: unable to retrieve unique values for mtb_page_max")
	}
	nzUniqueCols = queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), uniqueMap)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.retQuery)
		fmt.Fprintln(boil.DebugWriter, nzUniqueCols...)
	}

	err = exec.QueryRowContext(ctx, cache.retQuery, nzUniqueCols...).Scan(returns...)
	if err != nil {
		return errors.Wrap(err, "model: unable to populate default values for mtb_page_max")
	}

CacheNoHooks:
	if !cached {
		mtbPageMaxUpsertCacheMut.Lock()
		mtbPageMaxUpsertCache[key] = cache
		mtbPageMaxUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single MTBPageMax record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *MTBPageMax) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("model: no MTBPageMax provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), mtbPageMaxPrimaryKeyMapping)
	sql := "DELETE FROM `mtb_page_max` WHERE `id`=?"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "model: unable to delete from mtb_page_max")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "model: failed to get rows affected by delete for mtb_page_max")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q mtbPageMaxQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("model: no mtbPageMaxQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "model: unable to delete all from mtb_page_max")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "model: failed to get rows affected by deleteall for mtb_page_max")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o MTBPageMaxSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("model: no MTBPageMax slice provided for delete all")
	}

	if len(o) == 0 {
		return 0, nil
	}

	if len(mtbPageMaxBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), mtbPageMaxPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM `mtb_page_max` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, mtbPageMaxPrimaryKeyColumns, len(o))

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args)
	}

	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "model: unable to delete all from mtbPageMax slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "model: failed to get rows affected by deleteall for mtb_page_max")
	}

	if len(mtbPageMaxAfterDeleteHooks) != 0 {
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
func (o *MTBPageMax) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindMTBPageMax(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *MTBPageMaxSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := MTBPageMaxSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), mtbPageMaxPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT `mtb_page_max`.* FROM `mtb_page_max` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, mtbPageMaxPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "model: unable to reload all in MTBPageMaxSlice")
	}

	*o = slice

	return nil
}

// MTBPageMaxExists checks if the MTBPageMax row exists.
func MTBPageMaxExists(ctx context.Context, exec boil.ContextExecutor, iD uint16) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from `mtb_page_max` where `id`=? limit 1)"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, iD)
	}

	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "model: unable to check if mtb_page_max exists")
	}

	return exists, nil
}
