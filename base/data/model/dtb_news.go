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

// DTBNews is an object representing the database table.
type DTBNews struct {
	ID                uint        `boil:"id" json:"id" toml:"id" yaml:"id"`
	CreatorID         null.Uint   `boil:"creator_id" json:"creator_id,omitempty" toml:"creator_id" yaml:"creator_id,omitempty"`
	PublishDate       null.Time   `boil:"publish_date" json:"publish_date,omitempty" toml:"publish_date" yaml:"publish_date,omitempty"`
	Title             string      `boil:"title" json:"title" toml:"title" yaml:"title"`
	Description       null.String `boil:"description" json:"description,omitempty" toml:"description" yaml:"description,omitempty"`
	URL               null.String `boil:"url" json:"url,omitempty" toml:"url" yaml:"url,omitempty"`
	LinkMethod        bool        `boil:"link_method" json:"link_method" toml:"link_method" yaml:"link_method"`
	CreateDate        time.Time   `boil:"create_date" json:"create_date" toml:"create_date" yaml:"create_date"`
	UpdateDate        time.Time   `boil:"update_date" json:"update_date" toml:"update_date" yaml:"update_date"`
	Visible           bool        `boil:"visible" json:"visible" toml:"visible" yaml:"visible"`
	DiscriminatorType string      `boil:"discriminator_type" json:"discriminator_type" toml:"discriminator_type" yaml:"discriminator_type"`

	R *dtbNewsR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L dtbNewsL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var DTBNewsColumns = struct {
	ID                string
	CreatorID         string
	PublishDate       string
	Title             string
	Description       string
	URL               string
	LinkMethod        string
	CreateDate        string
	UpdateDate        string
	Visible           string
	DiscriminatorType string
}{
	ID:                "id",
	CreatorID:         "creator_id",
	PublishDate:       "publish_date",
	Title:             "title",
	Description:       "description",
	URL:               "url",
	LinkMethod:        "link_method",
	CreateDate:        "create_date",
	UpdateDate:        "update_date",
	Visible:           "visible",
	DiscriminatorType: "discriminator_type",
}

// Generated where

var DTBNewsWhere = struct {
	ID                whereHelperuint
	CreatorID         whereHelpernull_Uint
	PublishDate       whereHelpernull_Time
	Title             whereHelperstring
	Description       whereHelpernull_String
	URL               whereHelpernull_String
	LinkMethod        whereHelperbool
	CreateDate        whereHelpertime_Time
	UpdateDate        whereHelpertime_Time
	Visible           whereHelperbool
	DiscriminatorType whereHelperstring
}{
	ID:                whereHelperuint{field: `id`},
	CreatorID:         whereHelpernull_Uint{field: `creator_id`},
	PublishDate:       whereHelpernull_Time{field: `publish_date`},
	Title:             whereHelperstring{field: `title`},
	Description:       whereHelpernull_String{field: `description`},
	URL:               whereHelpernull_String{field: `url`},
	LinkMethod:        whereHelperbool{field: `link_method`},
	CreateDate:        whereHelpertime_Time{field: `create_date`},
	UpdateDate:        whereHelpertime_Time{field: `update_date`},
	Visible:           whereHelperbool{field: `visible`},
	DiscriminatorType: whereHelperstring{field: `discriminator_type`},
}

// DTBNewsRels is where relationship names are stored.
var DTBNewsRels = struct {
	Creator string
}{
	Creator: "Creator",
}

// dtbNewsR is where relationships are stored.
type dtbNewsR struct {
	Creator *DTBMember
}

// NewStruct creates a new relationship struct
func (*dtbNewsR) NewStruct() *dtbNewsR {
	return &dtbNewsR{}
}

// dtbNewsL is where Load methods for each relationship are stored.
type dtbNewsL struct{}

var (
	dtbNewsColumns               = []string{"id", "creator_id", "publish_date", "title", "description", "url", "link_method", "create_date", "update_date", "visible", "discriminator_type"}
	dtbNewsColumnsWithoutDefault = []string{"creator_id", "publish_date", "title", "description", "url", "create_date", "update_date", "discriminator_type"}
	dtbNewsColumnsWithDefault    = []string{"id", "link_method", "visible"}
	dtbNewsPrimaryKeyColumns     = []string{"id"}
)

type (
	// DTBNewsSlice is an alias for a slice of pointers to DTBNews.
	// This should generally be used opposed to []DTBNews.
	DTBNewsSlice []*DTBNews
	// DTBNewsHook is the signature for custom DTBNews hook methods
	DTBNewsHook func(context.Context, boil.ContextExecutor, *DTBNews) error

	dtbNewsQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	dtbNewsType                 = reflect.TypeOf(&DTBNews{})
	dtbNewsMapping              = queries.MakeStructMapping(dtbNewsType)
	dtbNewsPrimaryKeyMapping, _ = queries.BindMapping(dtbNewsType, dtbNewsMapping, dtbNewsPrimaryKeyColumns)
	dtbNewsInsertCacheMut       sync.RWMutex
	dtbNewsInsertCache          = make(map[string]insertCache)
	dtbNewsUpdateCacheMut       sync.RWMutex
	dtbNewsUpdateCache          = make(map[string]updateCache)
	dtbNewsUpsertCacheMut       sync.RWMutex
	dtbNewsUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var dtbNewsBeforeInsertHooks []DTBNewsHook
var dtbNewsBeforeUpdateHooks []DTBNewsHook
var dtbNewsBeforeDeleteHooks []DTBNewsHook
var dtbNewsBeforeUpsertHooks []DTBNewsHook

var dtbNewsAfterInsertHooks []DTBNewsHook
var dtbNewsAfterSelectHooks []DTBNewsHook
var dtbNewsAfterUpdateHooks []DTBNewsHook
var dtbNewsAfterDeleteHooks []DTBNewsHook
var dtbNewsAfterUpsertHooks []DTBNewsHook

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *DTBNews) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range dtbNewsBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *DTBNews) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range dtbNewsBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *DTBNews) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range dtbNewsBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *DTBNews) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range dtbNewsBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *DTBNews) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range dtbNewsAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterSelectHooks executes all "after Select" hooks.
func (o *DTBNews) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range dtbNewsAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *DTBNews) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range dtbNewsAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *DTBNews) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range dtbNewsAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *DTBNews) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range dtbNewsAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddDTBNewsHook registers your hook function for all future operations.
func AddDTBNewsHook(hookPoint boil.HookPoint, dtbNewsHook DTBNewsHook) {
	switch hookPoint {
	case boil.BeforeInsertHook:
		dtbNewsBeforeInsertHooks = append(dtbNewsBeforeInsertHooks, dtbNewsHook)
	case boil.BeforeUpdateHook:
		dtbNewsBeforeUpdateHooks = append(dtbNewsBeforeUpdateHooks, dtbNewsHook)
	case boil.BeforeDeleteHook:
		dtbNewsBeforeDeleteHooks = append(dtbNewsBeforeDeleteHooks, dtbNewsHook)
	case boil.BeforeUpsertHook:
		dtbNewsBeforeUpsertHooks = append(dtbNewsBeforeUpsertHooks, dtbNewsHook)
	case boil.AfterInsertHook:
		dtbNewsAfterInsertHooks = append(dtbNewsAfterInsertHooks, dtbNewsHook)
	case boil.AfterSelectHook:
		dtbNewsAfterSelectHooks = append(dtbNewsAfterSelectHooks, dtbNewsHook)
	case boil.AfterUpdateHook:
		dtbNewsAfterUpdateHooks = append(dtbNewsAfterUpdateHooks, dtbNewsHook)
	case boil.AfterDeleteHook:
		dtbNewsAfterDeleteHooks = append(dtbNewsAfterDeleteHooks, dtbNewsHook)
	case boil.AfterUpsertHook:
		dtbNewsAfterUpsertHooks = append(dtbNewsAfterUpsertHooks, dtbNewsHook)
	}
}

// One returns a single dtbNews record from the query.
func (q dtbNewsQuery) One(ctx context.Context, exec boil.ContextExecutor) (*DTBNews, error) {
	o := &DTBNews{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "model: failed to execute a one query for dtb_news")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all DTBNews records from the query.
func (q dtbNewsQuery) All(ctx context.Context, exec boil.ContextExecutor) (DTBNewsSlice, error) {
	var o []*DTBNews

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "model: failed to assign all query results to DTBNews slice")
	}

	if len(dtbNewsAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all DTBNews records in the query.
func (q dtbNewsQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "model: failed to count dtb_news rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q dtbNewsQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "model: failed to check if dtb_news exists")
	}

	return count > 0, nil
}

// Creator pointed to by the foreign key.
func (o *DTBNews) Creator(mods ...qm.QueryMod) dtbMemberQuery {
	queryMods := []qm.QueryMod{
		qm.Where("id=?", o.CreatorID),
	}

	queryMods = append(queryMods, mods...)

	query := DTBMembers(queryMods...)
	queries.SetFrom(query.Query, "`dtb_member`")

	return query
}

// LoadCreator allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (dtbNewsL) LoadCreator(ctx context.Context, e boil.ContextExecutor, singular bool, maybeDTBNews interface{}, mods queries.Applicator) error {
	var slice []*DTBNews
	var object *DTBNews

	if singular {
		object = maybeDTBNews.(*DTBNews)
	} else {
		slice = *maybeDTBNews.(*[]*DTBNews)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &dtbNewsR{}
		}
		if !queries.IsNil(object.CreatorID) {
			args = append(args, object.CreatorID)
		}

	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &dtbNewsR{}
			}

			for _, a := range args {
				if queries.Equal(a, obj.CreatorID) {
					continue Outer
				}
			}

			if !queries.IsNil(obj.CreatorID) {
				args = append(args, obj.CreatorID)
			}

		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(qm.From(`dtb_member`), qm.WhereIn(`id in ?`, args...))
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load DTBMember")
	}

	var resultSlice []*DTBMember
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice DTBMember")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results of eager load for dtb_member")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for dtb_member")
	}

	if len(dtbNewsAfterSelectHooks) != 0 {
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
		object.R.Creator = foreign
		if foreign.R == nil {
			foreign.R = &dtbMemberR{}
		}
		foreign.R.CreatorDTBNews = append(foreign.R.CreatorDTBNews, object)
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if queries.Equal(local.CreatorID, foreign.ID) {
				local.R.Creator = foreign
				if foreign.R == nil {
					foreign.R = &dtbMemberR{}
				}
				foreign.R.CreatorDTBNews = append(foreign.R.CreatorDTBNews, local)
				break
			}
		}
	}

	return nil
}

// SetCreator of the dtbNews to the related item.
// Sets o.R.Creator to related.
// Adds o to related.R.CreatorDTBNews.
func (o *DTBNews) SetCreator(ctx context.Context, exec boil.ContextExecutor, insert bool, related *DTBMember) error {
	var err error
	if insert {
		if err = related.Insert(ctx, exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE `dtb_news` SET %s WHERE %s",
		strmangle.SetParamNames("`", "`", 0, []string{"creator_id"}),
		strmangle.WhereClause("`", "`", 0, dtbNewsPrimaryKeyColumns),
	)
	values := []interface{}{related.ID, o.ID}

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, updateQuery)
		fmt.Fprintln(boil.DebugWriter, values)
	}

	if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	queries.Assign(&o.CreatorID, related.ID)
	if o.R == nil {
		o.R = &dtbNewsR{
			Creator: related,
		}
	} else {
		o.R.Creator = related
	}

	if related.R == nil {
		related.R = &dtbMemberR{
			CreatorDTBNews: DTBNewsSlice{o},
		}
	} else {
		related.R.CreatorDTBNews = append(related.R.CreatorDTBNews, o)
	}

	return nil
}

// RemoveCreator relationship.
// Sets o.R.Creator to nil.
// Removes o from all passed in related items' relationships struct (Optional).
func (o *DTBNews) RemoveCreator(ctx context.Context, exec boil.ContextExecutor, related *DTBMember) error {
	var err error

	queries.SetScanner(&o.CreatorID, nil)
	if _, err = o.Update(ctx, exec, boil.Whitelist("creator_id")); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	o.R.Creator = nil
	if related == nil || related.R == nil {
		return nil
	}

	for i, ri := range related.R.CreatorDTBNews {
		if queries.Equal(o.CreatorID, ri.CreatorID) {
			continue
		}

		ln := len(related.R.CreatorDTBNews)
		if ln > 1 && i < ln-1 {
			related.R.CreatorDTBNews[i] = related.R.CreatorDTBNews[ln-1]
		}
		related.R.CreatorDTBNews = related.R.CreatorDTBNews[:ln-1]
		break
	}
	return nil
}

// DTBNews retrieves all the records using an executor.
func DTBNews(mods ...qm.QueryMod) dtbNewsQuery {
	mods = append(mods, qm.From("`dtb_news`"))
	return dtbNewsQuery{NewQuery(mods...)}
}

// FindDTBNews retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindDTBNews(ctx context.Context, exec boil.ContextExecutor, iD uint, selectCols ...string) (*DTBNews, error) {
	dtbNewsObj := &DTBNews{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from `dtb_news` where `id`=?", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, dtbNewsObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "model: unable to select from dtb_news")
	}

	return dtbNewsObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *DTBNews) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("model: no dtb_news provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(dtbNewsColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	dtbNewsInsertCacheMut.RLock()
	cache, cached := dtbNewsInsertCache[key]
	dtbNewsInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			dtbNewsColumns,
			dtbNewsColumnsWithDefault,
			dtbNewsColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(dtbNewsType, dtbNewsMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(dtbNewsType, dtbNewsMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO `dtb_news` (`%s`) %%sVALUES (%s)%%s", strings.Join(wl, "`,`"), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO `dtb_news` () VALUES ()%s%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			cache.retQuery = fmt.Sprintf("SELECT `%s` FROM `dtb_news` WHERE %s", strings.Join(returnColumns, "`,`"), strmangle.WhereClause("`", "`", 0, dtbNewsPrimaryKeyColumns))
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
		return errors.Wrap(err, "model: unable to insert into dtb_news")
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
	if lastID != 0 && len(cache.retMapping) == 1 && cache.retMapping[0] == dtbNewsMapping["ID"] {
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
		return errors.Wrap(err, "model: unable to populate default values for dtb_news")
	}

CacheNoHooks:
	if !cached {
		dtbNewsInsertCacheMut.Lock()
		dtbNewsInsertCache[key] = cache
		dtbNewsInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the DTBNews.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *DTBNews) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	dtbNewsUpdateCacheMut.RLock()
	cache, cached := dtbNewsUpdateCache[key]
	dtbNewsUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			dtbNewsColumns,
			dtbNewsPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("model: unable to update dtb_news, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE `dtb_news` SET %s WHERE %s",
			strmangle.SetParamNames("`", "`", 0, wl),
			strmangle.WhereClause("`", "`", 0, dtbNewsPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(dtbNewsType, dtbNewsMapping, append(wl, dtbNewsPrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "model: unable to update dtb_news row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "model: failed to get rows affected by update for dtb_news")
	}

	if !cached {
		dtbNewsUpdateCacheMut.Lock()
		dtbNewsUpdateCache[key] = cache
		dtbNewsUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q dtbNewsQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "model: unable to update all for dtb_news")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "model: unable to retrieve rows affected for dtb_news")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o DTBNewsSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), dtbNewsPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE `dtb_news` SET %s WHERE %s",
		strmangle.SetParamNames("`", "`", 0, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, dtbNewsPrimaryKeyColumns, len(o)))

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "model: unable to update all in dtbNews slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "model: unable to retrieve rows affected all in update all dtbNews")
	}
	return rowsAff, nil
}

var mySQLDTBNewsUniqueColumns = []string{
	"id",
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *DTBNews) Upsert(ctx context.Context, exec boil.ContextExecutor, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("model: no dtb_news provided for upsert")
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(dtbNewsColumnsWithDefault, o)
	nzUniques := queries.NonZeroDefaultSet(mySQLDTBNewsUniqueColumns, o)

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

	dtbNewsUpsertCacheMut.RLock()
	cache, cached := dtbNewsUpsertCache[key]
	dtbNewsUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			dtbNewsColumns,
			dtbNewsColumnsWithDefault,
			dtbNewsColumnsWithoutDefault,
			nzDefaults,
		)
		update := updateColumns.UpdateColumnSet(
			dtbNewsColumns,
			dtbNewsPrimaryKeyColumns,
		)

		if len(update) == 0 {
			return errors.New("model: unable to upsert dtb_news, could not build update column list")
		}

		ret = strmangle.SetComplement(ret, nzUniques)
		cache.query = buildUpsertQueryMySQL(dialect, "dtb_news", update, insert)
		cache.retQuery = fmt.Sprintf(
			"SELECT %s FROM `dtb_news` WHERE %s",
			strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, ret), ","),
			strmangle.WhereClause("`", "`", 0, nzUniques),
		)

		cache.valueMapping, err = queries.BindMapping(dtbNewsType, dtbNewsMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(dtbNewsType, dtbNewsMapping, ret)
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
		return errors.Wrap(err, "model: unable to upsert for dtb_news")
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
	if lastID != 0 && len(cache.retMapping) == 1 && cache.retMapping[0] == dtbNewsMapping["id"] {
		goto CacheNoHooks
	}

	uniqueMap, err = queries.BindMapping(dtbNewsType, dtbNewsMapping, nzUniques)
	if err != nil {
		return errors.Wrap(err, "model: unable to retrieve unique values for dtb_news")
	}
	nzUniqueCols = queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), uniqueMap)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.retQuery)
		fmt.Fprintln(boil.DebugWriter, nzUniqueCols...)
	}

	err = exec.QueryRowContext(ctx, cache.retQuery, nzUniqueCols...).Scan(returns...)
	if err != nil {
		return errors.Wrap(err, "model: unable to populate default values for dtb_news")
	}

CacheNoHooks:
	if !cached {
		dtbNewsUpsertCacheMut.Lock()
		dtbNewsUpsertCache[key] = cache
		dtbNewsUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single DTBNews record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *DTBNews) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("model: no DTBNews provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), dtbNewsPrimaryKeyMapping)
	sql := "DELETE FROM `dtb_news` WHERE `id`=?"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "model: unable to delete from dtb_news")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "model: failed to get rows affected by delete for dtb_news")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q dtbNewsQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("model: no dtbNewsQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "model: unable to delete all from dtb_news")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "model: failed to get rows affected by deleteall for dtb_news")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o DTBNewsSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("model: no DTBNews slice provided for delete all")
	}

	if len(o) == 0 {
		return 0, nil
	}

	if len(dtbNewsBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), dtbNewsPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM `dtb_news` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, dtbNewsPrimaryKeyColumns, len(o))

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args)
	}

	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "model: unable to delete all from dtbNews slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "model: failed to get rows affected by deleteall for dtb_news")
	}

	if len(dtbNewsAfterDeleteHooks) != 0 {
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
func (o *DTBNews) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindDTBNews(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *DTBNewsSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := DTBNewsSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), dtbNewsPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT `dtb_news`.* FROM `dtb_news` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, dtbNewsPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "model: unable to reload all in DTBNewsSlice")
	}

	*o = slice

	return nil
}

// DTBNewsExists checks if the DTBNews row exists.
func DTBNewsExists(ctx context.Context, exec boil.ContextExecutor, iD uint) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from `dtb_news` where `id`=? limit 1)"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, iD)
	}

	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "model: unable to check if dtb_news exists")
	}

	return exists, nil
}
