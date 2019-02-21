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

func testMTBSaleTypes(t *testing.T) {
	t.Parallel()

	query := MTBSaleTypes()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testMTBSaleTypesDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &MTBSaleType{}
	if err = randomize.Struct(seed, o, mtbSaleTypeDBTypes, true, mtbSaleTypeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize MTBSaleType struct: %s", err)
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

	count, err := MTBSaleTypes().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testMTBSaleTypesQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &MTBSaleType{}
	if err = randomize.Struct(seed, o, mtbSaleTypeDBTypes, true, mtbSaleTypeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize MTBSaleType struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := MTBSaleTypes().DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := MTBSaleTypes().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testMTBSaleTypesSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &MTBSaleType{}
	if err = randomize.Struct(seed, o, mtbSaleTypeDBTypes, true, mtbSaleTypeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize MTBSaleType struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := MTBSaleTypeSlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := MTBSaleTypes().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testMTBSaleTypesExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &MTBSaleType{}
	if err = randomize.Struct(seed, o, mtbSaleTypeDBTypes, true, mtbSaleTypeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize MTBSaleType struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := MTBSaleTypeExists(ctx, tx, o.ID)
	if err != nil {
		t.Errorf("Unable to check if MTBSaleType exists: %s", err)
	}
	if !e {
		t.Errorf("Expected MTBSaleTypeExists to return true, but got false.")
	}
}

func testMTBSaleTypesFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &MTBSaleType{}
	if err = randomize.Struct(seed, o, mtbSaleTypeDBTypes, true, mtbSaleTypeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize MTBSaleType struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	mtbSaleTypeFound, err := FindMTBSaleType(ctx, tx, o.ID)
	if err != nil {
		t.Error(err)
	}

	if mtbSaleTypeFound == nil {
		t.Error("want a record, got nil")
	}
}

func testMTBSaleTypesBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &MTBSaleType{}
	if err = randomize.Struct(seed, o, mtbSaleTypeDBTypes, true, mtbSaleTypeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize MTBSaleType struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = MTBSaleTypes().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testMTBSaleTypesOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &MTBSaleType{}
	if err = randomize.Struct(seed, o, mtbSaleTypeDBTypes, true, mtbSaleTypeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize MTBSaleType struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := MTBSaleTypes().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testMTBSaleTypesAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	mtbSaleTypeOne := &MTBSaleType{}
	mtbSaleTypeTwo := &MTBSaleType{}
	if err = randomize.Struct(seed, mtbSaleTypeOne, mtbSaleTypeDBTypes, false, mtbSaleTypeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize MTBSaleType struct: %s", err)
	}
	if err = randomize.Struct(seed, mtbSaleTypeTwo, mtbSaleTypeDBTypes, false, mtbSaleTypeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize MTBSaleType struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = mtbSaleTypeOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = mtbSaleTypeTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := MTBSaleTypes().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testMTBSaleTypesCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	mtbSaleTypeOne := &MTBSaleType{}
	mtbSaleTypeTwo := &MTBSaleType{}
	if err = randomize.Struct(seed, mtbSaleTypeOne, mtbSaleTypeDBTypes, false, mtbSaleTypeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize MTBSaleType struct: %s", err)
	}
	if err = randomize.Struct(seed, mtbSaleTypeTwo, mtbSaleTypeDBTypes, false, mtbSaleTypeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize MTBSaleType struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = mtbSaleTypeOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = mtbSaleTypeTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := MTBSaleTypes().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func mtbSaleTypeBeforeInsertHook(ctx context.Context, e boil.ContextExecutor, o *MTBSaleType) error {
	*o = MTBSaleType{}
	return nil
}

func mtbSaleTypeAfterInsertHook(ctx context.Context, e boil.ContextExecutor, o *MTBSaleType) error {
	*o = MTBSaleType{}
	return nil
}

func mtbSaleTypeAfterSelectHook(ctx context.Context, e boil.ContextExecutor, o *MTBSaleType) error {
	*o = MTBSaleType{}
	return nil
}

func mtbSaleTypeBeforeUpdateHook(ctx context.Context, e boil.ContextExecutor, o *MTBSaleType) error {
	*o = MTBSaleType{}
	return nil
}

func mtbSaleTypeAfterUpdateHook(ctx context.Context, e boil.ContextExecutor, o *MTBSaleType) error {
	*o = MTBSaleType{}
	return nil
}

func mtbSaleTypeBeforeDeleteHook(ctx context.Context, e boil.ContextExecutor, o *MTBSaleType) error {
	*o = MTBSaleType{}
	return nil
}

func mtbSaleTypeAfterDeleteHook(ctx context.Context, e boil.ContextExecutor, o *MTBSaleType) error {
	*o = MTBSaleType{}
	return nil
}

func mtbSaleTypeBeforeUpsertHook(ctx context.Context, e boil.ContextExecutor, o *MTBSaleType) error {
	*o = MTBSaleType{}
	return nil
}

func mtbSaleTypeAfterUpsertHook(ctx context.Context, e boil.ContextExecutor, o *MTBSaleType) error {
	*o = MTBSaleType{}
	return nil
}

func testMTBSaleTypesHooks(t *testing.T) {
	t.Parallel()

	var err error

	ctx := context.Background()
	empty := &MTBSaleType{}
	o := &MTBSaleType{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, mtbSaleTypeDBTypes, false); err != nil {
		t.Errorf("Unable to randomize MTBSaleType object: %s", err)
	}

	AddMTBSaleTypeHook(boil.BeforeInsertHook, mtbSaleTypeBeforeInsertHook)
	if err = o.doBeforeInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	mtbSaleTypeBeforeInsertHooks = []MTBSaleTypeHook{}

	AddMTBSaleTypeHook(boil.AfterInsertHook, mtbSaleTypeAfterInsertHook)
	if err = o.doAfterInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	mtbSaleTypeAfterInsertHooks = []MTBSaleTypeHook{}

	AddMTBSaleTypeHook(boil.AfterSelectHook, mtbSaleTypeAfterSelectHook)
	if err = o.doAfterSelectHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	mtbSaleTypeAfterSelectHooks = []MTBSaleTypeHook{}

	AddMTBSaleTypeHook(boil.BeforeUpdateHook, mtbSaleTypeBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	mtbSaleTypeBeforeUpdateHooks = []MTBSaleTypeHook{}

	AddMTBSaleTypeHook(boil.AfterUpdateHook, mtbSaleTypeAfterUpdateHook)
	if err = o.doAfterUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	mtbSaleTypeAfterUpdateHooks = []MTBSaleTypeHook{}

	AddMTBSaleTypeHook(boil.BeforeDeleteHook, mtbSaleTypeBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	mtbSaleTypeBeforeDeleteHooks = []MTBSaleTypeHook{}

	AddMTBSaleTypeHook(boil.AfterDeleteHook, mtbSaleTypeAfterDeleteHook)
	if err = o.doAfterDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	mtbSaleTypeAfterDeleteHooks = []MTBSaleTypeHook{}

	AddMTBSaleTypeHook(boil.BeforeUpsertHook, mtbSaleTypeBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	mtbSaleTypeBeforeUpsertHooks = []MTBSaleTypeHook{}

	AddMTBSaleTypeHook(boil.AfterUpsertHook, mtbSaleTypeAfterUpsertHook)
	if err = o.doAfterUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	mtbSaleTypeAfterUpsertHooks = []MTBSaleTypeHook{}
}

func testMTBSaleTypesInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &MTBSaleType{}
	if err = randomize.Struct(seed, o, mtbSaleTypeDBTypes, true, mtbSaleTypeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize MTBSaleType struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := MTBSaleTypes().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testMTBSaleTypesInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &MTBSaleType{}
	if err = randomize.Struct(seed, o, mtbSaleTypeDBTypes, true); err != nil {
		t.Errorf("Unable to randomize MTBSaleType struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(mtbSaleTypeColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := MTBSaleTypes().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testMTBSaleTypeToManySaleTypeDTBDeliveries(t *testing.T) {
	var err error
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a MTBSaleType
	var b, c DTBDelivery

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, mtbSaleTypeDBTypes, true, mtbSaleTypeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize MTBSaleType struct: %s", err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	if err = randomize.Struct(seed, &b, dtbDeliveryDBTypes, false, dtbDeliveryColumnsWithDefault...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, dtbDeliveryDBTypes, false, dtbDeliveryColumnsWithDefault...); err != nil {
		t.Fatal(err)
	}

	queries.Assign(&b.SaleTypeID, a.ID)
	queries.Assign(&c.SaleTypeID, a.ID)
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := a.SaleTypeDTBDeliveries().All(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	bFound, cFound := false, false
	for _, v := range check {
		if queries.Equal(v.SaleTypeID, b.SaleTypeID) {
			bFound = true
		}
		if queries.Equal(v.SaleTypeID, c.SaleTypeID) {
			cFound = true
		}
	}

	if !bFound {
		t.Error("expected to find b")
	}
	if !cFound {
		t.Error("expected to find c")
	}

	slice := MTBSaleTypeSlice{&a}
	if err = a.L.LoadSaleTypeDTBDeliveries(ctx, tx, false, (*[]*MTBSaleType)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.SaleTypeDTBDeliveries); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	a.R.SaleTypeDTBDeliveries = nil
	if err = a.L.LoadSaleTypeDTBDeliveries(ctx, tx, true, &a, nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.SaleTypeDTBDeliveries); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	if t.Failed() {
		t.Logf("%#v", check)
	}
}

func testMTBSaleTypeToManyAddOpSaleTypeDTBDeliveries(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a MTBSaleType
	var b, c, d, e DTBDelivery

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, mtbSaleTypeDBTypes, false, strmangle.SetComplement(mtbSaleTypePrimaryKeyColumns, mtbSaleTypeColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*DTBDelivery{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, dtbDeliveryDBTypes, false, strmangle.SetComplement(dtbDeliveryPrimaryKeyColumns, dtbDeliveryColumnsWithoutDefault)...); err != nil {
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

	foreignersSplitByInsertion := [][]*DTBDelivery{
		{&b, &c},
		{&d, &e},
	}

	for i, x := range foreignersSplitByInsertion {
		err = a.AddSaleTypeDTBDeliveries(ctx, tx, i != 0, x...)
		if err != nil {
			t.Fatal(err)
		}

		first := x[0]
		second := x[1]

		if !queries.Equal(a.ID, first.SaleTypeID) {
			t.Error("foreign key was wrong value", a.ID, first.SaleTypeID)
		}
		if !queries.Equal(a.ID, second.SaleTypeID) {
			t.Error("foreign key was wrong value", a.ID, second.SaleTypeID)
		}

		if first.R.SaleType != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}
		if second.R.SaleType != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}

		if a.R.SaleTypeDTBDeliveries[i*2] != first {
			t.Error("relationship struct slice not set to correct value")
		}
		if a.R.SaleTypeDTBDeliveries[i*2+1] != second {
			t.Error("relationship struct slice not set to correct value")
		}

		count, err := a.SaleTypeDTBDeliveries().Count(ctx, tx)
		if err != nil {
			t.Fatal(err)
		}
		if want := int64((i + 1) * 2); count != want {
			t.Error("want", want, "got", count)
		}
	}
}

func testMTBSaleTypeToManySetOpSaleTypeDTBDeliveries(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a MTBSaleType
	var b, c, d, e DTBDelivery

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, mtbSaleTypeDBTypes, false, strmangle.SetComplement(mtbSaleTypePrimaryKeyColumns, mtbSaleTypeColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*DTBDelivery{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, dtbDeliveryDBTypes, false, strmangle.SetComplement(dtbDeliveryPrimaryKeyColumns, dtbDeliveryColumnsWithoutDefault)...); err != nil {
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

	err = a.SetSaleTypeDTBDeliveries(ctx, tx, false, &b, &c)
	if err != nil {
		t.Fatal(err)
	}

	count, err := a.SaleTypeDTBDeliveries().Count(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}
	if count != 2 {
		t.Error("count was wrong:", count)
	}

	err = a.SetSaleTypeDTBDeliveries(ctx, tx, true, &d, &e)
	if err != nil {
		t.Fatal(err)
	}

	count, err = a.SaleTypeDTBDeliveries().Count(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}
	if count != 2 {
		t.Error("count was wrong:", count)
	}

	if !queries.IsValuerNil(b.SaleTypeID) {
		t.Error("want b's foreign key value to be nil")
	}
	if !queries.IsValuerNil(c.SaleTypeID) {
		t.Error("want c's foreign key value to be nil")
	}
	if !queries.Equal(a.ID, d.SaleTypeID) {
		t.Error("foreign key was wrong value", a.ID, d.SaleTypeID)
	}
	if !queries.Equal(a.ID, e.SaleTypeID) {
		t.Error("foreign key was wrong value", a.ID, e.SaleTypeID)
	}

	if b.R.SaleType != nil {
		t.Error("relationship was not removed properly from the foreign struct")
	}
	if c.R.SaleType != nil {
		t.Error("relationship was not removed properly from the foreign struct")
	}
	if d.R.SaleType != &a {
		t.Error("relationship was not added properly to the foreign struct")
	}
	if e.R.SaleType != &a {
		t.Error("relationship was not added properly to the foreign struct")
	}

	if a.R.SaleTypeDTBDeliveries[0] != &d {
		t.Error("relationship struct slice not set to correct value")
	}
	if a.R.SaleTypeDTBDeliveries[1] != &e {
		t.Error("relationship struct slice not set to correct value")
	}
}

func testMTBSaleTypeToManyRemoveOpSaleTypeDTBDeliveries(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a MTBSaleType
	var b, c, d, e DTBDelivery

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, mtbSaleTypeDBTypes, false, strmangle.SetComplement(mtbSaleTypePrimaryKeyColumns, mtbSaleTypeColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*DTBDelivery{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, dtbDeliveryDBTypes, false, strmangle.SetComplement(dtbDeliveryPrimaryKeyColumns, dtbDeliveryColumnsWithoutDefault)...); err != nil {
			t.Fatal(err)
		}
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	err = a.AddSaleTypeDTBDeliveries(ctx, tx, true, foreigners...)
	if err != nil {
		t.Fatal(err)
	}

	count, err := a.SaleTypeDTBDeliveries().Count(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}
	if count != 4 {
		t.Error("count was wrong:", count)
	}

	err = a.RemoveSaleTypeDTBDeliveries(ctx, tx, foreigners[:2]...)
	if err != nil {
		t.Fatal(err)
	}

	count, err = a.SaleTypeDTBDeliveries().Count(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}
	if count != 2 {
		t.Error("count was wrong:", count)
	}

	if !queries.IsValuerNil(b.SaleTypeID) {
		t.Error("want b's foreign key value to be nil")
	}
	if !queries.IsValuerNil(c.SaleTypeID) {
		t.Error("want c's foreign key value to be nil")
	}

	if b.R.SaleType != nil {
		t.Error("relationship was not removed properly from the foreign struct")
	}
	if c.R.SaleType != nil {
		t.Error("relationship was not removed properly from the foreign struct")
	}
	if d.R.SaleType != &a {
		t.Error("relationship to a should have been preserved")
	}
	if e.R.SaleType != &a {
		t.Error("relationship to a should have been preserved")
	}

	if len(a.R.SaleTypeDTBDeliveries) != 2 {
		t.Error("should have preserved two relationships")
	}

	// Removal doesn't do a stable deletion for performance so we have to flip the order
	if a.R.SaleTypeDTBDeliveries[1] != &d {
		t.Error("relationship to d should have been preserved")
	}
	if a.R.SaleTypeDTBDeliveries[0] != &e {
		t.Error("relationship to e should have been preserved")
	}
}

func testMTBSaleTypesReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &MTBSaleType{}
	if err = randomize.Struct(seed, o, mtbSaleTypeDBTypes, true, mtbSaleTypeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize MTBSaleType struct: %s", err)
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

func testMTBSaleTypesReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &MTBSaleType{}
	if err = randomize.Struct(seed, o, mtbSaleTypeDBTypes, true, mtbSaleTypeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize MTBSaleType struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := MTBSaleTypeSlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testMTBSaleTypesSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &MTBSaleType{}
	if err = randomize.Struct(seed, o, mtbSaleTypeDBTypes, true, mtbSaleTypeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize MTBSaleType struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := MTBSaleTypes().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	mtbSaleTypeDBTypes = map[string]string{`ID`: `smallint`, `Name`: `varchar`, `SortNo`: `smallint`, `DiscriminatorType`: `varchar`}
	_                  = bytes.MinRead
)

func testMTBSaleTypesUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(mtbSaleTypePrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(mtbSaleTypeColumns) == len(mtbSaleTypePrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &MTBSaleType{}
	if err = randomize.Struct(seed, o, mtbSaleTypeDBTypes, true, mtbSaleTypeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize MTBSaleType struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := MTBSaleTypes().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, mtbSaleTypeDBTypes, true, mtbSaleTypePrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize MTBSaleType struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testMTBSaleTypesSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(mtbSaleTypeColumns) == len(mtbSaleTypePrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &MTBSaleType{}
	if err = randomize.Struct(seed, o, mtbSaleTypeDBTypes, true, mtbSaleTypeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize MTBSaleType struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := MTBSaleTypes().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, mtbSaleTypeDBTypes, true, mtbSaleTypePrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize MTBSaleType struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(mtbSaleTypeColumns, mtbSaleTypePrimaryKeyColumns) {
		fields = mtbSaleTypeColumns
	} else {
		fields = strmangle.SetComplement(
			mtbSaleTypeColumns,
			mtbSaleTypePrimaryKeyColumns,
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

	slice := MTBSaleTypeSlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testMTBSaleTypesUpsert(t *testing.T) {
	t.Parallel()

	if len(mtbSaleTypeColumns) == len(mtbSaleTypePrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}
	if len(mySQLMTBSaleTypeUniqueColumns) == 0 {
		t.Skip("Skipping table with no unique columns to conflict on")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := MTBSaleType{}
	if err = randomize.Struct(seed, &o, mtbSaleTypeDBTypes, false); err != nil {
		t.Errorf("Unable to randomize MTBSaleType struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(ctx, tx, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert MTBSaleType: %s", err)
	}

	count, err := MTBSaleTypes().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, mtbSaleTypeDBTypes, false, mtbSaleTypePrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize MTBSaleType struct: %s", err)
	}

	if err = o.Upsert(ctx, tx, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert MTBSaleType: %s", err)
	}

	count, err = MTBSaleTypes().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}