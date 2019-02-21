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

func testDTBProductTags(t *testing.T) {
	t.Parallel()

	query := DTBProductTags()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testDTBProductTagsDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &DTBProductTag{}
	if err = randomize.Struct(seed, o, dtbProductTagDBTypes, true, dtbProductTagColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DTBProductTag struct: %s", err)
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

	count, err := DTBProductTags().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testDTBProductTagsQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &DTBProductTag{}
	if err = randomize.Struct(seed, o, dtbProductTagDBTypes, true, dtbProductTagColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DTBProductTag struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := DTBProductTags().DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := DTBProductTags().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testDTBProductTagsSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &DTBProductTag{}
	if err = randomize.Struct(seed, o, dtbProductTagDBTypes, true, dtbProductTagColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DTBProductTag struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := DTBProductTagSlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := DTBProductTags().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testDTBProductTagsExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &DTBProductTag{}
	if err = randomize.Struct(seed, o, dtbProductTagDBTypes, true, dtbProductTagColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DTBProductTag struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := DTBProductTagExists(ctx, tx, o.ID)
	if err != nil {
		t.Errorf("Unable to check if DTBProductTag exists: %s", err)
	}
	if !e {
		t.Errorf("Expected DTBProductTagExists to return true, but got false.")
	}
}

func testDTBProductTagsFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &DTBProductTag{}
	if err = randomize.Struct(seed, o, dtbProductTagDBTypes, true, dtbProductTagColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DTBProductTag struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	dtbProductTagFound, err := FindDTBProductTag(ctx, tx, o.ID)
	if err != nil {
		t.Error(err)
	}

	if dtbProductTagFound == nil {
		t.Error("want a record, got nil")
	}
}

func testDTBProductTagsBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &DTBProductTag{}
	if err = randomize.Struct(seed, o, dtbProductTagDBTypes, true, dtbProductTagColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DTBProductTag struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = DTBProductTags().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testDTBProductTagsOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &DTBProductTag{}
	if err = randomize.Struct(seed, o, dtbProductTagDBTypes, true, dtbProductTagColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DTBProductTag struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := DTBProductTags().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testDTBProductTagsAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	dtbProductTagOne := &DTBProductTag{}
	dtbProductTagTwo := &DTBProductTag{}
	if err = randomize.Struct(seed, dtbProductTagOne, dtbProductTagDBTypes, false, dtbProductTagColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DTBProductTag struct: %s", err)
	}
	if err = randomize.Struct(seed, dtbProductTagTwo, dtbProductTagDBTypes, false, dtbProductTagColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DTBProductTag struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = dtbProductTagOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = dtbProductTagTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := DTBProductTags().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testDTBProductTagsCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	dtbProductTagOne := &DTBProductTag{}
	dtbProductTagTwo := &DTBProductTag{}
	if err = randomize.Struct(seed, dtbProductTagOne, dtbProductTagDBTypes, false, dtbProductTagColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DTBProductTag struct: %s", err)
	}
	if err = randomize.Struct(seed, dtbProductTagTwo, dtbProductTagDBTypes, false, dtbProductTagColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DTBProductTag struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = dtbProductTagOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = dtbProductTagTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := DTBProductTags().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func dtbProductTagBeforeInsertHook(ctx context.Context, e boil.ContextExecutor, o *DTBProductTag) error {
	*o = DTBProductTag{}
	return nil
}

func dtbProductTagAfterInsertHook(ctx context.Context, e boil.ContextExecutor, o *DTBProductTag) error {
	*o = DTBProductTag{}
	return nil
}

func dtbProductTagAfterSelectHook(ctx context.Context, e boil.ContextExecutor, o *DTBProductTag) error {
	*o = DTBProductTag{}
	return nil
}

func dtbProductTagBeforeUpdateHook(ctx context.Context, e boil.ContextExecutor, o *DTBProductTag) error {
	*o = DTBProductTag{}
	return nil
}

func dtbProductTagAfterUpdateHook(ctx context.Context, e boil.ContextExecutor, o *DTBProductTag) error {
	*o = DTBProductTag{}
	return nil
}

func dtbProductTagBeforeDeleteHook(ctx context.Context, e boil.ContextExecutor, o *DTBProductTag) error {
	*o = DTBProductTag{}
	return nil
}

func dtbProductTagAfterDeleteHook(ctx context.Context, e boil.ContextExecutor, o *DTBProductTag) error {
	*o = DTBProductTag{}
	return nil
}

func dtbProductTagBeforeUpsertHook(ctx context.Context, e boil.ContextExecutor, o *DTBProductTag) error {
	*o = DTBProductTag{}
	return nil
}

func dtbProductTagAfterUpsertHook(ctx context.Context, e boil.ContextExecutor, o *DTBProductTag) error {
	*o = DTBProductTag{}
	return nil
}

func testDTBProductTagsHooks(t *testing.T) {
	t.Parallel()

	var err error

	ctx := context.Background()
	empty := &DTBProductTag{}
	o := &DTBProductTag{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, dtbProductTagDBTypes, false); err != nil {
		t.Errorf("Unable to randomize DTBProductTag object: %s", err)
	}

	AddDTBProductTagHook(boil.BeforeInsertHook, dtbProductTagBeforeInsertHook)
	if err = o.doBeforeInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	dtbProductTagBeforeInsertHooks = []DTBProductTagHook{}

	AddDTBProductTagHook(boil.AfterInsertHook, dtbProductTagAfterInsertHook)
	if err = o.doAfterInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	dtbProductTagAfterInsertHooks = []DTBProductTagHook{}

	AddDTBProductTagHook(boil.AfterSelectHook, dtbProductTagAfterSelectHook)
	if err = o.doAfterSelectHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	dtbProductTagAfterSelectHooks = []DTBProductTagHook{}

	AddDTBProductTagHook(boil.BeforeUpdateHook, dtbProductTagBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	dtbProductTagBeforeUpdateHooks = []DTBProductTagHook{}

	AddDTBProductTagHook(boil.AfterUpdateHook, dtbProductTagAfterUpdateHook)
	if err = o.doAfterUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	dtbProductTagAfterUpdateHooks = []DTBProductTagHook{}

	AddDTBProductTagHook(boil.BeforeDeleteHook, dtbProductTagBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	dtbProductTagBeforeDeleteHooks = []DTBProductTagHook{}

	AddDTBProductTagHook(boil.AfterDeleteHook, dtbProductTagAfterDeleteHook)
	if err = o.doAfterDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	dtbProductTagAfterDeleteHooks = []DTBProductTagHook{}

	AddDTBProductTagHook(boil.BeforeUpsertHook, dtbProductTagBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	dtbProductTagBeforeUpsertHooks = []DTBProductTagHook{}

	AddDTBProductTagHook(boil.AfterUpsertHook, dtbProductTagAfterUpsertHook)
	if err = o.doAfterUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	dtbProductTagAfterUpsertHooks = []DTBProductTagHook{}
}

func testDTBProductTagsInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &DTBProductTag{}
	if err = randomize.Struct(seed, o, dtbProductTagDBTypes, true, dtbProductTagColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DTBProductTag struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := DTBProductTags().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testDTBProductTagsInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &DTBProductTag{}
	if err = randomize.Struct(seed, o, dtbProductTagDBTypes, true); err != nil {
		t.Errorf("Unable to randomize DTBProductTag struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(dtbProductTagColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := DTBProductTags().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testDTBProductTagToOneDTBProductUsingProduct(t *testing.T) {
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var local DTBProductTag
	var foreign DTBProduct

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, dtbProductTagDBTypes, true, dtbProductTagColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DTBProductTag struct: %s", err)
	}
	if err := randomize.Struct(seed, &foreign, dtbProductDBTypes, false, dtbProductColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DTBProduct struct: %s", err)
	}

	if err := foreign.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	queries.Assign(&local.ProductID, foreign.ID)
	if err := local.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := local.Product().One(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	if !queries.Equal(check.ID, foreign.ID) {
		t.Errorf("want: %v, got %v", foreign.ID, check.ID)
	}

	slice := DTBProductTagSlice{&local}
	if err = local.L.LoadProduct(ctx, tx, false, (*[]*DTBProductTag)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if local.R.Product == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.Product = nil
	if err = local.L.LoadProduct(ctx, tx, true, &local, nil); err != nil {
		t.Fatal(err)
	}
	if local.R.Product == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testDTBProductTagToOneDTBTagUsingTag(t *testing.T) {
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var local DTBProductTag
	var foreign DTBTag

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, dtbProductTagDBTypes, true, dtbProductTagColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DTBProductTag struct: %s", err)
	}
	if err := randomize.Struct(seed, &foreign, dtbTagDBTypes, false, dtbTagColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DTBTag struct: %s", err)
	}

	if err := foreign.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	queries.Assign(&local.TagID, foreign.ID)
	if err := local.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := local.Tag().One(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	if !queries.Equal(check.ID, foreign.ID) {
		t.Errorf("want: %v, got %v", foreign.ID, check.ID)
	}

	slice := DTBProductTagSlice{&local}
	if err = local.L.LoadTag(ctx, tx, false, (*[]*DTBProductTag)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if local.R.Tag == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.Tag = nil
	if err = local.L.LoadTag(ctx, tx, true, &local, nil); err != nil {
		t.Fatal(err)
	}
	if local.R.Tag == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testDTBProductTagToOneSetOpDTBProductUsingProduct(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a DTBProductTag
	var b, c DTBProduct

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, dtbProductTagDBTypes, false, strmangle.SetComplement(dtbProductTagPrimaryKeyColumns, dtbProductTagColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, dtbProductDBTypes, false, strmangle.SetComplement(dtbProductPrimaryKeyColumns, dtbProductColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, dtbProductDBTypes, false, strmangle.SetComplement(dtbProductPrimaryKeyColumns, dtbProductColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*DTBProduct{&b, &c} {
		err = a.SetProduct(ctx, tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.Product != x {
			t.Error("relationship struct not set to correct value")
		}

		if x.R.ProductDTBProductTags[0] != &a {
			t.Error("failed to append to foreign relationship struct")
		}
		if !queries.Equal(a.ProductID, x.ID) {
			t.Error("foreign key was wrong value", a.ProductID)
		}

		zero := reflect.Zero(reflect.TypeOf(a.ProductID))
		reflect.Indirect(reflect.ValueOf(&a.ProductID)).Set(zero)

		if err = a.Reload(ctx, tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if !queries.Equal(a.ProductID, x.ID) {
			t.Error("foreign key was wrong value", a.ProductID, x.ID)
		}
	}
}

func testDTBProductTagToOneRemoveOpDTBProductUsingProduct(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a DTBProductTag
	var b DTBProduct

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, dtbProductTagDBTypes, false, strmangle.SetComplement(dtbProductTagPrimaryKeyColumns, dtbProductTagColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, dtbProductDBTypes, false, strmangle.SetComplement(dtbProductPrimaryKeyColumns, dtbProductColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err = a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	if err = a.SetProduct(ctx, tx, true, &b); err != nil {
		t.Fatal(err)
	}

	if err = a.RemoveProduct(ctx, tx, &b); err != nil {
		t.Error("failed to remove relationship")
	}

	count, err := a.Product().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 0 {
		t.Error("want no relationships remaining")
	}

	if a.R.Product != nil {
		t.Error("R struct entry should be nil")
	}

	if !queries.IsValuerNil(a.ProductID) {
		t.Error("foreign key value should be nil")
	}

	if len(b.R.ProductDTBProductTags) != 0 {
		t.Error("failed to remove a from b's relationships")
	}
}

func testDTBProductTagToOneSetOpDTBTagUsingTag(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a DTBProductTag
	var b, c DTBTag

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, dtbProductTagDBTypes, false, strmangle.SetComplement(dtbProductTagPrimaryKeyColumns, dtbProductTagColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, dtbTagDBTypes, false, strmangle.SetComplement(dtbTagPrimaryKeyColumns, dtbTagColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, dtbTagDBTypes, false, strmangle.SetComplement(dtbTagPrimaryKeyColumns, dtbTagColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*DTBTag{&b, &c} {
		err = a.SetTag(ctx, tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.Tag != x {
			t.Error("relationship struct not set to correct value")
		}

		if x.R.TagDTBProductTags[0] != &a {
			t.Error("failed to append to foreign relationship struct")
		}
		if !queries.Equal(a.TagID, x.ID) {
			t.Error("foreign key was wrong value", a.TagID)
		}

		zero := reflect.Zero(reflect.TypeOf(a.TagID))
		reflect.Indirect(reflect.ValueOf(&a.TagID)).Set(zero)

		if err = a.Reload(ctx, tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if !queries.Equal(a.TagID, x.ID) {
			t.Error("foreign key was wrong value", a.TagID, x.ID)
		}
	}
}

func testDTBProductTagToOneRemoveOpDTBTagUsingTag(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a DTBProductTag
	var b DTBTag

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, dtbProductTagDBTypes, false, strmangle.SetComplement(dtbProductTagPrimaryKeyColumns, dtbProductTagColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, dtbTagDBTypes, false, strmangle.SetComplement(dtbTagPrimaryKeyColumns, dtbTagColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err = a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	if err = a.SetTag(ctx, tx, true, &b); err != nil {
		t.Fatal(err)
	}

	if err = a.RemoveTag(ctx, tx, &b); err != nil {
		t.Error("failed to remove relationship")
	}

	count, err := a.Tag().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 0 {
		t.Error("want no relationships remaining")
	}

	if a.R.Tag != nil {
		t.Error("R struct entry should be nil")
	}

	if !queries.IsValuerNil(a.TagID) {
		t.Error("foreign key value should be nil")
	}

	if len(b.R.TagDTBProductTags) != 0 {
		t.Error("failed to remove a from b's relationships")
	}
}

func testDTBProductTagsReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &DTBProductTag{}
	if err = randomize.Struct(seed, o, dtbProductTagDBTypes, true, dtbProductTagColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DTBProductTag struct: %s", err)
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

func testDTBProductTagsReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &DTBProductTag{}
	if err = randomize.Struct(seed, o, dtbProductTagDBTypes, true, dtbProductTagColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DTBProductTag struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := DTBProductTagSlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testDTBProductTagsSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &DTBProductTag{}
	if err = randomize.Struct(seed, o, dtbProductTagDBTypes, true, dtbProductTagColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DTBProductTag struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := DTBProductTags().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	dtbProductTagDBTypes = map[string]string{`ID`: `int`, `ProductID`: `int`, `TagID`: `int`, `CreatorID`: `int`, `CreateDate`: `datetime`, `DiscriminatorType`: `varchar`}
	_                    = bytes.MinRead
)

func testDTBProductTagsUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(dtbProductTagPrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(dtbProductTagColumns) == len(dtbProductTagPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &DTBProductTag{}
	if err = randomize.Struct(seed, o, dtbProductTagDBTypes, true, dtbProductTagColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DTBProductTag struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := DTBProductTags().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, dtbProductTagDBTypes, true, dtbProductTagPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize DTBProductTag struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testDTBProductTagsSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(dtbProductTagColumns) == len(dtbProductTagPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &DTBProductTag{}
	if err = randomize.Struct(seed, o, dtbProductTagDBTypes, true, dtbProductTagColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DTBProductTag struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := DTBProductTags().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, dtbProductTagDBTypes, true, dtbProductTagPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize DTBProductTag struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(dtbProductTagColumns, dtbProductTagPrimaryKeyColumns) {
		fields = dtbProductTagColumns
	} else {
		fields = strmangle.SetComplement(
			dtbProductTagColumns,
			dtbProductTagPrimaryKeyColumns,
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

	slice := DTBProductTagSlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testDTBProductTagsUpsert(t *testing.T) {
	t.Parallel()

	if len(dtbProductTagColumns) == len(dtbProductTagPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}
	if len(mySQLDTBProductTagUniqueColumns) == 0 {
		t.Skip("Skipping table with no unique columns to conflict on")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := DTBProductTag{}
	if err = randomize.Struct(seed, &o, dtbProductTagDBTypes, false); err != nil {
		t.Errorf("Unable to randomize DTBProductTag struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(ctx, tx, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert DTBProductTag: %s", err)
	}

	count, err := DTBProductTags().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, dtbProductTagDBTypes, false, dtbProductTagPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize DTBProductTag struct: %s", err)
	}

	if err = o.Upsert(ctx, tx, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert DTBProductTag: %s", err)
	}

	count, err = DTBProductTags().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
