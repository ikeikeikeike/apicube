// Code generated by SQLBoiler (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package model

import (
	"bytes"
	"context"
	"reflect"
	"testing"

	"github.com/volatiletech/sqlboiler/boil"
	"github.com/volatiletech/sqlboiler/queries"
	"github.com/volatiletech/sqlboiler/randomize"
	"github.com/volatiletech/sqlboiler/strmangle"
)

var (
	// Relationships sometimes use the reflection helper queries.Equal/queries.Assign
	// so force a package dependency in case they don't.
	_ = queries.Equal
)

func testMTBProductListOrderBies(t *testing.T) {
	t.Parallel()

	query := MTBProductListOrderBies()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testMTBProductListOrderBiesDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &MTBProductListOrderBy{}
	if err = randomize.Struct(seed, o, mtbProductListOrderByDBTypes, true, mtbProductListOrderByColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize MTBProductListOrderBy struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := o.Delete(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := MTBProductListOrderBies().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testMTBProductListOrderBiesQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &MTBProductListOrderBy{}
	if err = randomize.Struct(seed, o, mtbProductListOrderByDBTypes, true, mtbProductListOrderByColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize MTBProductListOrderBy struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := MTBProductListOrderBies().DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := MTBProductListOrderBies().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testMTBProductListOrderBiesSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &MTBProductListOrderBy{}
	if err = randomize.Struct(seed, o, mtbProductListOrderByDBTypes, true, mtbProductListOrderByColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize MTBProductListOrderBy struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := MTBProductListOrderBySlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := MTBProductListOrderBies().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testMTBProductListOrderBiesExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &MTBProductListOrderBy{}
	if err = randomize.Struct(seed, o, mtbProductListOrderByDBTypes, true, mtbProductListOrderByColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize MTBProductListOrderBy struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := MTBProductListOrderByExists(ctx, tx, o.ID)
	if err != nil {
		t.Errorf("Unable to check if MTBProductListOrderBy exists: %s", err)
	}
	if !e {
		t.Errorf("Expected MTBProductListOrderByExists to return true, but got false.")
	}
}

func testMTBProductListOrderBiesFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &MTBProductListOrderBy{}
	if err = randomize.Struct(seed, o, mtbProductListOrderByDBTypes, true, mtbProductListOrderByColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize MTBProductListOrderBy struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	mtbProductListOrderByFound, err := FindMTBProductListOrderBy(ctx, tx, o.ID)
	if err != nil {
		t.Error(err)
	}

	if mtbProductListOrderByFound == nil {
		t.Error("want a record, got nil")
	}
}

func testMTBProductListOrderBiesBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &MTBProductListOrderBy{}
	if err = randomize.Struct(seed, o, mtbProductListOrderByDBTypes, true, mtbProductListOrderByColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize MTBProductListOrderBy struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = MTBProductListOrderBies().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testMTBProductListOrderBiesOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &MTBProductListOrderBy{}
	if err = randomize.Struct(seed, o, mtbProductListOrderByDBTypes, true, mtbProductListOrderByColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize MTBProductListOrderBy struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := MTBProductListOrderBies().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testMTBProductListOrderBiesAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	mtbProductListOrderByOne := &MTBProductListOrderBy{}
	mtbProductListOrderByTwo := &MTBProductListOrderBy{}
	if err = randomize.Struct(seed, mtbProductListOrderByOne, mtbProductListOrderByDBTypes, false, mtbProductListOrderByColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize MTBProductListOrderBy struct: %s", err)
	}
	if err = randomize.Struct(seed, mtbProductListOrderByTwo, mtbProductListOrderByDBTypes, false, mtbProductListOrderByColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize MTBProductListOrderBy struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = mtbProductListOrderByOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = mtbProductListOrderByTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := MTBProductListOrderBies().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testMTBProductListOrderBiesCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	mtbProductListOrderByOne := &MTBProductListOrderBy{}
	mtbProductListOrderByTwo := &MTBProductListOrderBy{}
	if err = randomize.Struct(seed, mtbProductListOrderByOne, mtbProductListOrderByDBTypes, false, mtbProductListOrderByColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize MTBProductListOrderBy struct: %s", err)
	}
	if err = randomize.Struct(seed, mtbProductListOrderByTwo, mtbProductListOrderByDBTypes, false, mtbProductListOrderByColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize MTBProductListOrderBy struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = mtbProductListOrderByOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = mtbProductListOrderByTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := MTBProductListOrderBies().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func mtbProductListOrderByBeforeInsertHook(ctx context.Context, e boil.ContextExecutor, o *MTBProductListOrderBy) error {
	*o = MTBProductListOrderBy{}
	return nil
}

func mtbProductListOrderByAfterInsertHook(ctx context.Context, e boil.ContextExecutor, o *MTBProductListOrderBy) error {
	*o = MTBProductListOrderBy{}
	return nil
}

func mtbProductListOrderByAfterSelectHook(ctx context.Context, e boil.ContextExecutor, o *MTBProductListOrderBy) error {
	*o = MTBProductListOrderBy{}
	return nil
}

func mtbProductListOrderByBeforeUpdateHook(ctx context.Context, e boil.ContextExecutor, o *MTBProductListOrderBy) error {
	*o = MTBProductListOrderBy{}
	return nil
}

func mtbProductListOrderByAfterUpdateHook(ctx context.Context, e boil.ContextExecutor, o *MTBProductListOrderBy) error {
	*o = MTBProductListOrderBy{}
	return nil
}

func mtbProductListOrderByBeforeDeleteHook(ctx context.Context, e boil.ContextExecutor, o *MTBProductListOrderBy) error {
	*o = MTBProductListOrderBy{}
	return nil
}

func mtbProductListOrderByAfterDeleteHook(ctx context.Context, e boil.ContextExecutor, o *MTBProductListOrderBy) error {
	*o = MTBProductListOrderBy{}
	return nil
}

func mtbProductListOrderByBeforeUpsertHook(ctx context.Context, e boil.ContextExecutor, o *MTBProductListOrderBy) error {
	*o = MTBProductListOrderBy{}
	return nil
}

func mtbProductListOrderByAfterUpsertHook(ctx context.Context, e boil.ContextExecutor, o *MTBProductListOrderBy) error {
	*o = MTBProductListOrderBy{}
	return nil
}

func testMTBProductListOrderBiesHooks(t *testing.T) {
	t.Parallel()

	var err error

	ctx := context.Background()
	empty := &MTBProductListOrderBy{}
	o := &MTBProductListOrderBy{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, mtbProductListOrderByDBTypes, false); err != nil {
		t.Errorf("Unable to randomize MTBProductListOrderBy object: %s", err)
	}

	AddMTBProductListOrderByHook(boil.BeforeInsertHook, mtbProductListOrderByBeforeInsertHook)
	if err = o.doBeforeInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	mtbProductListOrderByBeforeInsertHooks = []MTBProductListOrderByHook{}

	AddMTBProductListOrderByHook(boil.AfterInsertHook, mtbProductListOrderByAfterInsertHook)
	if err = o.doAfterInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	mtbProductListOrderByAfterInsertHooks = []MTBProductListOrderByHook{}

	AddMTBProductListOrderByHook(boil.AfterSelectHook, mtbProductListOrderByAfterSelectHook)
	if err = o.doAfterSelectHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	mtbProductListOrderByAfterSelectHooks = []MTBProductListOrderByHook{}

	AddMTBProductListOrderByHook(boil.BeforeUpdateHook, mtbProductListOrderByBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	mtbProductListOrderByBeforeUpdateHooks = []MTBProductListOrderByHook{}

	AddMTBProductListOrderByHook(boil.AfterUpdateHook, mtbProductListOrderByAfterUpdateHook)
	if err = o.doAfterUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	mtbProductListOrderByAfterUpdateHooks = []MTBProductListOrderByHook{}

	AddMTBProductListOrderByHook(boil.BeforeDeleteHook, mtbProductListOrderByBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	mtbProductListOrderByBeforeDeleteHooks = []MTBProductListOrderByHook{}

	AddMTBProductListOrderByHook(boil.AfterDeleteHook, mtbProductListOrderByAfterDeleteHook)
	if err = o.doAfterDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	mtbProductListOrderByAfterDeleteHooks = []MTBProductListOrderByHook{}

	AddMTBProductListOrderByHook(boil.BeforeUpsertHook, mtbProductListOrderByBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	mtbProductListOrderByBeforeUpsertHooks = []MTBProductListOrderByHook{}

	AddMTBProductListOrderByHook(boil.AfterUpsertHook, mtbProductListOrderByAfterUpsertHook)
	if err = o.doAfterUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	mtbProductListOrderByAfterUpsertHooks = []MTBProductListOrderByHook{}
}

func testMTBProductListOrderBiesInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &MTBProductListOrderBy{}
	if err = randomize.Struct(seed, o, mtbProductListOrderByDBTypes, true, mtbProductListOrderByColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize MTBProductListOrderBy struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := MTBProductListOrderBies().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testMTBProductListOrderBiesInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &MTBProductListOrderBy{}
	if err = randomize.Struct(seed, o, mtbProductListOrderByDBTypes, true); err != nil {
		t.Errorf("Unable to randomize MTBProductListOrderBy struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(mtbProductListOrderByColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := MTBProductListOrderBies().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testMTBProductListOrderBiesReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &MTBProductListOrderBy{}
	if err = randomize.Struct(seed, o, mtbProductListOrderByDBTypes, true, mtbProductListOrderByColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize MTBProductListOrderBy struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = o.Reload(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testMTBProductListOrderBiesReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &MTBProductListOrderBy{}
	if err = randomize.Struct(seed, o, mtbProductListOrderByDBTypes, true, mtbProductListOrderByColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize MTBProductListOrderBy struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := MTBProductListOrderBySlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testMTBProductListOrderBiesSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &MTBProductListOrderBy{}
	if err = randomize.Struct(seed, o, mtbProductListOrderByDBTypes, true, mtbProductListOrderByColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize MTBProductListOrderBy struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := MTBProductListOrderBies().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	mtbProductListOrderByDBTypes = map[string]string{`ID`: `smallint`, `Name`: `varchar`, `SortNo`: `smallint`, `DiscriminatorType`: `varchar`}
	_                            = bytes.MinRead
)

func testMTBProductListOrderBiesUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(mtbProductListOrderByPrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(mtbProductListOrderByColumns) == len(mtbProductListOrderByPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &MTBProductListOrderBy{}
	if err = randomize.Struct(seed, o, mtbProductListOrderByDBTypes, true, mtbProductListOrderByColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize MTBProductListOrderBy struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := MTBProductListOrderBies().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, mtbProductListOrderByDBTypes, true, mtbProductListOrderByPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize MTBProductListOrderBy struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testMTBProductListOrderBiesSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(mtbProductListOrderByColumns) == len(mtbProductListOrderByPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &MTBProductListOrderBy{}
	if err = randomize.Struct(seed, o, mtbProductListOrderByDBTypes, true, mtbProductListOrderByColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize MTBProductListOrderBy struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := MTBProductListOrderBies().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, mtbProductListOrderByDBTypes, true, mtbProductListOrderByPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize MTBProductListOrderBy struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(mtbProductListOrderByColumns, mtbProductListOrderByPrimaryKeyColumns) {
		fields = mtbProductListOrderByColumns
	} else {
		fields = strmangle.SetComplement(
			mtbProductListOrderByColumns,
			mtbProductListOrderByPrimaryKeyColumns,
		)
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	typ := reflect.TypeOf(o).Elem()
	n := typ.NumField()

	updateMap := M{}
	for _, col := range fields {
		for i := 0; i < n; i++ {
			f := typ.Field(i)
			if f.Tag.Get("boil") == col {
				updateMap[col] = value.Field(i).Interface()
			}
		}
	}

	slice := MTBProductListOrderBySlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testMTBProductListOrderBiesUpsert(t *testing.T) {
	t.Parallel()

	if len(mtbProductListOrderByColumns) == len(mtbProductListOrderByPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}
	if len(mySQLMTBProductListOrderByUniqueColumns) == 0 {
		t.Skip("Skipping table with no unique columns to conflict on")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := MTBProductListOrderBy{}
	if err = randomize.Struct(seed, &o, mtbProductListOrderByDBTypes, false); err != nil {
		t.Errorf("Unable to randomize MTBProductListOrderBy struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(ctx, tx, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert MTBProductListOrderBy: %s", err)
	}

	count, err := MTBProductListOrderBies().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, mtbProductListOrderByDBTypes, false, mtbProductListOrderByPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize MTBProductListOrderBy struct: %s", err)
	}

	if err = o.Upsert(ctx, tx, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert MTBProductListOrderBy: %s", err)
	}

	count, err = MTBProductListOrderBies().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
