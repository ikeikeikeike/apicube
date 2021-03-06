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

func testMTBJobs(t *testing.T) {
	t.Parallel()

	query := MTBJobs()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testMTBJobsDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &MTBJob{}
	if err = randomize.Struct(seed, o, mtbJobDBTypes, true, mtbJobColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize MTBJob struct: %s", err)
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

	count, err := MTBJobs().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testMTBJobsQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &MTBJob{}
	if err = randomize.Struct(seed, o, mtbJobDBTypes, true, mtbJobColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize MTBJob struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := MTBJobs().DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := MTBJobs().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testMTBJobsSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &MTBJob{}
	if err = randomize.Struct(seed, o, mtbJobDBTypes, true, mtbJobColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize MTBJob struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := MTBJobSlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := MTBJobs().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testMTBJobsExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &MTBJob{}
	if err = randomize.Struct(seed, o, mtbJobDBTypes, true, mtbJobColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize MTBJob struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := MTBJobExists(ctx, tx, o.ID)
	if err != nil {
		t.Errorf("Unable to check if MTBJob exists: %s", err)
	}
	if !e {
		t.Errorf("Expected MTBJobExists to return true, but got false.")
	}
}

func testMTBJobsFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &MTBJob{}
	if err = randomize.Struct(seed, o, mtbJobDBTypes, true, mtbJobColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize MTBJob struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	mtbJobFound, err := FindMTBJob(ctx, tx, o.ID)
	if err != nil {
		t.Error(err)
	}

	if mtbJobFound == nil {
		t.Error("want a record, got nil")
	}
}

func testMTBJobsBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &MTBJob{}
	if err = randomize.Struct(seed, o, mtbJobDBTypes, true, mtbJobColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize MTBJob struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = MTBJobs().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testMTBJobsOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &MTBJob{}
	if err = randomize.Struct(seed, o, mtbJobDBTypes, true, mtbJobColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize MTBJob struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := MTBJobs().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testMTBJobsAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	mtbJobOne := &MTBJob{}
	mtbJobTwo := &MTBJob{}
	if err = randomize.Struct(seed, mtbJobOne, mtbJobDBTypes, false, mtbJobColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize MTBJob struct: %s", err)
	}
	if err = randomize.Struct(seed, mtbJobTwo, mtbJobDBTypes, false, mtbJobColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize MTBJob struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = mtbJobOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = mtbJobTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := MTBJobs().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testMTBJobsCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	mtbJobOne := &MTBJob{}
	mtbJobTwo := &MTBJob{}
	if err = randomize.Struct(seed, mtbJobOne, mtbJobDBTypes, false, mtbJobColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize MTBJob struct: %s", err)
	}
	if err = randomize.Struct(seed, mtbJobTwo, mtbJobDBTypes, false, mtbJobColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize MTBJob struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = mtbJobOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = mtbJobTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := MTBJobs().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func mtbJobBeforeInsertHook(ctx context.Context, e boil.ContextExecutor, o *MTBJob) error {
	*o = MTBJob{}
	return nil
}

func mtbJobAfterInsertHook(ctx context.Context, e boil.ContextExecutor, o *MTBJob) error {
	*o = MTBJob{}
	return nil
}

func mtbJobAfterSelectHook(ctx context.Context, e boil.ContextExecutor, o *MTBJob) error {
	*o = MTBJob{}
	return nil
}

func mtbJobBeforeUpdateHook(ctx context.Context, e boil.ContextExecutor, o *MTBJob) error {
	*o = MTBJob{}
	return nil
}

func mtbJobAfterUpdateHook(ctx context.Context, e boil.ContextExecutor, o *MTBJob) error {
	*o = MTBJob{}
	return nil
}

func mtbJobBeforeDeleteHook(ctx context.Context, e boil.ContextExecutor, o *MTBJob) error {
	*o = MTBJob{}
	return nil
}

func mtbJobAfterDeleteHook(ctx context.Context, e boil.ContextExecutor, o *MTBJob) error {
	*o = MTBJob{}
	return nil
}

func mtbJobBeforeUpsertHook(ctx context.Context, e boil.ContextExecutor, o *MTBJob) error {
	*o = MTBJob{}
	return nil
}

func mtbJobAfterUpsertHook(ctx context.Context, e boil.ContextExecutor, o *MTBJob) error {
	*o = MTBJob{}
	return nil
}

func testMTBJobsHooks(t *testing.T) {
	t.Parallel()

	var err error

	ctx := context.Background()
	empty := &MTBJob{}
	o := &MTBJob{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, mtbJobDBTypes, false); err != nil {
		t.Errorf("Unable to randomize MTBJob object: %s", err)
	}

	AddMTBJobHook(boil.BeforeInsertHook, mtbJobBeforeInsertHook)
	if err = o.doBeforeInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	mtbJobBeforeInsertHooks = []MTBJobHook{}

	AddMTBJobHook(boil.AfterInsertHook, mtbJobAfterInsertHook)
	if err = o.doAfterInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	mtbJobAfterInsertHooks = []MTBJobHook{}

	AddMTBJobHook(boil.AfterSelectHook, mtbJobAfterSelectHook)
	if err = o.doAfterSelectHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	mtbJobAfterSelectHooks = []MTBJobHook{}

	AddMTBJobHook(boil.BeforeUpdateHook, mtbJobBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	mtbJobBeforeUpdateHooks = []MTBJobHook{}

	AddMTBJobHook(boil.AfterUpdateHook, mtbJobAfterUpdateHook)
	if err = o.doAfterUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	mtbJobAfterUpdateHooks = []MTBJobHook{}

	AddMTBJobHook(boil.BeforeDeleteHook, mtbJobBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	mtbJobBeforeDeleteHooks = []MTBJobHook{}

	AddMTBJobHook(boil.AfterDeleteHook, mtbJobAfterDeleteHook)
	if err = o.doAfterDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	mtbJobAfterDeleteHooks = []MTBJobHook{}

	AddMTBJobHook(boil.BeforeUpsertHook, mtbJobBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	mtbJobBeforeUpsertHooks = []MTBJobHook{}

	AddMTBJobHook(boil.AfterUpsertHook, mtbJobAfterUpsertHook)
	if err = o.doAfterUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	mtbJobAfterUpsertHooks = []MTBJobHook{}
}

func testMTBJobsInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &MTBJob{}
	if err = randomize.Struct(seed, o, mtbJobDBTypes, true, mtbJobColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize MTBJob struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := MTBJobs().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testMTBJobsInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &MTBJob{}
	if err = randomize.Struct(seed, o, mtbJobDBTypes, true); err != nil {
		t.Errorf("Unable to randomize MTBJob struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(mtbJobColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := MTBJobs().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testMTBJobToManyJobDTBCustomers(t *testing.T) {
	var err error
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a MTBJob
	var b, c DTBCustomer

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, mtbJobDBTypes, true, mtbJobColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize MTBJob struct: %s", err)
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

	queries.Assign(&b.JobID, a.ID)
	queries.Assign(&c.JobID, a.ID)
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := a.JobDTBCustomers().All(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	bFound, cFound := false, false
	for _, v := range check {
		if queries.Equal(v.JobID, b.JobID) {
			bFound = true
		}
		if queries.Equal(v.JobID, c.JobID) {
			cFound = true
		}
	}

	if !bFound {
		t.Error("expected to find b")
	}
	if !cFound {
		t.Error("expected to find c")
	}

	slice := MTBJobSlice{&a}
	if err = a.L.LoadJobDTBCustomers(ctx, tx, false, (*[]*MTBJob)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.JobDTBCustomers); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	a.R.JobDTBCustomers = nil
	if err = a.L.LoadJobDTBCustomers(ctx, tx, true, &a, nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.JobDTBCustomers); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	if t.Failed() {
		t.Logf("%#v", check)
	}
}

func testMTBJobToManyJobDTBOrders(t *testing.T) {
	var err error
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a MTBJob
	var b, c DTBOrder

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, mtbJobDBTypes, true, mtbJobColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize MTBJob struct: %s", err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	if err = randomize.Struct(seed, &b, dtbOrderDBTypes, false, dtbOrderColumnsWithDefault...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, dtbOrderDBTypes, false, dtbOrderColumnsWithDefault...); err != nil {
		t.Fatal(err)
	}

	queries.Assign(&b.JobID, a.ID)
	queries.Assign(&c.JobID, a.ID)
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := a.JobDTBOrders().All(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	bFound, cFound := false, false
	for _, v := range check {
		if queries.Equal(v.JobID, b.JobID) {
			bFound = true
		}
		if queries.Equal(v.JobID, c.JobID) {
			cFound = true
		}
	}

	if !bFound {
		t.Error("expected to find b")
	}
	if !cFound {
		t.Error("expected to find c")
	}

	slice := MTBJobSlice{&a}
	if err = a.L.LoadJobDTBOrders(ctx, tx, false, (*[]*MTBJob)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.JobDTBOrders); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	a.R.JobDTBOrders = nil
	if err = a.L.LoadJobDTBOrders(ctx, tx, true, &a, nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.JobDTBOrders); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	if t.Failed() {
		t.Logf("%#v", check)
	}
}

func testMTBJobToManyAddOpJobDTBCustomers(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a MTBJob
	var b, c, d, e DTBCustomer

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, mtbJobDBTypes, false, strmangle.SetComplement(mtbJobPrimaryKeyColumns, mtbJobColumnsWithoutDefault)...); err != nil {
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
		err = a.AddJobDTBCustomers(ctx, tx, i != 0, x...)
		if err != nil {
			t.Fatal(err)
		}

		first := x[0]
		second := x[1]

		if !queries.Equal(a.ID, first.JobID) {
			t.Error("foreign key was wrong value", a.ID, first.JobID)
		}
		if !queries.Equal(a.ID, second.JobID) {
			t.Error("foreign key was wrong value", a.ID, second.JobID)
		}

		if first.R.Job != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}
		if second.R.Job != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}

		if a.R.JobDTBCustomers[i*2] != first {
			t.Error("relationship struct slice not set to correct value")
		}
		if a.R.JobDTBCustomers[i*2+1] != second {
			t.Error("relationship struct slice not set to correct value")
		}

		count, err := a.JobDTBCustomers().Count(ctx, tx)
		if err != nil {
			t.Fatal(err)
		}
		if want := int64((i + 1) * 2); count != want {
			t.Error("want", want, "got", count)
		}
	}
}

func testMTBJobToManySetOpJobDTBCustomers(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a MTBJob
	var b, c, d, e DTBCustomer

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, mtbJobDBTypes, false, strmangle.SetComplement(mtbJobPrimaryKeyColumns, mtbJobColumnsWithoutDefault)...); err != nil {
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

	err = a.SetJobDTBCustomers(ctx, tx, false, &b, &c)
	if err != nil {
		t.Fatal(err)
	}

	count, err := a.JobDTBCustomers().Count(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}
	if count != 2 {
		t.Error("count was wrong:", count)
	}

	err = a.SetJobDTBCustomers(ctx, tx, true, &d, &e)
	if err != nil {
		t.Fatal(err)
	}

	count, err = a.JobDTBCustomers().Count(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}
	if count != 2 {
		t.Error("count was wrong:", count)
	}

	if !queries.IsValuerNil(b.JobID) {
		t.Error("want b's foreign key value to be nil")
	}
	if !queries.IsValuerNil(c.JobID) {
		t.Error("want c's foreign key value to be nil")
	}
	if !queries.Equal(a.ID, d.JobID) {
		t.Error("foreign key was wrong value", a.ID, d.JobID)
	}
	if !queries.Equal(a.ID, e.JobID) {
		t.Error("foreign key was wrong value", a.ID, e.JobID)
	}

	if b.R.Job != nil {
		t.Error("relationship was not removed properly from the foreign struct")
	}
	if c.R.Job != nil {
		t.Error("relationship was not removed properly from the foreign struct")
	}
	if d.R.Job != &a {
		t.Error("relationship was not added properly to the foreign struct")
	}
	if e.R.Job != &a {
		t.Error("relationship was not added properly to the foreign struct")
	}

	if a.R.JobDTBCustomers[0] != &d {
		t.Error("relationship struct slice not set to correct value")
	}
	if a.R.JobDTBCustomers[1] != &e {
		t.Error("relationship struct slice not set to correct value")
	}
}

func testMTBJobToManyRemoveOpJobDTBCustomers(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a MTBJob
	var b, c, d, e DTBCustomer

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, mtbJobDBTypes, false, strmangle.SetComplement(mtbJobPrimaryKeyColumns, mtbJobColumnsWithoutDefault)...); err != nil {
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

	err = a.AddJobDTBCustomers(ctx, tx, true, foreigners...)
	if err != nil {
		t.Fatal(err)
	}

	count, err := a.JobDTBCustomers().Count(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}
	if count != 4 {
		t.Error("count was wrong:", count)
	}

	err = a.RemoveJobDTBCustomers(ctx, tx, foreigners[:2]...)
	if err != nil {
		t.Fatal(err)
	}

	count, err = a.JobDTBCustomers().Count(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}
	if count != 2 {
		t.Error("count was wrong:", count)
	}

	if !queries.IsValuerNil(b.JobID) {
		t.Error("want b's foreign key value to be nil")
	}
	if !queries.IsValuerNil(c.JobID) {
		t.Error("want c's foreign key value to be nil")
	}

	if b.R.Job != nil {
		t.Error("relationship was not removed properly from the foreign struct")
	}
	if c.R.Job != nil {
		t.Error("relationship was not removed properly from the foreign struct")
	}
	if d.R.Job != &a {
		t.Error("relationship to a should have been preserved")
	}
	if e.R.Job != &a {
		t.Error("relationship to a should have been preserved")
	}

	if len(a.R.JobDTBCustomers) != 2 {
		t.Error("should have preserved two relationships")
	}

	// Removal doesn't do a stable deletion for performance so we have to flip the order
	if a.R.JobDTBCustomers[1] != &d {
		t.Error("relationship to d should have been preserved")
	}
	if a.R.JobDTBCustomers[0] != &e {
		t.Error("relationship to e should have been preserved")
	}
}

func testMTBJobToManyAddOpJobDTBOrders(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a MTBJob
	var b, c, d, e DTBOrder

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, mtbJobDBTypes, false, strmangle.SetComplement(mtbJobPrimaryKeyColumns, mtbJobColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*DTBOrder{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, dtbOrderDBTypes, false, strmangle.SetComplement(dtbOrderPrimaryKeyColumns, dtbOrderColumnsWithoutDefault)...); err != nil {
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

	foreignersSplitByInsertion := [][]*DTBOrder{
		{&b, &c},
		{&d, &e},
	}

	for i, x := range foreignersSplitByInsertion {
		err = a.AddJobDTBOrders(ctx, tx, i != 0, x...)
		if err != nil {
			t.Fatal(err)
		}

		first := x[0]
		second := x[1]

		if !queries.Equal(a.ID, first.JobID) {
			t.Error("foreign key was wrong value", a.ID, first.JobID)
		}
		if !queries.Equal(a.ID, second.JobID) {
			t.Error("foreign key was wrong value", a.ID, second.JobID)
		}

		if first.R.Job != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}
		if second.R.Job != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}

		if a.R.JobDTBOrders[i*2] != first {
			t.Error("relationship struct slice not set to correct value")
		}
		if a.R.JobDTBOrders[i*2+1] != second {
			t.Error("relationship struct slice not set to correct value")
		}

		count, err := a.JobDTBOrders().Count(ctx, tx)
		if err != nil {
			t.Fatal(err)
		}
		if want := int64((i + 1) * 2); count != want {
			t.Error("want", want, "got", count)
		}
	}
}

func testMTBJobToManySetOpJobDTBOrders(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a MTBJob
	var b, c, d, e DTBOrder

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, mtbJobDBTypes, false, strmangle.SetComplement(mtbJobPrimaryKeyColumns, mtbJobColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*DTBOrder{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, dtbOrderDBTypes, false, strmangle.SetComplement(dtbOrderPrimaryKeyColumns, dtbOrderColumnsWithoutDefault)...); err != nil {
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

	err = a.SetJobDTBOrders(ctx, tx, false, &b, &c)
	if err != nil {
		t.Fatal(err)
	}

	count, err := a.JobDTBOrders().Count(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}
	if count != 2 {
		t.Error("count was wrong:", count)
	}

	err = a.SetJobDTBOrders(ctx, tx, true, &d, &e)
	if err != nil {
		t.Fatal(err)
	}

	count, err = a.JobDTBOrders().Count(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}
	if count != 2 {
		t.Error("count was wrong:", count)
	}

	if !queries.IsValuerNil(b.JobID) {
		t.Error("want b's foreign key value to be nil")
	}
	if !queries.IsValuerNil(c.JobID) {
		t.Error("want c's foreign key value to be nil")
	}
	if !queries.Equal(a.ID, d.JobID) {
		t.Error("foreign key was wrong value", a.ID, d.JobID)
	}
	if !queries.Equal(a.ID, e.JobID) {
		t.Error("foreign key was wrong value", a.ID, e.JobID)
	}

	if b.R.Job != nil {
		t.Error("relationship was not removed properly from the foreign struct")
	}
	if c.R.Job != nil {
		t.Error("relationship was not removed properly from the foreign struct")
	}
	if d.R.Job != &a {
		t.Error("relationship was not added properly to the foreign struct")
	}
	if e.R.Job != &a {
		t.Error("relationship was not added properly to the foreign struct")
	}

	if a.R.JobDTBOrders[0] != &d {
		t.Error("relationship struct slice not set to correct value")
	}
	if a.R.JobDTBOrders[1] != &e {
		t.Error("relationship struct slice not set to correct value")
	}
}

func testMTBJobToManyRemoveOpJobDTBOrders(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a MTBJob
	var b, c, d, e DTBOrder

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, mtbJobDBTypes, false, strmangle.SetComplement(mtbJobPrimaryKeyColumns, mtbJobColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*DTBOrder{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, dtbOrderDBTypes, false, strmangle.SetComplement(dtbOrderPrimaryKeyColumns, dtbOrderColumnsWithoutDefault)...); err != nil {
			t.Fatal(err)
		}
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	err = a.AddJobDTBOrders(ctx, tx, true, foreigners...)
	if err != nil {
		t.Fatal(err)
	}

	count, err := a.JobDTBOrders().Count(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}
	if count != 4 {
		t.Error("count was wrong:", count)
	}

	err = a.RemoveJobDTBOrders(ctx, tx, foreigners[:2]...)
	if err != nil {
		t.Fatal(err)
	}

	count, err = a.JobDTBOrders().Count(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}
	if count != 2 {
		t.Error("count was wrong:", count)
	}

	if !queries.IsValuerNil(b.JobID) {
		t.Error("want b's foreign key value to be nil")
	}
	if !queries.IsValuerNil(c.JobID) {
		t.Error("want c's foreign key value to be nil")
	}

	if b.R.Job != nil {
		t.Error("relationship was not removed properly from the foreign struct")
	}
	if c.R.Job != nil {
		t.Error("relationship was not removed properly from the foreign struct")
	}
	if d.R.Job != &a {
		t.Error("relationship to a should have been preserved")
	}
	if e.R.Job != &a {
		t.Error("relationship to a should have been preserved")
	}

	if len(a.R.JobDTBOrders) != 2 {
		t.Error("should have preserved two relationships")
	}

	// Removal doesn't do a stable deletion for performance so we have to flip the order
	if a.R.JobDTBOrders[1] != &d {
		t.Error("relationship to d should have been preserved")
	}
	if a.R.JobDTBOrders[0] != &e {
		t.Error("relationship to e should have been preserved")
	}
}

func testMTBJobsReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &MTBJob{}
	if err = randomize.Struct(seed, o, mtbJobDBTypes, true, mtbJobColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize MTBJob struct: %s", err)
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

func testMTBJobsReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &MTBJob{}
	if err = randomize.Struct(seed, o, mtbJobDBTypes, true, mtbJobColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize MTBJob struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := MTBJobSlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testMTBJobsSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &MTBJob{}
	if err = randomize.Struct(seed, o, mtbJobDBTypes, true, mtbJobColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize MTBJob struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := MTBJobs().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	mtbJobDBTypes = map[string]string{`ID`: `smallint`, `Name`: `varchar`, `SortNo`: `smallint`, `DiscriminatorType`: `varchar`}
	_             = bytes.MinRead
)

func testMTBJobsUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(mtbJobPrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(mtbJobColumns) == len(mtbJobPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &MTBJob{}
	if err = randomize.Struct(seed, o, mtbJobDBTypes, true, mtbJobColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize MTBJob struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := MTBJobs().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, mtbJobDBTypes, true, mtbJobPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize MTBJob struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testMTBJobsSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(mtbJobColumns) == len(mtbJobPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &MTBJob{}
	if err = randomize.Struct(seed, o, mtbJobDBTypes, true, mtbJobColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize MTBJob struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := MTBJobs().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, mtbJobDBTypes, true, mtbJobPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize MTBJob struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(mtbJobColumns, mtbJobPrimaryKeyColumns) {
		fields = mtbJobColumns
	} else {
		fields = strmangle.SetComplement(
			mtbJobColumns,
			mtbJobPrimaryKeyColumns,
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

	slice := MTBJobSlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testMTBJobsUpsert(t *testing.T) {
	t.Parallel()

	if len(mtbJobColumns) == len(mtbJobPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}
	if len(mySQLMTBJobUniqueColumns) == 0 {
		t.Skip("Skipping table with no unique columns to conflict on")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := MTBJob{}
	if err = randomize.Struct(seed, &o, mtbJobDBTypes, false); err != nil {
		t.Errorf("Unable to randomize MTBJob struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(ctx, tx, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert MTBJob: %s", err)
	}

	count, err := MTBJobs().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, mtbJobDBTypes, false, mtbJobPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize MTBJob struct: %s", err)
	}

	if err = o.Upsert(ctx, tx, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert MTBJob: %s", err)
	}

	count, err = MTBJobs().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
