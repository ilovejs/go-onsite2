// Code generated by SQLBoiler (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

import (
	"bytes"
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

func testClaimHistories(t *testing.T) {
	t.Parallel()

	query := ClaimHistories()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testClaimHistoriesDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ClaimHistory{}
	if err = randomize.Struct(seed, o, claimHistoryDBTypes, true, claimHistoryColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ClaimHistory struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := o.Delete(tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := ClaimHistories().Count(tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testClaimHistoriesQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ClaimHistory{}
	if err = randomize.Struct(seed, o, claimHistoryDBTypes, true, claimHistoryColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ClaimHistory struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := ClaimHistories().DeleteAll(tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := ClaimHistories().Count(tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testClaimHistoriesSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ClaimHistory{}
	if err = randomize.Struct(seed, o, claimHistoryDBTypes, true, claimHistoryColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ClaimHistory struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := ClaimHistorySlice{o}

	if rowsAff, err := slice.DeleteAll(tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := ClaimHistories().Count(tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testClaimHistoriesExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ClaimHistory{}
	if err = randomize.Struct(seed, o, claimHistoryDBTypes, true, claimHistoryColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ClaimHistory struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := ClaimHistoryExists(tx, o.ID)
	if err != nil {
		t.Errorf("Unable to check if ClaimHistory exists: %s", err)
	}
	if !e {
		t.Errorf("Expected ClaimHistoryExists to return true, but got false.")
	}
}

func testClaimHistoriesFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ClaimHistory{}
	if err = randomize.Struct(seed, o, claimHistoryDBTypes, true, claimHistoryColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ClaimHistory struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	claimHistoryFound, err := FindClaimHistory(tx, o.ID)
	if err != nil {
		t.Error(err)
	}

	if claimHistoryFound == nil {
		t.Error("want a record, got nil")
	}
}

func testClaimHistoriesBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ClaimHistory{}
	if err = randomize.Struct(seed, o, claimHistoryDBTypes, true, claimHistoryColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ClaimHistory struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = ClaimHistories().Bind(nil, tx, o); err != nil {
		t.Error(err)
	}
}

func testClaimHistoriesOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ClaimHistory{}
	if err = randomize.Struct(seed, o, claimHistoryDBTypes, true, claimHistoryColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ClaimHistory struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := ClaimHistories().One(tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testClaimHistoriesAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	claimHistoryOne := &ClaimHistory{}
	claimHistoryTwo := &ClaimHistory{}
	if err = randomize.Struct(seed, claimHistoryOne, claimHistoryDBTypes, false, claimHistoryColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ClaimHistory struct: %s", err)
	}
	if err = randomize.Struct(seed, claimHistoryTwo, claimHistoryDBTypes, false, claimHistoryColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ClaimHistory struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer func() { _ = tx.Rollback() }()
	if err = claimHistoryOne.Insert(tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = claimHistoryTwo.Insert(tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := ClaimHistories().All(tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testClaimHistoriesCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	claimHistoryOne := &ClaimHistory{}
	claimHistoryTwo := &ClaimHistory{}
	if err = randomize.Struct(seed, claimHistoryOne, claimHistoryDBTypes, false, claimHistoryColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ClaimHistory struct: %s", err)
	}
	if err = randomize.Struct(seed, claimHistoryTwo, claimHistoryDBTypes, false, claimHistoryColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ClaimHistory struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer func() { _ = tx.Rollback() }()
	if err = claimHistoryOne.Insert(tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = claimHistoryTwo.Insert(tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := ClaimHistories().Count(tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func testClaimHistoriesInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ClaimHistory{}
	if err = randomize.Struct(seed, o, claimHistoryDBTypes, true, claimHistoryColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ClaimHistory struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := ClaimHistories().Count(tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testClaimHistoriesInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ClaimHistory{}
	if err = randomize.Struct(seed, o, claimHistoryDBTypes, true); err != nil {
		t.Errorf("Unable to randomize ClaimHistory struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(tx, boil.Whitelist(claimHistoryColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := ClaimHistories().Count(tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testClaimHistoryToOneClaimUsingClaim(t *testing.T) {

	tx := MustTx(boil.Begin())
	defer func() { _ = tx.Rollback() }()

	var local ClaimHistory
	var foreign Claim

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, claimHistoryDBTypes, false, claimHistoryColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ClaimHistory struct: %s", err)
	}
	if err := randomize.Struct(seed, &foreign, claimDBTypes, false, claimColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Claim struct: %s", err)
	}

	if err := foreign.Insert(tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	local.ClaimID = foreign.ID
	if err := local.Insert(tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := local.Claim().One(tx)
	if err != nil {
		t.Fatal(err)
	}

	if check.ID != foreign.ID {
		t.Errorf("want: %v, got %v", foreign.ID, check.ID)
	}

	slice := ClaimHistorySlice{&local}
	if err = local.L.LoadClaim(tx, false, (*[]*ClaimHistory)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if local.R.Claim == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.Claim = nil
	if err = local.L.LoadClaim(tx, true, &local, nil); err != nil {
		t.Fatal(err)
	}
	if local.R.Claim == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testClaimHistoryToOneProfileUsingProfile(t *testing.T) {

	tx := MustTx(boil.Begin())
	defer func() { _ = tx.Rollback() }()

	var local ClaimHistory
	var foreign Profile

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, claimHistoryDBTypes, false, claimHistoryColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ClaimHistory struct: %s", err)
	}
	if err := randomize.Struct(seed, &foreign, profileDBTypes, false, profileColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Profile struct: %s", err)
	}

	if err := foreign.Insert(tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	local.ProfileID = foreign.ID
	if err := local.Insert(tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := local.Profile().One(tx)
	if err != nil {
		t.Fatal(err)
	}

	if check.ID != foreign.ID {
		t.Errorf("want: %v, got %v", foreign.ID, check.ID)
	}

	slice := ClaimHistorySlice{&local}
	if err = local.L.LoadProfile(tx, false, (*[]*ClaimHistory)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if local.R.Profile == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.Profile = nil
	if err = local.L.LoadProfile(tx, true, &local, nil); err != nil {
		t.Fatal(err)
	}
	if local.R.Profile == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testClaimHistoryToOneTradeUsingTrade(t *testing.T) {

	tx := MustTx(boil.Begin())
	defer func() { _ = tx.Rollback() }()

	var local ClaimHistory
	var foreign Trade

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, claimHistoryDBTypes, false, claimHistoryColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ClaimHistory struct: %s", err)
	}
	if err := randomize.Struct(seed, &foreign, tradeDBTypes, false, tradeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Trade struct: %s", err)
	}

	if err := foreign.Insert(tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	local.TradeID = foreign.ID
	if err := local.Insert(tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := local.Trade().One(tx)
	if err != nil {
		t.Fatal(err)
	}

	if check.ID != foreign.ID {
		t.Errorf("want: %v, got %v", foreign.ID, check.ID)
	}

	slice := ClaimHistorySlice{&local}
	if err = local.L.LoadTrade(tx, false, (*[]*ClaimHistory)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if local.R.Trade == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.Trade = nil
	if err = local.L.LoadTrade(tx, true, &local, nil); err != nil {
		t.Fatal(err)
	}
	if local.R.Trade == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testClaimHistoryToOneSetOpClaimUsingClaim(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer func() { _ = tx.Rollback() }()

	var a ClaimHistory
	var b, c Claim

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, claimHistoryDBTypes, false, strmangle.SetComplement(claimHistoryPrimaryKeyColumns, claimHistoryColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, claimDBTypes, false, strmangle.SetComplement(claimPrimaryKeyColumns, claimColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, claimDBTypes, false, strmangle.SetComplement(claimPrimaryKeyColumns, claimColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*Claim{&b, &c} {
		err = a.SetClaim(tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.Claim != x {
			t.Error("relationship struct not set to correct value")
		}

		if x.R.ClaimHistories[0] != &a {
			t.Error("failed to append to foreign relationship struct")
		}
		if a.ClaimID != x.ID {
			t.Error("foreign key was wrong value", a.ClaimID)
		}

		zero := reflect.Zero(reflect.TypeOf(a.ClaimID))
		reflect.Indirect(reflect.ValueOf(&a.ClaimID)).Set(zero)

		if err = a.Reload(tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if a.ClaimID != x.ID {
			t.Error("foreign key was wrong value", a.ClaimID, x.ID)
		}
	}
}
func testClaimHistoryToOneSetOpProfileUsingProfile(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer func() { _ = tx.Rollback() }()

	var a ClaimHistory
	var b, c Profile

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, claimHistoryDBTypes, false, strmangle.SetComplement(claimHistoryPrimaryKeyColumns, claimHistoryColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, profileDBTypes, false, strmangle.SetComplement(profilePrimaryKeyColumns, profileColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, profileDBTypes, false, strmangle.SetComplement(profilePrimaryKeyColumns, profileColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*Profile{&b, &c} {
		err = a.SetProfile(tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.Profile != x {
			t.Error("relationship struct not set to correct value")
		}

		if x.R.ClaimHistories[0] != &a {
			t.Error("failed to append to foreign relationship struct")
		}
		if a.ProfileID != x.ID {
			t.Error("foreign key was wrong value", a.ProfileID)
		}

		zero := reflect.Zero(reflect.TypeOf(a.ProfileID))
		reflect.Indirect(reflect.ValueOf(&a.ProfileID)).Set(zero)

		if err = a.Reload(tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if a.ProfileID != x.ID {
			t.Error("foreign key was wrong value", a.ProfileID, x.ID)
		}
	}
}
func testClaimHistoryToOneSetOpTradeUsingTrade(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer func() { _ = tx.Rollback() }()

	var a ClaimHistory
	var b, c Trade

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, claimHistoryDBTypes, false, strmangle.SetComplement(claimHistoryPrimaryKeyColumns, claimHistoryColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, tradeDBTypes, false, strmangle.SetComplement(tradePrimaryKeyColumns, tradeColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, tradeDBTypes, false, strmangle.SetComplement(tradePrimaryKeyColumns, tradeColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*Trade{&b, &c} {
		err = a.SetTrade(tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.Trade != x {
			t.Error("relationship struct not set to correct value")
		}

		if x.R.ClaimHistories[0] != &a {
			t.Error("failed to append to foreign relationship struct")
		}
		if a.TradeID != x.ID {
			t.Error("foreign key was wrong value", a.TradeID)
		}

		zero := reflect.Zero(reflect.TypeOf(a.TradeID))
		reflect.Indirect(reflect.ValueOf(&a.TradeID)).Set(zero)

		if err = a.Reload(tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if a.TradeID != x.ID {
			t.Error("foreign key was wrong value", a.TradeID, x.ID)
		}
	}
}

func testClaimHistoriesReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ClaimHistory{}
	if err = randomize.Struct(seed, o, claimHistoryDBTypes, true, claimHistoryColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ClaimHistory struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = o.Reload(tx); err != nil {
		t.Error(err)
	}
}

func testClaimHistoriesReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ClaimHistory{}
	if err = randomize.Struct(seed, o, claimHistoryDBTypes, true, claimHistoryColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ClaimHistory struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := ClaimHistorySlice{o}

	if err = slice.ReloadAll(tx); err != nil {
		t.Error(err)
	}
}

func testClaimHistoriesSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ClaimHistory{}
	if err = randomize.Struct(seed, o, claimHistoryDBTypes, true, claimHistoryColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ClaimHistory struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := ClaimHistories().All(tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	claimHistoryDBTypes = map[string]string{`ID`: `int`, `TradeID`: `int`, `ClaimID`: `int`, `ProfileID`: `int`, `PreviousClaimed`: `float`, `Created`: `datetime`, `Updated`: `datetime`}
	_                   = bytes.MinRead
)

func testClaimHistoriesUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(claimHistoryPrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(claimHistoryAllColumns) == len(claimHistoryPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &ClaimHistory{}
	if err = randomize.Struct(seed, o, claimHistoryDBTypes, true, claimHistoryColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ClaimHistory struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := ClaimHistories().Count(tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, claimHistoryDBTypes, true, claimHistoryPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize ClaimHistory struct: %s", err)
	}

	if rowsAff, err := o.Update(tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testClaimHistoriesSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(claimHistoryAllColumns) == len(claimHistoryPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &ClaimHistory{}
	if err = randomize.Struct(seed, o, claimHistoryDBTypes, true, claimHistoryColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ClaimHistory struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := ClaimHistories().Count(tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, claimHistoryDBTypes, true, claimHistoryPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize ClaimHistory struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(claimHistoryAllColumns, claimHistoryPrimaryKeyColumns) {
		fields = claimHistoryAllColumns
	} else {
		fields = strmangle.SetComplement(
			claimHistoryAllColumns,
			claimHistoryPrimaryKeyColumns,
		)
		fields = strmangle.SetComplement(
			fields,
			claimHistoryColumnsWithAuto,
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

	slice := ClaimHistorySlice{o}
	if rowsAff, err := slice.UpdateAll(tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testClaimHistoriesUpsert(t *testing.T) {
	t.Parallel()

	if len(claimHistoryAllColumns) == len(claimHistoryPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := ClaimHistory{}
	if err = randomize.Struct(seed, &o, claimHistoryDBTypes, true); err != nil {
		t.Errorf("Unable to randomize ClaimHistory struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(tx, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert ClaimHistory: %s", err)
	}

	count, err := ClaimHistories().Count(tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, claimHistoryDBTypes, false, claimHistoryPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize ClaimHistory struct: %s", err)
	}

	if err = o.Upsert(tx, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert ClaimHistory: %s", err)
	}

	count, err = ClaimHistories().Count(tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
