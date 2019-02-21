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

func testMTBTaxDisplayTypes(t *testing.T) {
	t.Parallel()

	query := MTBTaxDisplayTypes()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testMTBTaxDisplayTypesDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &MTBTaxDisplayType{}
	if err = randomize.Struct(seed, o, mtbTaxDisplayTypeDBTypes, true, mtbTaxDisplayTypeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize MTBTaxDisplayType struct: %s", err)
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

	count, err := MTBTaxDisplayTypes().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testMTBTaxDisplayTypesQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &MTBTaxDisplayType{}
	if err = randomize.Struct(seed, o, mtbTaxDisplayTypeDBTypes, true, mtbTaxDisplayTypeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize MTBTaxDisplayType struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := MTBTaxDisplayTypes().DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := MTBTaxDisplayTypes().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testMTBTaxDisplayTypesSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &MTBTaxDisplayType{}
	if err = randomize.Struct(seed, o, mtbTaxDisplayTypeDBTypes, true, mtbTaxDisplayTypeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize MTBTaxDisplayType struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := MTBTaxDisplayTypeSlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := MTBTaxDisplayTypes().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testMTBTaxDisplayTypesExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &MTBTaxDisplayType{}
	if err = randomize.Struct(seed, o, mtbTaxDisplayTypeDBTypes, true, mtbTaxDisplayTypeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize MTBTaxDisplayType struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := MTBTaxDisplayTypeExists(ctx, tx, o.ID)
	if err != nil {
		t.Errorf("Unable to check if MTBTaxDisplayType exists: %s", err)
	}
	if !e {
		t.Errorf("Expected MTBTaxDisplayTypeExists to return true, but got false.")
	}
}

func testMTBTaxDisplayTypesFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &MTBTaxDisplayType{}
	if err = randomize.Struct(seed, o, mtbTaxDisplayTypeDBTypes, true, mtbTaxDisplayTypeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize MTBTaxDisplayType struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	mtbTaxDisplayTypeFound, err := FindMTBTaxDisplayType(ctx, tx, o.ID)
	if err != nil {
		t.Error(err)
	}

	if mtbTaxDisplayTypeFound == nil {
		t.Error("want a record, got nil")
	}
}

func testMTBTaxDisplayTypesBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &MTBTaxDisplayType{}
	if err = randomize.Struct(seed, o, mtbTaxDisplayTypeDBTypes, true, mtbTaxDisplayTypeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize MTBTaxDisplayType struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = MTBTaxDisplayTypes().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testMTBTaxDisplayTypesOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &MTBTaxDisplayType{}
	if err = randomize.Struct(seed, o, mtbTaxDisplayTypeDBTypes, true, mtbTaxDisplayTypeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize MTBTaxDisplayType struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := MTBTaxDisplayTypes().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testMTBTaxDisplayTypesAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	mtbTaxDisplayTypeOne := &MTBTaxDisplayType{}
	mtbTaxDisplayTypeTwo := &MTBTaxDisplayType{}
	if err = randomize.Struct(seed, mtbTaxDisplayTypeOne, mtbTaxDisplayTypeDBTypes, false, mtbTaxDisplayTypeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize MTBTaxDisplayType struct: %s", err)
	}
	if err = randomize.Struct(seed, mtbTaxDisplayTypeTwo, mtbTaxDisplayTypeDBTypes, false, mtbTaxDisplayTypeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize MTBTaxDisplayType struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = mtbTaxDisplayTypeOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = mtbTaxDisplayTypeTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := MTBTaxDisplayTypes().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testMTBTaxDisplayTypesCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	mtbTaxDisplayTypeOne := &MTBTaxDisplayType{}
	mtbTaxDisplayTypeTwo := &MTBTaxDisplayType{}
	if err = randomize.Struct(seed, mtbTaxDisplayTypeOne, mtbTaxDisplayTypeDBTypes, false, mtbTaxDisplayTypeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize MTBTaxDisplayType struct: %s", err)
	}
	if err = randomize.Struct(seed, mtbTaxDisplayTypeTwo, mtbTaxDisplayTypeDBTypes, false, mtbTaxDisplayTypeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize MTBTaxDisplayType struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = mtbTaxDisplayTypeOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = mtbTaxDisplayTypeTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := MTBTaxDisplayTypes().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func mtbTaxDisplayTypeBeforeInsertHook(ctx context.Context, e boil.ContextExecutor, o *MTBTaxDisplayType) error {
	*o = MTBTaxDisplayType{}
	return nil
}

func mtbTaxDisplayTypeAfterInsertHook(ctx context.Context, e boil.ContextExecutor, o *MTBTaxDisplayType) error {
	*o = MTBTaxDisplayType{}
	return nil
}

func mtbTaxDisplayTypeAfterSelectHook(ctx context.Context, e boil.ContextExecutor, o *MTBTaxDisplayType) error {
	*o = MTBTaxDisplayType{}
	return nil
}

func mtbTaxDisplayTypeBeforeUpdateHook(ctx context.Context, e boil.ContextExecutor, o *MTBTaxDisplayType) error {
	*o = MTBTaxDisplayType{}
	return nil
}

func mtbTaxDisplayTypeAfterUpdateHook(ctx context.Context, e boil.ContextExecutor, o *MTBTaxDisplayType) error {
	*o = MTBTaxDisplayType{}
	return nil
}

func mtbTaxDisplayTypeBeforeDeleteHook(ctx context.Context, e boil.ContextExecutor, o *MTBTaxDisplayType) error {
	*o = MTBTaxDisplayType{}
	return nil
}

func mtbTaxDisplayTypeAfterDeleteHook(ctx context.Context, e boil.ContextExecutor, o *MTBTaxDisplayType) error {
	*o = MTBTaxDisplayType{}
	return nil
}

func mtbTaxDisplayTypeBeforeUpsertHook(ctx context.Context, e boil.ContextExecutor, o *MTBTaxDisplayType) error {
	*o = MTBTaxDisplayType{}
	return nil
}

func mtbTaxDisplayTypeAfterUpsertHook(ctx context.Context, e boil.ContextExecutor, o *MTBTaxDisplayType) error {
	*o = MTBTaxDisplayType{}
	return nil
}

func testMTBTaxDisplayTypesHooks(t *testing.T) {
	t.Parallel()

	var err error

	ctx := context.Background()
	empty := &MTBTaxDisplayType{}
	o := &MTBTaxDisplayType{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, mtbTaxDisplayTypeDBTypes, false); err != nil {
		t.Errorf("Unable to randomize MTBTaxDisplayType object: %s", err)
	}

	AddMTBTaxDisplayTypeHook(boil.BeforeInsertHook, mtbTaxDisplayTypeBeforeInsertHook)
	if err = o.doBeforeInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	mtbTaxDisplayTypeBeforeInsertHooks = []MTBTaxDisplayTypeHook{}

	AddMTBTaxDisplayTypeHook(boil.AfterInsertHook, mtbTaxDisplayTypeAfterInsertHook)
	if err = o.doAfterInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	mtbTaxDisplayTypeAfterInsertHooks = []MTBTaxDisplayTypeHook{}

	AddMTBTaxDisplayTypeHook(boil.AfterSelectHook, mtbTaxDisplayTypeAfterSelectHook)
	if err = o.doAfterSelectHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	mtbTaxDisplayTypeAfterSelectHooks = []MTBTaxDisplayTypeHook{}

	AddMTBTaxDisplayTypeHook(boil.BeforeUpdateHook, mtbTaxDisplayTypeBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	mtbTaxDisplayTypeBeforeUpdateHooks = []MTBTaxDisplayTypeHook{}

	AddMTBTaxDisplayTypeHook(boil.AfterUpdateHook, mtbTaxDisplayTypeAfterUpdateHook)
	if err = o.doAfterUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	mtbTaxDisplayTypeAfterUpdateHooks = []MTBTaxDisplayTypeHook{}

	AddMTBTaxDisplayTypeHook(boil.BeforeDeleteHook, mtbTaxDisplayTypeBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	mtbTaxDisplayTypeBeforeDeleteHooks = []MTBTaxDisplayTypeHook{}

	AddMTBTaxDisplayTypeHook(boil.AfterDeleteHook, mtbTaxDisplayTypeAfterDeleteHook)
	if err = o.doAfterDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	mtbTaxDisplayTypeAfterDeleteHooks = []MTBTaxDisplayTypeHook{}

	AddMTBTaxDisplayTypeHook(boil.BeforeUpsertHook, mtbTaxDisplayTypeBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	mtbTaxDisplayTypeBeforeUpsertHooks = []MTBTaxDisplayTypeHook{}

	AddMTBTaxDisplayTypeHook(boil.AfterUpsertHook, mtbTaxDisplayTypeAfterUpsertHook)
	if err = o.doAfterUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	mtbTaxDisplayTypeAfterUpsertHooks = []MTBTaxDisplayTypeHook{}
}

func testMTBTaxDisplayTypesInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &MTBTaxDisplayType{}
	if err = randomize.Struct(seed, o, mtbTaxDisplayTypeDBTypes, true, mtbTaxDisplayTypeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize MTBTaxDisplayType struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := MTBTaxDisplayTypes().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testMTBTaxDisplayTypesInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &MTBTaxDisplayType{}
	if err = randomize.Struct(seed, o, mtbTaxDisplayTypeDBTypes, true); err != nil {
		t.Errorf("Unable to randomize MTBTaxDisplayType struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(mtbTaxDisplayTypeColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := MTBTaxDisplayTypes().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testMTBTaxDisplayTypeToManyTaxDisplayTypeDTBOrderItems(t *testing.T) {
	var err error
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a MTBTaxDisplayType
	var b, c DTBOrderItem

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, mtbTaxDisplayTypeDBTypes, true, mtbTaxDisplayTypeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize MTBTaxDisplayType struct: %s", err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	if err = randomize.Struct(seed, &b, dtbOrderItemDBTypes, false, dtbOrderItemColumnsWithDefault...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, dtbOrderItemDBTypes, false, dtbOrderItemColumnsWithDefault...); err != nil {
		t.Fatal(err)
	}

	queries.Assign(&b.TaxDisplayTypeID, a.ID)
	queries.Assign(&c.TaxDisplayTypeID, a.ID)
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := a.TaxDisplayTypeDTBOrderItems().All(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	bFound, cFound := false, false
	for _, v := range check {
		if queries.Equal(v.TaxDisplayTypeID, b.TaxDisplayTypeID) {
			bFound = true
		}
		if queries.Equal(v.TaxDisplayTypeID, c.TaxDisplayTypeID) {
			cFound = true
		}
	}

	if !bFound {
		t.Error("expected to find b")
	}
	if !cFound {
		t.Error("expected to find c")
	}

	slice := MTBTaxDisplayTypeSlice{&a}
	if err = a.L.LoadTaxDisplayTypeDTBOrderItems(ctx, tx, false, (*[]*MTBTaxDisplayType)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.TaxDisplayTypeDTBOrderItems); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	a.R.TaxDisplayTypeDTBOrderItems = nil
	if err = a.L.LoadTaxDisplayTypeDTBOrderItems(ctx, tx, true, &a, nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.TaxDisplayTypeDTBOrderItems); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	if t.Failed() {
		t.Logf("%#v", check)
	}
}

func testMTBTaxDisplayTypeToManyAddOpTaxDisplayTypeDTBOrderItems(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a MTBTaxDisplayType
	var b, c, d, e DTBOrderItem

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, mtbTaxDisplayTypeDBTypes, false, strmangle.SetComplement(mtbTaxDisplayTypePrimaryKeyColumns, mtbTaxDisplayTypeColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*DTBOrderItem{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, dtbOrderItemDBTypes, false, strmangle.SetComplement(dtbOrderItemPrimaryKeyColumns, dtbOrderItemColumnsWithoutDefault)...); err != nil {
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

	foreignersSplitByInsertion := [][]*DTBOrderItem{
		{&b, &c},
		{&d, &e},
	}

	for i, x := range foreignersSplitByInsertion {
		err = a.AddTaxDisplayTypeDTBOrderItems(ctx, tx, i != 0, x...)
		if err != nil {
			t.Fatal(err)
		}

		first := x[0]
		second := x[1]

		if !queries.Equal(a.ID, first.TaxDisplayTypeID) {
			t.Error("foreign key was wrong value", a.ID, first.TaxDisplayTypeID)
		}
		if !queries.Equal(a.ID, second.TaxDisplayTypeID) {
			t.Error("foreign key was wrong value", a.ID, second.TaxDisplayTypeID)
		}

		if first.R.TaxDisplayType != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}
		if second.R.TaxDisplayType != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}

		if a.R.TaxDisplayTypeDTBOrderItems[i*2] != first {
			t.Error("relationship struct slice not set to correct value")
		}
		if a.R.TaxDisplayTypeDTBOrderItems[i*2+1] != second {
			t.Error("relationship struct slice not set to correct value")
		}

		count, err := a.TaxDisplayTypeDTBOrderItems().Count(ctx, tx)
		if err != nil {
			t.Fatal(err)
		}
		if want := int64((i + 1) * 2); count != want {
			t.Error("want", want, "got", count)
		}
	}
}

func testMTBTaxDisplayTypeToManySetOpTaxDisplayTypeDTBOrderItems(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a MTBTaxDisplayType
	var b, c, d, e DTBOrderItem

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, mtbTaxDisplayTypeDBTypes, false, strmangle.SetComplement(mtbTaxDisplayTypePrimaryKeyColumns, mtbTaxDisplayTypeColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*DTBOrderItem{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, dtbOrderItemDBTypes, false, strmangle.SetComplement(dtbOrderItemPrimaryKeyColumns, dtbOrderItemColumnsWithoutDefault)...); err != nil {
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

	err = a.SetTaxDisplayTypeDTBOrderItems(ctx, tx, false, &b, &c)
	if err != nil {
		t.Fatal(err)
	}

	count, err := a.TaxDisplayTypeDTBOrderItems().Count(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}
	if count != 2 {
		t.Error("count was wrong:", count)
	}

	err = a.SetTaxDisplayTypeDTBOrderItems(ctx, tx, true, &d, &e)
	if err != nil {
		t.Fatal(err)
	}

	count, err = a.TaxDisplayTypeDTBOrderItems().Count(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}
	if count != 2 {
		t.Error("count was wrong:", count)
	}

	if !queries.IsValuerNil(b.TaxDisplayTypeID) {
		t.Error("want b's foreign key value to be nil")
	}
	if !queries.IsValuerNil(c.TaxDisplayTypeID) {
		t.Error("want c's foreign key value to be nil")
	}
	if !queries.Equal(a.ID, d.TaxDisplayTypeID) {
		t.Error("foreign key was wrong value", a.ID, d.TaxDisplayTypeID)
	}
	if !queries.Equal(a.ID, e.TaxDisplayTypeID) {
		t.Error("foreign key was wrong value", a.ID, e.TaxDisplayTypeID)
	}

	if b.R.TaxDisplayType != nil {
		t.Error("relationship was not removed properly from the foreign struct")
	}
	if c.R.TaxDisplayType != nil {
		t.Error("relationship was not removed properly from the foreign struct")
	}
	if d.R.TaxDisplayType != &a {
		t.Error("relationship was not added properly to the foreign struct")
	}
	if e.R.TaxDisplayType != &a {
		t.Error("relationship was not added properly to the foreign struct")
	}

	if a.R.TaxDisplayTypeDTBOrderItems[0] != &d {
		t.Error("relationship struct slice not set to correct value")
	}
	if a.R.TaxDisplayTypeDTBOrderItems[1] != &e {
		t.Error("relationship struct slice not set to correct value")
	}
}

func testMTBTaxDisplayTypeToManyRemoveOpTaxDisplayTypeDTBOrderItems(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a MTBTaxDisplayType
	var b, c, d, e DTBOrderItem

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, mtbTaxDisplayTypeDBTypes, false, strmangle.SetComplement(mtbTaxDisplayTypePrimaryKeyColumns, mtbTaxDisplayTypeColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*DTBOrderItem{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, dtbOrderItemDBTypes, false, strmangle.SetComplement(dtbOrderItemPrimaryKeyColumns, dtbOrderItemColumnsWithoutDefault)...); err != nil {
			t.Fatal(err)
		}
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	err = a.AddTaxDisplayTypeDTBOrderItems(ctx, tx, true, foreigners...)
	if err != nil {
		t.Fatal(err)
	}

	count, err := a.TaxDisplayTypeDTBOrderItems().Count(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}
	if count != 4 {
		t.Error("count was wrong:", count)
	}

	err = a.RemoveTaxDisplayTypeDTBOrderItems(ctx, tx, foreigners[:2]...)
	if err != nil {
		t.Fatal(err)
	}

	count, err = a.TaxDisplayTypeDTBOrderItems().Count(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}
	if count != 2 {
		t.Error("count was wrong:", count)
	}

	if !queries.IsValuerNil(b.TaxDisplayTypeID) {
		t.Error("want b's foreign key value to be nil")
	}
	if !queries.IsValuerNil(c.TaxDisplayTypeID) {
		t.Error("want c's foreign key value to be nil")
	}

	if b.R.TaxDisplayType != nil {
		t.Error("relationship was not removed properly from the foreign struct")
	}
	if c.R.TaxDisplayType != nil {
		t.Error("relationship was not removed properly from the foreign struct")
	}
	if d.R.TaxDisplayType != &a {
		t.Error("relationship to a should have been preserved")
	}
	if e.R.TaxDisplayType != &a {
		t.Error("relationship to a should have been preserved")
	}

	if len(a.R.TaxDisplayTypeDTBOrderItems) != 2 {
		t.Error("should have preserved two relationships")
	}

	// Removal doesn't do a stable deletion for performance so we have to flip the order
	if a.R.TaxDisplayTypeDTBOrderItems[1] != &d {
		t.Error("relationship to d should have been preserved")
	}
	if a.R.TaxDisplayTypeDTBOrderItems[0] != &e {
		t.Error("relationship to e should have been preserved")
	}
}

func testMTBTaxDisplayTypesReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &MTBTaxDisplayType{}
	if err = randomize.Struct(seed, o, mtbTaxDisplayTypeDBTypes, true, mtbTaxDisplayTypeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize MTBTaxDisplayType struct: %s", err)
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

func testMTBTaxDisplayTypesReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &MTBTaxDisplayType{}
	if err = randomize.Struct(seed, o, mtbTaxDisplayTypeDBTypes, true, mtbTaxDisplayTypeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize MTBTaxDisplayType struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := MTBTaxDisplayTypeSlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testMTBTaxDisplayTypesSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &MTBTaxDisplayType{}
	if err = randomize.Struct(seed, o, mtbTaxDisplayTypeDBTypes, true, mtbTaxDisplayTypeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize MTBTaxDisplayType struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := MTBTaxDisplayTypes().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	mtbTaxDisplayTypeDBTypes = map[string]string{`ID`: `smallint`, `Name`: `varchar`, `SortNo`: `smallint`, `DiscriminatorType`: `varchar`}
	_                        = bytes.MinRead
)

func testMTBTaxDisplayTypesUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(mtbTaxDisplayTypePrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(mtbTaxDisplayTypeColumns) == len(mtbTaxDisplayTypePrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &MTBTaxDisplayType{}
	if err = randomize.Struct(seed, o, mtbTaxDisplayTypeDBTypes, true, mtbTaxDisplayTypeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize MTBTaxDisplayType struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := MTBTaxDisplayTypes().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, mtbTaxDisplayTypeDBTypes, true, mtbTaxDisplayTypePrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize MTBTaxDisplayType struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testMTBTaxDisplayTypesSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(mtbTaxDisplayTypeColumns) == len(mtbTaxDisplayTypePrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &MTBTaxDisplayType{}
	if err = randomize.Struct(seed, o, mtbTaxDisplayTypeDBTypes, true, mtbTaxDisplayTypeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize MTBTaxDisplayType struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := MTBTaxDisplayTypes().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, mtbTaxDisplayTypeDBTypes, true, mtbTaxDisplayTypePrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize MTBTaxDisplayType struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(mtbTaxDisplayTypeColumns, mtbTaxDisplayTypePrimaryKeyColumns) {
		fields = mtbTaxDisplayTypeColumns
	} else {
		fields = strmangle.SetComplement(
			mtbTaxDisplayTypeColumns,
			mtbTaxDisplayTypePrimaryKeyColumns,
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

	slice := MTBTaxDisplayTypeSlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testMTBTaxDisplayTypesUpsert(t *testing.T) {
	t.Parallel()

	if len(mtbTaxDisplayTypeColumns) == len(mtbTaxDisplayTypePrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}
	if len(mySQLMTBTaxDisplayTypeUniqueColumns) == 0 {
		t.Skip("Skipping table with no unique columns to conflict on")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := MTBTaxDisplayType{}
	if err = randomize.Struct(seed, &o, mtbTaxDisplayTypeDBTypes, false); err != nil {
		t.Errorf("Unable to randomize MTBTaxDisplayType struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(ctx, tx, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert MTBTaxDisplayType: %s", err)
	}

	count, err := MTBTaxDisplayTypes().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, mtbTaxDisplayTypeDBTypes, false, mtbTaxDisplayTypePrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize MTBTaxDisplayType struct: %s", err)
	}

	if err = o.Upsert(ctx, tx, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert MTBTaxDisplayType: %s", err)
	}

	count, err = MTBTaxDisplayTypes().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}