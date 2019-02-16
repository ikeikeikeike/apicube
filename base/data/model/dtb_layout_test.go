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

func testDTBLayouts(t *testing.T) {
	t.Parallel()

	query := DTBLayouts()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testDTBLayoutsDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &DTBLayout{}
	if err = randomize.Struct(seed, o, dtbLayoutDBTypes, true, dtbLayoutColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DTBLayout struct: %s", err)
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

	count, err := DTBLayouts().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testDTBLayoutsQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &DTBLayout{}
	if err = randomize.Struct(seed, o, dtbLayoutDBTypes, true, dtbLayoutColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DTBLayout struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := DTBLayouts().DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := DTBLayouts().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testDTBLayoutsSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &DTBLayout{}
	if err = randomize.Struct(seed, o, dtbLayoutDBTypes, true, dtbLayoutColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DTBLayout struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := DTBLayoutSlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := DTBLayouts().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testDTBLayoutsExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &DTBLayout{}
	if err = randomize.Struct(seed, o, dtbLayoutDBTypes, true, dtbLayoutColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DTBLayout struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := DTBLayoutExists(ctx, tx, o.ID)
	if err != nil {
		t.Errorf("Unable to check if DTBLayout exists: %s", err)
	}
	if !e {
		t.Errorf("Expected DTBLayoutExists to return true, but got false.")
	}
}

func testDTBLayoutsFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &DTBLayout{}
	if err = randomize.Struct(seed, o, dtbLayoutDBTypes, true, dtbLayoutColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DTBLayout struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	dtbLayoutFound, err := FindDTBLayout(ctx, tx, o.ID)
	if err != nil {
		t.Error(err)
	}

	if dtbLayoutFound == nil {
		t.Error("want a record, got nil")
	}
}

func testDTBLayoutsBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &DTBLayout{}
	if err = randomize.Struct(seed, o, dtbLayoutDBTypes, true, dtbLayoutColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DTBLayout struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = DTBLayouts().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testDTBLayoutsOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &DTBLayout{}
	if err = randomize.Struct(seed, o, dtbLayoutDBTypes, true, dtbLayoutColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DTBLayout struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := DTBLayouts().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testDTBLayoutsAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	dtbLayoutOne := &DTBLayout{}
	dtbLayoutTwo := &DTBLayout{}
	if err = randomize.Struct(seed, dtbLayoutOne, dtbLayoutDBTypes, false, dtbLayoutColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DTBLayout struct: %s", err)
	}
	if err = randomize.Struct(seed, dtbLayoutTwo, dtbLayoutDBTypes, false, dtbLayoutColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DTBLayout struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = dtbLayoutOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = dtbLayoutTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := DTBLayouts().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testDTBLayoutsCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	dtbLayoutOne := &DTBLayout{}
	dtbLayoutTwo := &DTBLayout{}
	if err = randomize.Struct(seed, dtbLayoutOne, dtbLayoutDBTypes, false, dtbLayoutColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DTBLayout struct: %s", err)
	}
	if err = randomize.Struct(seed, dtbLayoutTwo, dtbLayoutDBTypes, false, dtbLayoutColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DTBLayout struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = dtbLayoutOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = dtbLayoutTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := DTBLayouts().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func dtbLayoutBeforeInsertHook(ctx context.Context, e boil.ContextExecutor, o *DTBLayout) error {
	*o = DTBLayout{}
	return nil
}

func dtbLayoutAfterInsertHook(ctx context.Context, e boil.ContextExecutor, o *DTBLayout) error {
	*o = DTBLayout{}
	return nil
}

func dtbLayoutAfterSelectHook(ctx context.Context, e boil.ContextExecutor, o *DTBLayout) error {
	*o = DTBLayout{}
	return nil
}

func dtbLayoutBeforeUpdateHook(ctx context.Context, e boil.ContextExecutor, o *DTBLayout) error {
	*o = DTBLayout{}
	return nil
}

func dtbLayoutAfterUpdateHook(ctx context.Context, e boil.ContextExecutor, o *DTBLayout) error {
	*o = DTBLayout{}
	return nil
}

func dtbLayoutBeforeDeleteHook(ctx context.Context, e boil.ContextExecutor, o *DTBLayout) error {
	*o = DTBLayout{}
	return nil
}

func dtbLayoutAfterDeleteHook(ctx context.Context, e boil.ContextExecutor, o *DTBLayout) error {
	*o = DTBLayout{}
	return nil
}

func dtbLayoutBeforeUpsertHook(ctx context.Context, e boil.ContextExecutor, o *DTBLayout) error {
	*o = DTBLayout{}
	return nil
}

func dtbLayoutAfterUpsertHook(ctx context.Context, e boil.ContextExecutor, o *DTBLayout) error {
	*o = DTBLayout{}
	return nil
}

func testDTBLayoutsHooks(t *testing.T) {
	t.Parallel()

	var err error

	ctx := context.Background()
	empty := &DTBLayout{}
	o := &DTBLayout{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, dtbLayoutDBTypes, false); err != nil {
		t.Errorf("Unable to randomize DTBLayout object: %s", err)
	}

	AddDTBLayoutHook(boil.BeforeInsertHook, dtbLayoutBeforeInsertHook)
	if err = o.doBeforeInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	dtbLayoutBeforeInsertHooks = []DTBLayoutHook{}

	AddDTBLayoutHook(boil.AfterInsertHook, dtbLayoutAfterInsertHook)
	if err = o.doAfterInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	dtbLayoutAfterInsertHooks = []DTBLayoutHook{}

	AddDTBLayoutHook(boil.AfterSelectHook, dtbLayoutAfterSelectHook)
	if err = o.doAfterSelectHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	dtbLayoutAfterSelectHooks = []DTBLayoutHook{}

	AddDTBLayoutHook(boil.BeforeUpdateHook, dtbLayoutBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	dtbLayoutBeforeUpdateHooks = []DTBLayoutHook{}

	AddDTBLayoutHook(boil.AfterUpdateHook, dtbLayoutAfterUpdateHook)
	if err = o.doAfterUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	dtbLayoutAfterUpdateHooks = []DTBLayoutHook{}

	AddDTBLayoutHook(boil.BeforeDeleteHook, dtbLayoutBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	dtbLayoutBeforeDeleteHooks = []DTBLayoutHook{}

	AddDTBLayoutHook(boil.AfterDeleteHook, dtbLayoutAfterDeleteHook)
	if err = o.doAfterDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	dtbLayoutAfterDeleteHooks = []DTBLayoutHook{}

	AddDTBLayoutHook(boil.BeforeUpsertHook, dtbLayoutBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	dtbLayoutBeforeUpsertHooks = []DTBLayoutHook{}

	AddDTBLayoutHook(boil.AfterUpsertHook, dtbLayoutAfterUpsertHook)
	if err = o.doAfterUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	dtbLayoutAfterUpsertHooks = []DTBLayoutHook{}
}

func testDTBLayoutsInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &DTBLayout{}
	if err = randomize.Struct(seed, o, dtbLayoutDBTypes, true, dtbLayoutColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DTBLayout struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := DTBLayouts().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testDTBLayoutsInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &DTBLayout{}
	if err = randomize.Struct(seed, o, dtbLayoutDBTypes, true); err != nil {
		t.Errorf("Unable to randomize DTBLayout struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(dtbLayoutColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := DTBLayouts().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testDTBLayoutToManyLayoutDTBBlockPositions(t *testing.T) {
	var err error
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a DTBLayout
	var b, c DTBBlockPosition

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, dtbLayoutDBTypes, true, dtbLayoutColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DTBLayout struct: %s", err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	if err = randomize.Struct(seed, &b, dtbBlockPositionDBTypes, false, dtbBlockPositionColumnsWithDefault...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, dtbBlockPositionDBTypes, false, dtbBlockPositionColumnsWithDefault...); err != nil {
		t.Fatal(err)
	}

	b.LayoutID = a.ID
	c.LayoutID = a.ID

	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := a.LayoutDTBBlockPositions().All(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	bFound, cFound := false, false
	for _, v := range check {
		if v.LayoutID == b.LayoutID {
			bFound = true
		}
		if v.LayoutID == c.LayoutID {
			cFound = true
		}
	}

	if !bFound {
		t.Error("expected to find b")
	}
	if !cFound {
		t.Error("expected to find c")
	}

	slice := DTBLayoutSlice{&a}
	if err = a.L.LoadLayoutDTBBlockPositions(ctx, tx, false, (*[]*DTBLayout)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.LayoutDTBBlockPositions); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	a.R.LayoutDTBBlockPositions = nil
	if err = a.L.LoadLayoutDTBBlockPositions(ctx, tx, true, &a, nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.LayoutDTBBlockPositions); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	if t.Failed() {
		t.Logf("%#v", check)
	}
}

func testDTBLayoutToManyLayoutDTBPageLayouts(t *testing.T) {
	var err error
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a DTBLayout
	var b, c DTBPageLayout

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, dtbLayoutDBTypes, true, dtbLayoutColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DTBLayout struct: %s", err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	if err = randomize.Struct(seed, &b, dtbPageLayoutDBTypes, false, dtbPageLayoutColumnsWithDefault...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, dtbPageLayoutDBTypes, false, dtbPageLayoutColumnsWithDefault...); err != nil {
		t.Fatal(err)
	}

	b.LayoutID = a.ID
	c.LayoutID = a.ID

	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := a.LayoutDTBPageLayouts().All(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	bFound, cFound := false, false
	for _, v := range check {
		if v.LayoutID == b.LayoutID {
			bFound = true
		}
		if v.LayoutID == c.LayoutID {
			cFound = true
		}
	}

	if !bFound {
		t.Error("expected to find b")
	}
	if !cFound {
		t.Error("expected to find c")
	}

	slice := DTBLayoutSlice{&a}
	if err = a.L.LoadLayoutDTBPageLayouts(ctx, tx, false, (*[]*DTBLayout)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.LayoutDTBPageLayouts); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	a.R.LayoutDTBPageLayouts = nil
	if err = a.L.LoadLayoutDTBPageLayouts(ctx, tx, true, &a, nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.LayoutDTBPageLayouts); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	if t.Failed() {
		t.Logf("%#v", check)
	}
}

func testDTBLayoutToManyAddOpLayoutDTBBlockPositions(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a DTBLayout
	var b, c, d, e DTBBlockPosition

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, dtbLayoutDBTypes, false, strmangle.SetComplement(dtbLayoutPrimaryKeyColumns, dtbLayoutColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*DTBBlockPosition{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, dtbBlockPositionDBTypes, false, strmangle.SetComplement(dtbBlockPositionPrimaryKeyColumns, dtbBlockPositionColumnsWithoutDefault)...); err != nil {
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

	foreignersSplitByInsertion := [][]*DTBBlockPosition{
		{&b, &c},
		{&d, &e},
	}

	for i, x := range foreignersSplitByInsertion {
		err = a.AddLayoutDTBBlockPositions(ctx, tx, i != 0, x...)
		if err != nil {
			t.Fatal(err)
		}

		first := x[0]
		second := x[1]

		if a.ID != first.LayoutID {
			t.Error("foreign key was wrong value", a.ID, first.LayoutID)
		}
		if a.ID != second.LayoutID {
			t.Error("foreign key was wrong value", a.ID, second.LayoutID)
		}

		if first.R.Layout != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}
		if second.R.Layout != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}

		if a.R.LayoutDTBBlockPositions[i*2] != first {
			t.Error("relationship struct slice not set to correct value")
		}
		if a.R.LayoutDTBBlockPositions[i*2+1] != second {
			t.Error("relationship struct slice not set to correct value")
		}

		count, err := a.LayoutDTBBlockPositions().Count(ctx, tx)
		if err != nil {
			t.Fatal(err)
		}
		if want := int64((i + 1) * 2); count != want {
			t.Error("want", want, "got", count)
		}
	}
}
func testDTBLayoutToManyAddOpLayoutDTBPageLayouts(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a DTBLayout
	var b, c, d, e DTBPageLayout

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, dtbLayoutDBTypes, false, strmangle.SetComplement(dtbLayoutPrimaryKeyColumns, dtbLayoutColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*DTBPageLayout{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, dtbPageLayoutDBTypes, false, strmangle.SetComplement(dtbPageLayoutPrimaryKeyColumns, dtbPageLayoutColumnsWithoutDefault)...); err != nil {
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

	foreignersSplitByInsertion := [][]*DTBPageLayout{
		{&b, &c},
		{&d, &e},
	}

	for i, x := range foreignersSplitByInsertion {
		err = a.AddLayoutDTBPageLayouts(ctx, tx, i != 0, x...)
		if err != nil {
			t.Fatal(err)
		}

		first := x[0]
		second := x[1]

		if a.ID != first.LayoutID {
			t.Error("foreign key was wrong value", a.ID, first.LayoutID)
		}
		if a.ID != second.LayoutID {
			t.Error("foreign key was wrong value", a.ID, second.LayoutID)
		}

		if first.R.Layout != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}
		if second.R.Layout != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}

		if a.R.LayoutDTBPageLayouts[i*2] != first {
			t.Error("relationship struct slice not set to correct value")
		}
		if a.R.LayoutDTBPageLayouts[i*2+1] != second {
			t.Error("relationship struct slice not set to correct value")
		}

		count, err := a.LayoutDTBPageLayouts().Count(ctx, tx)
		if err != nil {
			t.Fatal(err)
		}
		if want := int64((i + 1) * 2); count != want {
			t.Error("want", want, "got", count)
		}
	}
}
func testDTBLayoutToOneMTBDeviceTypeUsingDeviceType(t *testing.T) {
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var local DTBLayout
	var foreign MTBDeviceType

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, dtbLayoutDBTypes, true, dtbLayoutColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DTBLayout struct: %s", err)
	}
	if err := randomize.Struct(seed, &foreign, mtbDeviceTypeDBTypes, false, mtbDeviceTypeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize MTBDeviceType struct: %s", err)
	}

	if err := foreign.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	queries.Assign(&local.DeviceTypeID, foreign.ID)
	if err := local.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := local.DeviceType().One(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	if !queries.Equal(check.ID, foreign.ID) {
		t.Errorf("want: %v, got %v", foreign.ID, check.ID)
	}

	slice := DTBLayoutSlice{&local}
	if err = local.L.LoadDeviceType(ctx, tx, false, (*[]*DTBLayout)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if local.R.DeviceType == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.DeviceType = nil
	if err = local.L.LoadDeviceType(ctx, tx, true, &local, nil); err != nil {
		t.Fatal(err)
	}
	if local.R.DeviceType == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testDTBLayoutToOneSetOpMTBDeviceTypeUsingDeviceType(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a DTBLayout
	var b, c MTBDeviceType

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, dtbLayoutDBTypes, false, strmangle.SetComplement(dtbLayoutPrimaryKeyColumns, dtbLayoutColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, mtbDeviceTypeDBTypes, false, strmangle.SetComplement(mtbDeviceTypePrimaryKeyColumns, mtbDeviceTypeColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, mtbDeviceTypeDBTypes, false, strmangle.SetComplement(mtbDeviceTypePrimaryKeyColumns, mtbDeviceTypeColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*MTBDeviceType{&b, &c} {
		err = a.SetDeviceType(ctx, tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.DeviceType != x {
			t.Error("relationship struct not set to correct value")
		}

		if x.R.DeviceTypeDTBLayouts[0] != &a {
			t.Error("failed to append to foreign relationship struct")
		}
		if !queries.Equal(a.DeviceTypeID, x.ID) {
			t.Error("foreign key was wrong value", a.DeviceTypeID)
		}

		zero := reflect.Zero(reflect.TypeOf(a.DeviceTypeID))
		reflect.Indirect(reflect.ValueOf(&a.DeviceTypeID)).Set(zero)

		if err = a.Reload(ctx, tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if !queries.Equal(a.DeviceTypeID, x.ID) {
			t.Error("foreign key was wrong value", a.DeviceTypeID, x.ID)
		}
	}
}

func testDTBLayoutToOneRemoveOpMTBDeviceTypeUsingDeviceType(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a DTBLayout
	var b MTBDeviceType

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, dtbLayoutDBTypes, false, strmangle.SetComplement(dtbLayoutPrimaryKeyColumns, dtbLayoutColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, mtbDeviceTypeDBTypes, false, strmangle.SetComplement(mtbDeviceTypePrimaryKeyColumns, mtbDeviceTypeColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err = a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	if err = a.SetDeviceType(ctx, tx, true, &b); err != nil {
		t.Fatal(err)
	}

	if err = a.RemoveDeviceType(ctx, tx, &b); err != nil {
		t.Error("failed to remove relationship")
	}

	count, err := a.DeviceType().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 0 {
		t.Error("want no relationships remaining")
	}

	if a.R.DeviceType != nil {
		t.Error("R struct entry should be nil")
	}

	if !queries.IsValuerNil(a.DeviceTypeID) {
		t.Error("foreign key value should be nil")
	}

	if len(b.R.DeviceTypeDTBLayouts) != 0 {
		t.Error("failed to remove a from b's relationships")
	}
}

func testDTBLayoutsReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &DTBLayout{}
	if err = randomize.Struct(seed, o, dtbLayoutDBTypes, true, dtbLayoutColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DTBLayout struct: %s", err)
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

func testDTBLayoutsReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &DTBLayout{}
	if err = randomize.Struct(seed, o, dtbLayoutDBTypes, true, dtbLayoutColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DTBLayout struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := DTBLayoutSlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testDTBLayoutsSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &DTBLayout{}
	if err = randomize.Struct(seed, o, dtbLayoutDBTypes, true, dtbLayoutColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DTBLayout struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := DTBLayouts().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	dtbLayoutDBTypes = map[string]string{`ID`: `int`, `DeviceTypeID`: `smallint`, `LayoutName`: `varchar`, `CreateDate`: `datetime`, `UpdateDate`: `datetime`, `DiscriminatorType`: `varchar`}
	_                = bytes.MinRead
)

func testDTBLayoutsUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(dtbLayoutPrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(dtbLayoutColumns) == len(dtbLayoutPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &DTBLayout{}
	if err = randomize.Struct(seed, o, dtbLayoutDBTypes, true, dtbLayoutColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DTBLayout struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := DTBLayouts().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, dtbLayoutDBTypes, true, dtbLayoutPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize DTBLayout struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testDTBLayoutsSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(dtbLayoutColumns) == len(dtbLayoutPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &DTBLayout{}
	if err = randomize.Struct(seed, o, dtbLayoutDBTypes, true, dtbLayoutColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DTBLayout struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := DTBLayouts().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, dtbLayoutDBTypes, true, dtbLayoutPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize DTBLayout struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(dtbLayoutColumns, dtbLayoutPrimaryKeyColumns) {
		fields = dtbLayoutColumns
	} else {
		fields = strmangle.SetComplement(
			dtbLayoutColumns,
			dtbLayoutPrimaryKeyColumns,
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

	slice := DTBLayoutSlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testDTBLayoutsUpsert(t *testing.T) {
	t.Parallel()

	if len(dtbLayoutColumns) == len(dtbLayoutPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}
	if len(mySQLDTBLayoutUniqueColumns) == 0 {
		t.Skip("Skipping table with no unique columns to conflict on")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := DTBLayout{}
	if err = randomize.Struct(seed, &o, dtbLayoutDBTypes, false); err != nil {
		t.Errorf("Unable to randomize DTBLayout struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(ctx, tx, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert DTBLayout: %s", err)
	}

	count, err := DTBLayouts().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, dtbLayoutDBTypes, false, dtbLayoutPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize DTBLayout struct: %s", err)
	}

	if err = o.Upsert(ctx, tx, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert DTBLayout: %s", err)
	}

	count, err = DTBLayouts().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
