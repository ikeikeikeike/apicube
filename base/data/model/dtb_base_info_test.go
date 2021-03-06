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

func testDTBBaseInfos(t *testing.T) {
	t.Parallel()

	query := DTBBaseInfos()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testDTBBaseInfosDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &DTBBaseInfo{}
	if err = randomize.Struct(seed, o, dtbBaseInfoDBTypes, true, dtbBaseInfoColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DTBBaseInfo struct: %s", err)
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

	count, err := DTBBaseInfos().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testDTBBaseInfosQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &DTBBaseInfo{}
	if err = randomize.Struct(seed, o, dtbBaseInfoDBTypes, true, dtbBaseInfoColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DTBBaseInfo struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := DTBBaseInfos().DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := DTBBaseInfos().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testDTBBaseInfosSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &DTBBaseInfo{}
	if err = randomize.Struct(seed, o, dtbBaseInfoDBTypes, true, dtbBaseInfoColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DTBBaseInfo struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := DTBBaseInfoSlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := DTBBaseInfos().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testDTBBaseInfosExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &DTBBaseInfo{}
	if err = randomize.Struct(seed, o, dtbBaseInfoDBTypes, true, dtbBaseInfoColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DTBBaseInfo struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := DTBBaseInfoExists(ctx, tx, o.ID)
	if err != nil {
		t.Errorf("Unable to check if DTBBaseInfo exists: %s", err)
	}
	if !e {
		t.Errorf("Expected DTBBaseInfoExists to return true, but got false.")
	}
}

func testDTBBaseInfosFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &DTBBaseInfo{}
	if err = randomize.Struct(seed, o, dtbBaseInfoDBTypes, true, dtbBaseInfoColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DTBBaseInfo struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	dtbBaseInfoFound, err := FindDTBBaseInfo(ctx, tx, o.ID)
	if err != nil {
		t.Error(err)
	}

	if dtbBaseInfoFound == nil {
		t.Error("want a record, got nil")
	}
}

func testDTBBaseInfosBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &DTBBaseInfo{}
	if err = randomize.Struct(seed, o, dtbBaseInfoDBTypes, true, dtbBaseInfoColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DTBBaseInfo struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = DTBBaseInfos().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testDTBBaseInfosOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &DTBBaseInfo{}
	if err = randomize.Struct(seed, o, dtbBaseInfoDBTypes, true, dtbBaseInfoColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DTBBaseInfo struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := DTBBaseInfos().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testDTBBaseInfosAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	dtbBaseInfoOne := &DTBBaseInfo{}
	dtbBaseInfoTwo := &DTBBaseInfo{}
	if err = randomize.Struct(seed, dtbBaseInfoOne, dtbBaseInfoDBTypes, false, dtbBaseInfoColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DTBBaseInfo struct: %s", err)
	}
	if err = randomize.Struct(seed, dtbBaseInfoTwo, dtbBaseInfoDBTypes, false, dtbBaseInfoColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DTBBaseInfo struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = dtbBaseInfoOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = dtbBaseInfoTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := DTBBaseInfos().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testDTBBaseInfosCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	dtbBaseInfoOne := &DTBBaseInfo{}
	dtbBaseInfoTwo := &DTBBaseInfo{}
	if err = randomize.Struct(seed, dtbBaseInfoOne, dtbBaseInfoDBTypes, false, dtbBaseInfoColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DTBBaseInfo struct: %s", err)
	}
	if err = randomize.Struct(seed, dtbBaseInfoTwo, dtbBaseInfoDBTypes, false, dtbBaseInfoColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DTBBaseInfo struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = dtbBaseInfoOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = dtbBaseInfoTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := DTBBaseInfos().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func dtbBaseInfoBeforeInsertHook(ctx context.Context, e boil.ContextExecutor, o *DTBBaseInfo) error {
	*o = DTBBaseInfo{}
	return nil
}

func dtbBaseInfoAfterInsertHook(ctx context.Context, e boil.ContextExecutor, o *DTBBaseInfo) error {
	*o = DTBBaseInfo{}
	return nil
}

func dtbBaseInfoAfterSelectHook(ctx context.Context, e boil.ContextExecutor, o *DTBBaseInfo) error {
	*o = DTBBaseInfo{}
	return nil
}

func dtbBaseInfoBeforeUpdateHook(ctx context.Context, e boil.ContextExecutor, o *DTBBaseInfo) error {
	*o = DTBBaseInfo{}
	return nil
}

func dtbBaseInfoAfterUpdateHook(ctx context.Context, e boil.ContextExecutor, o *DTBBaseInfo) error {
	*o = DTBBaseInfo{}
	return nil
}

func dtbBaseInfoBeforeDeleteHook(ctx context.Context, e boil.ContextExecutor, o *DTBBaseInfo) error {
	*o = DTBBaseInfo{}
	return nil
}

func dtbBaseInfoAfterDeleteHook(ctx context.Context, e boil.ContextExecutor, o *DTBBaseInfo) error {
	*o = DTBBaseInfo{}
	return nil
}

func dtbBaseInfoBeforeUpsertHook(ctx context.Context, e boil.ContextExecutor, o *DTBBaseInfo) error {
	*o = DTBBaseInfo{}
	return nil
}

func dtbBaseInfoAfterUpsertHook(ctx context.Context, e boil.ContextExecutor, o *DTBBaseInfo) error {
	*o = DTBBaseInfo{}
	return nil
}

func testDTBBaseInfosHooks(t *testing.T) {
	t.Parallel()

	var err error

	ctx := context.Background()
	empty := &DTBBaseInfo{}
	o := &DTBBaseInfo{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, dtbBaseInfoDBTypes, false); err != nil {
		t.Errorf("Unable to randomize DTBBaseInfo object: %s", err)
	}

	AddDTBBaseInfoHook(boil.BeforeInsertHook, dtbBaseInfoBeforeInsertHook)
	if err = o.doBeforeInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	dtbBaseInfoBeforeInsertHooks = []DTBBaseInfoHook{}

	AddDTBBaseInfoHook(boil.AfterInsertHook, dtbBaseInfoAfterInsertHook)
	if err = o.doAfterInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	dtbBaseInfoAfterInsertHooks = []DTBBaseInfoHook{}

	AddDTBBaseInfoHook(boil.AfterSelectHook, dtbBaseInfoAfterSelectHook)
	if err = o.doAfterSelectHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	dtbBaseInfoAfterSelectHooks = []DTBBaseInfoHook{}

	AddDTBBaseInfoHook(boil.BeforeUpdateHook, dtbBaseInfoBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	dtbBaseInfoBeforeUpdateHooks = []DTBBaseInfoHook{}

	AddDTBBaseInfoHook(boil.AfterUpdateHook, dtbBaseInfoAfterUpdateHook)
	if err = o.doAfterUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	dtbBaseInfoAfterUpdateHooks = []DTBBaseInfoHook{}

	AddDTBBaseInfoHook(boil.BeforeDeleteHook, dtbBaseInfoBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	dtbBaseInfoBeforeDeleteHooks = []DTBBaseInfoHook{}

	AddDTBBaseInfoHook(boil.AfterDeleteHook, dtbBaseInfoAfterDeleteHook)
	if err = o.doAfterDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	dtbBaseInfoAfterDeleteHooks = []DTBBaseInfoHook{}

	AddDTBBaseInfoHook(boil.BeforeUpsertHook, dtbBaseInfoBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	dtbBaseInfoBeforeUpsertHooks = []DTBBaseInfoHook{}

	AddDTBBaseInfoHook(boil.AfterUpsertHook, dtbBaseInfoAfterUpsertHook)
	if err = o.doAfterUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	dtbBaseInfoAfterUpsertHooks = []DTBBaseInfoHook{}
}

func testDTBBaseInfosInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &DTBBaseInfo{}
	if err = randomize.Struct(seed, o, dtbBaseInfoDBTypes, true, dtbBaseInfoColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DTBBaseInfo struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := DTBBaseInfos().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testDTBBaseInfosInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &DTBBaseInfo{}
	if err = randomize.Struct(seed, o, dtbBaseInfoDBTypes, true); err != nil {
		t.Errorf("Unable to randomize DTBBaseInfo struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(dtbBaseInfoColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := DTBBaseInfos().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testDTBBaseInfoToOneMTBPrefUsingPref(t *testing.T) {
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var local DTBBaseInfo
	var foreign MTBPref

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, dtbBaseInfoDBTypes, true, dtbBaseInfoColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DTBBaseInfo struct: %s", err)
	}
	if err := randomize.Struct(seed, &foreign, mtbPrefDBTypes, false, mtbPrefColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize MTBPref struct: %s", err)
	}

	if err := foreign.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	queries.Assign(&local.PrefID, foreign.ID)
	if err := local.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := local.Pref().One(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	if !queries.Equal(check.ID, foreign.ID) {
		t.Errorf("want: %v, got %v", foreign.ID, check.ID)
	}

	slice := DTBBaseInfoSlice{&local}
	if err = local.L.LoadPref(ctx, tx, false, (*[]*DTBBaseInfo)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if local.R.Pref == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.Pref = nil
	if err = local.L.LoadPref(ctx, tx, true, &local, nil); err != nil {
		t.Fatal(err)
	}
	if local.R.Pref == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testDTBBaseInfoToOneMTBCountryUsingCountry(t *testing.T) {
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var local DTBBaseInfo
	var foreign MTBCountry

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, dtbBaseInfoDBTypes, true, dtbBaseInfoColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DTBBaseInfo struct: %s", err)
	}
	if err := randomize.Struct(seed, &foreign, mtbCountryDBTypes, false, mtbCountryColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize MTBCountry struct: %s", err)
	}

	if err := foreign.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	queries.Assign(&local.CountryID, foreign.ID)
	if err := local.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := local.Country().One(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	if !queries.Equal(check.ID, foreign.ID) {
		t.Errorf("want: %v, got %v", foreign.ID, check.ID)
	}

	slice := DTBBaseInfoSlice{&local}
	if err = local.L.LoadCountry(ctx, tx, false, (*[]*DTBBaseInfo)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if local.R.Country == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.Country = nil
	if err = local.L.LoadCountry(ctx, tx, true, &local, nil); err != nil {
		t.Fatal(err)
	}
	if local.R.Country == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testDTBBaseInfoToOneSetOpMTBPrefUsingPref(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a DTBBaseInfo
	var b, c MTBPref

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, dtbBaseInfoDBTypes, false, strmangle.SetComplement(dtbBaseInfoPrimaryKeyColumns, dtbBaseInfoColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, mtbPrefDBTypes, false, strmangle.SetComplement(mtbPrefPrimaryKeyColumns, mtbPrefColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, mtbPrefDBTypes, false, strmangle.SetComplement(mtbPrefPrimaryKeyColumns, mtbPrefColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*MTBPref{&b, &c} {
		err = a.SetPref(ctx, tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.Pref != x {
			t.Error("relationship struct not set to correct value")
		}

		if x.R.PrefDTBBaseInfos[0] != &a {
			t.Error("failed to append to foreign relationship struct")
		}
		if !queries.Equal(a.PrefID, x.ID) {
			t.Error("foreign key was wrong value", a.PrefID)
		}

		zero := reflect.Zero(reflect.TypeOf(a.PrefID))
		reflect.Indirect(reflect.ValueOf(&a.PrefID)).Set(zero)

		if err = a.Reload(ctx, tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if !queries.Equal(a.PrefID, x.ID) {
			t.Error("foreign key was wrong value", a.PrefID, x.ID)
		}
	}
}

func testDTBBaseInfoToOneRemoveOpMTBPrefUsingPref(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a DTBBaseInfo
	var b MTBPref

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, dtbBaseInfoDBTypes, false, strmangle.SetComplement(dtbBaseInfoPrimaryKeyColumns, dtbBaseInfoColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, mtbPrefDBTypes, false, strmangle.SetComplement(mtbPrefPrimaryKeyColumns, mtbPrefColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err = a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	if err = a.SetPref(ctx, tx, true, &b); err != nil {
		t.Fatal(err)
	}

	if err = a.RemovePref(ctx, tx, &b); err != nil {
		t.Error("failed to remove relationship")
	}

	count, err := a.Pref().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 0 {
		t.Error("want no relationships remaining")
	}

	if a.R.Pref != nil {
		t.Error("R struct entry should be nil")
	}

	if !queries.IsValuerNil(a.PrefID) {
		t.Error("foreign key value should be nil")
	}

	if len(b.R.PrefDTBBaseInfos) != 0 {
		t.Error("failed to remove a from b's relationships")
	}
}

func testDTBBaseInfoToOneSetOpMTBCountryUsingCountry(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a DTBBaseInfo
	var b, c MTBCountry

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, dtbBaseInfoDBTypes, false, strmangle.SetComplement(dtbBaseInfoPrimaryKeyColumns, dtbBaseInfoColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, mtbCountryDBTypes, false, strmangle.SetComplement(mtbCountryPrimaryKeyColumns, mtbCountryColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, mtbCountryDBTypes, false, strmangle.SetComplement(mtbCountryPrimaryKeyColumns, mtbCountryColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*MTBCountry{&b, &c} {
		err = a.SetCountry(ctx, tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.Country != x {
			t.Error("relationship struct not set to correct value")
		}

		if x.R.CountryDTBBaseInfos[0] != &a {
			t.Error("failed to append to foreign relationship struct")
		}
		if !queries.Equal(a.CountryID, x.ID) {
			t.Error("foreign key was wrong value", a.CountryID)
		}

		zero := reflect.Zero(reflect.TypeOf(a.CountryID))
		reflect.Indirect(reflect.ValueOf(&a.CountryID)).Set(zero)

		if err = a.Reload(ctx, tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if !queries.Equal(a.CountryID, x.ID) {
			t.Error("foreign key was wrong value", a.CountryID, x.ID)
		}
	}
}

func testDTBBaseInfoToOneRemoveOpMTBCountryUsingCountry(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a DTBBaseInfo
	var b MTBCountry

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, dtbBaseInfoDBTypes, false, strmangle.SetComplement(dtbBaseInfoPrimaryKeyColumns, dtbBaseInfoColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, mtbCountryDBTypes, false, strmangle.SetComplement(mtbCountryPrimaryKeyColumns, mtbCountryColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err = a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	if err = a.SetCountry(ctx, tx, true, &b); err != nil {
		t.Fatal(err)
	}

	if err = a.RemoveCountry(ctx, tx, &b); err != nil {
		t.Error("failed to remove relationship")
	}

	count, err := a.Country().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 0 {
		t.Error("want no relationships remaining")
	}

	if a.R.Country != nil {
		t.Error("R struct entry should be nil")
	}

	if !queries.IsValuerNil(a.CountryID) {
		t.Error("foreign key value should be nil")
	}

	if len(b.R.CountryDTBBaseInfos) != 0 {
		t.Error("failed to remove a from b's relationships")
	}
}

func testDTBBaseInfosReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &DTBBaseInfo{}
	if err = randomize.Struct(seed, o, dtbBaseInfoDBTypes, true, dtbBaseInfoColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DTBBaseInfo struct: %s", err)
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

func testDTBBaseInfosReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &DTBBaseInfo{}
	if err = randomize.Struct(seed, o, dtbBaseInfoDBTypes, true, dtbBaseInfoColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DTBBaseInfo struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := DTBBaseInfoSlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testDTBBaseInfosSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &DTBBaseInfo{}
	if err = randomize.Struct(seed, o, dtbBaseInfoDBTypes, true, dtbBaseInfoColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DTBBaseInfo struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := DTBBaseInfos().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	dtbBaseInfoDBTypes = map[string]string{`ID`: `int`, `CountryID`: `smallint`, `PrefID`: `smallint`, `CompanyName`: `varchar`, `CompanyKana`: `varchar`, `PostalCode`: `varchar`, `Addr01`: `varchar`, `Addr02`: `varchar`, `PhoneNumber`: `varchar`, `BusinessHour`: `varchar`, `Email01`: `varchar`, `Email02`: `varchar`, `Email03`: `varchar`, `Email04`: `varchar`, `ShopName`: `varchar`, `ShopKana`: `varchar`, `ShopNameEng`: `varchar`, `UpdateDate`: `datetime`, `GoodTraded`: `varchar`, `Message`: `varchar`, `DeliveryFreeAmount`: `decimal`, `DeliveryFreeQuantity`: `int`, `OptionMypageOrderStatusDisplay`: `tinyint`, `OptionNostockHidden`: `tinyint`, `OptionFavoriteProduct`: `tinyint`, `OptionProductDeliveryFee`: `tinyint`, `OptionProductTaxRule`: `tinyint`, `OptionCustomerActivate`: `tinyint`, `OptionRememberMe`: `tinyint`, `AuthenticationKey`: `varchar`, `PHPPath`: `varchar`, `OptionPoint`: `tinyint`, `BasicPointRate`: `decimal`, `PointConversionRate`: `decimal`, `DiscriminatorType`: `varchar`}
	_                  = bytes.MinRead
)

func testDTBBaseInfosUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(dtbBaseInfoPrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(dtbBaseInfoColumns) == len(dtbBaseInfoPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &DTBBaseInfo{}
	if err = randomize.Struct(seed, o, dtbBaseInfoDBTypes, true, dtbBaseInfoColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DTBBaseInfo struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := DTBBaseInfos().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, dtbBaseInfoDBTypes, true, dtbBaseInfoPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize DTBBaseInfo struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testDTBBaseInfosSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(dtbBaseInfoColumns) == len(dtbBaseInfoPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &DTBBaseInfo{}
	if err = randomize.Struct(seed, o, dtbBaseInfoDBTypes, true, dtbBaseInfoColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DTBBaseInfo struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := DTBBaseInfos().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, dtbBaseInfoDBTypes, true, dtbBaseInfoPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize DTBBaseInfo struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(dtbBaseInfoColumns, dtbBaseInfoPrimaryKeyColumns) {
		fields = dtbBaseInfoColumns
	} else {
		fields = strmangle.SetComplement(
			dtbBaseInfoColumns,
			dtbBaseInfoPrimaryKeyColumns,
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

	slice := DTBBaseInfoSlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testDTBBaseInfosUpsert(t *testing.T) {
	t.Parallel()

	if len(dtbBaseInfoColumns) == len(dtbBaseInfoPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}
	if len(mySQLDTBBaseInfoUniqueColumns) == 0 {
		t.Skip("Skipping table with no unique columns to conflict on")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := DTBBaseInfo{}
	if err = randomize.Struct(seed, &o, dtbBaseInfoDBTypes, false); err != nil {
		t.Errorf("Unable to randomize DTBBaseInfo struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(ctx, tx, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert DTBBaseInfo: %s", err)
	}

	count, err := DTBBaseInfos().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, dtbBaseInfoDBTypes, false, dtbBaseInfoPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize DTBBaseInfo struct: %s", err)
	}

	if err = o.Upsert(ctx, tx, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert DTBBaseInfo: %s", err)
	}

	count, err = DTBBaseInfos().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
