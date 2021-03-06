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

func testDTBCartItems(t *testing.T) {
	t.Parallel()

	query := DTBCartItems()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testDTBCartItemsDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &DTBCartItem{}
	if err = randomize.Struct(seed, o, dtbCartItemDBTypes, true, dtbCartItemColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DTBCartItem struct: %s", err)
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

	count, err := DTBCartItems().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testDTBCartItemsQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &DTBCartItem{}
	if err = randomize.Struct(seed, o, dtbCartItemDBTypes, true, dtbCartItemColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DTBCartItem struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := DTBCartItems().DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := DTBCartItems().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testDTBCartItemsSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &DTBCartItem{}
	if err = randomize.Struct(seed, o, dtbCartItemDBTypes, true, dtbCartItemColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DTBCartItem struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := DTBCartItemSlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := DTBCartItems().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testDTBCartItemsExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &DTBCartItem{}
	if err = randomize.Struct(seed, o, dtbCartItemDBTypes, true, dtbCartItemColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DTBCartItem struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := DTBCartItemExists(ctx, tx, o.ID)
	if err != nil {
		t.Errorf("Unable to check if DTBCartItem exists: %s", err)
	}
	if !e {
		t.Errorf("Expected DTBCartItemExists to return true, but got false.")
	}
}

func testDTBCartItemsFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &DTBCartItem{}
	if err = randomize.Struct(seed, o, dtbCartItemDBTypes, true, dtbCartItemColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DTBCartItem struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	dtbCartItemFound, err := FindDTBCartItem(ctx, tx, o.ID)
	if err != nil {
		t.Error(err)
	}

	if dtbCartItemFound == nil {
		t.Error("want a record, got nil")
	}
}

func testDTBCartItemsBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &DTBCartItem{}
	if err = randomize.Struct(seed, o, dtbCartItemDBTypes, true, dtbCartItemColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DTBCartItem struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = DTBCartItems().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testDTBCartItemsOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &DTBCartItem{}
	if err = randomize.Struct(seed, o, dtbCartItemDBTypes, true, dtbCartItemColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DTBCartItem struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := DTBCartItems().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testDTBCartItemsAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	dtbCartItemOne := &DTBCartItem{}
	dtbCartItemTwo := &DTBCartItem{}
	if err = randomize.Struct(seed, dtbCartItemOne, dtbCartItemDBTypes, false, dtbCartItemColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DTBCartItem struct: %s", err)
	}
	if err = randomize.Struct(seed, dtbCartItemTwo, dtbCartItemDBTypes, false, dtbCartItemColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DTBCartItem struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = dtbCartItemOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = dtbCartItemTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := DTBCartItems().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testDTBCartItemsCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	dtbCartItemOne := &DTBCartItem{}
	dtbCartItemTwo := &DTBCartItem{}
	if err = randomize.Struct(seed, dtbCartItemOne, dtbCartItemDBTypes, false, dtbCartItemColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DTBCartItem struct: %s", err)
	}
	if err = randomize.Struct(seed, dtbCartItemTwo, dtbCartItemDBTypes, false, dtbCartItemColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DTBCartItem struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = dtbCartItemOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = dtbCartItemTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := DTBCartItems().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func dtbCartItemBeforeInsertHook(ctx context.Context, e boil.ContextExecutor, o *DTBCartItem) error {
	*o = DTBCartItem{}
	return nil
}

func dtbCartItemAfterInsertHook(ctx context.Context, e boil.ContextExecutor, o *DTBCartItem) error {
	*o = DTBCartItem{}
	return nil
}

func dtbCartItemAfterSelectHook(ctx context.Context, e boil.ContextExecutor, o *DTBCartItem) error {
	*o = DTBCartItem{}
	return nil
}

func dtbCartItemBeforeUpdateHook(ctx context.Context, e boil.ContextExecutor, o *DTBCartItem) error {
	*o = DTBCartItem{}
	return nil
}

func dtbCartItemAfterUpdateHook(ctx context.Context, e boil.ContextExecutor, o *DTBCartItem) error {
	*o = DTBCartItem{}
	return nil
}

func dtbCartItemBeforeDeleteHook(ctx context.Context, e boil.ContextExecutor, o *DTBCartItem) error {
	*o = DTBCartItem{}
	return nil
}

func dtbCartItemAfterDeleteHook(ctx context.Context, e boil.ContextExecutor, o *DTBCartItem) error {
	*o = DTBCartItem{}
	return nil
}

func dtbCartItemBeforeUpsertHook(ctx context.Context, e boil.ContextExecutor, o *DTBCartItem) error {
	*o = DTBCartItem{}
	return nil
}

func dtbCartItemAfterUpsertHook(ctx context.Context, e boil.ContextExecutor, o *DTBCartItem) error {
	*o = DTBCartItem{}
	return nil
}

func testDTBCartItemsHooks(t *testing.T) {
	t.Parallel()

	var err error

	ctx := context.Background()
	empty := &DTBCartItem{}
	o := &DTBCartItem{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, dtbCartItemDBTypes, false); err != nil {
		t.Errorf("Unable to randomize DTBCartItem object: %s", err)
	}

	AddDTBCartItemHook(boil.BeforeInsertHook, dtbCartItemBeforeInsertHook)
	if err = o.doBeforeInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	dtbCartItemBeforeInsertHooks = []DTBCartItemHook{}

	AddDTBCartItemHook(boil.AfterInsertHook, dtbCartItemAfterInsertHook)
	if err = o.doAfterInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	dtbCartItemAfterInsertHooks = []DTBCartItemHook{}

	AddDTBCartItemHook(boil.AfterSelectHook, dtbCartItemAfterSelectHook)
	if err = o.doAfterSelectHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	dtbCartItemAfterSelectHooks = []DTBCartItemHook{}

	AddDTBCartItemHook(boil.BeforeUpdateHook, dtbCartItemBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	dtbCartItemBeforeUpdateHooks = []DTBCartItemHook{}

	AddDTBCartItemHook(boil.AfterUpdateHook, dtbCartItemAfterUpdateHook)
	if err = o.doAfterUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	dtbCartItemAfterUpdateHooks = []DTBCartItemHook{}

	AddDTBCartItemHook(boil.BeforeDeleteHook, dtbCartItemBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	dtbCartItemBeforeDeleteHooks = []DTBCartItemHook{}

	AddDTBCartItemHook(boil.AfterDeleteHook, dtbCartItemAfterDeleteHook)
	if err = o.doAfterDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	dtbCartItemAfterDeleteHooks = []DTBCartItemHook{}

	AddDTBCartItemHook(boil.BeforeUpsertHook, dtbCartItemBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	dtbCartItemBeforeUpsertHooks = []DTBCartItemHook{}

	AddDTBCartItemHook(boil.AfterUpsertHook, dtbCartItemAfterUpsertHook)
	if err = o.doAfterUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	dtbCartItemAfterUpsertHooks = []DTBCartItemHook{}
}

func testDTBCartItemsInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &DTBCartItem{}
	if err = randomize.Struct(seed, o, dtbCartItemDBTypes, true, dtbCartItemColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DTBCartItem struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := DTBCartItems().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testDTBCartItemsInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &DTBCartItem{}
	if err = randomize.Struct(seed, o, dtbCartItemDBTypes, true); err != nil {
		t.Errorf("Unable to randomize DTBCartItem struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(dtbCartItemColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := DTBCartItems().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testDTBCartItemToOneDTBCartUsingCart(t *testing.T) {
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var local DTBCartItem
	var foreign DTBCart

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, dtbCartItemDBTypes, true, dtbCartItemColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DTBCartItem struct: %s", err)
	}
	if err := randomize.Struct(seed, &foreign, dtbCartDBTypes, false, dtbCartColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DTBCart struct: %s", err)
	}

	if err := foreign.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	queries.Assign(&local.CartID, foreign.ID)
	if err := local.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := local.Cart().One(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	if !queries.Equal(check.ID, foreign.ID) {
		t.Errorf("want: %v, got %v", foreign.ID, check.ID)
	}

	slice := DTBCartItemSlice{&local}
	if err = local.L.LoadCart(ctx, tx, false, (*[]*DTBCartItem)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if local.R.Cart == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.Cart = nil
	if err = local.L.LoadCart(ctx, tx, true, &local, nil); err != nil {
		t.Fatal(err)
	}
	if local.R.Cart == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testDTBCartItemToOneSetOpDTBCartUsingCart(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a DTBCartItem
	var b, c DTBCart

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, dtbCartItemDBTypes, false, strmangle.SetComplement(dtbCartItemPrimaryKeyColumns, dtbCartItemColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, dtbCartDBTypes, false, strmangle.SetComplement(dtbCartPrimaryKeyColumns, dtbCartColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, dtbCartDBTypes, false, strmangle.SetComplement(dtbCartPrimaryKeyColumns, dtbCartColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*DTBCart{&b, &c} {
		err = a.SetCart(ctx, tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.Cart != x {
			t.Error("relationship struct not set to correct value")
		}

		if x.R.CartDTBCartItems[0] != &a {
			t.Error("failed to append to foreign relationship struct")
		}
		if !queries.Equal(a.CartID, x.ID) {
			t.Error("foreign key was wrong value", a.CartID)
		}

		zero := reflect.Zero(reflect.TypeOf(a.CartID))
		reflect.Indirect(reflect.ValueOf(&a.CartID)).Set(zero)

		if err = a.Reload(ctx, tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if !queries.Equal(a.CartID, x.ID) {
			t.Error("foreign key was wrong value", a.CartID, x.ID)
		}
	}
}

func testDTBCartItemToOneRemoveOpDTBCartUsingCart(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a DTBCartItem
	var b DTBCart

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, dtbCartItemDBTypes, false, strmangle.SetComplement(dtbCartItemPrimaryKeyColumns, dtbCartItemColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, dtbCartDBTypes, false, strmangle.SetComplement(dtbCartPrimaryKeyColumns, dtbCartColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err = a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	if err = a.SetCart(ctx, tx, true, &b); err != nil {
		t.Fatal(err)
	}

	if err = a.RemoveCart(ctx, tx, &b); err != nil {
		t.Error("failed to remove relationship")
	}

	count, err := a.Cart().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 0 {
		t.Error("want no relationships remaining")
	}

	if a.R.Cart != nil {
		t.Error("R struct entry should be nil")
	}

	if !queries.IsValuerNil(a.CartID) {
		t.Error("foreign key value should be nil")
	}

	if len(b.R.CartDTBCartItems) != 0 {
		t.Error("failed to remove a from b's relationships")
	}
}

func testDTBCartItemsReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &DTBCartItem{}
	if err = randomize.Struct(seed, o, dtbCartItemDBTypes, true, dtbCartItemColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DTBCartItem struct: %s", err)
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

func testDTBCartItemsReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &DTBCartItem{}
	if err = randomize.Struct(seed, o, dtbCartItemDBTypes, true, dtbCartItemColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DTBCartItem struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := DTBCartItemSlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testDTBCartItemsSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &DTBCartItem{}
	if err = randomize.Struct(seed, o, dtbCartItemDBTypes, true, dtbCartItemColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DTBCartItem struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := DTBCartItems().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	dtbCartItemDBTypes = map[string]string{`ID`: `int`, `ProductClassID`: `int`, `CartID`: `int`, `Price`: `decimal`, `Quantity`: `decimal`, `PointRate`: `decimal`, `DiscriminatorType`: `varchar`}
	_                  = bytes.MinRead
)

func testDTBCartItemsUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(dtbCartItemPrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(dtbCartItemColumns) == len(dtbCartItemPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &DTBCartItem{}
	if err = randomize.Struct(seed, o, dtbCartItemDBTypes, true, dtbCartItemColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DTBCartItem struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := DTBCartItems().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, dtbCartItemDBTypes, true, dtbCartItemPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize DTBCartItem struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testDTBCartItemsSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(dtbCartItemColumns) == len(dtbCartItemPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &DTBCartItem{}
	if err = randomize.Struct(seed, o, dtbCartItemDBTypes, true, dtbCartItemColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DTBCartItem struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := DTBCartItems().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, dtbCartItemDBTypes, true, dtbCartItemPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize DTBCartItem struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(dtbCartItemColumns, dtbCartItemPrimaryKeyColumns) {
		fields = dtbCartItemColumns
	} else {
		fields = strmangle.SetComplement(
			dtbCartItemColumns,
			dtbCartItemPrimaryKeyColumns,
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

	slice := DTBCartItemSlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testDTBCartItemsUpsert(t *testing.T) {
	t.Parallel()

	if len(dtbCartItemColumns) == len(dtbCartItemPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}
	if len(mySQLDTBCartItemUniqueColumns) == 0 {
		t.Skip("Skipping table with no unique columns to conflict on")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := DTBCartItem{}
	if err = randomize.Struct(seed, &o, dtbCartItemDBTypes, false); err != nil {
		t.Errorf("Unable to randomize DTBCartItem struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(ctx, tx, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert DTBCartItem: %s", err)
	}

	count, err := DTBCartItems().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, dtbCartItemDBTypes, false, dtbCartItemPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize DTBCartItem struct: %s", err)
	}

	if err = o.Upsert(ctx, tx, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert DTBCartItem: %s", err)
	}

	count, err = DTBCartItems().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
