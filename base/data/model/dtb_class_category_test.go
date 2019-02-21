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

func testDTBClassCategories(t *testing.T) {
	t.Parallel()

	query := DTBClassCategories()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testDTBClassCategoriesDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &DTBClassCategory{}
	if err = randomize.Struct(seed, o, dtbClassCategoryDBTypes, true, dtbClassCategoryColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DTBClassCategory struct: %s", err)
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

	count, err := DTBClassCategories().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testDTBClassCategoriesQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &DTBClassCategory{}
	if err = randomize.Struct(seed, o, dtbClassCategoryDBTypes, true, dtbClassCategoryColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DTBClassCategory struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := DTBClassCategories().DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := DTBClassCategories().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testDTBClassCategoriesSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &DTBClassCategory{}
	if err = randomize.Struct(seed, o, dtbClassCategoryDBTypes, true, dtbClassCategoryColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DTBClassCategory struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := DTBClassCategorySlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := DTBClassCategories().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testDTBClassCategoriesExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &DTBClassCategory{}
	if err = randomize.Struct(seed, o, dtbClassCategoryDBTypes, true, dtbClassCategoryColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DTBClassCategory struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := DTBClassCategoryExists(ctx, tx, o.ID)
	if err != nil {
		t.Errorf("Unable to check if DTBClassCategory exists: %s", err)
	}
	if !e {
		t.Errorf("Expected DTBClassCategoryExists to return true, but got false.")
	}
}

func testDTBClassCategoriesFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &DTBClassCategory{}
	if err = randomize.Struct(seed, o, dtbClassCategoryDBTypes, true, dtbClassCategoryColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DTBClassCategory struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	dtbClassCategoryFound, err := FindDTBClassCategory(ctx, tx, o.ID)
	if err != nil {
		t.Error(err)
	}

	if dtbClassCategoryFound == nil {
		t.Error("want a record, got nil")
	}
}

func testDTBClassCategoriesBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &DTBClassCategory{}
	if err = randomize.Struct(seed, o, dtbClassCategoryDBTypes, true, dtbClassCategoryColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DTBClassCategory struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = DTBClassCategories().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testDTBClassCategoriesOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &DTBClassCategory{}
	if err = randomize.Struct(seed, o, dtbClassCategoryDBTypes, true, dtbClassCategoryColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DTBClassCategory struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := DTBClassCategories().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testDTBClassCategoriesAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	dtbClassCategoryOne := &DTBClassCategory{}
	dtbClassCategoryTwo := &DTBClassCategory{}
	if err = randomize.Struct(seed, dtbClassCategoryOne, dtbClassCategoryDBTypes, false, dtbClassCategoryColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DTBClassCategory struct: %s", err)
	}
	if err = randomize.Struct(seed, dtbClassCategoryTwo, dtbClassCategoryDBTypes, false, dtbClassCategoryColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DTBClassCategory struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = dtbClassCategoryOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = dtbClassCategoryTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := DTBClassCategories().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testDTBClassCategoriesCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	dtbClassCategoryOne := &DTBClassCategory{}
	dtbClassCategoryTwo := &DTBClassCategory{}
	if err = randomize.Struct(seed, dtbClassCategoryOne, dtbClassCategoryDBTypes, false, dtbClassCategoryColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DTBClassCategory struct: %s", err)
	}
	if err = randomize.Struct(seed, dtbClassCategoryTwo, dtbClassCategoryDBTypes, false, dtbClassCategoryColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DTBClassCategory struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = dtbClassCategoryOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = dtbClassCategoryTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := DTBClassCategories().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func dtbClassCategoryBeforeInsertHook(ctx context.Context, e boil.ContextExecutor, o *DTBClassCategory) error {
	*o = DTBClassCategory{}
	return nil
}

func dtbClassCategoryAfterInsertHook(ctx context.Context, e boil.ContextExecutor, o *DTBClassCategory) error {
	*o = DTBClassCategory{}
	return nil
}

func dtbClassCategoryAfterSelectHook(ctx context.Context, e boil.ContextExecutor, o *DTBClassCategory) error {
	*o = DTBClassCategory{}
	return nil
}

func dtbClassCategoryBeforeUpdateHook(ctx context.Context, e boil.ContextExecutor, o *DTBClassCategory) error {
	*o = DTBClassCategory{}
	return nil
}

func dtbClassCategoryAfterUpdateHook(ctx context.Context, e boil.ContextExecutor, o *DTBClassCategory) error {
	*o = DTBClassCategory{}
	return nil
}

func dtbClassCategoryBeforeDeleteHook(ctx context.Context, e boil.ContextExecutor, o *DTBClassCategory) error {
	*o = DTBClassCategory{}
	return nil
}

func dtbClassCategoryAfterDeleteHook(ctx context.Context, e boil.ContextExecutor, o *DTBClassCategory) error {
	*o = DTBClassCategory{}
	return nil
}

func dtbClassCategoryBeforeUpsertHook(ctx context.Context, e boil.ContextExecutor, o *DTBClassCategory) error {
	*o = DTBClassCategory{}
	return nil
}

func dtbClassCategoryAfterUpsertHook(ctx context.Context, e boil.ContextExecutor, o *DTBClassCategory) error {
	*o = DTBClassCategory{}
	return nil
}

func testDTBClassCategoriesHooks(t *testing.T) {
	t.Parallel()

	var err error

	ctx := context.Background()
	empty := &DTBClassCategory{}
	o := &DTBClassCategory{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, dtbClassCategoryDBTypes, false); err != nil {
		t.Errorf("Unable to randomize DTBClassCategory object: %s", err)
	}

	AddDTBClassCategoryHook(boil.BeforeInsertHook, dtbClassCategoryBeforeInsertHook)
	if err = o.doBeforeInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	dtbClassCategoryBeforeInsertHooks = []DTBClassCategoryHook{}

	AddDTBClassCategoryHook(boil.AfterInsertHook, dtbClassCategoryAfterInsertHook)
	if err = o.doAfterInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	dtbClassCategoryAfterInsertHooks = []DTBClassCategoryHook{}

	AddDTBClassCategoryHook(boil.AfterSelectHook, dtbClassCategoryAfterSelectHook)
	if err = o.doAfterSelectHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	dtbClassCategoryAfterSelectHooks = []DTBClassCategoryHook{}

	AddDTBClassCategoryHook(boil.BeforeUpdateHook, dtbClassCategoryBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	dtbClassCategoryBeforeUpdateHooks = []DTBClassCategoryHook{}

	AddDTBClassCategoryHook(boil.AfterUpdateHook, dtbClassCategoryAfterUpdateHook)
	if err = o.doAfterUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	dtbClassCategoryAfterUpdateHooks = []DTBClassCategoryHook{}

	AddDTBClassCategoryHook(boil.BeforeDeleteHook, dtbClassCategoryBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	dtbClassCategoryBeforeDeleteHooks = []DTBClassCategoryHook{}

	AddDTBClassCategoryHook(boil.AfterDeleteHook, dtbClassCategoryAfterDeleteHook)
	if err = o.doAfterDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	dtbClassCategoryAfterDeleteHooks = []DTBClassCategoryHook{}

	AddDTBClassCategoryHook(boil.BeforeUpsertHook, dtbClassCategoryBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	dtbClassCategoryBeforeUpsertHooks = []DTBClassCategoryHook{}

	AddDTBClassCategoryHook(boil.AfterUpsertHook, dtbClassCategoryAfterUpsertHook)
	if err = o.doAfterUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	dtbClassCategoryAfterUpsertHooks = []DTBClassCategoryHook{}
}

func testDTBClassCategoriesInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &DTBClassCategory{}
	if err = randomize.Struct(seed, o, dtbClassCategoryDBTypes, true, dtbClassCategoryColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DTBClassCategory struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := DTBClassCategories().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testDTBClassCategoriesInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &DTBClassCategory{}
	if err = randomize.Struct(seed, o, dtbClassCategoryDBTypes, true); err != nil {
		t.Errorf("Unable to randomize DTBClassCategory struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(dtbClassCategoryColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := DTBClassCategories().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testDTBClassCategoryToOneDTBClassNameUsingClassName(t *testing.T) {
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var local DTBClassCategory
	var foreign DTBClassName

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, dtbClassCategoryDBTypes, true, dtbClassCategoryColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DTBClassCategory struct: %s", err)
	}
	if err := randomize.Struct(seed, &foreign, dtbClassNameDBTypes, false, dtbClassNameColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DTBClassName struct: %s", err)
	}

	if err := foreign.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	queries.Assign(&local.ClassNameID, foreign.ID)
	if err := local.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := local.ClassName().One(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	if !queries.Equal(check.ID, foreign.ID) {
		t.Errorf("want: %v, got %v", foreign.ID, check.ID)
	}

	slice := DTBClassCategorySlice{&local}
	if err = local.L.LoadClassName(ctx, tx, false, (*[]*DTBClassCategory)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if local.R.ClassName == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.ClassName = nil
	if err = local.L.LoadClassName(ctx, tx, true, &local, nil); err != nil {
		t.Fatal(err)
	}
	if local.R.ClassName == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testDTBClassCategoryToOneSetOpDTBClassNameUsingClassName(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a DTBClassCategory
	var b, c DTBClassName

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, dtbClassCategoryDBTypes, false, strmangle.SetComplement(dtbClassCategoryPrimaryKeyColumns, dtbClassCategoryColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, dtbClassNameDBTypes, false, strmangle.SetComplement(dtbClassNamePrimaryKeyColumns, dtbClassNameColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, dtbClassNameDBTypes, false, strmangle.SetComplement(dtbClassNamePrimaryKeyColumns, dtbClassNameColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*DTBClassName{&b, &c} {
		err = a.SetClassName(ctx, tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.ClassName != x {
			t.Error("relationship struct not set to correct value")
		}

		if x.R.ClassNameDTBClassCategories[0] != &a {
			t.Error("failed to append to foreign relationship struct")
		}
		if !queries.Equal(a.ClassNameID, x.ID) {
			t.Error("foreign key was wrong value", a.ClassNameID)
		}

		zero := reflect.Zero(reflect.TypeOf(a.ClassNameID))
		reflect.Indirect(reflect.ValueOf(&a.ClassNameID)).Set(zero)

		if err = a.Reload(ctx, tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if !queries.Equal(a.ClassNameID, x.ID) {
			t.Error("foreign key was wrong value", a.ClassNameID, x.ID)
		}
	}
}

func testDTBClassCategoryToOneRemoveOpDTBClassNameUsingClassName(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a DTBClassCategory
	var b DTBClassName

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, dtbClassCategoryDBTypes, false, strmangle.SetComplement(dtbClassCategoryPrimaryKeyColumns, dtbClassCategoryColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, dtbClassNameDBTypes, false, strmangle.SetComplement(dtbClassNamePrimaryKeyColumns, dtbClassNameColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err = a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	if err = a.SetClassName(ctx, tx, true, &b); err != nil {
		t.Fatal(err)
	}

	if err = a.RemoveClassName(ctx, tx, &b); err != nil {
		t.Error("failed to remove relationship")
	}

	count, err := a.ClassName().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 0 {
		t.Error("want no relationships remaining")
	}

	if a.R.ClassName != nil {
		t.Error("R struct entry should be nil")
	}

	if !queries.IsValuerNil(a.ClassNameID) {
		t.Error("foreign key value should be nil")
	}

	if len(b.R.ClassNameDTBClassCategories) != 0 {
		t.Error("failed to remove a from b's relationships")
	}
}

func testDTBClassCategoriesReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &DTBClassCategory{}
	if err = randomize.Struct(seed, o, dtbClassCategoryDBTypes, true, dtbClassCategoryColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DTBClassCategory struct: %s", err)
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

func testDTBClassCategoriesReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &DTBClassCategory{}
	if err = randomize.Struct(seed, o, dtbClassCategoryDBTypes, true, dtbClassCategoryColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DTBClassCategory struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := DTBClassCategorySlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testDTBClassCategoriesSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &DTBClassCategory{}
	if err = randomize.Struct(seed, o, dtbClassCategoryDBTypes, true, dtbClassCategoryColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DTBClassCategory struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := DTBClassCategories().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	dtbClassCategoryDBTypes = map[string]string{`ID`: `int`, `ClassNameID`: `int`, `CreatorID`: `int`, `BackendName`: `varchar`, `Name`: `varchar`, `SortNo`: `int`, `Visible`: `tinyint`, `CreateDate`: `datetime`, `UpdateDate`: `datetime`, `DiscriminatorType`: `varchar`}
	_                       = bytes.MinRead
)

func testDTBClassCategoriesUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(dtbClassCategoryPrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(dtbClassCategoryColumns) == len(dtbClassCategoryPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &DTBClassCategory{}
	if err = randomize.Struct(seed, o, dtbClassCategoryDBTypes, true, dtbClassCategoryColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DTBClassCategory struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := DTBClassCategories().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, dtbClassCategoryDBTypes, true, dtbClassCategoryPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize DTBClassCategory struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testDTBClassCategoriesSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(dtbClassCategoryColumns) == len(dtbClassCategoryPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &DTBClassCategory{}
	if err = randomize.Struct(seed, o, dtbClassCategoryDBTypes, true, dtbClassCategoryColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DTBClassCategory struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := DTBClassCategories().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, dtbClassCategoryDBTypes, true, dtbClassCategoryPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize DTBClassCategory struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(dtbClassCategoryColumns, dtbClassCategoryPrimaryKeyColumns) {
		fields = dtbClassCategoryColumns
	} else {
		fields = strmangle.SetComplement(
			dtbClassCategoryColumns,
			dtbClassCategoryPrimaryKeyColumns,
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

	slice := DTBClassCategorySlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testDTBClassCategoriesUpsert(t *testing.T) {
	t.Parallel()

	if len(dtbClassCategoryColumns) == len(dtbClassCategoryPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}
	if len(mySQLDTBClassCategoryUniqueColumns) == 0 {
		t.Skip("Skipping table with no unique columns to conflict on")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := DTBClassCategory{}
	if err = randomize.Struct(seed, &o, dtbClassCategoryDBTypes, false); err != nil {
		t.Errorf("Unable to randomize DTBClassCategory struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(ctx, tx, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert DTBClassCategory: %s", err)
	}

	count, err := DTBClassCategories().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, dtbClassCategoryDBTypes, false, dtbClassCategoryPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize DTBClassCategory struct: %s", err)
	}

	if err = o.Upsert(ctx, tx, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert DTBClassCategory: %s", err)
	}

	count, err = DTBClassCategories().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
