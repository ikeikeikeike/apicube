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

func testMTBProductListMaxes(t *testing.T) {
	t.Parallel()

	query := MTBProductListMaxes()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testMTBProductListMaxesDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &MTBProductListMax{}
	if err = randomize.Struct(seed, o, mtbProductListMaxDBTypes, true, mtbProductListMaxColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize MTBProductListMax struct: %s", err)
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

	count, err := MTBProductListMaxes().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testMTBProductListMaxesQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &MTBProductListMax{}
	if err = randomize.Struct(seed, o, mtbProductListMaxDBTypes, true, mtbProductListMaxColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize MTBProductListMax struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := MTBProductListMaxes().DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := MTBProductListMaxes().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testMTBProductListMaxesSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &MTBProductListMax{}
	if err = randomize.Struct(seed, o, mtbProductListMaxDBTypes, true, mtbProductListMaxColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize MTBProductListMax struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := MTBProductListMaxSlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := MTBProductListMaxes().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testMTBProductListMaxesExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &MTBProductListMax{}
	if err = randomize.Struct(seed, o, mtbProductListMaxDBTypes, true, mtbProductListMaxColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize MTBProductListMax struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := MTBProductListMaxExists(ctx, tx, o.ID)
	if err != nil {
		t.Errorf("Unable to check if MTBProductListMax exists: %s", err)
	}
	if !e {
		t.Errorf("Expected MTBProductListMaxExists to return true, but got false.")
	}
}

func testMTBProductListMaxesFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &MTBProductListMax{}
	if err = randomize.Struct(seed, o, mtbProductListMaxDBTypes, true, mtbProductListMaxColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize MTBProductListMax struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	mtbProductListMaxFound, err := FindMTBProductListMax(ctx, tx, o.ID)
	if err != nil {
		t.Error(err)
	}

	if mtbProductListMaxFound == nil {
		t.Error("want a record, got nil")
	}
}

func testMTBProductListMaxesBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &MTBProductListMax{}
	if err = randomize.Struct(seed, o, mtbProductListMaxDBTypes, true, mtbProductListMaxColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize MTBProductListMax struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = MTBProductListMaxes().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testMTBProductListMaxesOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &MTBProductListMax{}
	if err = randomize.Struct(seed, o, mtbProductListMaxDBTypes, true, mtbProductListMaxColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize MTBProductListMax struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := MTBProductListMaxes().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testMTBProductListMaxesAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	mtbProductListMaxOne := &MTBProductListMax{}
	mtbProductListMaxTwo := &MTBProductListMax{}
	if err = randomize.Struct(seed, mtbProductListMaxOne, mtbProductListMaxDBTypes, false, mtbProductListMaxColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize MTBProductListMax struct: %s", err)
	}
	if err = randomize.Struct(seed, mtbProductListMaxTwo, mtbProductListMaxDBTypes, false, mtbProductListMaxColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize MTBProductListMax struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = mtbProductListMaxOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = mtbProductListMaxTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := MTBProductListMaxes().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testMTBProductListMaxesCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	mtbProductListMaxOne := &MTBProductListMax{}
	mtbProductListMaxTwo := &MTBProductListMax{}
	if err = randomize.Struct(seed, mtbProductListMaxOne, mtbProductListMaxDBTypes, false, mtbProductListMaxColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize MTBProductListMax struct: %s", err)
	}
	if err = randomize.Struct(seed, mtbProductListMaxTwo, mtbProductListMaxDBTypes, false, mtbProductListMaxColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize MTBProductListMax struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = mtbProductListMaxOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = mtbProductListMaxTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := MTBProductListMaxes().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func mtbProductListMaxBeforeInsertHook(ctx context.Context, e boil.ContextExecutor, o *MTBProductListMax) error {
	*o = MTBProductListMax{}
	return nil
}

func mtbProductListMaxAfterInsertHook(ctx context.Context, e boil.ContextExecutor, o *MTBProductListMax) error {
	*o = MTBProductListMax{}
	return nil
}

func mtbProductListMaxAfterSelectHook(ctx context.Context, e boil.ContextExecutor, o *MTBProductListMax) error {
	*o = MTBProductListMax{}
	return nil
}

func mtbProductListMaxBeforeUpdateHook(ctx context.Context, e boil.ContextExecutor, o *MTBProductListMax) error {
	*o = MTBProductListMax{}
	return nil
}

func mtbProductListMaxAfterUpdateHook(ctx context.Context, e boil.ContextExecutor, o *MTBProductListMax) error {
	*o = MTBProductListMax{}
	return nil
}

func mtbProductListMaxBeforeDeleteHook(ctx context.Context, e boil.ContextExecutor, o *MTBProductListMax) error {
	*o = MTBProductListMax{}
	return nil
}

func mtbProductListMaxAfterDeleteHook(ctx context.Context, e boil.ContextExecutor, o *MTBProductListMax) error {
	*o = MTBProductListMax{}
	return nil
}

func mtbProductListMaxBeforeUpsertHook(ctx context.Context, e boil.ContextExecutor, o *MTBProductListMax) error {
	*o = MTBProductListMax{}
	return nil
}

func mtbProductListMaxAfterUpsertHook(ctx context.Context, e boil.ContextExecutor, o *MTBProductListMax) error {
	*o = MTBProductListMax{}
	return nil
}

func testMTBProductListMaxesHooks(t *testing.T) {
	t.Parallel()

	var err error

	ctx := context.Background()
	empty := &MTBProductListMax{}
	o := &MTBProductListMax{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, mtbProductListMaxDBTypes, false); err != nil {
		t.Errorf("Unable to randomize MTBProductListMax object: %s", err)
	}

	AddMTBProductListMaxHook(boil.BeforeInsertHook, mtbProductListMaxBeforeInsertHook)
	if err = o.doBeforeInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	mtbProductListMaxBeforeInsertHooks = []MTBProductListMaxHook{}

	AddMTBProductListMaxHook(boil.AfterInsertHook, mtbProductListMaxAfterInsertHook)
	if err = o.doAfterInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	mtbProductListMaxAfterInsertHooks = []MTBProductListMaxHook{}

	AddMTBProductListMaxHook(boil.AfterSelectHook, mtbProductListMaxAfterSelectHook)
	if err = o.doAfterSelectHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	mtbProductListMaxAfterSelectHooks = []MTBProductListMaxHook{}

	AddMTBProductListMaxHook(boil.BeforeUpdateHook, mtbProductListMaxBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	mtbProductListMaxBeforeUpdateHooks = []MTBProductListMaxHook{}

	AddMTBProductListMaxHook(boil.AfterUpdateHook, mtbProductListMaxAfterUpdateHook)
	if err = o.doAfterUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	mtbProductListMaxAfterUpdateHooks = []MTBProductListMaxHook{}

	AddMTBProductListMaxHook(boil.BeforeDeleteHook, mtbProductListMaxBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	mtbProductListMaxBeforeDeleteHooks = []MTBProductListMaxHook{}

	AddMTBProductListMaxHook(boil.AfterDeleteHook, mtbProductListMaxAfterDeleteHook)
	if err = o.doAfterDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	mtbProductListMaxAfterDeleteHooks = []MTBProductListMaxHook{}

	AddMTBProductListMaxHook(boil.BeforeUpsertHook, mtbProductListMaxBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	mtbProductListMaxBeforeUpsertHooks = []MTBProductListMaxHook{}

	AddMTBProductListMaxHook(boil.AfterUpsertHook, mtbProductListMaxAfterUpsertHook)
	if err = o.doAfterUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	mtbProductListMaxAfterUpsertHooks = []MTBProductListMaxHook{}
}

func testMTBProductListMaxesInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &MTBProductListMax{}
	if err = randomize.Struct(seed, o, mtbProductListMaxDBTypes, true, mtbProductListMaxColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize MTBProductListMax struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := MTBProductListMaxes().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testMTBProductListMaxesInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &MTBProductListMax{}
	if err = randomize.Struct(seed, o, mtbProductListMaxDBTypes, true); err != nil {
		t.Errorf("Unable to randomize MTBProductListMax struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(mtbProductListMaxColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := MTBProductListMaxes().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testMTBProductListMaxesReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &MTBProductListMax{}
	if err = randomize.Struct(seed, o, mtbProductListMaxDBTypes, true, mtbProductListMaxColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize MTBProductListMax struct: %s", err)
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

func testMTBProductListMaxesReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &MTBProductListMax{}
	if err = randomize.Struct(seed, o, mtbProductListMaxDBTypes, true, mtbProductListMaxColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize MTBProductListMax struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := MTBProductListMaxSlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testMTBProductListMaxesSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &MTBProductListMax{}
	if err = randomize.Struct(seed, o, mtbProductListMaxDBTypes, true, mtbProductListMaxColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize MTBProductListMax struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := MTBProductListMaxes().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	mtbProductListMaxDBTypes = map[string]string{`ID`: `smallint`, `Name`: `varchar`, `SortNo`: `smallint`, `DiscriminatorType`: `varchar`}
	_                        = bytes.MinRead
)

func testMTBProductListMaxesUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(mtbProductListMaxPrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(mtbProductListMaxColumns) == len(mtbProductListMaxPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &MTBProductListMax{}
	if err = randomize.Struct(seed, o, mtbProductListMaxDBTypes, true, mtbProductListMaxColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize MTBProductListMax struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := MTBProductListMaxes().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, mtbProductListMaxDBTypes, true, mtbProductListMaxPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize MTBProductListMax struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testMTBProductListMaxesSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(mtbProductListMaxColumns) == len(mtbProductListMaxPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &MTBProductListMax{}
	if err = randomize.Struct(seed, o, mtbProductListMaxDBTypes, true, mtbProductListMaxColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize MTBProductListMax struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := MTBProductListMaxes().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, mtbProductListMaxDBTypes, true, mtbProductListMaxPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize MTBProductListMax struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(mtbProductListMaxColumns, mtbProductListMaxPrimaryKeyColumns) {
		fields = mtbProductListMaxColumns
	} else {
		fields = strmangle.SetComplement(
			mtbProductListMaxColumns,
			mtbProductListMaxPrimaryKeyColumns,
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

	slice := MTBProductListMaxSlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testMTBProductListMaxesUpsert(t *testing.T) {
	t.Parallel()

	if len(mtbProductListMaxColumns) == len(mtbProductListMaxPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}
	if len(mySQLMTBProductListMaxUniqueColumns) == 0 {
		t.Skip("Skipping table with no unique columns to conflict on")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := MTBProductListMax{}
	if err = randomize.Struct(seed, &o, mtbProductListMaxDBTypes, false); err != nil {
		t.Errorf("Unable to randomize MTBProductListMax struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(ctx, tx, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert MTBProductListMax: %s", err)
	}

	count, err := MTBProductListMaxes().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, mtbProductListMaxDBTypes, false, mtbProductListMaxPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize MTBProductListMax struct: %s", err)
	}

	if err = o.Upsert(ctx, tx, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert MTBProductListMax: %s", err)
	}

	count, err = MTBProductListMaxes().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
