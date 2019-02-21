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

// DTBBlockPosition is an object representing the database table.
type DTBBlockPosition struct {
	Section           uint      `boil:"section" json:"section" toml:"section" yaml:"section"`
	BlockID           uint      `boil:"block_id" json:"block_id" toml:"block_id" yaml:"block_id"`
	LayoutID          uint      `boil:"layout_id" json:"layout_id" toml:"layout_id" yaml:"layout_id"`
	BlockRow          null.Uint `boil:"block_row" json:"block_row,omitempty" toml:"block_row" yaml:"block_row,omitempty"`
	DiscriminatorType string    `boil:"discriminator_type" json:"discriminator_type" toml:"discriminator_type" yaml:"discriminator_type"`

	R *dtbBlockPositionR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L dtbBlockPositionL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var DTBBlockPositionColumns = struct {
	Section           string
	BlockID           string
	LayoutID          string
	BlockRow          string
	DiscriminatorType string
}{
	Section:           "section",
	BlockID:           "block_id",
	LayoutID:          "layout_id",
	BlockRow:          "block_row",
	DiscriminatorType: "discriminator_type",
}

// Generated where

var DTBBlockPositionWhere = struct {
	Section           whereHelperuint
	BlockID           whereHelperuint
	LayoutID          whereHelperuint
	BlockRow          whereHelpernull_Uint
	DiscriminatorType whereHelperstring
}{
	Section:           whereHelperuint{field: `section`},
	BlockID:           whereHelperuint{field: `block_id`},
	LayoutID:          whereHelperuint{field: `layout_id`},
	BlockRow:          whereHelpernull_Uint{field: `block_row`},
	DiscriminatorType: whereHelperstring{field: `discriminator_type`},
}

// DTBBlockPositionRels is where relationship names are stored.
var DTBBlockPositionRels = struct {
	Layout string
	Block  string
}{
	Layout: "Layout",
	Block:  "Block",
}

// dtbBlockPositionR is where relationships are stored.
type dtbBlockPositionR struct {
	Layout *DTBLayout
	Block  *DTBBlock
}

// NewStruct creates a new relationship struct
func (*dtbBlockPositionR) NewStruct() *dtbBlockPositionR {
	return &dtbBlockPositionR{}
}

// dtbBlockPositionL is where Load methods for each relationship are stored.
type dtbBlockPositionL struct{}

var (
	dtbBlockPositionColumns               = []string{"section", "block_id", "layout_id", "block_row", "discriminator_type"}
	dtbBlockPositionColumnsWithoutDefault = []string{"section", "block_id", "layout_id", "block_row", "discriminator_type"}
	dtbBlockPositionColumnsWithDefault    = []string{}
	dtbBlockPositionPrimaryKeyColumns     = []string{"section", "block_id", "layout_id"}
)

type (
	// DTBBlockPositionSlice is an alias for a slice of pointers to DTBBlockPosition.
	// This should generally be used opposed to []DTBBlockPosition.
	DTBBlockPositionSlice []*DTBBlockPosition
	// DTBBlockPositionHook is the signature for custom DTBBlockPosition hook methods
	DTBBlockPositionHook func(context.Context, boil.ContextExecutor, *DTBBlockPosition) error

	dtbBlockPositionQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	dtbBlockPositionType                 = reflect.TypeOf(&DTBBlockPosition{})
	dtbBlockPositionMapping              = queries.MakeStructMapping(dtbBlockPositionType)
	dtbBlockPositionPrimaryKeyMapping, _ = queries.BindMapping(dtbBlockPositionType, dtbBlockPositionMapping, dtbBlockPositionPrimaryKeyColumns)
	dtbBlockPositionInsertCacheMut       sync.RWMutex
	dtbBlockPositionInsertCache          = make(map[string]insertCache)
	dtbBlockPositionUpdateCacheMut       sync.RWMutex
	dtbBlockPositionUpdateCache          = make(map[string]updateCache)
	dtbBlockPositionUpsertCacheMut       sync.RWMutex
	dtbBlockPositionUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var dtbBlockPositionBeforeInsertHooks []DTBBlockPositionHook
var dtbBlockPositionBeforeUpdateHooks []DTBBlockPositionHook
var dtbBlockPositionBeforeDeleteHooks []DTBBlockPositionHook
var dtbBlockPositionBeforeUpsertHooks []DTBBlockPositionHook

var dtbBlockPositionAfterInsertHooks []DTBBlockPositionHook
var dtbBlockPositionAfterSelectHooks []DTBBlockPositionHook
var dtbBlockPositionAfterUpdateHooks []DTBBlockPositionHook
var dtbBlockPositionAfterDeleteHooks []DTBBlockPositionHook
var dtbBlockPositionAfterUpsertHooks []DTBBlockPositionHook

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *DTBBlockPosition) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range dtbBlockPositionBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *DTBBlockPosition) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range dtbBlockPositionBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *DTBBlockPosition) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range dtbBlockPositionBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *DTBBlockPosition) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range dtbBlockPositionBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *DTBBlockPosition) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range dtbBlockPositionAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterSelectHooks executes all "after Select" hooks.
func (o *DTBBlockPosition) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range dtbBlockPositionAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *DTBBlockPosition) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range dtbBlockPositionAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *DTBBlockPosition) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range dtbBlockPositionAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *DTBBlockPosition) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range dtbBlockPositionAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddDTBBlockPositionHook registers your hook function for all future operations.
func AddDTBBlockPositionHook(hookPoint boil.HookPoint, dtbBlockPositionHook DTBBlockPositionHook) {
	switch hookPoint {
	case boil.BeforeInsertHook:
		dtbBlockPositionBeforeInsertHooks = append(dtbBlockPositionBeforeInsertHooks, dtbBlockPositionHook)
	case boil.BeforeUpdateHook:
		dtbBlockPositionBeforeUpdateHooks = append(dtbBlockPositionBeforeUpdateHooks, dtbBlockPositionHook)
	case boil.BeforeDeleteHook:
		dtbBlockPositionBeforeDeleteHooks = append(dtbBlockPositionBeforeDeleteHooks, dtbBlockPositionHook)
	case boil.BeforeUpsertHook:
		dtbBlockPositionBeforeUpsertHooks = append(dtbBlockPositionBeforeUpsertHooks, dtbBlockPositionHook)
	case boil.AfterInsertHook:
		dtbBlockPositionAfterInsertHooks = append(dtbBlockPositionAfterInsertHooks, dtbBlockPositionHook)
	case boil.AfterSelectHook:
		dtbBlockPositionAfterSelectHooks = append(dtbBlockPositionAfterSelectHooks, dtbBlockPositionHook)
	case boil.AfterUpdateHook:
		dtbBlockPositionAfterUpdateHooks = append(dtbBlockPositionAfterUpdateHooks, dtbBlockPositionHook)
	case boil.AfterDeleteHook:
		dtbBlockPositionAfterDeleteHooks = append(dtbBlockPositionAfterDeleteHooks, dtbBlockPositionHook)
	case boil.AfterUpsertHook:
		dtbBlockPositionAfterUpsertHooks = append(dtbBlockPositionAfterUpsertHooks, dtbBlockPositionHook)
	}
}

// One returns a single dtbBlockPosition record from the query.
func (q dtbBlockPositionQuery) One(ctx context.Context, exec boil.ContextExecutor) (*DTBBlockPosition, error) {
	o := &DTBBlockPosition{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "model: failed to execute a one query for dtb_block_position")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all DTBBlockPosition records from the query.
func (q dtbBlockPositionQuery) All(ctx context.Context, exec boil.ContextExecutor) (DTBBlockPositionSlice, error) {
	var o []*DTBBlockPosition

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "model: failed to assign all query results to DTBBlockPosition slice")
	}

	if len(dtbBlockPositionAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all DTBBlockPosition records in the query.
func (q dtbBlockPositionQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "model: failed to count dtb_block_position rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q dtbBlockPositionQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "model: failed to check if dtb_block_position exists")
	}

	return count > 0, nil
}

// Layout pointed to by the foreign key.
func (o *DTBBlockPosition) Layout(mods ...qm.QueryMod) dtbLayoutQuery {
	queryMods := []qm.QueryMod{
		qm.Where("id=?", o.LayoutID),
	}

	queryMods = append(queryMods, mods...)

	query := DTBLayouts(queryMods...)
	queries.SetFrom(query.Query, "`dtb_layout`")

	return query
}

// Block pointed to by the foreign key.
func (o *DTBBlockPosition) Block(mods ...qm.QueryMod) dtbBlockQuery {
	queryMods := []qm.QueryMod{
		qm.Where("id=?", o.BlockID),
	}

	queryMods = append(queryMods, mods...)

	query := DTBBlocks(queryMods...)
	queries.SetFrom(query.Query, "`dtb_block`")

	return query
}

// LoadLayout allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (dtbBlockPositionL) LoadLayout(ctx context.Context, e boil.ContextExecutor, singular bool, maybeDTBBlockPosition interface{}, mods queries.Applicator) error {
	var slice []*DTBBlockPosition
	var object *DTBBlockPosition

	if singular {
		object = maybeDTBBlockPosition.(*DTBBlockPosition)
	} else {
		slice = *maybeDTBBlockPosition.(*[]*DTBBlockPosition)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &dtbBlockPositionR{}
		}
		args = append(args, object.LayoutID)

	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &dtbBlockPositionR{}
			}

			for _, a := range args {
				if a == obj.LayoutID {
					continue Outer
				}
			}

			args = append(args, obj.LayoutID)

		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(qm.From(`dtb_layout`), qm.WhereIn(`id in ?`, args...))
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load DTBLayout")
	}

	var resultSlice []*DTBLayout
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice DTBLayout")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results of eager load for dtb_layout")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for dtb_layout")
	}

	if len(dtbBlockPositionAfterSelectHooks) != 0 {
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
		object.R.Layout = foreign
		if foreign.R == nil {
			foreign.R = &dtbLayoutR{}
		}
		foreign.R.LayoutDTBBlockPositions = append(foreign.R.LayoutDTBBlockPositions, object)
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if local.LayoutID == foreign.ID {
				local.R.Layout = foreign
				if foreign.R == nil {
					foreign.R = &dtbLayoutR{}
				}
				foreign.R.LayoutDTBBlockPositions = append(foreign.R.LayoutDTBBlockPositions, local)
				break
			}
		}
	}

	return nil
}

// LoadBlock allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (dtbBlockPositionL) LoadBlock(ctx context.Context, e boil.ContextExecutor, singular bool, maybeDTBBlockPosition interface{}, mods queries.Applicator) error {
	var slice []*DTBBlockPosition
	var object *DTBBlockPosition

	if singular {
		object = maybeDTBBlockPosition.(*DTBBlockPosition)
	} else {
		slice = *maybeDTBBlockPosition.(*[]*DTBBlockPosition)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &dtbBlockPositionR{}
		}
		args = append(args, object.BlockID)

	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &dtbBlockPositionR{}
			}

			for _, a := range args {
				if a == obj.BlockID {
					continue Outer
				}
			}

			args = append(args, obj.BlockID)

		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(qm.From(`dtb_block`), qm.WhereIn(`id in ?`, args...))
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load DTBBlock")
	}

	var resultSlice []*DTBBlock
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice DTBBlock")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results of eager load for dtb_block")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for dtb_block")
	}

	if len(dtbBlockPositionAfterSelectHooks) != 0 {
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
		object.R.Block = foreign
		if foreign.R == nil {
			foreign.R = &dtbBlockR{}
		}
		foreign.R.BlockDTBBlockPositions = append(foreign.R.BlockDTBBlockPositions, object)
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if local.BlockID == foreign.ID {
				local.R.Block = foreign
				if foreign.R == nil {
					foreign.R = &dtbBlockR{}
				}
				foreign.R.BlockDTBBlockPositions = append(foreign.R.BlockDTBBlockPositions, local)
				break
			}
		}
	}

	return nil
}

// SetLayout of the dtbBlockPosition to the related item.
// Sets o.R.Layout to related.
// Adds o to related.R.LayoutDTBBlockPositions.
func (o *DTBBlockPosition) SetLayout(ctx context.Context, exec boil.ContextExecutor, insert bool, related *DTBLayout) error {
	var err error
	if insert {
		if err = related.Insert(ctx, exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE `dtb_block_position` SET %s WHERE %s",
		strmangle.SetParamNames("`", "`", 0, []string{"layout_id"}),
		strmangle.WhereClause("`", "`", 0, dtbBlockPositionPrimaryKeyColumns),
	)
	values := []interface{}{related.ID, o.Section, o.BlockID, o.LayoutID}

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, updateQuery)
		fmt.Fprintln(boil.DebugWriter, values)
	}

	if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	o.LayoutID = related.ID
	if o.R == nil {
		o.R = &dtbBlockPositionR{
			Layout: related,
		}
	} else {
		o.R.Layout = related
	}

	if related.R == nil {
		related.R = &dtbLayoutR{
			LayoutDTBBlockPositions: DTBBlockPositionSlice{o},
		}
	} else {
		related.R.LayoutDTBBlockPositions = append(related.R.LayoutDTBBlockPositions, o)
	}

	return nil
}

// SetBlock of the dtbBlockPosition to the related item.
// Sets o.R.Block to related.
// Adds o to related.R.BlockDTBBlockPositions.
func (o *DTBBlockPosition) SetBlock(ctx context.Context, exec boil.ContextExecutor, insert bool, related *DTBBlock) error {
	var err error
	if insert {
		if err = related.Insert(ctx, exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE `dtb_block_position` SET %s WHERE %s",
		strmangle.SetParamNames("`", "`", 0, []string{"block_id"}),
		strmangle.WhereClause("`", "`", 0, dtbBlockPositionPrimaryKeyColumns),
	)
	values := []interface{}{related.ID, o.Section, o.BlockID, o.LayoutID}

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, updateQuery)
		fmt.Fprintln(boil.DebugWriter, values)
	}

	if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	o.BlockID = related.ID
	if o.R == nil {
		o.R = &dtbBlockPositionR{
			Block: related,
		}
	} else {
		o.R.Block = related
	}

	if related.R == nil {
		related.R = &dtbBlockR{
			BlockDTBBlockPositions: DTBBlockPositionSlice{o},
		}
	} else {
		related.R.BlockDTBBlockPositions = append(related.R.BlockDTBBlockPositions, o)
	}

	return nil
}

// DTBBlockPositions retrieves all the records using an executor.
func DTBBlockPositions(mods ...qm.QueryMod) dtbBlockPositionQuery {
	mods = append(mods, qm.From("`dtb_block_position`"))
	return dtbBlockPositionQuery{NewQuery(mods...)}
}

// FindDTBBlockPosition retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindDTBBlockPosition(ctx context.Context, exec boil.ContextExecutor, section uint, blockID uint, layoutID uint, selectCols ...string) (*DTBBlockPosition, error) {
	dtbBlockPositionObj := &DTBBlockPosition{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from `dtb_block_position` where `section`=? AND `block_id`=? AND `layout_id`=?", sel,
	)

	q := queries.Raw(query, section, blockID, layoutID)

	err := q.Bind(ctx, exec, dtbBlockPositionObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "model: unable to select from dtb_block_position")
	}

	return dtbBlockPositionObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *DTBBlockPosition) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("model: no dtb_block_position provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(dtbBlockPositionColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	dtbBlockPositionInsertCacheMut.RLock()
	cache, cached := dtbBlockPositionInsertCache[key]
	dtbBlockPositionInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			dtbBlockPositionColumns,
			dtbBlockPositionColumnsWithDefault,
			dtbBlockPositionColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(dtbBlockPositionType, dtbBlockPositionMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(dtbBlockPositionType, dtbBlockPositionMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO `dtb_block_position` (`%s`) %%sVALUES (%s)%%s", strings.Join(wl, "`,`"), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO `dtb_block_position` () VALUES ()%s%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			cache.retQuery = fmt.Sprintf("SELECT `%s` FROM `dtb_block_position` WHERE %s", strings.Join(returnColumns, "`,`"), strmangle.WhereClause("`", "`", 0, dtbBlockPositionPrimaryKeyColumns))
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
		return errors.Wrap(err, "model: unable to insert into dtb_block_position")
	}

	var identifierCols []interface{}

	if len(cache.retMapping) == 0 {
		goto CacheNoHooks
	}

	identifierCols = []interface{}{
		o.Section,
		o.BlockID,
		o.LayoutID,
	}

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.retQuery)
		fmt.Fprintln(boil.DebugWriter, identifierCols...)
	}

	err = exec.QueryRowContext(ctx, cache.retQuery, identifierCols...).Scan(queries.PtrsFromMapping(value, cache.retMapping)...)
	if err != nil {
		return errors.Wrap(err, "model: unable to populate default values for dtb_block_position")
	}

CacheNoHooks:
	if !cached {
		dtbBlockPositionInsertCacheMut.Lock()
		dtbBlockPositionInsertCache[key] = cache
		dtbBlockPositionInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the DTBBlockPosition.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *DTBBlockPosition) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	dtbBlockPositionUpdateCacheMut.RLock()
	cache, cached := dtbBlockPositionUpdateCache[key]
	dtbBlockPositionUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			dtbBlockPositionColumns,
			dtbBlockPositionPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("model: unable to update dtb_block_position, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE `dtb_block_position` SET %s WHERE %s",
			strmangle.SetParamNames("`", "`", 0, wl),
			strmangle.WhereClause("`", "`", 0, dtbBlockPositionPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(dtbBlockPositionType, dtbBlockPositionMapping, append(wl, dtbBlockPositionPrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "model: unable to update dtb_block_position row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "model: failed to get rows affected by update for dtb_block_position")
	}

	if !cached {
		dtbBlockPositionUpdateCacheMut.Lock()
		dtbBlockPositionUpdateCache[key] = cache
		dtbBlockPositionUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q dtbBlockPositionQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "model: unable to update all for dtb_block_position")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "model: unable to retrieve rows affected for dtb_block_position")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o DTBBlockPositionSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), dtbBlockPositionPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE `dtb_block_position` SET %s WHERE %s",
		strmangle.SetParamNames("`", "`", 0, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, dtbBlockPositionPrimaryKeyColumns, len(o)))

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "model: unable to update all in dtbBlockPosition slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "model: unable to retrieve rows affected all in update all dtbBlockPosition")
	}
	return rowsAff, nil
}

var mySQLDTBBlockPositionUniqueColumns = []string{}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *DTBBlockPosition) Upsert(ctx context.Context, exec boil.ContextExecutor, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("model: no dtb_block_position provided for upsert")
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(dtbBlockPositionColumnsWithDefault, o)
	nzUniques := queries.NonZeroDefaultSet(mySQLDTBBlockPositionUniqueColumns, o)

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

	dtbBlockPositionUpsertCacheMut.RLock()
	cache, cached := dtbBlockPositionUpsertCache[key]
	dtbBlockPositionUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			dtbBlockPositionColumns,
			dtbBlockPositionColumnsWithDefault,
			dtbBlockPositionColumnsWithoutDefault,
			nzDefaults,
		)
		update := updateColumns.UpdateColumnSet(
			dtbBlockPositionColumns,
			dtbBlockPositionPrimaryKeyColumns,
		)

		if len(update) == 0 {
			return errors.New("model: unable to upsert dtb_block_position, could not build update column list")
		}

		ret = strmangle.SetComplement(ret, nzUniques)
		cache.query = buildUpsertQueryMySQL(dialect, "dtb_block_position", update, insert)
		cache.retQuery = fmt.Sprintf(
			"SELECT %s FROM `dtb_block_position` WHERE %s",
			strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, ret), ","),
			strmangle.WhereClause("`", "`", 0, nzUniques),
		)

		cache.valueMapping, err = queries.BindMapping(dtbBlockPositionType, dtbBlockPositionMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(dtbBlockPositionType, dtbBlockPositionMapping, ret)
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
		return errors.Wrap(err, "model: unable to upsert for dtb_block_position")
	}

	var uniqueMap []uint64
	var nzUniqueCols []interface{}

	if len(cache.retMapping) == 0 {
		goto CacheNoHooks
	}

	uniqueMap, err = queries.BindMapping(dtbBlockPositionType, dtbBlockPositionMapping, nzUniques)
	if err != nil {
		return errors.Wrap(err, "model: unable to retrieve unique values for dtb_block_position")
	}
	nzUniqueCols = queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), uniqueMap)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.retQuery)
		fmt.Fprintln(boil.DebugWriter, nzUniqueCols...)
	}

	err = exec.QueryRowContext(ctx, cache.retQuery, nzUniqueCols...).Scan(returns...)
	if err != nil {
		return errors.Wrap(err, "model: unable to populate default values for dtb_block_position")
	}

CacheNoHooks:
	if !cached {
		dtbBlockPositionUpsertCacheMut.Lock()
		dtbBlockPositionUpsertCache[key] = cache
		dtbBlockPositionUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single DTBBlockPosition record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *DTBBlockPosition) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("model: no DTBBlockPosition provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), dtbBlockPositionPrimaryKeyMapping)
	sql := "DELETE FROM `dtb_block_position` WHERE `section`=? AND `block_id`=? AND `layout_id`=?"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "model: unable to delete from dtb_block_position")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "model: failed to get rows affected by delete for dtb_block_position")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q dtbBlockPositionQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("model: no dtbBlockPositionQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "model: unable to delete all from dtb_block_position")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "model: failed to get rows affected by deleteall for dtb_block_position")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o DTBBlockPositionSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("model: no DTBBlockPosition slice provided for delete all")
	}

	if len(o) == 0 {
		return 0, nil
	}

	if len(dtbBlockPositionBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), dtbBlockPositionPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM `dtb_block_position` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, dtbBlockPositionPrimaryKeyColumns, len(o))

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args)
	}

	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "model: unable to delete all from dtbBlockPosition slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "model: failed to get rows affected by deleteall for dtb_block_position")
	}

	if len(dtbBlockPositionAfterDeleteHooks) != 0 {
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
func (o *DTBBlockPosition) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindDTBBlockPosition(ctx, exec, o.Section, o.BlockID, o.LayoutID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *DTBBlockPositionSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := DTBBlockPositionSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), dtbBlockPositionPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT `dtb_block_position`.* FROM `dtb_block_position` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, dtbBlockPositionPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "model: unable to reload all in DTBBlockPositionSlice")
	}

	*o = slice

	return nil
}

// DTBBlockPositionExists checks if the DTBBlockPosition row exists.
func DTBBlockPositionExists(ctx context.Context, exec boil.ContextExecutor, section uint, blockID uint, layoutID uint) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from `dtb_block_position` where `section`=? AND `block_id`=? AND `layout_id`=? limit 1)"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, section, blockID, layoutID)
	}

	row := exec.QueryRowContext(ctx, sql, section, blockID, layoutID)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "model: unable to check if dtb_block_position exists")
	}

	return exists, nil
}