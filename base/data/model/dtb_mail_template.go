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

// DTBMailTemplate is an object representing the database table.
type DTBMailTemplate struct {
	ID                uint        `boil:"id" json:"id" toml:"id" yaml:"id"`
	CreatorID         null.Uint   `boil:"creator_id" json:"creator_id,omitempty" toml:"creator_id" yaml:"creator_id,omitempty"`
	Name              null.String `boil:"name" json:"name,omitempty" toml:"name" yaml:"name,omitempty"`
	FileName          null.String `boil:"file_name" json:"file_name,omitempty" toml:"file_name" yaml:"file_name,omitempty"`
	MailSubject       null.String `boil:"mail_subject" json:"mail_subject,omitempty" toml:"mail_subject" yaml:"mail_subject,omitempty"`
	CreateDate        time.Time   `boil:"create_date" json:"create_date" toml:"create_date" yaml:"create_date"`
	UpdateDate        time.Time   `boil:"update_date" json:"update_date" toml:"update_date" yaml:"update_date"`
	DiscriminatorType string      `boil:"discriminator_type" json:"discriminator_type" toml:"discriminator_type" yaml:"discriminator_type"`

	R *dtbMailTemplateR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L dtbMailTemplateL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var DTBMailTemplateColumns = struct {
	ID                string
	CreatorID         string
	Name              string
	FileName          string
	MailSubject       string
	CreateDate        string
	UpdateDate        string
	DiscriminatorType string
}{
	ID:                "id",
	CreatorID:         "creator_id",
	Name:              "name",
	FileName:          "file_name",
	MailSubject:       "mail_subject",
	CreateDate:        "create_date",
	UpdateDate:        "update_date",
	DiscriminatorType: "discriminator_type",
}

// Generated where

var DTBMailTemplateWhere = struct {
	ID                whereHelperuint
	CreatorID         whereHelpernull_Uint
	Name              whereHelpernull_String
	FileName          whereHelpernull_String
	MailSubject       whereHelpernull_String
	CreateDate        whereHelpertime_Time
	UpdateDate        whereHelpertime_Time
	DiscriminatorType whereHelperstring
}{
	ID:                whereHelperuint{field: `id`},
	CreatorID:         whereHelpernull_Uint{field: `creator_id`},
	Name:              whereHelpernull_String{field: `name`},
	FileName:          whereHelpernull_String{field: `file_name`},
	MailSubject:       whereHelpernull_String{field: `mail_subject`},
	CreateDate:        whereHelpertime_Time{field: `create_date`},
	UpdateDate:        whereHelpertime_Time{field: `update_date`},
	DiscriminatorType: whereHelperstring{field: `discriminator_type`},
}

// DTBMailTemplateRels is where relationship names are stored.
var DTBMailTemplateRels = struct {
}{}

// dtbMailTemplateR is where relationships are stored.
type dtbMailTemplateR struct {
}

// NewStruct creates a new relationship struct
func (*dtbMailTemplateR) NewStruct() *dtbMailTemplateR {
	return &dtbMailTemplateR{}
}

// dtbMailTemplateL is where Load methods for each relationship are stored.
type dtbMailTemplateL struct{}

var (
	dtbMailTemplateColumns               = []string{"id", "creator_id", "name", "file_name", "mail_subject", "create_date", "update_date", "discriminator_type"}
	dtbMailTemplateColumnsWithoutDefault = []string{"creator_id", "name", "file_name", "mail_subject", "create_date", "update_date", "discriminator_type"}
	dtbMailTemplateColumnsWithDefault    = []string{"id"}
	dtbMailTemplatePrimaryKeyColumns     = []string{"id"}
)

type (
	// DTBMailTemplateSlice is an alias for a slice of pointers to DTBMailTemplate.
	// This should generally be used opposed to []DTBMailTemplate.
	DTBMailTemplateSlice []*DTBMailTemplate
	// DTBMailTemplateHook is the signature for custom DTBMailTemplate hook methods
	DTBMailTemplateHook func(context.Context, boil.ContextExecutor, *DTBMailTemplate) error

	dtbMailTemplateQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	dtbMailTemplateType                 = reflect.TypeOf(&DTBMailTemplate{})
	dtbMailTemplateMapping              = queries.MakeStructMapping(dtbMailTemplateType)
	dtbMailTemplatePrimaryKeyMapping, _ = queries.BindMapping(dtbMailTemplateType, dtbMailTemplateMapping, dtbMailTemplatePrimaryKeyColumns)
	dtbMailTemplateInsertCacheMut       sync.RWMutex
	dtbMailTemplateInsertCache          = make(map[string]insertCache)
	dtbMailTemplateUpdateCacheMut       sync.RWMutex
	dtbMailTemplateUpdateCache          = make(map[string]updateCache)
	dtbMailTemplateUpsertCacheMut       sync.RWMutex
	dtbMailTemplateUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var dtbMailTemplateBeforeInsertHooks []DTBMailTemplateHook
var dtbMailTemplateBeforeUpdateHooks []DTBMailTemplateHook
var dtbMailTemplateBeforeDeleteHooks []DTBMailTemplateHook
var dtbMailTemplateBeforeUpsertHooks []DTBMailTemplateHook

var dtbMailTemplateAfterInsertHooks []DTBMailTemplateHook
var dtbMailTemplateAfterSelectHooks []DTBMailTemplateHook
var dtbMailTemplateAfterUpdateHooks []DTBMailTemplateHook
var dtbMailTemplateAfterDeleteHooks []DTBMailTemplateHook
var dtbMailTemplateAfterUpsertHooks []DTBMailTemplateHook

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *DTBMailTemplate) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range dtbMailTemplateBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *DTBMailTemplate) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range dtbMailTemplateBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *DTBMailTemplate) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range dtbMailTemplateBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *DTBMailTemplate) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range dtbMailTemplateBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *DTBMailTemplate) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range dtbMailTemplateAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterSelectHooks executes all "after Select" hooks.
func (o *DTBMailTemplate) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range dtbMailTemplateAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *DTBMailTemplate) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range dtbMailTemplateAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *DTBMailTemplate) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range dtbMailTemplateAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *DTBMailTemplate) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range dtbMailTemplateAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddDTBMailTemplateHook registers your hook function for all future operations.
func AddDTBMailTemplateHook(hookPoint boil.HookPoint, dtbMailTemplateHook DTBMailTemplateHook) {
	switch hookPoint {
	case boil.BeforeInsertHook:
		dtbMailTemplateBeforeInsertHooks = append(dtbMailTemplateBeforeInsertHooks, dtbMailTemplateHook)
	case boil.BeforeUpdateHook:
		dtbMailTemplateBeforeUpdateHooks = append(dtbMailTemplateBeforeUpdateHooks, dtbMailTemplateHook)
	case boil.BeforeDeleteHook:
		dtbMailTemplateBeforeDeleteHooks = append(dtbMailTemplateBeforeDeleteHooks, dtbMailTemplateHook)
	case boil.BeforeUpsertHook:
		dtbMailTemplateBeforeUpsertHooks = append(dtbMailTemplateBeforeUpsertHooks, dtbMailTemplateHook)
	case boil.AfterInsertHook:
		dtbMailTemplateAfterInsertHooks = append(dtbMailTemplateAfterInsertHooks, dtbMailTemplateHook)
	case boil.AfterSelectHook:
		dtbMailTemplateAfterSelectHooks = append(dtbMailTemplateAfterSelectHooks, dtbMailTemplateHook)
	case boil.AfterUpdateHook:
		dtbMailTemplateAfterUpdateHooks = append(dtbMailTemplateAfterUpdateHooks, dtbMailTemplateHook)
	case boil.AfterDeleteHook:
		dtbMailTemplateAfterDeleteHooks = append(dtbMailTemplateAfterDeleteHooks, dtbMailTemplateHook)
	case boil.AfterUpsertHook:
		dtbMailTemplateAfterUpsertHooks = append(dtbMailTemplateAfterUpsertHooks, dtbMailTemplateHook)
	}
}

// One returns a single dtbMailTemplate record from the query.
func (q dtbMailTemplateQuery) One(ctx context.Context, exec boil.ContextExecutor) (*DTBMailTemplate, error) {
	o := &DTBMailTemplate{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "model: failed to execute a one query for dtb_mail_template")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all DTBMailTemplate records from the query.
func (q dtbMailTemplateQuery) All(ctx context.Context, exec boil.ContextExecutor) (DTBMailTemplateSlice, error) {
	var o []*DTBMailTemplate

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "model: failed to assign all query results to DTBMailTemplate slice")
	}

	if len(dtbMailTemplateAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all DTBMailTemplate records in the query.
func (q dtbMailTemplateQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "model: failed to count dtb_mail_template rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q dtbMailTemplateQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "model: failed to check if dtb_mail_template exists")
	}

	return count > 0, nil
}

// DTBMailTemplates retrieves all the records using an executor.
func DTBMailTemplates(mods ...qm.QueryMod) dtbMailTemplateQuery {
	mods = append(mods, qm.From("`dtb_mail_template`"))
	return dtbMailTemplateQuery{NewQuery(mods...)}
}

// FindDTBMailTemplate retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindDTBMailTemplate(ctx context.Context, exec boil.ContextExecutor, iD uint, selectCols ...string) (*DTBMailTemplate, error) {
	dtbMailTemplateObj := &DTBMailTemplate{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from `dtb_mail_template` where `id`=?", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, dtbMailTemplateObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "model: unable to select from dtb_mail_template")
	}

	return dtbMailTemplateObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *DTBMailTemplate) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("model: no dtb_mail_template provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(dtbMailTemplateColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	dtbMailTemplateInsertCacheMut.RLock()
	cache, cached := dtbMailTemplateInsertCache[key]
	dtbMailTemplateInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			dtbMailTemplateColumns,
			dtbMailTemplateColumnsWithDefault,
			dtbMailTemplateColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(dtbMailTemplateType, dtbMailTemplateMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(dtbMailTemplateType, dtbMailTemplateMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO `dtb_mail_template` (`%s`) %%sVALUES (%s)%%s", strings.Join(wl, "`,`"), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO `dtb_mail_template` () VALUES ()%s%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			cache.retQuery = fmt.Sprintf("SELECT `%s` FROM `dtb_mail_template` WHERE %s", strings.Join(returnColumns, "`,`"), strmangle.WhereClause("`", "`", 0, dtbMailTemplatePrimaryKeyColumns))
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
		return errors.Wrap(err, "model: unable to insert into dtb_mail_template")
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
	if lastID != 0 && len(cache.retMapping) == 1 && cache.retMapping[0] == dtbMailTemplateMapping["ID"] {
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
		return errors.Wrap(err, "model: unable to populate default values for dtb_mail_template")
	}

CacheNoHooks:
	if !cached {
		dtbMailTemplateInsertCacheMut.Lock()
		dtbMailTemplateInsertCache[key] = cache
		dtbMailTemplateInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the DTBMailTemplate.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *DTBMailTemplate) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	dtbMailTemplateUpdateCacheMut.RLock()
	cache, cached := dtbMailTemplateUpdateCache[key]
	dtbMailTemplateUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			dtbMailTemplateColumns,
			dtbMailTemplatePrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("model: unable to update dtb_mail_template, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE `dtb_mail_template` SET %s WHERE %s",
			strmangle.SetParamNames("`", "`", 0, wl),
			strmangle.WhereClause("`", "`", 0, dtbMailTemplatePrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(dtbMailTemplateType, dtbMailTemplateMapping, append(wl, dtbMailTemplatePrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "model: unable to update dtb_mail_template row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "model: failed to get rows affected by update for dtb_mail_template")
	}

	if !cached {
		dtbMailTemplateUpdateCacheMut.Lock()
		dtbMailTemplateUpdateCache[key] = cache
		dtbMailTemplateUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q dtbMailTemplateQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "model: unable to update all for dtb_mail_template")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "model: unable to retrieve rows affected for dtb_mail_template")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o DTBMailTemplateSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), dtbMailTemplatePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE `dtb_mail_template` SET %s WHERE %s",
		strmangle.SetParamNames("`", "`", 0, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, dtbMailTemplatePrimaryKeyColumns, len(o)))

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "model: unable to update all in dtbMailTemplate slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "model: unable to retrieve rows affected all in update all dtbMailTemplate")
	}
	return rowsAff, nil
}

var mySQLDTBMailTemplateUniqueColumns = []string{
	"id",
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *DTBMailTemplate) Upsert(ctx context.Context, exec boil.ContextExecutor, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("model: no dtb_mail_template provided for upsert")
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(dtbMailTemplateColumnsWithDefault, o)
	nzUniques := queries.NonZeroDefaultSet(mySQLDTBMailTemplateUniqueColumns, o)

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

	dtbMailTemplateUpsertCacheMut.RLock()
	cache, cached := dtbMailTemplateUpsertCache[key]
	dtbMailTemplateUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			dtbMailTemplateColumns,
			dtbMailTemplateColumnsWithDefault,
			dtbMailTemplateColumnsWithoutDefault,
			nzDefaults,
		)
		update := updateColumns.UpdateColumnSet(
			dtbMailTemplateColumns,
			dtbMailTemplatePrimaryKeyColumns,
		)

		if len(update) == 0 {
			return errors.New("model: unable to upsert dtb_mail_template, could not build update column list")
		}

		ret = strmangle.SetComplement(ret, nzUniques)
		cache.query = buildUpsertQueryMySQL(dialect, "dtb_mail_template", update, insert)
		cache.retQuery = fmt.Sprintf(
			"SELECT %s FROM `dtb_mail_template` WHERE %s",
			strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, ret), ","),
			strmangle.WhereClause("`", "`", 0, nzUniques),
		)

		cache.valueMapping, err = queries.BindMapping(dtbMailTemplateType, dtbMailTemplateMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(dtbMailTemplateType, dtbMailTemplateMapping, ret)
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
		return errors.Wrap(err, "model: unable to upsert for dtb_mail_template")
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
	if lastID != 0 && len(cache.retMapping) == 1 && cache.retMapping[0] == dtbMailTemplateMapping["id"] {
		goto CacheNoHooks
	}

	uniqueMap, err = queries.BindMapping(dtbMailTemplateType, dtbMailTemplateMapping, nzUniques)
	if err != nil {
		return errors.Wrap(err, "model: unable to retrieve unique values for dtb_mail_template")
	}
	nzUniqueCols = queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), uniqueMap)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.retQuery)
		fmt.Fprintln(boil.DebugWriter, nzUniqueCols...)
	}

	err = exec.QueryRowContext(ctx, cache.retQuery, nzUniqueCols...).Scan(returns...)
	if err != nil {
		return errors.Wrap(err, "model: unable to populate default values for dtb_mail_template")
	}

CacheNoHooks:
	if !cached {
		dtbMailTemplateUpsertCacheMut.Lock()
		dtbMailTemplateUpsertCache[key] = cache
		dtbMailTemplateUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single DTBMailTemplate record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *DTBMailTemplate) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("model: no DTBMailTemplate provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), dtbMailTemplatePrimaryKeyMapping)
	sql := "DELETE FROM `dtb_mail_template` WHERE `id`=?"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "model: unable to delete from dtb_mail_template")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "model: failed to get rows affected by delete for dtb_mail_template")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q dtbMailTemplateQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("model: no dtbMailTemplateQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "model: unable to delete all from dtb_mail_template")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "model: failed to get rows affected by deleteall for dtb_mail_template")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o DTBMailTemplateSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("model: no DTBMailTemplate slice provided for delete all")
	}

	if len(o) == 0 {
		return 0, nil
	}

	if len(dtbMailTemplateBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), dtbMailTemplatePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM `dtb_mail_template` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, dtbMailTemplatePrimaryKeyColumns, len(o))

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args)
	}

	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "model: unable to delete all from dtbMailTemplate slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "model: failed to get rows affected by deleteall for dtb_mail_template")
	}

	if len(dtbMailTemplateAfterDeleteHooks) != 0 {
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
func (o *DTBMailTemplate) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindDTBMailTemplate(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *DTBMailTemplateSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := DTBMailTemplateSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), dtbMailTemplatePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT `dtb_mail_template`.* FROM `dtb_mail_template` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, dtbMailTemplatePrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "model: unable to reload all in DTBMailTemplateSlice")
	}

	*o = slice

	return nil
}

// DTBMailTemplateExists checks if the DTBMailTemplate row exists.
func DTBMailTemplateExists(ctx context.Context, exec boil.ContextExecutor, iD uint) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from `dtb_mail_template` where `id`=? limit 1)"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, iD)
	}

	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "model: unable to check if dtb_mail_template exists")
	}

	return exists, nil
}
