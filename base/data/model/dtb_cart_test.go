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

func testDTBCarts(t *testing.T) {
	t.Parallel()

	query := DTBCarts()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testDTBCartsDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &DTBCart{}
	if err = randomize.Struct(seed, o, dtbCartDBTypes, true, dtbCartColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DTBCart struct: %s", err)
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

	count, err := DTBCarts().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testDTBCartsQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &DTBCart{}
	if err = randomize.Struct(seed, o, dtbCartDBTypes, true, dtbCartColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DTBCart struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := DTBCarts().DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := DTBCarts().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testDTBCartsSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &DTBCart{}
	if err = randomize.Struct(seed, o, dtbCartDBTypes, true, dtbCartColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DTBCart struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := DTBCartSlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := DTBCarts().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testDTBCartsExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &DTBCart{}
	if err = randomize.Struct(seed, o, dtbCartDBTypes, true, dtbCartColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DTBCart struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := DTBCartExists(ctx, tx, o.ID)
	if err != nil {
		t.Errorf("Unable to check if DTBCart exists: %s", err)
	}
	if !e {
		t.Errorf("Expected DTBCartExists to return true, but got false.")
	}
}

func testDTBCartsFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &DTBCart{}
	if err = randomize.Struct(seed, o, dtbCartDBTypes, true, dtbCartColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DTBCart struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	dtbCartFound, err := FindDTBCart(ctx, tx, o.ID)
	if err != nil {
		t.Error(err)
	}

	if dtbCartFound == nil {
		t.Error("want a record, got nil")
	}
}

func testDTBCartsBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &DTBCart{}
	if err = randomize.Struct(seed, o, dtbCartDBTypes, true, dtbCartColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DTBCart struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = DTBCarts().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testDTBCartsOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &DTBCart{}
	if err = randomize.Struct(seed, o, dtbCartDBTypes, true, dtbCartColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DTBCart struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := DTBCarts().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testDTBCartsAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	dtbCartOne := &DTBCart{}
	dtbCartTwo := &DTBCart{}
	if err = randomize.Struct(seed, dtbCartOne, dtbCartDBTypes, false, dtbCartColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DTBCart struct: %s", err)
	}
	if err = randomize.Struct(seed, dtbCartTwo, dtbCartDBTypes, false, dtbCartColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DTBCart struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = dtbCartOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = dtbCartTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := DTBCarts().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testDTBCartsCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	dtbCartOne := &DTBCart{}
	dtbCartTwo := &DTBCart{}
	if err = randomize.Struct(seed, dtbCartOne, dtbCartDBTypes, false, dtbCartColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DTBCart struct: %s", err)
	}
	if err = randomize.Struct(seed, dtbCartTwo, dtbCartDBTypes, false, dtbCartColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DTBCart struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = dtbCartOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = dtbCartTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := DTBCarts().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func dtbCartBeforeInsertHook(ctx context.Context, e boil.ContextExecutor, o *DTBCart) error {
	*o = DTBCart{}
	return nil
}

func dtbCartAfterInsertHook(ctx context.Context, e boil.ContextExecutor, o *DTBCart) error {
	*o = DTBCart{}
	return nil
}

func dtbCartAfterSelectHook(ctx context.Context, e boil.ContextExecutor, o *DTBCart) error {
	*o = DTBCart{}
	return nil
}

func dtbCartBeforeUpdateHook(ctx context.Context, e boil.ContextExecutor, o *DTBCart) error {
	*o = DTBCart{}
	return nil
}

func dtbCartAfterUpdateHook(ctx context.Context, e boil.ContextExecutor, o *DTBCart) error {
	*o = DTBCart{}
	return nil
}

func dtbCartBeforeDeleteHook(ctx context.Context, e boil.ContextExecutor, o *DTBCart) error {
	*o = DTBCart{}
	return nil
}

func dtbCartAfterDeleteHook(ctx context.Context, e boil.ContextExecutor, o *DTBCart) error {
	*o = DTBCart{}
	return nil
}

func dtbCartBeforeUpsertHook(ctx context.Context, e boil.ContextExecutor, o *DTBCart) error {
	*o = DTBCart{}
	return nil
}

func dtbCartAfterUpsertHook(ctx context.Context, e boil.ContextExecutor, o *DTBCart) error {
	*o = DTBCart{}
	return nil
}

func testDTBCartsHooks(t *testing.T) {
	t.Parallel()

	var err error

	ctx := context.Background()
	empty := &DTBCart{}
	o := &DTBCart{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, dtbCartDBTypes, false); err != nil {
		t.Errorf("Unable to randomize DTBCart object: %s", err)
	}

	AddDTBCartHook(boil.BeforeInsertHook, dtbCartBeforeInsertHook)
	if err = o.doBeforeInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	dtbCartBeforeInsertHooks = []DTBCartHook{}

	AddDTBCartHook(boil.AfterInsertHook, dtbCartAfterInsertHook)
	if err = o.doAfterInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	dtbCartAfterInsertHooks = []DTBCartHook{}

	AddDTBCartHook(boil.AfterSelectHook, dtbCartAfterSelectHook)
	if err = o.doAfterSelectHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	dtbCartAfterSelectHooks = []DTBCartHook{}

	AddDTBCartHook(boil.BeforeUpdateHook, dtbCartBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	dtbCartBeforeUpdateHooks = []DTBCartHook{}

	AddDTBCartHook(boil.AfterUpdateHook, dtbCartAfterUpdateHook)
	if err = o.doAfterUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	dtbCartAfterUpdateHooks = []DTBCartHook{}

	AddDTBCartHook(boil.BeforeDeleteHook, dtbCartBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	dtbCartBeforeDeleteHooks = []DTBCartHook{}

	AddDTBCartHook(boil.AfterDeleteHook, dtbCartAfterDeleteHook)
	if err = o.doAfterDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	dtbCartAfterDeleteHooks = []DTBCartHook{}

	AddDTBCartHook(boil.BeforeUpsertHook, dtbCartBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	dtbCartBeforeUpsertHooks = []DTBCartHook{}

	AddDTBCartHook(boil.AfterUpsertHook, dtbCartAfterUpsertHook)
	if err = o.doAfterUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	dtbCartAfterUpsertHooks = []DTBCartHook{}
}

func testDTBCartsInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &DTBCart{}
	if err = randomize.Struct(seed, o, dtbCartDBTypes, true, dtbCartColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DTBCart struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := DTBCarts().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testDTBCartsInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &DTBCart{}
	if err = randomize.Struct(seed, o, dtbCartDBTypes, true); err != nil {
		t.Errorf("Unable to randomize DTBCart struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(dtbCartColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := DTBCarts().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testDTBCartToManyCartDTBCartItems(t *testing.T) {
	var err error
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a DTBCart
	var b, c DTBCartItem

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, dtbCartDBTypes, true, dtbCartColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DTBCart struct: %s", err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	if err = randomize.Struct(seed, &b, dtbCartItemDBTypes, false, dtbCartItemColumnsWithDefault...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, dtbCartItemDBTypes, false, dtbCartItemColumnsWithDefault...); err != nil {
		t.Fatal(err)
	}

	queries.Assign(&b.CartID, a.ID)
	queries.Assign(&c.CartID, a.ID)
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := a.CartDTBCartItems().All(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	bFound, cFound := false, false
	for _, v := range check {
		if queries.Equal(v.CartID, b.CartID) {
			bFound = true
		}
		if queries.Equal(v.CartID, c.CartID) {
			cFound = true
		}
	}

	if !bFound {
		t.Error("expected to find b")
	}
	if !cFound {
		t.Error("expected to find c")
	}

	slice := DTBCartSlice{&a}
	if err = a.L.LoadCartDTBCartItems(ctx, tx, false, (*[]*DTBCart)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.CartDTBCartItems); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	a.R.CartDTBCartItems = nil
	if err = a.L.LoadCartDTBCartItems(ctx, tx, true, &a, nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.CartDTBCartItems); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	if t.Failed() {
		t.Logf("%#v", check)
	}
}

func testDTBCartToManyAddOpCartDTBCartItems(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a DTBCart
	var b, c, d, e DTBCartItem

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, dtbCartDBTypes, false, strmangle.SetComplement(dtbCartPrimaryKeyColumns, dtbCartColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*DTBCartItem{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, dtbCartItemDBTypes, false, strmangle.SetComplement(dtbCartItemPrimaryKeyColumns, dtbCartItemColumnsWithoutDefault)...); err != nil {
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

	foreignersSplitByInsertion := [][]*DTBCartItem{
		{&b, &c},
		{&d, &e},
	}

	for i, x := range foreignersSplitByInsertion {
		err = a.AddCartDTBCartItems(ctx, tx, i != 0, x...)
		if err != nil {
			t.Fatal(err)
		}

		first := x[0]
		second := x[1]

		if !queries.Equal(a.ID, first.CartID) {
			t.Error("foreign key was wrong value", a.ID, first.CartID)
		}
		if !queries.Equal(a.ID, second.CartID) {
			t.Error("foreign key was wrong value", a.ID, second.CartID)
		}

		if first.R.Cart != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}
		if second.R.Cart != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}

		if a.R.CartDTBCartItems[i*2] != first {
			t.Error("relationship struct slice not set to correct value")
		}
		if a.R.CartDTBCartItems[i*2+1] != second {
			t.Error("relationship struct slice not set to correct value")
		}

		count, err := a.CartDTBCartItems().Count(ctx, tx)
		if err != nil {
			t.Fatal(err)
		}
		if want := int64((i + 1) * 2); count != want {
			t.Error("want", want, "got", count)
		}
	}
}

func testDTBCartToManySetOpCartDTBCartItems(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a DTBCart
	var b, c, d, e DTBCartItem

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, dtbCartDBTypes, false, strmangle.SetComplement(dtbCartPrimaryKeyColumns, dtbCartColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*DTBCartItem{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, dtbCartItemDBTypes, false, strmangle.SetComplement(dtbCartItemPrimaryKeyColumns, dtbCartItemColumnsWithoutDefault)...); err != nil {
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

	err = a.SetCartDTBCartItems(ctx, tx, false, &b, &c)
	if err != nil {
		t.Fatal(err)
	}

	count, err := a.CartDTBCartItems().Count(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}
	if count != 2 {
		t.Error("count was wrong:", count)
	}

	err = a.SetCartDTBCartItems(ctx, tx, true, &d, &e)
	if err != nil {
		t.Fatal(err)
	}

	count, err = a.CartDTBCartItems().Count(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}
	if count != 2 {
		t.Error("count was wrong:", count)
	}

	if !queries.IsValuerNil(b.CartID) {
		t.Error("want b's foreign key value to be nil")
	}
	if !queries.IsValuerNil(c.CartID) {
		t.Error("want c's foreign key value to be nil")
	}
	if !queries.Equal(a.ID, d.CartID) {
		t.Error("foreign key was wrong value", a.ID, d.CartID)
	}
	if !queries.Equal(a.ID, e.CartID) {
		t.Error("foreign key was wrong value", a.ID, e.CartID)
	}

	if b.R.Cart != nil {
		t.Error("relationship was not removed properly from the foreign struct")
	}
	if c.R.Cart != nil {
		t.Error("relationship was not removed properly from the foreign struct")
	}
	if d.R.Cart != &a {
		t.Error("relationship was not added properly to the foreign struct")
	}
	if e.R.Cart != &a {
		t.Error("relationship was not added properly to the foreign struct")
	}

	if a.R.CartDTBCartItems[0] != &d {
		t.Error("relationship struct slice not set to correct value")
	}
	if a.R.CartDTBCartItems[1] != &e {
		t.Error("relationship struct slice not set to correct value")
	}
}

func testDTBCartToManyRemoveOpCartDTBCartItems(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a DTBCart
	var b, c, d, e DTBCartItem

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, dtbCartDBTypes, false, strmangle.SetComplement(dtbCartPrimaryKeyColumns, dtbCartColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*DTBCartItem{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, dtbCartItemDBTypes, false, strmangle.SetComplement(dtbCartItemPrimaryKeyColumns, dtbCartItemColumnsWithoutDefault)...); err != nil {
			t.Fatal(err)
		}
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	err = a.AddCartDTBCartItems(ctx, tx, true, foreigners...)
	if err != nil {
		t.Fatal(err)
	}

	count, err := a.CartDTBCartItems().Count(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}
	if count != 4 {
		t.Error("count was wrong:", count)
	}

	err = a.RemoveCartDTBCartItems(ctx, tx, foreigners[:2]...)
	if err != nil {
		t.Fatal(err)
	}

	count, err = a.CartDTBCartItems().Count(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}
	if count != 2 {
		t.Error("count was wrong:", count)
	}

	if !queries.IsValuerNil(b.CartID) {
		t.Error("want b's foreign key value to be nil")
	}
	if !queries.IsValuerNil(c.CartID) {
		t.Error("want c's foreign key value to be nil")
	}

	if b.R.Cart != nil {
		t.Error("relationship was not removed properly from the foreign struct")
	}
	if c.R.Cart != nil {
		t.Error("relationship was not removed properly from the foreign struct")
	}
	if d.R.Cart != &a {
		t.Error("relationship to a should have been preserved")
	}
	if e.R.Cart != &a {
		t.Error("relationship to a should have been preserved")
	}

	if len(a.R.CartDTBCartItems) != 2 {
		t.Error("should have preserved two relationships")
	}

	// Removal doesn't do a stable deletion for performance so we have to flip the order
	if a.R.CartDTBCartItems[1] != &d {
		t.Error("relationship to d should have been preserved")
	}
	if a.R.CartDTBCartItems[0] != &e {
		t.Error("relationship to e should have been preserved")
	}
}

func testDTBCartToOneDTBCustomerUsingCustomer(t *testing.T) {
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var local DTBCart
	var foreign DTBCustomer

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, dtbCartDBTypes, true, dtbCartColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DTBCart struct: %s", err)
	}
	if err := randomize.Struct(seed, &foreign, dtbCustomerDBTypes, false, dtbCustomerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DTBCustomer struct: %s", err)
	}

	if err := foreign.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	queries.Assign(&local.CustomerID, foreign.ID)
	if err := local.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := local.Customer().One(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	if !queries.Equal(check.ID, foreign.ID) {
		t.Errorf("want: %v, got %v", foreign.ID, check.ID)
	}

	slice := DTBCartSlice{&local}
	if err = local.L.LoadCustomer(ctx, tx, false, (*[]*DTBCart)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if local.R.Customer == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.Customer = nil
	if err = local.L.LoadCustomer(ctx, tx, true, &local, nil); err != nil {
		t.Fatal(err)
	}
	if local.R.Customer == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testDTBCartToOneSetOpDTBCustomerUsingCustomer(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a DTBCart
	var b, c DTBCustomer

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, dtbCartDBTypes, false, strmangle.SetComplement(dtbCartPrimaryKeyColumns, dtbCartColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, dtbCustomerDBTypes, false, strmangle.SetComplement(dtbCustomerPrimaryKeyColumns, dtbCustomerColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, dtbCustomerDBTypes, false, strmangle.SetComplement(dtbCustomerPrimaryKeyColumns, dtbCustomerColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*DTBCustomer{&b, &c} {
		err = a.SetCustomer(ctx, tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.Customer != x {
			t.Error("relationship struct not set to correct value")
		}

		if x.R.CustomerDTBCarts[0] != &a {
			t.Error("failed to append to foreign relationship struct")
		}
		if !queries.Equal(a.CustomerID, x.ID) {
			t.Error("foreign key was wrong value", a.CustomerID)
		}

		zero := reflect.Zero(reflect.TypeOf(a.CustomerID))
		reflect.Indirect(reflect.ValueOf(&a.CustomerID)).Set(zero)

		if err = a.Reload(ctx, tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if !queries.Equal(a.CustomerID, x.ID) {
			t.Error("foreign key was wrong value", a.CustomerID, x.ID)
		}
	}
}

func testDTBCartToOneRemoveOpDTBCustomerUsingCustomer(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a DTBCart
	var b DTBCustomer

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, dtbCartDBTypes, false, strmangle.SetComplement(dtbCartPrimaryKeyColumns, dtbCartColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, dtbCustomerDBTypes, false, strmangle.SetComplement(dtbCustomerPrimaryKeyColumns, dtbCustomerColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err = a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	if err = a.SetCustomer(ctx, tx, true, &b); err != nil {
		t.Fatal(err)
	}

	if err = a.RemoveCustomer(ctx, tx, &b); err != nil {
		t.Error("failed to remove relationship")
	}

	count, err := a.Customer().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 0 {
		t.Error("want no relationships remaining")
	}

	if a.R.Customer != nil {
		t.Error("R struct entry should be nil")
	}

	if !queries.IsValuerNil(a.CustomerID) {
		t.Error("foreign key value should be nil")
	}

	if len(b.R.CustomerDTBCarts) != 0 {
		t.Error("failed to remove a from b's relationships")
	}
}

func testDTBCartsReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &DTBCart{}
	if err = randomize.Struct(seed, o, dtbCartDBTypes, true, dtbCartColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DTBCart struct: %s", err)
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

func testDTBCartsReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &DTBCart{}
	if err = randomize.Struct(seed, o, dtbCartDBTypes, true, dtbCartColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DTBCart struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := DTBCartSlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testDTBCartsSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &DTBCart{}
	if err = randomize.Struct(seed, o, dtbCartDBTypes, true, dtbCartColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DTBCart struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := DTBCarts().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	dtbCartDBTypes = map[string]string{`ID`: `int`, `CustomerID`: `int`, `CartKey`: `varchar`, `PreOrderID`: `varchar`, `TotalPrice`: `decimal`, `DeliveryFeeTotal`: `decimal`, `SortNo`: `smallint`, `CreateDate`: `datetime`, `UpdateDate`: `datetime`, `AddPoint`: `decimal`, `UsePoint`: `decimal`, `DiscriminatorType`: `varchar`}
	_              = bytes.MinRead
)

func testDTBCartsUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(dtbCartPrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(dtbCartColumns) == len(dtbCartPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &DTBCart{}
	if err = randomize.Struct(seed, o, dtbCartDBTypes, true, dtbCartColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DTBCart struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := DTBCarts().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, dtbCartDBTypes, true, dtbCartPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize DTBCart struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testDTBCartsSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(dtbCartColumns) == len(dtbCartPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &DTBCart{}
	if err = randomize.Struct(seed, o, dtbCartDBTypes, true, dtbCartColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DTBCart struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := DTBCarts().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, dtbCartDBTypes, true, dtbCartPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize DTBCart struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(dtbCartColumns, dtbCartPrimaryKeyColumns) {
		fields = dtbCartColumns
	} else {
		fields = strmangle.SetComplement(
			dtbCartColumns,
			dtbCartPrimaryKeyColumns,
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

	slice := DTBCartSlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testDTBCartsUpsert(t *testing.T) {
	t.Parallel()

	if len(dtbCartColumns) == len(dtbCartPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}
	if len(mySQLDTBCartUniqueColumns) == 0 {
		t.Skip("Skipping table with no unique columns to conflict on")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := DTBCart{}
	if err = randomize.Struct(seed, &o, dtbCartDBTypes, false); err != nil {
		t.Errorf("Unable to randomize DTBCart struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(ctx, tx, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert DTBCart: %s", err)
	}

	count, err := DTBCarts().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, dtbCartDBTypes, false, dtbCartPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize DTBCart struct: %s", err)
	}

	if err = o.Upsert(ctx, tx, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert DTBCart: %s", err)
	}

	count, err = DTBCarts().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
