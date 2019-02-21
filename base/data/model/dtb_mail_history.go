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

// DTBMailHistory is an object representing the database table.
type DTBMailHistory struct {
	ID                uint        `boil:"id" json:"id" toml:"id" yaml:"id"`
	OrderID           null.Uint   `boil:"order_id" json:"order_id,omitempty" toml:"order_id" yaml:"order_id,omitempty"`
	CreatorID         null.Uint   `boil:"creator_id" json:"creator_id,omitempty" toml:"creator_id" yaml:"creator_id,omitempty"`
	SendDate          null.Time   `boil:"send_date" json:"send_date,omitempty" toml:"send_date" yaml:"send_date,omitempty"`
	MailSubject       null.String `boil:"mail_subject" json:"mail_subject,omitempty" toml:"mail_subject" yaml:"mail_subject,omitempty"`
	MailBody          null.String `boil:"mail_body" json:"mail_body,omitempty" toml:"mail_body" yaml:"mail_body,omitempty"`
	MailHTMLBody      null.String `boil:"mail_html_body" json:"mail_html_body,omitempty" toml:"mail_html_body" yaml:"mail_html_body,omitempty"`
	DiscriminatorType string      `boil:"discriminator_type" json:"discriminator_type" toml:"discriminator_type" yaml:"discriminator_type"`

	R *dtbMailHistoryR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L dtbMailHistoryL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var DTBMailHistoryColumns = struct {
	ID                string
	OrderID           string
	CreatorID         string
	SendDate          string
	MailSubject       string
	MailBody          string
	MailHTMLBody      string
	DiscriminatorType string
}{
	ID:                "id",
	OrderID:           "order_id",
	CreatorID:         "creator_id",
	SendDate:          "send_date",
	MailSubject:       "mail_subject",
	MailBody:          "mail_body",
	MailHTMLBody:      "mail_html_body",
	DiscriminatorType: "discriminator_type",
}

// Generated where

var DTBMailHistoryWhere = struct {
	ID                whereHelperuint
	OrderID           whereHelpernull_Uint
	CreatorID         whereHelpernull_Uint
	SendDate          whereHelpernull_Time
	MailSubject       whereHelpernull_String
	MailBody          whereHelpernull_String
	MailHTMLBody      whereHelpernull_String
	DiscriminatorType whereHelperstring
}{
	ID:                whereHelperuint{field: `id`},
	OrderID:           whereHelpernull_Uint{field: `order_id`},
	CreatorID:         whereHelpernull_Uint{field: `creator_id`},
	SendDate:          whereHelpernull_Time{field: `send_date`},
	MailSubject:       whereHelpernull_String{field: `mail_subject`},
	MailBody:          whereHelpernull_String{field: `mail_body`},
	MailHTMLBody:      whereHelpernull_String{field: `mail_html_body`},
	DiscriminatorType: whereHelperstring{field: `discriminator_type`},
}

// DTBMailHistoryRels is where relationship names are stored.
var DTBMailHistoryRels = struct {
	Order string
}{
	Order: "Order",
}

// dtbMailHistoryR is where relationships are stored.
type dtbMailHistoryR struct {
	Order *DTBOrder
}

// NewStruct creates a new relationship struct
func (*dtbMailHistoryR) NewStruct() *dtbMailHistoryR {
	return &dtbMailHistoryR{}
}

// dtbMailHistoryL is where Load methods for each relationship are stored.
type dtbMailHistoryL struct{}

var (
	dtbMailHistoryColumns               = []string{"id", "order_id", "creator_id", "send_date", "mail_subject", "mail_body", "mail_html_body", "discriminator_type"}
	dtbMailHistoryColumnsWithoutDefault = []string{"order_id", "creator_id", "send_date", "mail_subject", "mail_body", "mail_html_body", "discriminator_type"}
	dtbMailHistoryColumnsWithDefault    = []string{"id"}
	dtbMailHistoryPrimaryKeyColumns     = []string{"id"}
)

type (
	// DTBMailHistorySlice is an alias for a slice of pointers to DTBMailHistory.
	// This should generally be used opposed to []DTBMailHistory.
	DTBMailHistorySlice []*DTBMailHistory
	// DTBMailHistoryHook is the signature for custom DTBMailHistory hook methods
	DTBMailHistoryHook func(context.Context, boil.ContextExecutor, *DTBMailHistory) error

	dtbMailHistoryQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	dtbMailHistoryType                 = reflect.TypeOf(&DTBMailHistory{})
	dtbMailHistoryMapping              = queries.MakeStructMapping(dtbMailHistoryType)
	dtbMailHistoryPrimaryKeyMapping, _ = queries.BindMapping(dtbMailHistoryType, dtbMailHistoryMapping, dtbMailHistoryPrimaryKeyColumns)
	dtbMailHistoryInsertCacheMut       sync.RWMutex
	dtbMailHistoryInsertCache          = make(map[string]insertCache)
	dtbMailHistoryUpdateCacheMut       sync.RWMutex
	dtbMailHistoryUpdateCache          = make(map[string]updateCache)
	dtbMailHistoryUpsertCacheMut       sync.RWMutex
	dtbMailHistoryUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var dtbMailHistoryBeforeInsertHooks []DTBMailHistoryHook
var dtbMailHistoryBeforeUpdateHooks []DTBMailHistoryHook
var dtbMailHistoryBeforeDeleteHooks []DTBMailHistoryHook
var dtbMailHistoryBeforeUpsertHooks []DTBMailHistoryHook

var dtbMailHistoryAfterInsertHooks []DTBMailHistoryHook
var dtbMailHistoryAfterSelectHooks []DTBMailHistoryHook
var dtbMailHistoryAfterUpdateHooks []DTBMailHistoryHook
var dtbMailHistoryAfterDeleteHooks []DTBMailHistoryHook
var dtbMailHistoryAfterUpsertHooks []DTBMailHistoryHook

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *DTBMailHistory) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range dtbMailHistoryBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *DTBMailHistory) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range dtbMailHistoryBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *DTBMailHistory) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range dtbMailHistoryBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *DTBMailHistory) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range dtbMailHistoryBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *DTBMailHistory) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range dtbMailHistoryAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterSelectHooks executes all "after Select" hooks.
func (o *DTBMailHistory) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range dtbMailHistoryAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *DTBMailHistory) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range dtbMailHistoryAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *DTBMailHistory) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range dtbMailHistoryAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *DTBMailHistory) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range dtbMailHistoryAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddDTBMailHistoryHook registers your hook function for all future operations.
func AddDTBMailHistoryHook(hookPoint boil.HookPoint, dtbMailHistoryHook DTBMailHistoryHook) {
	switch hookPoint {
	case boil.BeforeInsertHook:
		dtbMailHistoryBeforeInsertHooks = append(dtbMailHistoryBeforeInsertHooks, dtbMailHistoryHook)
	case boil.BeforeUpdateHook:
		dtbMailHistoryBeforeUpdateHooks = append(dtbMailHistoryBeforeUpdateHooks, dtbMailHistoryHook)
	case boil.BeforeDeleteHook:
		dtbMailHistoryBeforeDeleteHooks = append(dtbMailHistoryBeforeDeleteHooks, dtbMailHistoryHook)
	case boil.BeforeUpsertHook:
		dtbMailHistoryBeforeUpsertHooks = append(dtbMailHistoryBeforeUpsertHooks, dtbMailHistoryHook)
	case boil.AfterInsertHook:
		dtbMailHistoryAfterInsertHooks = append(dtbMailHistoryAfterInsertHooks, dtbMailHistoryHook)
	case boil.AfterSelectHook:
		dtbMailHistoryAfterSelectHooks = append(dtbMailHistoryAfterSelectHooks, dtbMailHistoryHook)
	case boil.AfterUpdateHook:
		dtbMailHistoryAfterUpdateHooks = append(dtbMailHistoryAfterUpdateHooks, dtbMailHistoryHook)
	case boil.AfterDeleteHook:
		dtbMailHistoryAfterDeleteHooks = append(dtbMailHistoryAfterDeleteHooks, dtbMailHistoryHook)
	case boil.AfterUpsertHook:
		dtbMailHistoryAfterUpsertHooks = append(dtbMailHistoryAfterUpsertHooks, dtbMailHistoryHook)
	}
}

// One returns a single dtbMailHistory record from the query.
func (q dtbMailHistoryQuery) One(ctx context.Context, exec boil.ContextExecutor) (*DTBMailHistory, error) {
	o := &DTBMailHistory{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "model: failed to execute a one query for dtb_mail_history")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all DTBMailHistory records from the query.
func (q dtbMailHistoryQuery) All(ctx context.Context, exec boil.ContextExecutor) (DTBMailHistorySlice, error) {
	var o []*DTBMailHistory

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "model: failed to assign all query results to DTBMailHistory slice")
	}

	if len(dtbMailHistoryAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all DTBMailHistory records in the query.
func (q dtbMailHistoryQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "model: failed to count dtb_mail_history rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q dtbMailHistoryQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "model: failed to check if dtb_mail_history exists")
	}

	return count > 0, nil
}

// Order pointed to by the foreign key.
func (o *DTBMailHistory) Order(mods ...qm.QueryMod) dtbOrderQuery {
	queryMods := []qm.QueryMod{
		qm.Where("id=?", o.OrderID),
	}

	queryMods = append(queryMods, mods...)

	query := DTBOrders(queryMods...)
	queries.SetFrom(query.Query, "`dtb_order`")

	return query
}

// LoadOrder allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (dtbMailHistoryL) LoadOrder(ctx context.Context, e boil.ContextExecutor, singular bool, maybeDTBMailHistory interface{}, mods queries.Applicator) error {
	var slice []*DTBMailHistory
	var object *DTBMailHistory

	if singular {
		object = maybeDTBMailHistory.(*DTBMailHistory)
	} else {
		slice = *maybeDTBMailHistory.(*[]*DTBMailHistory)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &dtbMailHistoryR{}
		}
		if !queries.IsNil(object.OrderID) {
			args = append(args, object.OrderID)
		}

	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &dtbMailHistoryR{}
			}

			for _, a := range args {
				if queries.Equal(a, obj.OrderID) {
					continue Outer
				}
			}

			if !queries.IsNil(obj.OrderID) {
				args = append(args, obj.OrderID)
			}

		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(qm.From(`dtb_order`), qm.WhereIn(`id in ?`, args...))
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load DTBOrder")
	}

	var resultSlice []*DTBOrder
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice DTBOrder")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results of eager load for dtb_order")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for dtb_order")
	}

	if len(dtbMailHistoryAfterSelectHooks) != 0 {
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
		object.R.Order = foreign
		if foreign.R == nil {
			foreign.R = &dtbOrderR{}
		}
		foreign.R.OrderDTBMailHistories = append(foreign.R.OrderDTBMailHistories, object)
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if queries.Equal(local.OrderID, foreign.ID) {
				local.R.Order = foreign
				if foreign.R == nil {
					foreign.R = &dtbOrderR{}
				}
				foreign.R.OrderDTBMailHistories = append(foreign.R.OrderDTBMailHistories, local)
				break
			}
		}
	}

	return nil
}

// SetOrder of the dtbMailHistory to the related item.
// Sets o.R.Order to related.
// Adds o to related.R.OrderDTBMailHistories.
func (o *DTBMailHistory) SetOrder(ctx context.Context, exec boil.ContextExecutor, insert bool, related *DTBOrder) error {
	var err error
	if insert {
		if err = related.Insert(ctx, exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE `dtb_mail_history` SET %s WHERE %s",
		strmangle.SetParamNames("`", "`", 0, []string{"order_id"}),
		strmangle.WhereClause("`", "`", 0, dtbMailHistoryPrimaryKeyColumns),
	)
	values := []interface{}{related.ID, o.ID}

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, updateQuery)
		fmt.Fprintln(boil.DebugWriter, values)
	}

	if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	queries.Assign(&o.OrderID, related.ID)
	if o.R == nil {
		o.R = &dtbMailHistoryR{
			Order: related,
		}
	} else {
		o.R.Order = related
	}

	if related.R == nil {
		related.R = &dtbOrderR{
			OrderDTBMailHistories: DTBMailHistorySlice{o},
		}
	} else {
		related.R.OrderDTBMailHistories = append(related.R.OrderDTBMailHistories, o)
	}

	return nil
}

// RemoveOrder relationship.
// Sets o.R.Order to nil.
// Removes o from all passed in related items' relationships struct (Optional).
func (o *DTBMailHistory) RemoveOrder(ctx context.Context, exec boil.ContextExecutor, related *DTBOrder) error {
	var err error

	queries.SetScanner(&o.OrderID, nil)
	if _, err = o.Update(ctx, exec, boil.Whitelist("order_id")); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	o.R.Order = nil
	if related == nil || related.R == nil {
		return nil
	}

	for i, ri := range related.R.OrderDTBMailHistories {
		if queries.Equal(o.OrderID, ri.OrderID) {
			continue
		}

		ln := len(related.R.OrderDTBMailHistories)
		if ln > 1 && i < ln-1 {
			related.R.OrderDTBMailHistories[i] = related.R.OrderDTBMailHistories[ln-1]
		}
		related.R.OrderDTBMailHistories = related.R.OrderDTBMailHistories[:ln-1]
		break
	}
	return nil
}

// DTBMailHistories retrieves all the records using an executor.
func DTBMailHistories(mods ...qm.QueryMod) dtbMailHistoryQuery {
	mods = append(mods, qm.From("`dtb_mail_history`"))
	return dtbMailHistoryQuery{NewQuery(mods...)}
}

// FindDTBMailHistory retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindDTBMailHistory(ctx context.Context, exec boil.ContextExecutor, iD uint, selectCols ...string) (*DTBMailHistory, error) {
	dtbMailHistoryObj := &DTBMailHistory{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from `dtb_mail_history` where `id`=?", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, dtbMailHistoryObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "model: unable to select from dtb_mail_history")
	}

	return dtbMailHistoryObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *DTBMailHistory) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("model: no dtb_mail_history provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(dtbMailHistoryColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	dtbMailHistoryInsertCacheMut.RLock()
	cache, cached := dtbMailHistoryInsertCache[key]
	dtbMailHistoryInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			dtbMailHistoryColumns,
			dtbMailHistoryColumnsWithDefault,
			dtbMailHistoryColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(dtbMailHistoryType, dtbMailHistoryMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(dtbMailHistoryType, dtbMailHistoryMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO `dtb_mail_history` (`%s`) %%sVALUES (%s)%%s", strings.Join(wl, "`,`"), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO `dtb_mail_history` () VALUES ()%s%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			cache.retQuery = fmt.Sprintf("SELECT `%s` FROM `dtb_mail_history` WHERE %s", strings.Join(returnColumns, "`,`"), strmangle.WhereClause("`", "`", 0, dtbMailHistoryPrimaryKeyColumns))
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
		return errors.Wrap(err, "model: unable to insert into dtb_mail_history")
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
	if lastID != 0 && len(cache.retMapping) == 1 && cache.retMapping[0] == dtbMailHistoryMapping["ID"] {
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
		return errors.Wrap(err, "model: unable to populate default values for dtb_mail_history")
	}

CacheNoHooks:
	if !cached {
		dtbMailHistoryInsertCacheMut.Lock()
		dtbMailHistoryInsertCache[key] = cache
		dtbMailHistoryInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the DTBMailHistory.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *DTBMailHistory) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	dtbMailHistoryUpdateCacheMut.RLock()
	cache, cached := dtbMailHistoryUpdateCache[key]
	dtbMailHistoryUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			dtbMailHistoryColumns,
			dtbMailHistoryPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("model: unable to update dtb_mail_history, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE `dtb_mail_history` SET %s WHERE %s",
			strmangle.SetParamNames("`", "`", 0, wl),
			strmangle.WhereClause("`", "`", 0, dtbMailHistoryPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(dtbMailHistoryType, dtbMailHistoryMapping, append(wl, dtbMailHistoryPrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "model: unable to update dtb_mail_history row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "model: failed to get rows affected by update for dtb_mail_history")
	}

	if !cached {
		dtbMailHistoryUpdateCacheMut.Lock()
		dtbMailHistoryUpdateCache[key] = cache
		dtbMailHistoryUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q dtbMailHistoryQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "model: unable to update all for dtb_mail_history")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "model: unable to retrieve rows affected for dtb_mail_history")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o DTBMailHistorySlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), dtbMailHistoryPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE `dtb_mail_history` SET %s WHERE %s",
		strmangle.SetParamNames("`", "`", 0, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, dtbMailHistoryPrimaryKeyColumns, len(o)))

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "model: unable to update all in dtbMailHistory slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "model: unable to retrieve rows affected all in update all dtbMailHistory")
	}
	return rowsAff, nil
}

var mySQLDTBMailHistoryUniqueColumns = []string{
	"id",
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *DTBMailHistory) Upsert(ctx context.Context, exec boil.ContextExecutor, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("model: no dtb_mail_history provided for upsert")
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(dtbMailHistoryColumnsWithDefault, o)
	nzUniques := queries.NonZeroDefaultSet(mySQLDTBMailHistoryUniqueColumns, o)

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

	dtbMailHistoryUpsertCacheMut.RLock()
	cache, cached := dtbMailHistoryUpsertCache[key]
	dtbMailHistoryUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			dtbMailHistoryColumns,
			dtbMailHistoryColumnsWithDefault,
			dtbMailHistoryColumnsWithoutDefault,
			nzDefaults,
		)
		update := updateColumns.UpdateColumnSet(
			dtbMailHistoryColumns,
			dtbMailHistoryPrimaryKeyColumns,
		)

		if len(update) == 0 {
			return errors.New("model: unable to upsert dtb_mail_history, could not build update column list")
		}

		ret = strmangle.SetComplement(ret, nzUniques)
		cache.query = buildUpsertQueryMySQL(dialect, "dtb_mail_history", update, insert)
		cache.retQuery = fmt.Sprintf(
			"SELECT %s FROM `dtb_mail_history` WHERE %s",
			strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, ret), ","),
			strmangle.WhereClause("`", "`", 0, nzUniques),
		)

		cache.valueMapping, err = queries.BindMapping(dtbMailHistoryType, dtbMailHistoryMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(dtbMailHistoryType, dtbMailHistoryMapping, ret)
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
		return errors.Wrap(err, "model: unable to upsert for dtb_mail_history")
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
	if lastID != 0 && len(cache.retMapping) == 1 && cache.retMapping[0] == dtbMailHistoryMapping["id"] {
		goto CacheNoHooks
	}

	uniqueMap, err = queries.BindMapping(dtbMailHistoryType, dtbMailHistoryMapping, nzUniques)
	if err != nil {
		return errors.Wrap(err, "model: unable to retrieve unique values for dtb_mail_history")
	}
	nzUniqueCols = queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), uniqueMap)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.retQuery)
		fmt.Fprintln(boil.DebugWriter, nzUniqueCols...)
	}

	err = exec.QueryRowContext(ctx, cache.retQuery, nzUniqueCols...).Scan(returns...)
	if err != nil {
		return errors.Wrap(err, "model: unable to populate default values for dtb_mail_history")
	}

CacheNoHooks:
	if !cached {
		dtbMailHistoryUpsertCacheMut.Lock()
		dtbMailHistoryUpsertCache[key] = cache
		dtbMailHistoryUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single DTBMailHistory record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *DTBMailHistory) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("model: no DTBMailHistory provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), dtbMailHistoryPrimaryKeyMapping)
	sql := "DELETE FROM `dtb_mail_history` WHERE `id`=?"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "model: unable to delete from dtb_mail_history")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "model: failed to get rows affected by delete for dtb_mail_history")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q dtbMailHistoryQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("model: no dtbMailHistoryQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "model: unable to delete all from dtb_mail_history")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "model: failed to get rows affected by deleteall for dtb_mail_history")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o DTBMailHistorySlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("model: no DTBMailHistory slice provided for delete all")
	}

	if len(o) == 0 {
		return 0, nil
	}

	if len(dtbMailHistoryBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), dtbMailHistoryPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM `dtb_mail_history` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, dtbMailHistoryPrimaryKeyColumns, len(o))

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args)
	}

	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "model: unable to delete all from dtbMailHistory slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "model: failed to get rows affected by deleteall for dtb_mail_history")
	}

	if len(dtbMailHistoryAfterDeleteHooks) != 0 {
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
func (o *DTBMailHistory) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindDTBMailHistory(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *DTBMailHistorySlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := DTBMailHistorySlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), dtbMailHistoryPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT `dtb_mail_history`.* FROM `dtb_mail_history` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, dtbMailHistoryPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "model: unable to reload all in DTBMailHistorySlice")
	}

	*o = slice

	return nil
}

// DTBMailHistoryExists checks if the DTBMailHistory row exists.
func DTBMailHistoryExists(ctx context.Context, exec boil.ContextExecutor, iD uint) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from `dtb_mail_history` where `id`=? limit 1)"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, iD)
	}

	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "model: unable to check if dtb_mail_history exists")
	}

	return exists, nil
}