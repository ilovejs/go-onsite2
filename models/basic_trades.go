// Code generated by SQLBoiler (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

import (
	"database/sql"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/pkg/errors"
	"github.com/volatiletech/null"
	"github.com/volatiletech/sqlboiler/boil"
	"github.com/volatiletech/sqlboiler/queries"
	"github.com/volatiletech/sqlboiler/queries/qm"
	"github.com/volatiletech/sqlboiler/queries/qmhelper"
	"github.com/volatiletech/sqlboiler/strmangle"
)

// BasicTrade is an object representing the database table.
type BasicTrade struct {
	ID        int       `boil:"id" json:"id" toml:"id" yaml:"id"`
	Name      string    `boil:"name" json:"name" toml:"name" yaml:"name"`
	IsActive  bool      `boil:"is_active" json:"is_active" toml:"is_active" yaml:"is_active"`
	IsDeleted bool      `boil:"is_deleted" json:"is_deleted" toml:"is_deleted" yaml:"is_deleted"`
	Created   null.Time `boil:"created" json:"created,omitempty" toml:"created" yaml:"created,omitempty"`
	Updated   null.Time `boil:"updated" json:"updated,omitempty" toml:"updated" yaml:"updated,omitempty"`

	R *basicTradeR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L basicTradeL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var BasicTradeColumns = struct {
	ID        string
	Name      string
	IsActive  string
	IsDeleted string
	Created   string
	Updated   string
}{
	ID:        "id",
	Name:      "name",
	IsActive:  "is_active",
	IsDeleted: "is_deleted",
	Created:   "created",
	Updated:   "updated",
}

// Generated where

type whereHelperint struct{ field string }

func (w whereHelperint) EQ(x int) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.EQ, x) }
func (w whereHelperint) NEQ(x int) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.NEQ, x) }
func (w whereHelperint) LT(x int) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.LT, x) }
func (w whereHelperint) LTE(x int) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.LTE, x) }
func (w whereHelperint) GT(x int) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.GT, x) }
func (w whereHelperint) GTE(x int) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.GTE, x) }

type whereHelperstring struct{ field string }

func (w whereHelperstring) EQ(x string) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.EQ, x) }
func (w whereHelperstring) NEQ(x string) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.NEQ, x) }
func (w whereHelperstring) LT(x string) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.LT, x) }
func (w whereHelperstring) LTE(x string) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.LTE, x) }
func (w whereHelperstring) GT(x string) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.GT, x) }
func (w whereHelperstring) GTE(x string) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.GTE, x) }

type whereHelperbool struct{ field string }

func (w whereHelperbool) EQ(x bool) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.EQ, x) }
func (w whereHelperbool) NEQ(x bool) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.NEQ, x) }
func (w whereHelperbool) LT(x bool) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.LT, x) }
func (w whereHelperbool) LTE(x bool) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.LTE, x) }
func (w whereHelperbool) GT(x bool) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.GT, x) }
func (w whereHelperbool) GTE(x bool) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.GTE, x) }

type whereHelpernull_Time struct{ field string }

func (w whereHelpernull_Time) EQ(x null.Time) qm.QueryMod {
	return qmhelper.WhereNullEQ(w.field, false, x)
}
func (w whereHelpernull_Time) NEQ(x null.Time) qm.QueryMod {
	return qmhelper.WhereNullEQ(w.field, true, x)
}
func (w whereHelpernull_Time) IsNull() qm.QueryMod    { return qmhelper.WhereIsNull(w.field) }
func (w whereHelpernull_Time) IsNotNull() qm.QueryMod { return qmhelper.WhereIsNotNull(w.field) }
func (w whereHelpernull_Time) LT(x null.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LT, x)
}
func (w whereHelpernull_Time) LTE(x null.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LTE, x)
}
func (w whereHelpernull_Time) GT(x null.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GT, x)
}
func (w whereHelpernull_Time) GTE(x null.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GTE, x)
}

var BasicTradeWhere = struct {
	ID        whereHelperint
	Name      whereHelperstring
	IsActive  whereHelperbool
	IsDeleted whereHelperbool
	Created   whereHelpernull_Time
	Updated   whereHelpernull_Time
}{
	ID:        whereHelperint{field: "[dbo].[basic_trades].[id]"},
	Name:      whereHelperstring{field: "[dbo].[basic_trades].[name]"},
	IsActive:  whereHelperbool{field: "[dbo].[basic_trades].[is_active]"},
	IsDeleted: whereHelperbool{field: "[dbo].[basic_trades].[is_deleted]"},
	Created:   whereHelpernull_Time{field: "[dbo].[basic_trades].[created]"},
	Updated:   whereHelpernull_Time{field: "[dbo].[basic_trades].[updated]"},
}

// BasicTradeRels is where relationship names are stored.
var BasicTradeRels = struct {
	Claims string
	Trades string
}{
	Claims: "Claims",
	Trades: "Trades",
}

// basicTradeR is where relationships are stored.
type basicTradeR struct {
	Claims ClaimSlice
	Trades TradeSlice
}

// NewStruct creates a new relationship struct
func (*basicTradeR) NewStruct() *basicTradeR {
	return &basicTradeR{}
}

// basicTradeL is where Load methods for each relationship are stored.
type basicTradeL struct{}

var (
	basicTradeAllColumns            = []string{"id", "name", "is_active", "is_deleted", "created", "updated"}
	basicTradeColumnsWithAuto       = []string{}
	basicTradeColumnsWithoutDefault = []string{"name", "created"}
	basicTradeColumnsWithDefault    = []string{"id", "is_active", "is_deleted", "updated"}
	basicTradePrimaryKeyColumns     = []string{"id"}
)

type (
	// BasicTradeSlice is an alias for a slice of pointers to BasicTrade.
	// This should generally be used opposed to []BasicTrade.
	BasicTradeSlice []*BasicTrade

	basicTradeQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	basicTradeType                 = reflect.TypeOf(&BasicTrade{})
	basicTradeMapping              = queries.MakeStructMapping(basicTradeType)
	basicTradePrimaryKeyMapping, _ = queries.BindMapping(basicTradeType, basicTradeMapping, basicTradePrimaryKeyColumns)
	basicTradeInsertCacheMut       sync.RWMutex
	basicTradeInsertCache          = make(map[string]insertCache)
	basicTradeUpdateCacheMut       sync.RWMutex
	basicTradeUpdateCache          = make(map[string]updateCache)
	basicTradeUpsertCacheMut       sync.RWMutex
	basicTradeUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

// One returns a single basicTrade record from the query.
func (q basicTradeQuery) One(exec boil.Executor) (*BasicTrade, error) {
	o := &BasicTrade{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(nil, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for basic_trades")
	}

	return o, nil
}

// All returns all BasicTrade records from the query.
func (q basicTradeQuery) All(exec boil.Executor) (BasicTradeSlice, error) {
	var o []*BasicTrade

	err := q.Bind(nil, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to BasicTrade slice")
	}

	return o, nil
}

// Count returns the count of all BasicTrade records in the query.
func (q basicTradeQuery) Count(exec boil.Executor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRow(exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count basic_trades rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q basicTradeQuery) Exists(exec boil.Executor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRow(exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if basic_trades exists")
	}

	return count > 0, nil
}

// Claims retrieves all the claim's Claims with an executor.
func (o *BasicTrade) Claims(mods ...qm.QueryMod) claimQuery {
	var queryMods []qm.QueryMod
	if len(mods) != 0 {
		queryMods = append(queryMods, mods...)
	}

	queryMods = append(queryMods,
		qm.Where("[dbo].[claims].[basic_trade_id]=?", o.ID),
	)

	query := Claims(queryMods...)
	queries.SetFrom(query.Query, "[dbo].[claims]")

	if len(queries.GetSelect(query.Query)) == 0 {
		queries.SetSelect(query.Query, []string{"[dbo].[claims].*"})
	}

	return query
}

// Trades retrieves all the trade's Trades with an executor.
func (o *BasicTrade) Trades(mods ...qm.QueryMod) tradeQuery {
	var queryMods []qm.QueryMod
	if len(mods) != 0 {
		queryMods = append(queryMods, mods...)
	}

	queryMods = append(queryMods,
		qm.Where("[dbo].[trades].[basic_trade_id]=?", o.ID),
	)

	query := Trades(queryMods...)
	queries.SetFrom(query.Query, "[dbo].[trades]")

	if len(queries.GetSelect(query.Query)) == 0 {
		queries.SetSelect(query.Query, []string{"[dbo].[trades].*"})
	}

	return query
}

// LoadClaims allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for a 1-M or N-M relationship.
func (basicTradeL) LoadClaims(e boil.Executor, singular bool, maybeBasicTrade interface{}, mods queries.Applicator) error {
	var slice []*BasicTrade
	var object *BasicTrade

	if singular {
		object = maybeBasicTrade.(*BasicTrade)
	} else {
		slice = *maybeBasicTrade.(*[]*BasicTrade)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &basicTradeR{}
		}
		args = append(args, object.ID)
	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &basicTradeR{}
			}

			for _, a := range args {
				if a == obj.ID {
					continue Outer
				}
			}

			args = append(args, obj.ID)
		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(qm.From(`dbo.claims`), qm.WhereIn(`basic_trade_id in ?`, args...))
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.Query(e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load claims")
	}

	var resultSlice []*Claim
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice claims")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results in eager load on claims")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for claims")
	}

	if singular {
		object.R.Claims = resultSlice
		for _, foreign := range resultSlice {
			if foreign.R == nil {
				foreign.R = &claimR{}
			}
			foreign.R.BasicTrade = object
		}
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if local.ID == foreign.BasicTradeID {
				local.R.Claims = append(local.R.Claims, foreign)
				if foreign.R == nil {
					foreign.R = &claimR{}
				}
				foreign.R.BasicTrade = local
				break
			}
		}
	}

	return nil
}

// LoadTrades allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for a 1-M or N-M relationship.
func (basicTradeL) LoadTrades(e boil.Executor, singular bool, maybeBasicTrade interface{}, mods queries.Applicator) error {
	var slice []*BasicTrade
	var object *BasicTrade

	if singular {
		object = maybeBasicTrade.(*BasicTrade)
	} else {
		slice = *maybeBasicTrade.(*[]*BasicTrade)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &basicTradeR{}
		}
		args = append(args, object.ID)
	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &basicTradeR{}
			}

			for _, a := range args {
				if a == obj.ID {
					continue Outer
				}
			}

			args = append(args, obj.ID)
		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(qm.From(`dbo.trades`), qm.WhereIn(`basic_trade_id in ?`, args...))
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.Query(e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load trades")
	}

	var resultSlice []*Trade
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice trades")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results in eager load on trades")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for trades")
	}

	if singular {
		object.R.Trades = resultSlice
		for _, foreign := range resultSlice {
			if foreign.R == nil {
				foreign.R = &tradeR{}
			}
			foreign.R.BasicTrade = object
		}
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if local.ID == foreign.BasicTradeID {
				local.R.Trades = append(local.R.Trades, foreign)
				if foreign.R == nil {
					foreign.R = &tradeR{}
				}
				foreign.R.BasicTrade = local
				break
			}
		}
	}

	return nil
}

// AddClaims adds the given related objects to the existing relationships
// of the basic_trade, optionally inserting them as new records.
// Appends related to o.R.Claims.
// Sets related.R.BasicTrade appropriately.
func (o *BasicTrade) AddClaims(exec boil.Executor, insert bool, related ...*Claim) error {
	var err error
	for _, rel := range related {
		if insert {
			rel.BasicTradeID = o.ID
			if err = rel.Insert(exec, boil.Infer()); err != nil {
				return errors.Wrap(err, "failed to insert into foreign table")
			}
		} else {
			updateQuery := fmt.Sprintf(
				"UPDATE [dbo].[claims] SET %s WHERE %s",
				strmangle.SetParamNames("[", "]", 1, []string{"basic_trade_id"}),
				strmangle.WhereClause("[", "]", 2, claimPrimaryKeyColumns),
			)
			values := []interface{}{o.ID, rel.ID}

			if boil.DebugMode {
				fmt.Fprintln(boil.DebugWriter, updateQuery)
				fmt.Fprintln(boil.DebugWriter, values)
			}

			if _, err = exec.Exec(updateQuery, values...); err != nil {
				return errors.Wrap(err, "failed to update foreign table")
			}

			rel.BasicTradeID = o.ID
		}
	}

	if o.R == nil {
		o.R = &basicTradeR{
			Claims: related,
		}
	} else {
		o.R.Claims = append(o.R.Claims, related...)
	}

	for _, rel := range related {
		if rel.R == nil {
			rel.R = &claimR{
				BasicTrade: o,
			}
		} else {
			rel.R.BasicTrade = o
		}
	}
	return nil
}

// AddTrades adds the given related objects to the existing relationships
// of the basic_trade, optionally inserting them as new records.
// Appends related to o.R.Trades.
// Sets related.R.BasicTrade appropriately.
func (o *BasicTrade) AddTrades(exec boil.Executor, insert bool, related ...*Trade) error {
	var err error
	for _, rel := range related {
		if insert {
			rel.BasicTradeID = o.ID
			if err = rel.Insert(exec, boil.Infer()); err != nil {
				return errors.Wrap(err, "failed to insert into foreign table")
			}
		} else {
			updateQuery := fmt.Sprintf(
				"UPDATE [dbo].[trades] SET %s WHERE %s",
				strmangle.SetParamNames("[", "]", 1, []string{"basic_trade_id"}),
				strmangle.WhereClause("[", "]", 2, tradePrimaryKeyColumns),
			)
			values := []interface{}{o.ID, rel.ID}

			if boil.DebugMode {
				fmt.Fprintln(boil.DebugWriter, updateQuery)
				fmt.Fprintln(boil.DebugWriter, values)
			}

			if _, err = exec.Exec(updateQuery, values...); err != nil {
				return errors.Wrap(err, "failed to update foreign table")
			}

			rel.BasicTradeID = o.ID
		}
	}

	if o.R == nil {
		o.R = &basicTradeR{
			Trades: related,
		}
	} else {
		o.R.Trades = append(o.R.Trades, related...)
	}

	for _, rel := range related {
		if rel.R == nil {
			rel.R = &tradeR{
				BasicTrade: o,
			}
		} else {
			rel.R.BasicTrade = o
		}
	}
	return nil
}

// BasicTrades retrieves all the records using an executor.
func BasicTrades(mods ...qm.QueryMod) basicTradeQuery {
	mods = append(mods, qm.From("[dbo].[basic_trades]"))
	return basicTradeQuery{NewQuery(mods...)}
}

// FindBasicTrade retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindBasicTrade(exec boil.Executor, iD int, selectCols ...string) (*BasicTrade, error) {
	basicTradeObj := &BasicTrade{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from [dbo].[basic_trades] where [id]=$1", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(nil, exec, basicTradeObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from basic_trades")
	}

	return basicTradeObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *BasicTrade) Insert(exec boil.Executor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no basic_trades provided for insertion")
	}

	var err error

	nzDefaults := queries.NonZeroDefaultSet(basicTradeColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	basicTradeInsertCacheMut.RLock()
	cache, cached := basicTradeInsertCache[key]
	basicTradeInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			basicTradeAllColumns,
			basicTradeColumnsWithDefault,
			basicTradeColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(basicTradeType, basicTradeMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(basicTradeType, basicTradeMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO [dbo].[basic_trades] ([%s]) %%sVALUES (%s)%%s", strings.Join(wl, "],["), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO [dbo].[basic_trades] %sDEFAULT VALUES%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			queryOutput = fmt.Sprintf("OUTPUT INSERTED.[%s] ", strings.Join(returnColumns, "],INSERTED.["))
		}

		cache.query = fmt.Sprintf(cache.query, queryOutput, queryReturning)
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.query)
		fmt.Fprintln(boil.DebugWriter, vals)
	}

	if len(cache.retMapping) != 0 {
		err = exec.QueryRow(cache.query, vals...).Scan(queries.PtrsFromMapping(value, cache.retMapping)...)
	} else {
		_, err = exec.Exec(cache.query, vals...)
	}

	if err != nil {
		return errors.Wrap(err, "models: unable to insert into basic_trades")
	}

	if !cached {
		basicTradeInsertCacheMut.Lock()
		basicTradeInsertCache[key] = cache
		basicTradeInsertCacheMut.Unlock()
	}

	return nil
}

// Update uses an executor to update the BasicTrade.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *BasicTrade) Update(exec boil.Executor, columns boil.Columns) (int64, error) {
	var err error
	key := makeCacheKey(columns, nil)
	basicTradeUpdateCacheMut.RLock()
	cache, cached := basicTradeUpdateCache[key]
	basicTradeUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			basicTradeAllColumns,
			basicTradePrimaryKeyColumns,
		)
		wl = strmangle.SetComplement(wl, basicTradeColumnsWithAuto)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update basic_trades, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE [dbo].[basic_trades] SET %s WHERE %s",
			strmangle.SetParamNames("[", "]", 1, wl),
			strmangle.WhereClause("[", "]", len(wl)+1, basicTradePrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(basicTradeType, basicTradeMapping, append(wl, basicTradePrimaryKeyColumns...))
		if err != nil {
			return 0, err
		}
	}

	values := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), cache.valueMapping)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.query)
		fmt.Fprintln(boil.DebugWriter, values)
	}

	var result sql.Result
	result, err = exec.Exec(cache.query, values...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update basic_trades row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for basic_trades")
	}

	if !cached {
		basicTradeUpdateCacheMut.Lock()
		basicTradeUpdateCache[key] = cache
		basicTradeUpdateCacheMut.Unlock()
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values.
func (q basicTradeQuery) UpdateAll(exec boil.Executor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.Exec(exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for basic_trades")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for basic_trades")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o BasicTradeSlice) UpdateAll(exec boil.Executor, cols M) (int64, error) {
	ln := int64(len(o))
	if ln == 0 {
		return 0, nil
	}

	if len(cols) == 0 {
		return 0, errors.New("models: update all requires at least one column argument")
	}

	colNames := make([]string, len(cols))
	args := make([]interface{}, len(cols))

	i := 0
	for name, value := range cols {
		colNames[i] = name
		args[i] = value
		i++
	}

	// Append all of the primary key values for each column
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), basicTradePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE [dbo].[basic_trades] SET %s WHERE %s",
		strmangle.SetParamNames("[", "]", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, basicTradePrimaryKeyColumns, len(o)))

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	result, err := exec.Exec(sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in basicTrade slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all basicTrade")
	}
	return rowsAff, nil
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
func (o *BasicTrade) Upsert(exec boil.Executor, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no basic_trades provided for upsert")
	}

	nzDefaults := queries.NonZeroDefaultSet(basicTradeColumnsWithDefault, o)

	// Build cache key in-line uglily - mysql vs psql problems
	buf := strmangle.GetBuffer()
	buf.WriteString(strconv.Itoa(updateColumns.Kind))
	for _, c := range updateColumns.Cols {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	buf.WriteString(strconv.Itoa(insertColumns.Kind))
	for _, c := range insertColumns.Cols {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	for _, c := range nzDefaults {
		buf.WriteString(c)
	}
	key := buf.String()
	strmangle.PutBuffer(buf)

	basicTradeUpsertCacheMut.RLock()
	cache, cached := basicTradeUpsertCache[key]
	basicTradeUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			basicTradeAllColumns,
			basicTradeColumnsWithDefault,
			basicTradeColumnsWithoutDefault,
			nzDefaults,
		)
		insert = strmangle.SetComplement(insert, basicTradeColumnsWithAuto)
		for i, v := range insert {
			if strmangle.ContainsAny(basicTradePrimaryKeyColumns, v) && strmangle.ContainsAny(basicTradeColumnsWithDefault, v) {
				insert = append(insert[:i], insert[i+1:]...)
			}
		}
		if len(insert) == 0 {
			return errors.New("models: unable to upsert basic_trades, could not build insert column list")
		}

		ret = strmangle.SetMerge(ret, basicTradeColumnsWithAuto)
		ret = strmangle.SetMerge(ret, basicTradeColumnsWithDefault)

		update := updateColumns.UpdateColumnSet(
			basicTradeAllColumns,
			basicTradePrimaryKeyColumns,
		)
		update = strmangle.SetComplement(update, basicTradeColumnsWithAuto)

		if len(update) == 0 {
			return errors.New("models: unable to upsert basic_trades, could not build update column list")
		}

		cache.query = buildUpsertQueryMSSQL(dialect, "basic_trades", basicTradePrimaryKeyColumns, update, insert, ret)

		whitelist := make([]string, len(basicTradePrimaryKeyColumns))
		copy(whitelist, basicTradePrimaryKeyColumns)
		whitelist = append(whitelist, update...)
		whitelist = append(whitelist, insert...)

		cache.valueMapping, err = queries.BindMapping(basicTradeType, basicTradeMapping, whitelist)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(basicTradeType, basicTradeMapping, ret)
			if err != nil {
				return err
			}
		}
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)
	var returns []interface{}
	if len(cache.retMapping) != 0 {
		returns = queries.PtrsFromMapping(value, cache.retMapping)
	}

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.query)
		fmt.Fprintln(boil.DebugWriter, vals)
	}

	if len(cache.retMapping) != 0 {
		err = exec.QueryRow(cache.query, vals...).Scan(returns...)
		if err == sql.ErrNoRows {
			err = nil // MSSQL doesn't return anything when there's no update
		}
	} else {
		_, err = exec.Exec(cache.query, vals...)
	}
	if err != nil {
		return errors.Wrap(err, "models: unable to upsert basic_trades")
	}

	if !cached {
		basicTradeUpsertCacheMut.Lock()
		basicTradeUpsertCache[key] = cache
		basicTradeUpsertCacheMut.Unlock()
	}

	return nil
}

// Delete deletes a single BasicTrade record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *BasicTrade) Delete(exec boil.Executor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no BasicTrade provided for delete")
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), basicTradePrimaryKeyMapping)
	sql := "DELETE FROM [dbo].[basic_trades] WHERE [id]=$1"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	result, err := exec.Exec(sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from basic_trades")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for basic_trades")
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q basicTradeQuery) DeleteAll(exec boil.Executor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no basicTradeQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.Exec(exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from basic_trades")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for basic_trades")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o BasicTradeSlice) DeleteAll(exec boil.Executor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), basicTradePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM [dbo].[basic_trades] WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, basicTradePrimaryKeyColumns, len(o))

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args)
	}

	result, err := exec.Exec(sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from basicTrade slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for basic_trades")
	}

	return rowsAff, nil
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *BasicTrade) Reload(exec boil.Executor) error {
	ret, err := FindBasicTrade(exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *BasicTradeSlice) ReloadAll(exec boil.Executor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := BasicTradeSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), basicTradePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT [dbo].[basic_trades].* FROM [dbo].[basic_trades] WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, basicTradePrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(nil, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in BasicTradeSlice")
	}

	*o = slice

	return nil
}

// BasicTradeExists checks if the BasicTrade row exists.
func BasicTradeExists(exec boil.Executor, iD int) (bool, error) {
	var exists bool
	sql := "select case when exists(select top(1) 1 from [dbo].[basic_trades] where [id]=$1) then 1 else 0 end"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, iD)
	}

	row := exec.QueryRow(sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if basic_trades exists")
	}

	return exists, nil
}
