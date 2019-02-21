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

func testMTBCustomerStatuses(t *testing.T) {
	t.Parallel()

	query := MTBCustomerStatuses()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testMTBCustomerStatusesDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &MTBCustomerStatus{}
	if err = randomize.Struct(seed, o, mtbCustomerStatusDBTypes, true, mtbCustomerStatusColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize MTBCustomerStatus struct: %s", err)
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

	count, err := MTBCustomerStatuses().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testMTBCustomerStatusesQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &MTBCustomerStatus{}
	if err = randomize.Struct(seed, o, mtbCustomerStatusDBTypes, true, mtbCustomerStatusColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize MTBCustomerStatus struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := MTBCustomerStatuses().DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := MTBCustomerStatuses().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testMTBCustomerStatusesSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &MTBCustomerStatus{}
	if err = randomize.Struct(seed, o, mtbCustomerStatusDBTypes, true, mtbCustomerStatusColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize MTBCustomerStatus struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := MTBCustomerStatusSlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := MTBCustomerStatuses().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testMTBCustomerStatusesExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &MTBCustomerStatus{}
	if err = randomize.Struct(seed, o, mtbCustomerStatusDBTypes, true, mtbCustomerStatusColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize MTBCustomerStatus struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := MTBCustomerStatusExists(ctx, tx, o.ID)
	if err != nil {
		t.Errorf("Unable to check if MTBCustomerStatus exists: %s", err)
	}
	if !e {
		t.Errorf("Expected MTBCustomerStatusExists to return true, but got false.")
	}
}

func testMTBCustomerStatusesFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &MTBCustomerStatus{}
	if err = randomize.Struct(seed, o, mtbCustomerStatusDBTypes, true, mtbCustomerStatusColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize MTBCustomerStatus struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	mtbCustomerStatusFound, err := FindMTBCustomerStatus(ctx, tx, o.ID)
	if err != nil {
		t.Error(err)
	}

	if mtbCustomerStatusFound == nil {
		t.Error("want a record, got nil")
	}
}

func testMTBCustomerStatusesBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &MTBCustomerStatus{}
	if err = randomize.Struct(seed, o, mtbCustomerStatusDBTypes, true, mtbCustomerStatusColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize MTBCustomerStatus struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = MTBCustomerStatuses().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testMTBCustomerStatusesOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &MTBCustomerStatus{}
	if err = randomize.Struct(seed, o, mtbCustomerStatusDBTypes, true, mtbCustomerStatusColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize MTBCustomerStatus struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := MTBCustomerStatuses().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testMTBCustomerStatusesAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	mtbCustomerStatusOne := &MTBCustomerStatus{}
	mtbCustomerStatusTwo := &MTBCustomerStatus{}
	if err = randomize.Struct(seed, mtbCustomerStatusOne, mtbCustomerStatusDBTypes, false, mtbCustomerStatusColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize MTBCustomerStatus struct: %s", err)
	}
	if err = randomize.Struct(seed, mtbCustomerStatusTwo, mtbCustomerStatusDBTypes, false, mtbCustomerStatusColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize MTBCustomerStatus struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = mtbCustomerStatusOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = mtbCustomerStatusTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := MTBCustomerStatuses().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testMTBCustomerStatusesCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	mtbCustomerStatusOne := &MTBCustomerStatus{}
	mtbCustomerStatusTwo := &MTBCustomerStatus{}
	if err = randomize.Struct(seed, mtbCustomerStatusOne, mtbCustomerStatusDBTypes, false, mtbCustomerStatusColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize MTBCustomerStatus struct: %s", err)
	}
	if err = randomize.Struct(seed, mtbCustomerStatusTwo, mtbCustomerStatusDBTypes, false, mtbCustomerStatusColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize MTBCustomerStatus struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = mtbCustomerStatusOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = mtbCustomerStatusTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := MTBCustomerStatuses().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func mtbCustomerStatusBeforeInsertHook(ctx context.Context, e boil.ContextExecutor, o *MTBCustomerStatus) error {
	*o = MTBCustomerStatus{}
	return nil
}

func mtbCustomerStatusAfterInsertHook(ctx context.Context, e boil.ContextExecutor, o *MTBCustomerStatus) error {
	*o = MTBCustomerStatus{}
	return nil
}

func mtbCustomerStatusAfterSelectHook(ctx context.Context, e boil.ContextExecutor, o *MTBCustomerStatus) error {
	*o = MTBCustomerStatus{}
	return nil
}

func mtbCustomerStatusBeforeUpdateHook(ctx context.Context, e boil.ContextExecutor, o *MTBCustomerStatus) error {
	*o = MTBCustomerStatus{}
	return nil
}

func mtbCustomerStatusAfterUpdateHook(ctx context.Context, e boil.ContextExecutor, o *MTBCustomerStatus) error {
	*o = MTBCustomerStatus{}
	return nil
}

func mtbCustomerStatusBeforeDeleteHook(ctx context.Context, e boil.ContextExecutor, o *MTBCustomerStatus) error {
	*o = MTBCustomerStatus{}
	return nil
}

func mtbCustomerStatusAfterDeleteHook(ctx context.Context, e boil.ContextExecutor, o *MTBCustomerStatus) error {
	*o = MTBCustomerStatus{}
	return nil
}

func mtbCustomerStatusBeforeUpsertHook(ctx context.Context, e boil.ContextExecutor, o *MTBCustomerStatus) error {
	*o = MTBCustomerStatus{}
	return nil
}

func mtbCustomerStatusAfterUpsertHook(ctx context.Context, e boil.ContextExecutor, o *MTBCustomerStatus) error {
	*o = MTBCustomerStatus{}
	return nil
}

func testMTBCustomerStatusesHooks(t *testing.T) {
	t.Parallel()

	var err error

	ctx := context.Background()
	empty := &MTBCustomerStatus{}
	o := &MTBCustomerStatus{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, mtbCustomerStatusDBTypes, false); err != nil {
		t.Errorf("Unable to randomize MTBCustomerStatus object: %s", err)
	}

	AddMTBCustomerStatusHook(boil.BeforeInsertHook, mtbCustomerStatusBeforeInsertHook)
	if err = o.doBeforeInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	mtbCustomerStatusBeforeInsertHooks = []MTBCustomerStatusHook{}

	AddMTBCustomerStatusHook(boil.AfterInsertHook, mtbCustomerStatusAfterInsertHook)
	if err = o.doAfterInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	mtbCustomerStatusAfterInsertHooks = []MTBCustomerStatusHook{}

	AddMTBCustomerStatusHook(boil.AfterSelectHook, mtbCustomerStatusAfterSelectHook)
	if err = o.doAfterSelectHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	mtbCustomerStatusAfterSelectHooks = []MTBCustomerStatusHook{}

	AddMTBCustomerStatusHook(boil.BeforeUpdateHook, mtbCustomerStatusBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	mtbCustomerStatusBeforeUpdateHooks = []MTBCustomerStatusHook{}

	AddMTBCustomerStatusHook(boil.AfterUpdateHook, mtbCustomerStatusAfterUpdateHook)
	if err = o.doAfterUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	mtbCustomerStatusAfterUpdateHooks = []MTBCustomerStatusHook{}

	AddMTBCustomerStatusHook(boil.BeforeDeleteHook, mtbCustomerStatusBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	mtbCustomerStatusBeforeDeleteHooks = []MTBCustomerStatusHook{}

	AddMTBCustomerStatusHook(boil.AfterDeleteHook, mtbCustomerStatusAfterDeleteHook)
	if err = o.doAfterDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	mtbCustomerStatusAfterDeleteHooks = []MTBCustomerStatusHook{}

	AddMTBCustomerStatusHook(boil.BeforeUpsertHook, mtbCustomerStatusBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	mtbCustomerStatusBeforeUpsertHooks = []MTBCustomerStatusHook{}

	AddMTBCustomerStatusHook(boil.AfterUpsertHook, mtbCustomerStatusAfterUpsertHook)
	if err = o.doAfterUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	mtbCustomerStatusAfterUpsertHooks = []MTBCustomerStatusHook{}
}

func testMTBCustomerStatusesInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &MTBCustomerStatus{}
	if err = randomize.Struct(seed, o, mtbCustomerStatusDBTypes, true, mtbCustomerStatusColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize MTBCustomerStatus struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := MTBCustomerStatuses().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testMTBCustomerStatusesInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &MTBCustomerStatus{}
	if err = randomize.Struct(seed, o, mtbCustomerStatusDBTypes, true); err != nil {
		t.Errorf("Unable to randomize MTBCustomerStatus struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(mtbCustomerStatusColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := MTBCustomerStatuses().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testMTBCustomerStatusToManyCustomerStatusDTBCustomers(t *testing.T) {
	var err error
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a MTBCustomerStatus
	var b, c DTBCustomer

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, mtbCustomerStatusDBTypes, true, mtbCustomerStatusColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize MTBCustomerStatus struct: %s", err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	if err = randomize.Struct(seed, &b, dtbCustomerDBTypes, false, dtbCustomerColumnsWithDefault...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, dtbCustomerDBTypes, false, dtbCustomerColumnsWithDefault...); err != nil {
		t.Fatal(err)
	}

	queries.Assign(&b.CustomerStatusID, a.ID)
	queries.Assign(&c.CustomerStatusID, a.ID)
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := a.CustomerStatusDTBCustomers().All(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	bFound, cFound := false, false
	for _, v := range check {
		if queries.Equal(v.CustomerStatusID, b.CustomerStatusID) {
			bFound = true
		}
		if queries.Equal(v.CustomerStatusID, c.CustomerStatusID) {
			cFound = true
		}
	}

	if !bFound {
		t.Error("expected to find b")
	}
	if !cFound {
		t.Error("expected to find c")
	}

	slice := MTBCustomerStatusSlice{&a}
	if err = a.L.LoadCustomerStatusDTBCustomers(ctx, tx, false, (*[]*MTBCustomerStatus)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.CustomerStatusDTBCustomers); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	a.R.CustomerStatusDTBCustomers = nil
	if err = a.L.LoadCustomerStatusDTBCustomers(ctx, tx, true, &a, nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.CustomerStatusDTBCustomers); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	if t.Failed() {
		t.Logf("%#v", check)
	}
}

func testMTBCustomerStatusToManyAddOpCustomerStatusDTBCustomers(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a MTBCustomerStatus
	var b, c, d, e DTBCustomer

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, mtbCustomerStatusDBTypes, false, strmangle.SetComplement(mtbCustomerStatusPrimaryKeyColumns, mtbCustomerStatusColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*DTBCustomer{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, dtbCustomerDBTypes, false, strmangle.SetComplement(dtbCustomerPrimaryKeyColumns, dtbCustomerColumnsWithoutDefault)...); err != nil {
			t.Fatal(err)
		}
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	foreignersSplitByInsertion := [][]*DTBCustomer{
		{&b, &c},
		{&d, &e},
	}

	for i, x := range foreignersSplitByInsertion {
		err = a.AddCustomerStatusDTBCustomers(ctx, tx, i != 0, x...)
		if err != nil {
			t.Fatal(err)
		}

		first := x[0]
		second := x[1]

		if !queries.Equal(a.ID, first.CustomerStatusID) {
			t.Error("foreign key was wrong value", a.ID, first.CustomerStatusID)
		}
		if !queries.Equal(a.ID, second.CustomerStatusID) {
			t.Error("foreign key was wrong value", a.ID, second.CustomerStatusID)
		}

		if first.R.CustomerStatus != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}
		if second.R.CustomerStatus != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}

		if a.R.CustomerStatusDTBCustomers[i*2] != first {
			t.Error("relationship struct slice not set to correct value")
		}
		if a.R.CustomerStatusDTBCustomers[i*2+1] != second {
			t.Error("relationship struct slice not set to correct value")
		}

		count, err := a.CustomerStatusDTBCustomers().Count(ctx, tx)
		if err != nil {
			t.Fatal(err)
		}
		if want := int64((i + 1) * 2); count != want {
			t.Error("want", want, "got", count)
		}
	}
}

func testMTBCustomerStatusToManySetOpCustomerStatusDTBCustomers(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a MTBCustomerStatus
	var b, c, d, e DTBCustomer

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, mtbCustomerStatusDBTypes, false, strmangle.SetComplement(mtbCustomerStatusPrimaryKeyColumns, mtbCustomerStatusColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*DTBCustomer{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, dtbCustomerDBTypes, false, strmangle.SetComplement(dtbCustomerPrimaryKeyColumns, dtbCustomerColumnsWithoutDefault)...); err != nil {
			t.Fatal(err)
		}
	}

	if err = a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	err = a.SetCustomerStatusDTBCustomers(ctx, tx, false, &b, &c)
	if err != nil {
		t.Fatal(err)
	}

	count, err := a.CustomerStatusDTBCustomers().Count(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}
	if count != 2 {
		t.Error("count was wrong:", count)
	}

	err = a.SetCustomerStatusDTBCustomers(ctx, tx, true, &d, &e)
	if err != nil {
		t.Fatal(err)
	}

	count, err = a.CustomerStatusDTBCustomers().Count(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}
	if count != 2 {
		t.Error("count was wrong:", count)
	}

	if !queries.IsValuerNil(b.CustomerStatusID) {
		t.Error("want b's foreign key value to be nil")
	}
	if !queries.IsValuerNil(c.CustomerStatusID) {
		t.Error("want c's foreign key value to be nil")
	}
	if !queries.Equal(a.ID, d.CustomerStatusID) {
		t.Error("foreign key was wrong value", a.ID, d.CustomerStatusID)
	}
	if !queries.Equal(a.ID, e.CustomerStatusID) {
		t.Error("foreign key was wrong value", a.ID, e.CustomerStatusID)
	}

	if b.R.CustomerStatus != nil {
		t.Error("relationship was not removed properly from the foreign struct")
	}
	if c.R.CustomerStatus != nil {
		t.Error("relationship was not removed properly from the foreign struct")
	}
	if d.R.CustomerStatus != &a {
		t.Error("relationship was not added properly to the foreign struct")
	}
	if e.R.CustomerStatus != &a {
		t.Error("relationship was not added properly to the foreign struct")
	}

	if a.R.CustomerStatusDTBCustomers[0] != &d {
		t.Error("relationship struct slice not set to correct value")
	}
	if a.R.CustomerStatusDTBCustomers[1] != &e {
		t.Error("relationship struct slice not set to correct value")
	}
}

func testMTBCustomerStatusToManyRemoveOpCustomerStatusDTBCustomers(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a MTBCustomerStatus
	var b, c, d, e DTBCustomer

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, mtbCustomerStatusDBTypes, false, strmangle.SetComplement(mtbCustomerStatusPrimaryKeyColumns, mtbCustomerStatusColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*DTBCustomer{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, dtbCustomerDBTypes, false, strmangle.SetComplement(dtbCustomerPrimaryKeyColumns, dtbCustomerColumnsWithoutDefault)...); err != nil {
			t.Fatal(err)
		}
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	err = a.AddCustomerStatusDTBCustomers(ctx, tx, true, foreigners...)
	if err != nil {
		t.Fatal(err)
	}

	count, err := a.CustomerStatusDTBCustomers().Count(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}
	if count != 4 {
		t.Error("count was wrong:", count)
	}

	err = a.RemoveCustomerStatusDTBCustomers(ctx, tx, foreigners[:2]...)
	if err != nil {
		t.Fatal(err)
	}

	count, err = a.CustomerStatusDTBCustomers().Count(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}
	if count != 2 {
		t.Error("count was wrong:", count)
	}

	if !queries.IsValuerNil(b.CustomerStatusID) {
		t.Error("want b's foreign key value to be nil")
	}
	if !queries.IsValuerNil(c.CustomerStatusID) {
		t.Error("want c's foreign key value to be nil")
	}

	if b.R.CustomerStatus != nil {
		t.Error("relationship was not removed properly from the foreign struct")
	}
	if c.R.CustomerStatus != nil {
		t.Error("relationship was not removed properly from the foreign struct")
	}
	if d.R.CustomerStatus != &a {
		t.Error("relationship to a should have been preserved")
	}
	if e.R.CustomerStatus != &a {
		t.Error("relationship to a should have been preserved")
	}

	if len(a.R.CustomerStatusDTBCustomers) != 2 {
		t.Error("should have preserved two relationships")
	}

	// Removal doesn't do a stable deletion for performance so we have to flip the order
	if a.R.CustomerStatusDTBCustomers[1] != &d {
		t.Error("relationship to d should have been preserved")
	}
	if a.R.CustomerStatusDTBCustomers[0] != &e {
		t.Error("relationship to e should have been preserved")
	}
}

func testMTBCustomerStatusesReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &MTBCustomerStatus{}
	if err = randomize.Struct(seed, o, mtbCustomerStatusDBTypes, true, mtbCustomerStatusColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize MTBCustomerStatus struct: %s", err)
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

func testMTBCustomerStatusesReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &MTBCustomerStatus{}
	if err = randomize.Struct(seed, o, mtbCustomerStatusDBTypes, true, mtbCustomerStatusColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize MTBCustomerStatus struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := MTBCustomerStatusSlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testMTBCustomerStatusesSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &MTBCustomerStatus{}
	if err = randomize.Struct(seed, o, mtbCustomerStatusDBTypes, true, mtbCustomerStatusColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize MTBCustomerStatus struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := MTBCustomerStatuses().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	mtbCustomerStatusDBTypes = map[string]string{`ID`: `smallint`, `Name`: `varchar`, `SortNo`: `smallint`, `DiscriminatorType`: `varchar`}
	_                        = bytes.MinRead
)

func testMTBCustomerStatusesUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(mtbCustomerStatusPrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(mtbCustomerStatusColumns) == len(mtbCustomerStatusPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &MTBCustomerStatus{}
	if err = randomize.Struct(seed, o, mtbCustomerStatusDBTypes, true, mtbCustomerStatusColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize MTBCustomerStatus struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := MTBCustomerStatuses().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, mtbCustomerStatusDBTypes, true, mtbCustomerStatusPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize MTBCustomerStatus struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testMTBCustomerStatusesSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(mtbCustomerStatusColumns) == len(mtbCustomerStatusPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &MTBCustomerStatus{}
	if err = randomize.Struct(seed, o, mtbCustomerStatusDBTypes, true, mtbCustomerStatusColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize MTBCustomerStatus struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := MTBCustomerStatuses().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, mtbCustomerStatusDBTypes, true, mtbCustomerStatusPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize MTBCustomerStatus struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(mtbCustomerStatusColumns, mtbCustomerStatusPrimaryKeyColumns) {
		fields = mtbCustomerStatusColumns
	} else {
		fields = strmangle.SetComplement(
			mtbCustomerStatusColumns,
			mtbCustomerStatusPrimaryKeyColumns,
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

	slice := MTBCustomerStatusSlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testMTBCustomerStatusesUpsert(t *testing.T) {
	t.Parallel()

	if len(mtbCustomerStatusColumns) == len(mtbCustomerStatusPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}
	if len(mySQLMTBCustomerStatusUniqueColumns) == 0 {
		t.Skip("Skipping table with no unique columns to conflict on")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := MTBCustomerStatus{}
	if err = randomize.Struct(seed, &o, mtbCustomerStatusDBTypes, false); err != nil {
		t.Errorf("Unable to randomize MTBCustomerStatus struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(ctx, tx, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert MTBCustomerStatus: %s", err)
	}

	count, err := MTBCustomerStatuses().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, mtbCustomerStatusDBTypes, false, mtbCustomerStatusPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize MTBCustomerStatus struct: %s", err)
	}

	if err = o.Upsert(ctx, tx, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert MTBCustomerStatus: %s", err)
	}

	count, err = MTBCustomerStatuses().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}