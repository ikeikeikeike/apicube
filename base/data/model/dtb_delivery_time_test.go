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

func testDTBDeliveryTimes(t *testing.T) {
	t.Parallel()

	query := DTBDeliveryTimes()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testDTBDeliveryTimesDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &DTBDeliveryTime{}
	if err = randomize.Struct(seed, o, dtbDeliveryTimeDBTypes, true, dtbDeliveryTimeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DTBDeliveryTime struct: %s", err)
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

	count, err := DTBDeliveryTimes().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testDTBDeliveryTimesQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &DTBDeliveryTime{}
	if err = randomize.Struct(seed, o, dtbDeliveryTimeDBTypes, true, dtbDeliveryTimeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DTBDeliveryTime struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := DTBDeliveryTimes().DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := DTBDeliveryTimes().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testDTBDeliveryTimesSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &DTBDeliveryTime{}
	if err = randomize.Struct(seed, o, dtbDeliveryTimeDBTypes, true, dtbDeliveryTimeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DTBDeliveryTime struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := DTBDeliveryTimeSlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := DTBDeliveryTimes().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testDTBDeliveryTimesExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &DTBDeliveryTime{}
	if err = randomize.Struct(seed, o, dtbDeliveryTimeDBTypes, true, dtbDeliveryTimeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DTBDeliveryTime struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := DTBDeliveryTimeExists(ctx, tx, o.ID)
	if err != nil {
		t.Errorf("Unable to check if DTBDeliveryTime exists: %s", err)
	}
	if !e {
		t.Errorf("Expected DTBDeliveryTimeExists to return true, but got false.")
	}
}

func testDTBDeliveryTimesFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &DTBDeliveryTime{}
	if err = randomize.Struct(seed, o, dtbDeliveryTimeDBTypes, true, dtbDeliveryTimeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DTBDeliveryTime struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	dtbDeliveryTimeFound, err := FindDTBDeliveryTime(ctx, tx, o.ID)
	if err != nil {
		t.Error(err)
	}

	if dtbDeliveryTimeFound == nil {
		t.Error("want a record, got nil")
	}
}

func testDTBDeliveryTimesBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &DTBDeliveryTime{}
	if err = randomize.Struct(seed, o, dtbDeliveryTimeDBTypes, true, dtbDeliveryTimeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DTBDeliveryTime struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = DTBDeliveryTimes().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testDTBDeliveryTimesOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &DTBDeliveryTime{}
	if err = randomize.Struct(seed, o, dtbDeliveryTimeDBTypes, true, dtbDeliveryTimeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DTBDeliveryTime struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := DTBDeliveryTimes().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testDTBDeliveryTimesAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	dtbDeliveryTimeOne := &DTBDeliveryTime{}
	dtbDeliveryTimeTwo := &DTBDeliveryTime{}
	if err = randomize.Struct(seed, dtbDeliveryTimeOne, dtbDeliveryTimeDBTypes, false, dtbDeliveryTimeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DTBDeliveryTime struct: %s", err)
	}
	if err = randomize.Struct(seed, dtbDeliveryTimeTwo, dtbDeliveryTimeDBTypes, false, dtbDeliveryTimeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DTBDeliveryTime struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = dtbDeliveryTimeOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = dtbDeliveryTimeTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := DTBDeliveryTimes().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testDTBDeliveryTimesCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	dtbDeliveryTimeOne := &DTBDeliveryTime{}
	dtbDeliveryTimeTwo := &DTBDeliveryTime{}
	if err = randomize.Struct(seed, dtbDeliveryTimeOne, dtbDeliveryTimeDBTypes, false, dtbDeliveryTimeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DTBDeliveryTime struct: %s", err)
	}
	if err = randomize.Struct(seed, dtbDeliveryTimeTwo, dtbDeliveryTimeDBTypes, false, dtbDeliveryTimeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DTBDeliveryTime struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = dtbDeliveryTimeOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = dtbDeliveryTimeTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := DTBDeliveryTimes().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func dtbDeliveryTimeBeforeInsertHook(ctx context.Context, e boil.ContextExecutor, o *DTBDeliveryTime) error {
	*o = DTBDeliveryTime{}
	return nil
}

func dtbDeliveryTimeAfterInsertHook(ctx context.Context, e boil.ContextExecutor, o *DTBDeliveryTime) error {
	*o = DTBDeliveryTime{}
	return nil
}

func dtbDeliveryTimeAfterSelectHook(ctx context.Context, e boil.ContextExecutor, o *DTBDeliveryTime) error {
	*o = DTBDeliveryTime{}
	return nil
}

func dtbDeliveryTimeBeforeUpdateHook(ctx context.Context, e boil.ContextExecutor, o *DTBDeliveryTime) error {
	*o = DTBDeliveryTime{}
	return nil
}

func dtbDeliveryTimeAfterUpdateHook(ctx context.Context, e boil.ContextExecutor, o *DTBDeliveryTime) error {
	*o = DTBDeliveryTime{}
	return nil
}

func dtbDeliveryTimeBeforeDeleteHook(ctx context.Context, e boil.ContextExecutor, o *DTBDeliveryTime) error {
	*o = DTBDeliveryTime{}
	return nil
}

func dtbDeliveryTimeAfterDeleteHook(ctx context.Context, e boil.ContextExecutor, o *DTBDeliveryTime) error {
	*o = DTBDeliveryTime{}
	return nil
}

func dtbDeliveryTimeBeforeUpsertHook(ctx context.Context, e boil.ContextExecutor, o *DTBDeliveryTime) error {
	*o = DTBDeliveryTime{}
	return nil
}

func dtbDeliveryTimeAfterUpsertHook(ctx context.Context, e boil.ContextExecutor, o *DTBDeliveryTime) error {
	*o = DTBDeliveryTime{}
	return nil
}

func testDTBDeliveryTimesHooks(t *testing.T) {
	t.Parallel()

	var err error

	ctx := context.Background()
	empty := &DTBDeliveryTime{}
	o := &DTBDeliveryTime{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, dtbDeliveryTimeDBTypes, false); err != nil {
		t.Errorf("Unable to randomize DTBDeliveryTime object: %s", err)
	}

	AddDTBDeliveryTimeHook(boil.BeforeInsertHook, dtbDeliveryTimeBeforeInsertHook)
	if err = o.doBeforeInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	dtbDeliveryTimeBeforeInsertHooks = []DTBDeliveryTimeHook{}

	AddDTBDeliveryTimeHook(boil.AfterInsertHook, dtbDeliveryTimeAfterInsertHook)
	if err = o.doAfterInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	dtbDeliveryTimeAfterInsertHooks = []DTBDeliveryTimeHook{}

	AddDTBDeliveryTimeHook(boil.AfterSelectHook, dtbDeliveryTimeAfterSelectHook)
	if err = o.doAfterSelectHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	dtbDeliveryTimeAfterSelectHooks = []DTBDeliveryTimeHook{}

	AddDTBDeliveryTimeHook(boil.BeforeUpdateHook, dtbDeliveryTimeBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	dtbDeliveryTimeBeforeUpdateHooks = []DTBDeliveryTimeHook{}

	AddDTBDeliveryTimeHook(boil.AfterUpdateHook, dtbDeliveryTimeAfterUpdateHook)
	if err = o.doAfterUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	dtbDeliveryTimeAfterUpdateHooks = []DTBDeliveryTimeHook{}

	AddDTBDeliveryTimeHook(boil.BeforeDeleteHook, dtbDeliveryTimeBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	dtbDeliveryTimeBeforeDeleteHooks = []DTBDeliveryTimeHook{}

	AddDTBDeliveryTimeHook(boil.AfterDeleteHook, dtbDeliveryTimeAfterDeleteHook)
	if err = o.doAfterDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	dtbDeliveryTimeAfterDeleteHooks = []DTBDeliveryTimeHook{}

	AddDTBDeliveryTimeHook(boil.BeforeUpsertHook, dtbDeliveryTimeBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	dtbDeliveryTimeBeforeUpsertHooks = []DTBDeliveryTimeHook{}

	AddDTBDeliveryTimeHook(boil.AfterUpsertHook, dtbDeliveryTimeAfterUpsertHook)
	if err = o.doAfterUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	dtbDeliveryTimeAfterUpsertHooks = []DTBDeliveryTimeHook{}
}

func testDTBDeliveryTimesInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &DTBDeliveryTime{}
	if err = randomize.Struct(seed, o, dtbDeliveryTimeDBTypes, true, dtbDeliveryTimeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DTBDeliveryTime struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := DTBDeliveryTimes().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testDTBDeliveryTimesInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &DTBDeliveryTime{}
	if err = randomize.Struct(seed, o, dtbDeliveryTimeDBTypes, true); err != nil {
		t.Errorf("Unable to randomize DTBDeliveryTime struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(dtbDeliveryTimeColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := DTBDeliveryTimes().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testDTBDeliveryTimeToOneDTBDeliveryUsingDelivery(t *testing.T) {
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var local DTBDeliveryTime
	var foreign DTBDelivery

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, dtbDeliveryTimeDBTypes, true, dtbDeliveryTimeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DTBDeliveryTime struct: %s", err)
	}
	if err := randomize.Struct(seed, &foreign, dtbDeliveryDBTypes, false, dtbDeliveryColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DTBDelivery struct: %s", err)
	}

	if err := foreign.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	queries.Assign(&local.DeliveryID, foreign.ID)
	if err := local.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := local.Delivery().One(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	if !queries.Equal(check.ID, foreign.ID) {
		t.Errorf("want: %v, got %v", foreign.ID, check.ID)
	}

	slice := DTBDeliveryTimeSlice{&local}
	if err = local.L.LoadDelivery(ctx, tx, false, (*[]*DTBDeliveryTime)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if local.R.Delivery == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.Delivery = nil
	if err = local.L.LoadDelivery(ctx, tx, true, &local, nil); err != nil {
		t.Fatal(err)
	}
	if local.R.Delivery == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testDTBDeliveryTimeToOneSetOpDTBDeliveryUsingDelivery(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a DTBDeliveryTime
	var b, c DTBDelivery

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, dtbDeliveryTimeDBTypes, false, strmangle.SetComplement(dtbDeliveryTimePrimaryKeyColumns, dtbDeliveryTimeColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, dtbDeliveryDBTypes, false, strmangle.SetComplement(dtbDeliveryPrimaryKeyColumns, dtbDeliveryColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, dtbDeliveryDBTypes, false, strmangle.SetComplement(dtbDeliveryPrimaryKeyColumns, dtbDeliveryColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*DTBDelivery{&b, &c} {
		err = a.SetDelivery(ctx, tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.Delivery != x {
			t.Error("relationship struct not set to correct value")
		}

		if x.R.DeliveryDTBDeliveryTimes[0] != &a {
			t.Error("failed to append to foreign relationship struct")
		}
		if !queries.Equal(a.DeliveryID, x.ID) {
			t.Error("foreign key was wrong value", a.DeliveryID)
		}

		zero := reflect.Zero(reflect.TypeOf(a.DeliveryID))
		reflect.Indirect(reflect.ValueOf(&a.DeliveryID)).Set(zero)

		if err = a.Reload(ctx, tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if !queries.Equal(a.DeliveryID, x.ID) {
			t.Error("foreign key was wrong value", a.DeliveryID, x.ID)
		}
	}
}

func testDTBDeliveryTimeToOneRemoveOpDTBDeliveryUsingDelivery(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a DTBDeliveryTime
	var b DTBDelivery

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, dtbDeliveryTimeDBTypes, false, strmangle.SetComplement(dtbDeliveryTimePrimaryKeyColumns, dtbDeliveryTimeColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, dtbDeliveryDBTypes, false, strmangle.SetComplement(dtbDeliveryPrimaryKeyColumns, dtbDeliveryColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err = a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	if err = a.SetDelivery(ctx, tx, true, &b); err != nil {
		t.Fatal(err)
	}

	if err = a.RemoveDelivery(ctx, tx, &b); err != nil {
		t.Error("failed to remove relationship")
	}

	count, err := a.Delivery().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 0 {
		t.Error("want no relationships remaining")
	}

	if a.R.Delivery != nil {
		t.Error("R struct entry should be nil")
	}

	if !queries.IsValuerNil(a.DeliveryID) {
		t.Error("foreign key value should be nil")
	}

	if len(b.R.DeliveryDTBDeliveryTimes) != 0 {
		t.Error("failed to remove a from b's relationships")
	}
}

func testDTBDeliveryTimesReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &DTBDeliveryTime{}
	if err = randomize.Struct(seed, o, dtbDeliveryTimeDBTypes, true, dtbDeliveryTimeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DTBDeliveryTime struct: %s", err)
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

func testDTBDeliveryTimesReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &DTBDeliveryTime{}
	if err = randomize.Struct(seed, o, dtbDeliveryTimeDBTypes, true, dtbDeliveryTimeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DTBDeliveryTime struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := DTBDeliveryTimeSlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testDTBDeliveryTimesSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &DTBDeliveryTime{}
	if err = randomize.Struct(seed, o, dtbDeliveryTimeDBTypes, true, dtbDeliveryTimeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DTBDeliveryTime struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := DTBDeliveryTimes().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	dtbDeliveryTimeDBTypes = map[string]string{`ID`: `int`, `DeliveryID`: `int`, `DeliveryTime`: `varchar`, `SortNo`: `smallint`, `Visible`: `tinyint`, `CreateDate`: `datetime`, `UpdateDate`: `datetime`, `DiscriminatorType`: `varchar`}
	_                      = bytes.MinRead
)

func testDTBDeliveryTimesUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(dtbDeliveryTimePrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(dtbDeliveryTimeColumns) == len(dtbDeliveryTimePrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &DTBDeliveryTime{}
	if err = randomize.Struct(seed, o, dtbDeliveryTimeDBTypes, true, dtbDeliveryTimeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DTBDeliveryTime struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := DTBDeliveryTimes().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, dtbDeliveryTimeDBTypes, true, dtbDeliveryTimePrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize DTBDeliveryTime struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testDTBDeliveryTimesSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(dtbDeliveryTimeColumns) == len(dtbDeliveryTimePrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &DTBDeliveryTime{}
	if err = randomize.Struct(seed, o, dtbDeliveryTimeDBTypes, true, dtbDeliveryTimeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DTBDeliveryTime struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := DTBDeliveryTimes().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, dtbDeliveryTimeDBTypes, true, dtbDeliveryTimePrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize DTBDeliveryTime struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(dtbDeliveryTimeColumns, dtbDeliveryTimePrimaryKeyColumns) {
		fields = dtbDeliveryTimeColumns
	} else {
		fields = strmangle.SetComplement(
			dtbDeliveryTimeColumns,
			dtbDeliveryTimePrimaryKeyColumns,
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

	slice := DTBDeliveryTimeSlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testDTBDeliveryTimesUpsert(t *testing.T) {
	t.Parallel()

	if len(dtbDeliveryTimeColumns) == len(dtbDeliveryTimePrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}
	if len(mySQLDTBDeliveryTimeUniqueColumns) == 0 {
		t.Skip("Skipping table with no unique columns to conflict on")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := DTBDeliveryTime{}
	if err = randomize.Struct(seed, &o, dtbDeliveryTimeDBTypes, false); err != nil {
		t.Errorf("Unable to randomize DTBDeliveryTime struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(ctx, tx, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert DTBDeliveryTime: %s", err)
	}

	count, err := DTBDeliveryTimes().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, dtbDeliveryTimeDBTypes, false, dtbDeliveryTimePrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize DTBDeliveryTime struct: %s", err)
	}

	if err = o.Upsert(ctx, tx, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert DTBDeliveryTime: %s", err)
	}

	count, err = DTBDeliveryTimes().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
