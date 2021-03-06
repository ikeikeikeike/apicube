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

func testDTBTags(t *testing.T) {
	t.Parallel()

	query := DTBTags()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testDTBTagsDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &DTBTag{}
	if err = randomize.Struct(seed, o, dtbTagDBTypes, true, dtbTagColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DTBTag struct: %s", err)
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

	count, err := DTBTags().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testDTBTagsQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &DTBTag{}
	if err = randomize.Struct(seed, o, dtbTagDBTypes, true, dtbTagColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DTBTag struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := DTBTags().DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := DTBTags().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testDTBTagsSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &DTBTag{}
	if err = randomize.Struct(seed, o, dtbTagDBTypes, true, dtbTagColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DTBTag struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := DTBTagSlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := DTBTags().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testDTBTagsExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &DTBTag{}
	if err = randomize.Struct(seed, o, dtbTagDBTypes, true, dtbTagColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DTBTag struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := DTBTagExists(ctx, tx, o.ID)
	if err != nil {
		t.Errorf("Unable to check if DTBTag exists: %s", err)
	}
	if !e {
		t.Errorf("Expected DTBTagExists to return true, but got false.")
	}
}

func testDTBTagsFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &DTBTag{}
	if err = randomize.Struct(seed, o, dtbTagDBTypes, true, dtbTagColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DTBTag struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	dtbTagFound, err := FindDTBTag(ctx, tx, o.ID)
	if err != nil {
		t.Error(err)
	}

	if dtbTagFound == nil {
		t.Error("want a record, got nil")
	}
}

func testDTBTagsBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &DTBTag{}
	if err = randomize.Struct(seed, o, dtbTagDBTypes, true, dtbTagColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DTBTag struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = DTBTags().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testDTBTagsOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &DTBTag{}
	if err = randomize.Struct(seed, o, dtbTagDBTypes, true, dtbTagColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DTBTag struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := DTBTags().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testDTBTagsAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	dtbTagOne := &DTBTag{}
	dtbTagTwo := &DTBTag{}
	if err = randomize.Struct(seed, dtbTagOne, dtbTagDBTypes, false, dtbTagColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DTBTag struct: %s", err)
	}
	if err = randomize.Struct(seed, dtbTagTwo, dtbTagDBTypes, false, dtbTagColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DTBTag struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = dtbTagOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = dtbTagTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := DTBTags().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testDTBTagsCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	dtbTagOne := &DTBTag{}
	dtbTagTwo := &DTBTag{}
	if err = randomize.Struct(seed, dtbTagOne, dtbTagDBTypes, false, dtbTagColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DTBTag struct: %s", err)
	}
	if err = randomize.Struct(seed, dtbTagTwo, dtbTagDBTypes, false, dtbTagColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DTBTag struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = dtbTagOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = dtbTagTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := DTBTags().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func dtbTagBeforeInsertHook(ctx context.Context, e boil.ContextExecutor, o *DTBTag) error {
	*o = DTBTag{}
	return nil
}

func dtbTagAfterInsertHook(ctx context.Context, e boil.ContextExecutor, o *DTBTag) error {
	*o = DTBTag{}
	return nil
}

func dtbTagAfterSelectHook(ctx context.Context, e boil.ContextExecutor, o *DTBTag) error {
	*o = DTBTag{}
	return nil
}

func dtbTagBeforeUpdateHook(ctx context.Context, e boil.ContextExecutor, o *DTBTag) error {
	*o = DTBTag{}
	return nil
}

func dtbTagAfterUpdateHook(ctx context.Context, e boil.ContextExecutor, o *DTBTag) error {
	*o = DTBTag{}
	return nil
}

func dtbTagBeforeDeleteHook(ctx context.Context, e boil.ContextExecutor, o *DTBTag) error {
	*o = DTBTag{}
	return nil
}

func dtbTagAfterDeleteHook(ctx context.Context, e boil.ContextExecutor, o *DTBTag) error {
	*o = DTBTag{}
	return nil
}

func dtbTagBeforeUpsertHook(ctx context.Context, e boil.ContextExecutor, o *DTBTag) error {
	*o = DTBTag{}
	return nil
}

func dtbTagAfterUpsertHook(ctx context.Context, e boil.ContextExecutor, o *DTBTag) error {
	*o = DTBTag{}
	return nil
}

func testDTBTagsHooks(t *testing.T) {
	t.Parallel()

	var err error

	ctx := context.Background()
	empty := &DTBTag{}
	o := &DTBTag{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, dtbTagDBTypes, false); err != nil {
		t.Errorf("Unable to randomize DTBTag object: %s", err)
	}

	AddDTBTagHook(boil.BeforeInsertHook, dtbTagBeforeInsertHook)
	if err = o.doBeforeInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	dtbTagBeforeInsertHooks = []DTBTagHook{}

	AddDTBTagHook(boil.AfterInsertHook, dtbTagAfterInsertHook)
	if err = o.doAfterInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	dtbTagAfterInsertHooks = []DTBTagHook{}

	AddDTBTagHook(boil.AfterSelectHook, dtbTagAfterSelectHook)
	if err = o.doAfterSelectHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	dtbTagAfterSelectHooks = []DTBTagHook{}

	AddDTBTagHook(boil.BeforeUpdateHook, dtbTagBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	dtbTagBeforeUpdateHooks = []DTBTagHook{}

	AddDTBTagHook(boil.AfterUpdateHook, dtbTagAfterUpdateHook)
	if err = o.doAfterUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	dtbTagAfterUpdateHooks = []DTBTagHook{}

	AddDTBTagHook(boil.BeforeDeleteHook, dtbTagBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	dtbTagBeforeDeleteHooks = []DTBTagHook{}

	AddDTBTagHook(boil.AfterDeleteHook, dtbTagAfterDeleteHook)
	if err = o.doAfterDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	dtbTagAfterDeleteHooks = []DTBTagHook{}

	AddDTBTagHook(boil.BeforeUpsertHook, dtbTagBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	dtbTagBeforeUpsertHooks = []DTBTagHook{}

	AddDTBTagHook(boil.AfterUpsertHook, dtbTagAfterUpsertHook)
	if err = o.doAfterUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	dtbTagAfterUpsertHooks = []DTBTagHook{}
}

func testDTBTagsInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &DTBTag{}
	if err = randomize.Struct(seed, o, dtbTagDBTypes, true, dtbTagColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DTBTag struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := DTBTags().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testDTBTagsInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &DTBTag{}
	if err = randomize.Struct(seed, o, dtbTagDBTypes, true); err != nil {
		t.Errorf("Unable to randomize DTBTag struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(dtbTagColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := DTBTags().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testDTBTagToManyTagDTBProductTags(t *testing.T) {
	var err error
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a DTBTag
	var b, c DTBProductTag

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, dtbTagDBTypes, true, dtbTagColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DTBTag struct: %s", err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	if err = randomize.Struct(seed, &b, dtbProductTagDBTypes, false, dtbProductTagColumnsWithDefault...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, dtbProductTagDBTypes, false, dtbProductTagColumnsWithDefault...); err != nil {
		t.Fatal(err)
	}

	queries.Assign(&b.TagID, a.ID)
	queries.Assign(&c.TagID, a.ID)
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := a.TagDTBProductTags().All(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	bFound, cFound := false, false
	for _, v := range check {
		if queries.Equal(v.TagID, b.TagID) {
			bFound = true
		}
		if queries.Equal(v.TagID, c.TagID) {
			cFound = true
		}
	}

	if !bFound {
		t.Error("expected to find b")
	}
	if !cFound {
		t.Error("expected to find c")
	}

	slice := DTBTagSlice{&a}
	if err = a.L.LoadTagDTBProductTags(ctx, tx, false, (*[]*DTBTag)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.TagDTBProductTags); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	a.R.TagDTBProductTags = nil
	if err = a.L.LoadTagDTBProductTags(ctx, tx, true, &a, nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.TagDTBProductTags); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	if t.Failed() {
		t.Logf("%#v", check)
	}
}

func testDTBTagToManyAddOpTagDTBProductTags(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a DTBTag
	var b, c, d, e DTBProductTag

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, dtbTagDBTypes, false, strmangle.SetComplement(dtbTagPrimaryKeyColumns, dtbTagColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*DTBProductTag{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, dtbProductTagDBTypes, false, strmangle.SetComplement(dtbProductTagPrimaryKeyColumns, dtbProductTagColumnsWithoutDefault)...); err != nil {
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

	foreignersSplitByInsertion := [][]*DTBProductTag{
		{&b, &c},
		{&d, &e},
	}

	for i, x := range foreignersSplitByInsertion {
		err = a.AddTagDTBProductTags(ctx, tx, i != 0, x...)
		if err != nil {
			t.Fatal(err)
		}

		first := x[0]
		second := x[1]

		if !queries.Equal(a.ID, first.TagID) {
			t.Error("foreign key was wrong value", a.ID, first.TagID)
		}
		if !queries.Equal(a.ID, second.TagID) {
			t.Error("foreign key was wrong value", a.ID, second.TagID)
		}

		if first.R.Tag != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}
		if second.R.Tag != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}

		if a.R.TagDTBProductTags[i*2] != first {
			t.Error("relationship struct slice not set to correct value")
		}
		if a.R.TagDTBProductTags[i*2+1] != second {
			t.Error("relationship struct slice not set to correct value")
		}

		count, err := a.TagDTBProductTags().Count(ctx, tx)
		if err != nil {
			t.Fatal(err)
		}
		if want := int64((i + 1) * 2); count != want {
			t.Error("want", want, "got", count)
		}
	}
}

func testDTBTagToManySetOpTagDTBProductTags(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a DTBTag
	var b, c, d, e DTBProductTag

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, dtbTagDBTypes, false, strmangle.SetComplement(dtbTagPrimaryKeyColumns, dtbTagColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*DTBProductTag{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, dtbProductTagDBTypes, false, strmangle.SetComplement(dtbProductTagPrimaryKeyColumns, dtbProductTagColumnsWithoutDefault)...); err != nil {
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

	err = a.SetTagDTBProductTags(ctx, tx, false, &b, &c)
	if err != nil {
		t.Fatal(err)
	}

	count, err := a.TagDTBProductTags().Count(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}
	if count != 2 {
		t.Error("count was wrong:", count)
	}

	err = a.SetTagDTBProductTags(ctx, tx, true, &d, &e)
	if err != nil {
		t.Fatal(err)
	}

	count, err = a.TagDTBProductTags().Count(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}
	if count != 2 {
		t.Error("count was wrong:", count)
	}

	if !queries.IsValuerNil(b.TagID) {
		t.Error("want b's foreign key value to be nil")
	}
	if !queries.IsValuerNil(c.TagID) {
		t.Error("want c's foreign key value to be nil")
	}
	if !queries.Equal(a.ID, d.TagID) {
		t.Error("foreign key was wrong value", a.ID, d.TagID)
	}
	if !queries.Equal(a.ID, e.TagID) {
		t.Error("foreign key was wrong value", a.ID, e.TagID)
	}

	if b.R.Tag != nil {
		t.Error("relationship was not removed properly from the foreign struct")
	}
	if c.R.Tag != nil {
		t.Error("relationship was not removed properly from the foreign struct")
	}
	if d.R.Tag != &a {
		t.Error("relationship was not added properly to the foreign struct")
	}
	if e.R.Tag != &a {
		t.Error("relationship was not added properly to the foreign struct")
	}

	if a.R.TagDTBProductTags[0] != &d {
		t.Error("relationship struct slice not set to correct value")
	}
	if a.R.TagDTBProductTags[1] != &e {
		t.Error("relationship struct slice not set to correct value")
	}
}

func testDTBTagToManyRemoveOpTagDTBProductTags(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a DTBTag
	var b, c, d, e DTBProductTag

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, dtbTagDBTypes, false, strmangle.SetComplement(dtbTagPrimaryKeyColumns, dtbTagColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*DTBProductTag{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, dtbProductTagDBTypes, false, strmangle.SetComplement(dtbProductTagPrimaryKeyColumns, dtbProductTagColumnsWithoutDefault)...); err != nil {
			t.Fatal(err)
		}
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	err = a.AddTagDTBProductTags(ctx, tx, true, foreigners...)
	if err != nil {
		t.Fatal(err)
	}

	count, err := a.TagDTBProductTags().Count(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}
	if count != 4 {
		t.Error("count was wrong:", count)
	}

	err = a.RemoveTagDTBProductTags(ctx, tx, foreigners[:2]...)
	if err != nil {
		t.Fatal(err)
	}

	count, err = a.TagDTBProductTags().Count(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}
	if count != 2 {
		t.Error("count was wrong:", count)
	}

	if !queries.IsValuerNil(b.TagID) {
		t.Error("want b's foreign key value to be nil")
	}
	if !queries.IsValuerNil(c.TagID) {
		t.Error("want c's foreign key value to be nil")
	}

	if b.R.Tag != nil {
		t.Error("relationship was not removed properly from the foreign struct")
	}
	if c.R.Tag != nil {
		t.Error("relationship was not removed properly from the foreign struct")
	}
	if d.R.Tag != &a {
		t.Error("relationship to a should have been preserved")
	}
	if e.R.Tag != &a {
		t.Error("relationship to a should have been preserved")
	}

	if len(a.R.TagDTBProductTags) != 2 {
		t.Error("should have preserved two relationships")
	}

	// Removal doesn't do a stable deletion for performance so we have to flip the order
	if a.R.TagDTBProductTags[1] != &d {
		t.Error("relationship to d should have been preserved")
	}
	if a.R.TagDTBProductTags[0] != &e {
		t.Error("relationship to e should have been preserved")
	}
}

func testDTBTagsReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &DTBTag{}
	if err = randomize.Struct(seed, o, dtbTagDBTypes, true, dtbTagColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DTBTag struct: %s", err)
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

func testDTBTagsReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &DTBTag{}
	if err = randomize.Struct(seed, o, dtbTagDBTypes, true, dtbTagColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DTBTag struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := DTBTagSlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testDTBTagsSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &DTBTag{}
	if err = randomize.Struct(seed, o, dtbTagDBTypes, true, dtbTagColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DTBTag struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := DTBTags().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	dtbTagDBTypes = map[string]string{`ID`: `int`, `Name`: `varchar`, `SortNo`: `smallint`, `DiscriminatorType`: `varchar`}
	_             = bytes.MinRead
)

func testDTBTagsUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(dtbTagPrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(dtbTagColumns) == len(dtbTagPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &DTBTag{}
	if err = randomize.Struct(seed, o, dtbTagDBTypes, true, dtbTagColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DTBTag struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := DTBTags().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, dtbTagDBTypes, true, dtbTagPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize DTBTag struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testDTBTagsSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(dtbTagColumns) == len(dtbTagPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &DTBTag{}
	if err = randomize.Struct(seed, o, dtbTagDBTypes, true, dtbTagColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DTBTag struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := DTBTags().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, dtbTagDBTypes, true, dtbTagPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize DTBTag struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(dtbTagColumns, dtbTagPrimaryKeyColumns) {
		fields = dtbTagColumns
	} else {
		fields = strmangle.SetComplement(
			dtbTagColumns,
			dtbTagPrimaryKeyColumns,
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

	slice := DTBTagSlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testDTBTagsUpsert(t *testing.T) {
	t.Parallel()

	if len(dtbTagColumns) == len(dtbTagPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}
	if len(mySQLDTBTagUniqueColumns) == 0 {
		t.Skip("Skipping table with no unique columns to conflict on")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := DTBTag{}
	if err = randomize.Struct(seed, &o, dtbTagDBTypes, false); err != nil {
		t.Errorf("Unable to randomize DTBTag struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(ctx, tx, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert DTBTag: %s", err)
	}

	count, err := DTBTags().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, dtbTagDBTypes, false, dtbTagPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize DTBTag struct: %s", err)
	}

	if err = o.Upsert(ctx, tx, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert DTBTag: %s", err)
	}

	count, err = DTBTags().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
