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

func testDTBOrderPDFS(t *testing.T) {
	t.Parallel()

	query := DTBOrderPDFS()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testDTBOrderPDFSDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &DTBOrderPDF{}
	if err = randomize.Struct(seed, o, dtbOrderPDFDBTypes, true, dtbOrderPDFColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DTBOrderPDF struct: %s", err)
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

	count, err := DTBOrderPDFS().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testDTBOrderPDFSQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &DTBOrderPDF{}
	if err = randomize.Struct(seed, o, dtbOrderPDFDBTypes, true, dtbOrderPDFColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DTBOrderPDF struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := DTBOrderPDFS().DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := DTBOrderPDFS().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testDTBOrderPDFSSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &DTBOrderPDF{}
	if err = randomize.Struct(seed, o, dtbOrderPDFDBTypes, true, dtbOrderPDFColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DTBOrderPDF struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := DTBOrderPDFSlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := DTBOrderPDFS().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testDTBOrderPDFSExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &DTBOrderPDF{}
	if err = randomize.Struct(seed, o, dtbOrderPDFDBTypes, true, dtbOrderPDFColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DTBOrderPDF struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := DTBOrderPDFExists(ctx, tx, o.MemberID)
	if err != nil {
		t.Errorf("Unable to check if DTBOrderPDF exists: %s", err)
	}
	if !e {
		t.Errorf("Expected DTBOrderPDFExists to return true, but got false.")
	}
}

func testDTBOrderPDFSFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &DTBOrderPDF{}
	if err = randomize.Struct(seed, o, dtbOrderPDFDBTypes, true, dtbOrderPDFColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DTBOrderPDF struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	dtbOrderPDFFound, err := FindDTBOrderPDF(ctx, tx, o.MemberID)
	if err != nil {
		t.Error(err)
	}

	if dtbOrderPDFFound == nil {
		t.Error("want a record, got nil")
	}
}

func testDTBOrderPDFSBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &DTBOrderPDF{}
	if err = randomize.Struct(seed, o, dtbOrderPDFDBTypes, true, dtbOrderPDFColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DTBOrderPDF struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = DTBOrderPDFS().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testDTBOrderPDFSOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &DTBOrderPDF{}
	if err = randomize.Struct(seed, o, dtbOrderPDFDBTypes, true, dtbOrderPDFColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DTBOrderPDF struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := DTBOrderPDFS().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testDTBOrderPDFSAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	dtbOrderPDFOne := &DTBOrderPDF{}
	dtbOrderPDFTwo := &DTBOrderPDF{}
	if err = randomize.Struct(seed, dtbOrderPDFOne, dtbOrderPDFDBTypes, false, dtbOrderPDFColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DTBOrderPDF struct: %s", err)
	}
	if err = randomize.Struct(seed, dtbOrderPDFTwo, dtbOrderPDFDBTypes, false, dtbOrderPDFColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DTBOrderPDF struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = dtbOrderPDFOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = dtbOrderPDFTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := DTBOrderPDFS().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testDTBOrderPDFSCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	dtbOrderPDFOne := &DTBOrderPDF{}
	dtbOrderPDFTwo := &DTBOrderPDF{}
	if err = randomize.Struct(seed, dtbOrderPDFOne, dtbOrderPDFDBTypes, false, dtbOrderPDFColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DTBOrderPDF struct: %s", err)
	}
	if err = randomize.Struct(seed, dtbOrderPDFTwo, dtbOrderPDFDBTypes, false, dtbOrderPDFColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DTBOrderPDF struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = dtbOrderPDFOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = dtbOrderPDFTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := DTBOrderPDFS().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func dtbOrderPDFBeforeInsertHook(ctx context.Context, e boil.ContextExecutor, o *DTBOrderPDF) error {
	*o = DTBOrderPDF{}
	return nil
}

func dtbOrderPDFAfterInsertHook(ctx context.Context, e boil.ContextExecutor, o *DTBOrderPDF) error {
	*o = DTBOrderPDF{}
	return nil
}

func dtbOrderPDFAfterSelectHook(ctx context.Context, e boil.ContextExecutor, o *DTBOrderPDF) error {
	*o = DTBOrderPDF{}
	return nil
}

func dtbOrderPDFBeforeUpdateHook(ctx context.Context, e boil.ContextExecutor, o *DTBOrderPDF) error {
	*o = DTBOrderPDF{}
	return nil
}

func dtbOrderPDFAfterUpdateHook(ctx context.Context, e boil.ContextExecutor, o *DTBOrderPDF) error {
	*o = DTBOrderPDF{}
	return nil
}

func dtbOrderPDFBeforeDeleteHook(ctx context.Context, e boil.ContextExecutor, o *DTBOrderPDF) error {
	*o = DTBOrderPDF{}
	return nil
}

func dtbOrderPDFAfterDeleteHook(ctx context.Context, e boil.ContextExecutor, o *DTBOrderPDF) error {
	*o = DTBOrderPDF{}
	return nil
}

func dtbOrderPDFBeforeUpsertHook(ctx context.Context, e boil.ContextExecutor, o *DTBOrderPDF) error {
	*o = DTBOrderPDF{}
	return nil
}

func dtbOrderPDFAfterUpsertHook(ctx context.Context, e boil.ContextExecutor, o *DTBOrderPDF) error {
	*o = DTBOrderPDF{}
	return nil
}

func testDTBOrderPDFSHooks(t *testing.T) {
	t.Parallel()

	var err error

	ctx := context.Background()
	empty := &DTBOrderPDF{}
	o := &DTBOrderPDF{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, dtbOrderPDFDBTypes, false); err != nil {
		t.Errorf("Unable to randomize DTBOrderPDF object: %s", err)
	}

	AddDTBOrderPDFHook(boil.BeforeInsertHook, dtbOrderPDFBeforeInsertHook)
	if err = o.doBeforeInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	dtbOrderPDFBeforeInsertHooks = []DTBOrderPDFHook{}

	AddDTBOrderPDFHook(boil.AfterInsertHook, dtbOrderPDFAfterInsertHook)
	if err = o.doAfterInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	dtbOrderPDFAfterInsertHooks = []DTBOrderPDFHook{}

	AddDTBOrderPDFHook(boil.AfterSelectHook, dtbOrderPDFAfterSelectHook)
	if err = o.doAfterSelectHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	dtbOrderPDFAfterSelectHooks = []DTBOrderPDFHook{}

	AddDTBOrderPDFHook(boil.BeforeUpdateHook, dtbOrderPDFBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	dtbOrderPDFBeforeUpdateHooks = []DTBOrderPDFHook{}

	AddDTBOrderPDFHook(boil.AfterUpdateHook, dtbOrderPDFAfterUpdateHook)
	if err = o.doAfterUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	dtbOrderPDFAfterUpdateHooks = []DTBOrderPDFHook{}

	AddDTBOrderPDFHook(boil.BeforeDeleteHook, dtbOrderPDFBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	dtbOrderPDFBeforeDeleteHooks = []DTBOrderPDFHook{}

	AddDTBOrderPDFHook(boil.AfterDeleteHook, dtbOrderPDFAfterDeleteHook)
	if err = o.doAfterDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	dtbOrderPDFAfterDeleteHooks = []DTBOrderPDFHook{}

	AddDTBOrderPDFHook(boil.BeforeUpsertHook, dtbOrderPDFBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	dtbOrderPDFBeforeUpsertHooks = []DTBOrderPDFHook{}

	AddDTBOrderPDFHook(boil.AfterUpsertHook, dtbOrderPDFAfterUpsertHook)
	if err = o.doAfterUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	dtbOrderPDFAfterUpsertHooks = []DTBOrderPDFHook{}
}

func testDTBOrderPDFSInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &DTBOrderPDF{}
	if err = randomize.Struct(seed, o, dtbOrderPDFDBTypes, true, dtbOrderPDFColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DTBOrderPDF struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := DTBOrderPDFS().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testDTBOrderPDFSInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &DTBOrderPDF{}
	if err = randomize.Struct(seed, o, dtbOrderPDFDBTypes, true); err != nil {
		t.Errorf("Unable to randomize DTBOrderPDF struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(dtbOrderPDFColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := DTBOrderPDFS().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testDTBOrderPDFSReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &DTBOrderPDF{}
	if err = randomize.Struct(seed, o, dtbOrderPDFDBTypes, true, dtbOrderPDFColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DTBOrderPDF struct: %s", err)
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

func testDTBOrderPDFSReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &DTBOrderPDF{}
	if err = randomize.Struct(seed, o, dtbOrderPDFDBTypes, true, dtbOrderPDFColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DTBOrderPDF struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := DTBOrderPDFSlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testDTBOrderPDFSSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &DTBOrderPDF{}
	if err = randomize.Struct(seed, o, dtbOrderPDFDBTypes, true, dtbOrderPDFColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DTBOrderPDF struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := DTBOrderPDFS().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	dtbOrderPDFDBTypes = map[string]string{`MemberID`: `int`, `Title`: `varchar`, `Message1`: `varchar`, `Message2`: `varchar`, `Message3`: `varchar`, `Note1`: `varchar`, `Note2`: `varchar`, `Note3`: `varchar`, `CreateDate`: `datetime`, `UpdateDate`: `datetime`, `Visible`: `tinyint`, `DiscriminatorType`: `varchar`}
	_                  = bytes.MinRead
)

func testDTBOrderPDFSUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(dtbOrderPDFPrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(dtbOrderPDFColumns) == len(dtbOrderPDFPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &DTBOrderPDF{}
	if err = randomize.Struct(seed, o, dtbOrderPDFDBTypes, true, dtbOrderPDFColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DTBOrderPDF struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := DTBOrderPDFS().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, dtbOrderPDFDBTypes, true, dtbOrderPDFPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize DTBOrderPDF struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testDTBOrderPDFSSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(dtbOrderPDFColumns) == len(dtbOrderPDFPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &DTBOrderPDF{}
	if err = randomize.Struct(seed, o, dtbOrderPDFDBTypes, true, dtbOrderPDFColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DTBOrderPDF struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := DTBOrderPDFS().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, dtbOrderPDFDBTypes, true, dtbOrderPDFPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize DTBOrderPDF struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(dtbOrderPDFColumns, dtbOrderPDFPrimaryKeyColumns) {
		fields = dtbOrderPDFColumns
	} else {
		fields = strmangle.SetComplement(
			dtbOrderPDFColumns,
			dtbOrderPDFPrimaryKeyColumns,
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

	slice := DTBOrderPDFSlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testDTBOrderPDFSUpsert(t *testing.T) {
	t.Parallel()

	if len(dtbOrderPDFColumns) == len(dtbOrderPDFPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}
	if len(mySQLDTBOrderPDFUniqueColumns) == 0 {
		t.Skip("Skipping table with no unique columns to conflict on")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := DTBOrderPDF{}
	if err = randomize.Struct(seed, &o, dtbOrderPDFDBTypes, false); err != nil {
		t.Errorf("Unable to randomize DTBOrderPDF struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(ctx, tx, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert DTBOrderPDF: %s", err)
	}

	count, err := DTBOrderPDFS().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, dtbOrderPDFDBTypes, false, dtbOrderPDFPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize DTBOrderPDF struct: %s", err)
	}

	if err = o.Upsert(ctx, tx, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert DTBOrderPDF: %s", err)
	}

	count, err = DTBOrderPDFS().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
