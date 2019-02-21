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

// DTBClassCategory is an object representing the database table.
type DTBClassCategory struct {
	ID                uint        `boil:"id" json:"id" toml:"id" yaml:"id"`
	ClassNameID       null.Uint   `boil:"class_name_id" json:"class_name_id,omitempty" toml:"class_name_id" yaml:"class_name_id,omitempty"`
	CreatorID         null.Uint   `boil:"creator_id" json:"creator_id,omitempty" toml:"creator_id" yaml:"creator_id,omitempty"`
	BackendName       null.String `boil:"backend_name" json:"backend_name,omitempty" toml:"backend_name" yaml:"backend_name,omitempty"`
	Name              string      `boil:"name" json:"name" toml:"name" yaml:"name"`
	SortNo            uint        `boil:"sort_no" json:"sort_no" toml:"sort_no" yaml:"sort_no"`
	Visible           bool        `boil:"visible" json:"visible" toml:"visible" yaml:"visible"`
	CreateDate        time.Time   `boil:"create_date" json:"create_date" toml:"create_date" yaml:"create_date"`
	UpdateDate        time.Time   `boil:"update_date" json:"update_date" toml:"update_date" yaml:"update_date"`
	DiscriminatorType string      `boil:"discriminator_type" json:"discriminator_type" toml:"discriminator_type" yaml:"discriminator_type"`

	R *dtbClassCategoryR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L dtbClassCategoryL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var DTBClassCategoryColumns = struct {
	ID                string
	ClassNameID       string
	CreatorID         string
	BackendName       string
	Name              string
	SortNo            string
	Visible           string
	CreateDate        string
	UpdateDate        string
	DiscriminatorType string
}{
	ID:                "id",
	ClassNameID:       "class_name_id",
	CreatorID:         "creator_id",
	BackendName:       "backend_name",
	Name:              "name",
	SortNo:            "sort_no",
	Visible:           "visible",
	CreateDate:        "create_date",
	UpdateDate:        "update_date",
	DiscriminatorType: "discriminator_type",
}

// Generated where

var DTBClassCategoryWhere = struct {
	ID                whereHelperuint
	ClassNameID       whereHelpernull_Uint
	CreatorID         whereHelpernull_Uint
	BackendName       whereHelpernull_String
	Name              whereHelperstring
	SortNo            whereHelperuint
	Visible           whereHelperbool
	CreateDate        whereHelpertime_Time
	UpdateDate        whereHelpertime_Time
	DiscriminatorType whereHelperstring
}{
	ID:                whereHelperuint{field: `id`},
	ClassNameID:       whereHelpernull_Uint{field: `class_name_id`},
	CreatorID:         whereHelpernull_Uint{field: `creator_id`},
	BackendName:       whereHelpernull_String{field: `backend_name`},
	Name:              whereHelperstring{field: `name`},
	SortNo:            whereHelperuint{field: `sort_no`},
	Visible:           whereHelperbool{field: `visible`},
	CreateDate:        whereHelpertime_Time{field: `create_date`},
	UpdateDate:        whereHelpertime_Time{field: `update_date`},
	DiscriminatorType: whereHelperstring{field: `discriminator_type`},
}

// DTBClassCategoryRels is where relationship names are stored.
var DTBClassCategoryRels = struct {
	ClassName string
}{
	ClassName: "ClassName",
}

// dtbClassCategoryR is where relationships are stored.
type dtbClassCategoryR struct {
	ClassName *DTBClassName
}

// NewStruct creates a new relationship struct
func (*dtbClassCategoryR) NewStruct() *dtbClassCategoryR {
	return &dtbClassCategoryR{}
}

// dtbClassCategoryL is where Load methods for each relationship are stored.
type dtbClassCategoryL struct{}

var (
	dtbClassCategoryColumns               = []string{"id", "class_name_id", "creator_id", "backend_name", "name", "sort_no", "visible", "create_date", "update_date", "discriminator_type"}
	dtbClassCategoryColumnsWithoutDefault = []string{"class_name_id", "creator_id", "backend_name", "name", "sort_no", "create_date", "update_date", "discriminator_type"}
	dtbClassCategoryColumnsWithDefault    = []string{"id", "visible"}
	dtbClassCategoryPrimaryKeyColumns     = []string{"id"}
)

type (
	// DTBClassCategorySlice is an alias for a slice of pointers to DTBClassCategory.
	// This should generally be used opposed to []DTBClassCategory.
	DTBClassCategorySlice []*DTBClassCategory
	// DTBClassCategoryHook is the signature for custom DTBClassCategory hook methods
	DTBClassCategoryHook func(context.Context, boil.ContextExecutor, *DTBClassCategory) error

	dtbClassCategoryQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	dtbClassCategoryType                 = reflect.TypeOf(&DTBClassCategory{})
	dtbClassCategoryMapping              = queries.MakeStructMapping(dtbClassCategoryType)
	dtbClassCategoryPrimaryKeyMapping, _ = queries.BindMapping(dtbClassCategoryType, dtbClassCategoryMapping, dtbClassCategoryPrimaryKeyColumns)
	dtbClassCategoryInsertCacheMut       sync.RWMutex
	dtbClassCategoryInsertCache          = make(map[string]insertCache)
	dtbClassCategoryUpdateCacheMut       sync.RWMutex
	dtbClassCategoryUpdateCache          = make(map[string]updateCache)
	dtbClassCategoryUpsertCacheMut       sync.RWMutex
	dtbClassCategoryUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var dtbClassCategoryBeforeInsertHooks []DTBClassCategoryHook
var dtbClassCategoryBeforeUpdateHooks []DTBClassCategoryHook
var dtbClassCategoryBeforeDeleteHooks []DTBClassCategoryHook
var dtbClassCategoryBeforeUpsertHooks []DTBClassCategoryHook

var dtbClassCategoryAfterInsertHooks []DTBClassCategoryHook
var dtbClassCategoryAfterSelectHooks []DTBClassCategoryHook
var dtbClassCategoryAfterUpdateHooks []DTBClassCategoryHook
var dtbClassCategoryAfterDeleteHooks []DTBClassCategoryHook
var dtbClassCategoryAfterUpsertHooks []DTBClassCategoryHook

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *DTBClassCategory) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range dtbClassCategoryBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *DTBClassCategory) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range dtbClassCategoryBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *DTBClassCategory) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range dtbClassCategoryBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *DTBClassCategory) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range dtbClassCategoryBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *DTBClassCategory) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range dtbClassCategoryAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterSelectHooks executes all "after Select" hooks.
func (o *DTBClassCategory) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range dtbClassCategoryAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *DTBClassCategory) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range dtbClassCategoryAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *DTBClassCategory) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range dtbClassCategoryAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *DTBClassCategory) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range dtbClassCategoryAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddDTBClassCategoryHook registers your hook function for all future operations.
func AddDTBClassCategoryHook(hookPoint boil.HookPoint, dtbClassCategoryHook DTBClassCategoryHook) {
	switch hookPoint {
	case boil.BeforeInsertHook:
		dtbClassCategoryBeforeInsertHooks = append(dtbClassCategoryBeforeInsertHooks, dtbClassCategoryHook)
	case boil.BeforeUpdateHook:
		dtbClassCategoryBeforeUpdateHooks = append(dtbClassCategoryBeforeUpdateHooks, dtbClassCategoryHook)
	case boil.BeforeDeleteHook:
		dtbClassCategoryBeforeDeleteHooks = append(dtbClassCategoryBeforeDeleteHooks, dtbClassCategoryHook)
	case boil.BeforeUpsertHook:
		dtbClassCategoryBeforeUpsertHooks = append(dtbClassCategoryBeforeUpsertHooks, dtbClassCategoryHook)
	case boil.AfterInsertHook:
		dtbClassCategoryAfterInsertHooks = append(dtbClassCategoryAfterInsertHooks, dtbClassCategoryHook)
	case boil.AfterSelectHook:
		dtbClassCategoryAfterSelectHooks = append(dtbClassCategoryAfterSelectHooks, dtbClassCategoryHook)
	case boil.AfterUpdateHook:
		dtbClassCategoryAfterUpdateHooks = append(dtbClassCategoryAfterUpdateHooks, dtbClassCategoryHook)
	case boil.AfterDeleteHook:
		dtbClassCategoryAfterDeleteHooks = append(dtbClassCategoryAfterDeleteHooks, dtbClassCategoryHook)
	case boil.AfterUpsertHook:
		dtbClassCategoryAfterUpsertHooks = append(dtbClassCategoryAfterUpsertHooks, dtbClassCategoryHook)
	}
}

// One returns a single dtbClassCategory record from the query.
func (q dtbClassCategoryQuery) One(ctx context.Context, exec boil.ContextExecutor) (*DTBClassCategory, error) {
	o := &DTBClassCategory{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "model: failed to execute a one query for dtb_class_category")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all DTBClassCategory records from the query.
func (q dtbClassCategoryQuery) All(ctx context.Context, exec boil.ContextExecutor) (DTBClassCategorySlice, error) {
	var o []*DTBClassCategory

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "model: failed to assign all query results to DTBClassCategory slice")
	}

	if len(dtbClassCategoryAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all DTBClassCategory records in the query.
func (q dtbClassCategoryQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "model: failed to count dtb_class_category rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q dtbClassCategoryQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "model: failed to check if dtb_class_category exists")
	}

	return count > 0, nil
}

// ClassName pointed to by the foreign key.
func (o *DTBClassCategory) ClassName(mods ...qm.QueryMod) dtbClassNameQuery {
	queryMods := []qm.QueryMod{
		qm.Where("id=?", o.ClassNameID),
	}

	queryMods = append(queryMods, mods...)

	query := DTBClassNames(queryMods...)
	queries.SetFrom(query.Query, "`dtb_class_name`")

	return query
}

// LoadClassName allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (dtbClassCategoryL) LoadClassName(ctx context.Context, e boil.ContextExecutor, singular bool, maybeDTBClassCategory interface{}, mods queries.Applicator) error {
	var slice []*DTBClassCategory
	var object *DTBClassCategory

	if singular {
		object = maybeDTBClassCategory.(*DTBClassCategory)
	} else {
		slice = *maybeDTBClassCategory.(*[]*DTBClassCategory)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &dtbClassCategoryR{}
		}
		if !queries.IsNil(object.ClassNameID) {
			args = append(args, object.ClassNameID)
		}

	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &dtbClassCategoryR{}
			}

			for _, a := range args {
				if queries.Equal(a, obj.ClassNameID) {
					continue Outer
				}
			}

			if !queries.IsNil(obj.ClassNameID) {
				args = append(args, obj.ClassNameID)
			}

		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(qm.From(`dtb_class_name`), qm.WhereIn(`id in ?`, args...))
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load DTBClassName")
	}

	var resultSlice []*DTBClassName
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice DTBClassName")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results of eager load for dtb_class_name")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for dtb_class_name")
	}

	if len(dtbClassCategoryAfterSelectHooks) != 0 {
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
		object.R.ClassName = foreign
		if foreign.R == nil {
			foreign.R = &dtbClassNameR{}
		}
		foreign.R.ClassNameDTBClassCategories = append(foreign.R.ClassNameDTBClassCategories, object)
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if queries.Equal(local.ClassNameID, foreign.ID) {
				local.R.ClassName = foreign
				if foreign.R == nil {
					foreign.R = &dtbClassNameR{}
				}
				foreign.R.ClassNameDTBClassCategories = append(foreign.R.ClassNameDTBClassCategories, local)
				break
			}
		}
	}

	return nil
}

// SetClassName of the dtbClassCategory to the related item.
// Sets o.R.ClassName to related.
// Adds o to related.R.ClassNameDTBClassCategories.
func (o *DTBClassCategory) SetClassName(ctx context.Context, exec boil.ContextExecutor, insert bool, related *DTBClassName) error {
	var err error
	if insert {
		if err = related.Insert(ctx, exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE `dtb_class_category` SET %s WHERE %s",
		strmangle.SetParamNames("`", "`", 0, []string{"class_name_id"}),
		strmangle.WhereClause("`", "`", 0, dtbClassCategoryPrimaryKeyColumns),
	)
	values := []interface{}{related.ID, o.ID}

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, updateQuery)
		fmt.Fprintln(boil.DebugWriter, values)
	}

	if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	queries.Assign(&o.ClassNameID, related.ID)
	if o.R == nil {
		o.R = &dtbClassCategoryR{
			ClassName: related,
		}
	} else {
		o.R.ClassName = related
	}

	if related.R == nil {
		related.R = &dtbClassNameR{
			ClassNameDTBClassCategories: DTBClassCategorySlice{o},
		}
	} else {
		related.R.ClassNameDTBClassCategories = append(related.R.ClassNameDTBClassCategories, o)
	}

	return nil
}

// RemoveClassName relationship.
// Sets o.R.ClassName to nil.
// Removes o from all passed in related items' relationships struct (Optional).
func (o *DTBClassCategory) RemoveClassName(ctx context.Context, exec boil.ContextExecutor, related *DTBClassName) error {
	var err error

	queries.SetScanner(&o.ClassNameID, nil)
	if _, err = o.Update(ctx, exec, boil.Whitelist("class_name_id")); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	o.R.ClassName = nil
	if related == nil || related.R == nil {
		return nil
	}

	for i, ri := range related.R.ClassNameDTBClassCategories {
		if queries.Equal(o.ClassNameID, ri.ClassNameID) {
			continue
		}

		ln := len(related.R.ClassNameDTBClassCategories)
		if ln > 1 && i < ln-1 {
			related.R.ClassNameDTBClassCategories[i] = related.R.ClassNameDTBClassCategories[ln-1]
		}
		related.R.ClassNameDTBClassCategories = related.R.ClassNameDTBClassCategories[:ln-1]
		break
	}
	return nil
}

// DTBClassCategories retrieves all the records using an executor.
func DTBClassCategories(mods ...qm.QueryMod) dtbClassCategoryQuery {
	mods = append(mods, qm.From("`dtb_class_category`"))
	return dtbClassCategoryQuery{NewQuery(mods...)}
}

// FindDTBClassCategory retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindDTBClassCategory(ctx context.Context, exec boil.ContextExecutor, iD uint, selectCols ...string) (*DTBClassCategory, error) {
	dtbClassCategoryObj := &DTBClassCategory{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from `dtb_class_category` where `id`=?", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, dtbClassCategoryObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "model: unable to select from dtb_class_category")
	}

	return dtbClassCategoryObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *DTBClassCategory) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("model: no dtb_class_category provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(dtbClassCategoryColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	dtbClassCategoryInsertCacheMut.RLock()
	cache, cached := dtbClassCategoryInsertCache[key]
	dtbClassCategoryInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			dtbClassCategoryColumns,
			dtbClassCategoryColumnsWithDefault,
			dtbClassCategoryColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(dtbClassCategoryType, dtbClassCategoryMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(dtbClassCategoryType, dtbClassCategoryMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO `dtb_class_category` (`%s`) %%sVALUES (%s)%%s", strings.Join(wl, "`,`"), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO `dtb_class_category` () VALUES ()%s%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			cache.retQuery = fmt.Sprintf("SELECT `%s` FROM `dtb_class_category` WHERE %s", strings.Join(returnColumns, "`,`"), strmangle.WhereClause("`", "`", 0, dtbClassCategoryPrimaryKeyColumns))
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
		return errors.Wrap(err, "model: unable to insert into dtb_class_category")
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
	if lastID != 0 && len(cache.retMapping) == 1 && cache.retMapping[0] == dtbClassCategoryMapping["ID"] {
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
		return errors.Wrap(err, "model: unable to populate default values for dtb_class_category")
	}

CacheNoHooks:
	if !cached {
		dtbClassCategoryInsertCacheMut.Lock()
		dtbClassCategoryInsertCache[key] = cache
		dtbClassCategoryInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the DTBClassCategory.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *DTBClassCategory) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	dtbClassCategoryUpdateCacheMut.RLock()
	cache, cached := dtbClassCategoryUpdateCache[key]
	dtbClassCategoryUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			dtbClassCategoryColumns,
			dtbClassCategoryPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("model: unable to update dtb_class_category, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE `dtb_class_category` SET %s WHERE %s",
			strmangle.SetParamNames("`", "`", 0, wl),
			strmangle.WhereClause("`", "`", 0, dtbClassCategoryPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(dtbClassCategoryType, dtbClassCategoryMapping, append(wl, dtbClassCategoryPrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "model: unable to update dtb_class_category row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "model: failed to get rows affected by update for dtb_class_category")
	}

	if !cached {
		dtbClassCategoryUpdateCacheMut.Lock()
		dtbClassCategoryUpdateCache[key] = cache
		dtbClassCategoryUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q dtbClassCategoryQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "model: unable to update all for dtb_class_category")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "model: unable to retrieve rows affected for dtb_class_category")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o DTBClassCategorySlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), dtbClassCategoryPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE `dtb_class_category` SET %s WHERE %s",
		strmangle.SetParamNames("`", "`", 0, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, dtbClassCategoryPrimaryKeyColumns, len(o)))

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "model: unable to update all in dtbClassCategory slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "model: unable to retrieve rows affected all in update all dtbClassCategory")
	}
	return rowsAff, nil
}

var mySQLDTBClassCategoryUniqueColumns = []string{
	"id",
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *DTBClassCategory) Upsert(ctx context.Context, exec boil.ContextExecutor, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("model: no dtb_class_category provided for upsert")
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(dtbClassCategoryColumnsWithDefault, o)
	nzUniques := queries.NonZeroDefaultSet(mySQLDTBClassCategoryUniqueColumns, o)

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

	dtbClassCategoryUpsertCacheMut.RLock()
	cache, cached := dtbClassCategoryUpsertCache[key]
	dtbClassCategoryUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			dtbClassCategoryColumns,
			dtbClassCategoryColumnsWithDefault,
			dtbClassCategoryColumnsWithoutDefault,
			nzDefaults,
		)
		update := updateColumns.UpdateColumnSet(
			dtbClassCategoryColumns,
			dtbClassCategoryPrimaryKeyColumns,
		)

		if len(update) == 0 {
			return errors.New("model: unable to upsert dtb_class_category, could not build update column list")
		}

		ret = strmangle.SetComplement(ret, nzUniques)
		cache.query = buildUpsertQueryMySQL(dialect, "dtb_class_category", update, insert)
		cache.retQuery = fmt.Sprintf(
			"SELECT %s FROM `dtb_class_category` WHERE %s",
			strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, ret), ","),
			strmangle.WhereClause("`", "`", 0, nzUniques),
		)

		cache.valueMapping, err = queries.BindMapping(dtbClassCategoryType, dtbClassCategoryMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(dtbClassCategoryType, dtbClassCategoryMapping, ret)
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
		return errors.Wrap(err, "model: unable to upsert for dtb_class_category")
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
	if lastID != 0 && len(cache.retMapping) == 1 && cache.retMapping[0] == dtbClassCategoryMapping["id"] {
		goto CacheNoHooks
	}

	uniqueMap, err = queries.BindMapping(dtbClassCategoryType, dtbClassCategoryMapping, nzUniques)
	if err != nil {
		return errors.Wrap(err, "model: unable to retrieve unique values for dtb_class_category")
	}
	nzUniqueCols = queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), uniqueMap)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.retQuery)
		fmt.Fprintln(boil.DebugWriter, nzUniqueCols...)
	}

	err = exec.QueryRowContext(ctx, cache.retQuery, nzUniqueCols...).Scan(returns...)
	if err != nil {
		return errors.Wrap(err, "model: unable to populate default values for dtb_class_category")
	}

CacheNoHooks:
	if !cached {
		dtbClassCategoryUpsertCacheMut.Lock()
		dtbClassCategoryUpsertCache[key] = cache
		dtbClassCategoryUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single DTBClassCategory record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *DTBClassCategory) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("model: no DTBClassCategory provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), dtbClassCategoryPrimaryKeyMapping)
	sql := "DELETE FROM `dtb_class_category` WHERE `id`=?"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "model: unable to delete from dtb_class_category")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "model: failed to get rows affected by delete for dtb_class_category")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q dtbClassCategoryQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("model: no dtbClassCategoryQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "model: unable to delete all from dtb_class_category")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "model: failed to get rows affected by deleteall for dtb_class_category")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o DTBClassCategorySlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("model: no DTBClassCategory slice provided for delete all")
	}

	if len(o) == 0 {
		return 0, nil
	}

	if len(dtbClassCategoryBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), dtbClassCategoryPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM `dtb_class_category` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, dtbClassCategoryPrimaryKeyColumns, len(o))

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args)
	}

	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "model: unable to delete all from dtbClassCategory slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "model: failed to get rows affected by deleteall for dtb_class_category")
	}

	if len(dtbClassCategoryAfterDeleteHooks) != 0 {
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
func (o *DTBClassCategory) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindDTBClassCategory(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *DTBClassCategorySlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := DTBClassCategorySlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), dtbClassCategoryPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT `dtb_class_category`.* FROM `dtb_class_category` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, dtbClassCategoryPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "model: unable to reload all in DTBClassCategorySlice")
	}

	*o = slice

	return nil
}

// DTBClassCategoryExists checks if the DTBClassCategory row exists.
func DTBClassCategoryExists(ctx context.Context, exec boil.ContextExecutor, iD uint) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from `dtb_class_category` where `id`=? limit 1)"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, iD)
	}

	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "model: unable to check if dtb_class_category exists")
	}

	return exists, nil
}
