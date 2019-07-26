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

// TradeCategory is an object representing the database table.
type TradeCategory struct {
	ID        int       `boil:"id" json:"id" toml:"id" yaml:"id"`
	Name      string    `boil:"name" json:"name" toml:"name" yaml:"name"`
	IsActive  bool      `boil:"is_active" json:"is_active" toml:"is_active" yaml:"is_active"`
	IsDeleted bool      `boil:"is_deleted" json:"is_deleted" toml:"is_deleted" yaml:"is_deleted"`
	Created   null.Time `boil:"created" json:"created,omitempty" toml:"created" yaml:"created,omitempty"`
	Updated   null.Time `boil:"updated" json:"updated,omitempty" toml:"updated" yaml:"updated,omitempty"`

	R *tradeCategoryR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L tradeCategoryL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var TradeCategoryColumns = struct {
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

var TradeCategoryWhere = struct {
	ID        whereHelperint
	Name      whereHelperstring
	IsActive  whereHelperbool
	IsDeleted whereHelperbool
	Created   whereHelpernull_Time
	Updated   whereHelpernull_Time
}{
	ID:        whereHelperint{field: "[dbo].[trade_categories].[id]"},
	Name:      whereHelperstring{field: "[dbo].[trade_categories].[name]"},
	IsActive:  whereHelperbool{field: "[dbo].[trade_categories].[is_active]"},
	IsDeleted: whereHelperbool{field: "[dbo].[trade_categories].[is_deleted]"},
	Created:   whereHelpernull_Time{field: "[dbo].[trade_categories].[created]"},
	Updated:   whereHelpernull_Time{field: "[dbo].[trade_categories].[updated]"},
}

// TradeCategoryRels is where relationship names are stored.
var TradeCategoryRels = struct {
	Claims string
	Trades string
}{
	Claims: "Claims",
	Trades: "Trades",
}

// tradeCategoryR is where relationships are stored.
type tradeCategoryR struct {
	Claims ClaimSlice
	Trades TradeSlice
}

// NewStruct creates a new relationship struct
func (*tradeCategoryR) NewStruct() *tradeCategoryR {
	return &tradeCategoryR{}
}

// tradeCategoryL is where Load methods for each relationship are stored.
type tradeCategoryL struct{}

var (
	tradeCategoryAllColumns            = []string{"id", "name", "is_active", "is_deleted", "created", "updated"}
	tradeCategoryColumnsWithAuto       = []string{}
	tradeCategoryColumnsWithoutDefault = []string{"name", "created"}
	tradeCategoryColumnsWithDefault    = []string{"id", "is_active", "is_deleted", "updated"}
	tradeCategoryPrimaryKeyColumns     = []string{"id"}
)

type (
	// TradeCategorySlice is an alias for a slice of pointers to TradeCategory.
	// This should generally be used opposed to []TradeCategory.
	TradeCategorySlice []*TradeCategory

	tradeCategoryQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	tradeCategoryType                 = reflect.TypeOf(&TradeCategory{})
	tradeCategoryMapping              = queries.MakeStructMapping(tradeCategoryType)
	tradeCategoryPrimaryKeyMapping, _ = queries.BindMapping(tradeCategoryType, tradeCategoryMapping, tradeCategoryPrimaryKeyColumns)
	tradeCategoryInsertCacheMut       sync.RWMutex
	tradeCategoryInsertCache          = make(map[string]insertCache)
	tradeCategoryUpdateCacheMut       sync.RWMutex
	tradeCategoryUpdateCache          = make(map[string]updateCache)
	tradeCategoryUpsertCacheMut       sync.RWMutex
	tradeCategoryUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

// One returns a single tradeCategory record from the query.
func (q tradeCategoryQuery) One(exec boil.Executor) (*TradeCategory, error) {
	o := &TradeCategory{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(nil, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for trade_categories")
	}

	return o, nil
}

// All returns all TradeCategory records from the query.
func (q tradeCategoryQuery) All(exec boil.Executor) (TradeCategorySlice, error) {
	var o []*TradeCategory

	err := q.Bind(nil, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to TradeCategory slice")
	}

	return o, nil
}

// Count returns the count of all TradeCategory records in the query.
func (q tradeCategoryQuery) Count(exec boil.Executor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRow(exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count trade_categories rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q tradeCategoryQuery) Exists(exec boil.Executor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRow(exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if trade_categories exists")
	}

	return count > 0, nil
}

// Claims retrieves all the claim's Claims with an executor.
func (o *TradeCategory) Claims(mods ...qm.QueryMod) claimQuery {
	var queryMods []qm.QueryMod
	if len(mods) != 0 {
		queryMods = append(queryMods, mods...)
	}

	queryMods = append(queryMods,
		qm.Where("[dbo].[claims].[trade_category_id]=?", o.ID),
	)

	query := Claims(queryMods...)
	queries.SetFrom(query.Query, "[dbo].[claims]")

	if len(queries.GetSelect(query.Query)) == 0 {
		queries.SetSelect(query.Query, []string{"[dbo].[claims].*"})
	}

	return query
}

// Trades retrieves all the trade's Trades with an executor.
func (o *TradeCategory) Trades(mods ...qm.QueryMod) tradeQuery {
	var queryMods []qm.QueryMod
	if len(mods) != 0 {
		queryMods = append(queryMods, mods...)
	}

	queryMods = append(queryMods,
		qm.Where("[dbo].[trades].[trade_category_id]=?", o.ID),
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
func (tradeCategoryL) LoadClaims(e boil.Executor, singular bool, maybeTradeCategory interface{}, mods queries.Applicator) error {
	var slice []*TradeCategory
	var object *TradeCategory

	if singular {
		object = maybeTradeCategory.(*TradeCategory)
	} else {
		slice = *maybeTradeCategory.(*[]*TradeCategory)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &tradeCategoryR{}
		}
		args = append(args, object.ID)
	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &tradeCategoryR{}
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

	query := NewQuery(qm.From(`dbo.claims`), qm.WhereIn(`trade_category_id in ?`, args...))
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
			foreign.R.TradeCategory = object
		}
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if local.ID == foreign.TradeCategoryID {
				local.R.Claims = append(local.R.Claims, foreign)
				if foreign.R == nil {
					foreign.R = &claimR{}
				}
				foreign.R.TradeCategory = local
				break
			}
		}
	}

	return nil
}

// LoadTrades allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for a 1-M or N-M relationship.
func (tradeCategoryL) LoadTrades(e boil.Executor, singular bool, maybeTradeCategory interface{}, mods queries.Applicator) error {
	var slice []*TradeCategory
	var object *TradeCategory

	if singular {
		object = maybeTradeCategory.(*TradeCategory)
	} else {
		slice = *maybeTradeCategory.(*[]*TradeCategory)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &tradeCategoryR{}
		}
		args = append(args, object.ID)
	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &tradeCategoryR{}
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

	query := NewQuery(qm.From(`dbo.trades`), qm.WhereIn(`trade_category_id in ?`, args...))
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
			foreign.R.TradeCategory = object
		}
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if local.ID == foreign.TradeCategoryID {
				local.R.Trades = append(local.R.Trades, foreign)
				if foreign.R == nil {
					foreign.R = &tradeR{}
				}
				foreign.R.TradeCategory = local
				break
			}
		}
	}

	return nil
}

// AddClaims adds the given related objects to the existing relationships
// of the trade_category, optionally inserting them as new records.
// Appends related to o.R.Claims.
// Sets related.R.TradeCategory appropriately.
func (o *TradeCategory) AddClaims(exec boil.Executor, insert bool, related ...*Claim) error {
	var err error
	for _, rel := range related {
		if insert {
			rel.TradeCategoryID = o.ID
			if err = rel.Insert(exec, boil.Infer()); err != nil {
				return errors.Wrap(err, "failed to insert into foreign table")
			}
		} else {
			updateQuery := fmt.Sprintf(
				"UPDATE [dbo].[claims] SET %s WHERE %s",
				strmangle.SetParamNames("[", "]", 1, []string{"trade_category_id"}),
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

			rel.TradeCategoryID = o.ID
		}
	}

	if o.R == nil {
		o.R = &tradeCategoryR{
			Claims: related,
		}
	} else {
		o.R.Claims = append(o.R.Claims, related...)
	}

	for _, rel := range related {
		if rel.R == nil {
			rel.R = &claimR{
				TradeCategory: o,
			}
		} else {
			rel.R.TradeCategory = o
		}
	}
	return nil
}

// AddTrades adds the given related objects to the existing relationships
// of the trade_category, optionally inserting them as new records.
// Appends related to o.R.Trades.
// Sets related.R.TradeCategory appropriately.
func (o *TradeCategory) AddTrades(exec boil.Executor, insert bool, related ...*Trade) error {
	var err error
	for _, rel := range related {
		if insert {
			rel.TradeCategoryID = o.ID
			if err = rel.Insert(exec, boil.Infer()); err != nil {
				return errors.Wrap(err, "failed to insert into foreign table")
			}
		} else {
			updateQuery := fmt.Sprintf(
				"UPDATE [dbo].[trades] SET %s WHERE %s",
				strmangle.SetParamNames("[", "]", 1, []string{"trade_category_id"}),
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

			rel.TradeCategoryID = o.ID
		}
	}

	if o.R == nil {
		o.R = &tradeCategoryR{
			Trades: related,
		}
	} else {
		o.R.Trades = append(o.R.Trades, related...)
	}

	for _, rel := range related {
		if rel.R == nil {
			rel.R = &tradeR{
				TradeCategory: o,
			}
		} else {
			rel.R.TradeCategory = o
		}
	}
	return nil
}

// TradeCategories retrieves all the records using an executor.
func TradeCategories(mods ...qm.QueryMod) tradeCategoryQuery {
	mods = append(mods, qm.From("[dbo].[trade_categories]"))
	return tradeCategoryQuery{NewQuery(mods...)}
}

// FindTradeCategory retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindTradeCategory(exec boil.Executor, iD int, selectCols ...string) (*TradeCategory, error) {
	tradeCategoryObj := &TradeCategory{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from [dbo].[trade_categories] where [id]=$1", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(nil, exec, tradeCategoryObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from trade_categories")
	}

	return tradeCategoryObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *TradeCategory) Insert(exec boil.Executor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no trade_categories provided for insertion")
	}

	var err error

	nzDefaults := queries.NonZeroDefaultSet(tradeCategoryColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	tradeCategoryInsertCacheMut.RLock()
	cache, cached := tradeCategoryInsertCache[key]
	tradeCategoryInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			tradeCategoryAllColumns,
			tradeCategoryColumnsWithDefault,
			tradeCategoryColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(tradeCategoryType, tradeCategoryMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(tradeCategoryType, tradeCategoryMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO [dbo].[trade_categories] ([%s]) %%sVALUES (%s)%%s", strings.Join(wl, "],["), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO [dbo].[trade_categories] %sDEFAULT VALUES%s"
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
		return errors.Wrap(err, "models: unable to insert into trade_categories")
	}

	if !cached {
		tradeCategoryInsertCacheMut.Lock()
		tradeCategoryInsertCache[key] = cache
		tradeCategoryInsertCacheMut.Unlock()
	}

	return nil
}

// Update uses an executor to update the TradeCategory.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *TradeCategory) Update(exec boil.Executor, columns boil.Columns) (int64, error) {
	var err error
	key := makeCacheKey(columns, nil)
	tradeCategoryUpdateCacheMut.RLock()
	cache, cached := tradeCategoryUpdateCache[key]
	tradeCategoryUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			tradeCategoryAllColumns,
			tradeCategoryPrimaryKeyColumns,
		)
		wl = strmangle.SetComplement(wl, tradeCategoryColumnsWithAuto)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update trade_categories, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE [dbo].[trade_categories] SET %s WHERE %s",
			strmangle.SetParamNames("[", "]", 1, wl),
			strmangle.WhereClause("[", "]", len(wl)+1, tradeCategoryPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(tradeCategoryType, tradeCategoryMapping, append(wl, tradeCategoryPrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "models: unable to update trade_categories row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for trade_categories")
	}

	if !cached {
		tradeCategoryUpdateCacheMut.Lock()
		tradeCategoryUpdateCache[key] = cache
		tradeCategoryUpdateCacheMut.Unlock()
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values.
func (q tradeCategoryQuery) UpdateAll(exec boil.Executor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.Exec(exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for trade_categories")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for trade_categories")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o TradeCategorySlice) UpdateAll(exec boil.Executor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), tradeCategoryPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE [dbo].[trade_categories] SET %s WHERE %s",
		strmangle.SetParamNames("[", "]", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, tradeCategoryPrimaryKeyColumns, len(o)))

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	result, err := exec.Exec(sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in tradeCategory slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all tradeCategory")
	}
	return rowsAff, nil
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
func (o *TradeCategory) Upsert(exec boil.Executor, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no trade_categories provided for upsert")
	}

	nzDefaults := queries.NonZeroDefaultSet(tradeCategoryColumnsWithDefault, o)

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

	tradeCategoryUpsertCacheMut.RLock()
	cache, cached := tradeCategoryUpsertCache[key]
	tradeCategoryUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			tradeCategoryAllColumns,
			tradeCategoryColumnsWithDefault,
			tradeCategoryColumnsWithoutDefault,
			nzDefaults,
		)
		insert = strmangle.SetComplement(insert, tradeCategoryColumnsWithAuto)
		for i, v := range insert {
			if strmangle.ContainsAny(tradeCategoryPrimaryKeyColumns, v) && strmangle.ContainsAny(tradeCategoryColumnsWithDefault, v) {
				insert = append(insert[:i], insert[i+1:]...)
			}
		}
		if len(insert) == 0 {
			return errors.New("models: unable to upsert trade_categories, could not build insert column list")
		}

		ret = strmangle.SetMerge(ret, tradeCategoryColumnsWithAuto)
		ret = strmangle.SetMerge(ret, tradeCategoryColumnsWithDefault)

		update := updateColumns.UpdateColumnSet(
			tradeCategoryAllColumns,
			tradeCategoryPrimaryKeyColumns,
		)
		update = strmangle.SetComplement(update, tradeCategoryColumnsWithAuto)

		if len(update) == 0 {
			return errors.New("models: unable to upsert trade_categories, could not build update column list")
		}

		cache.query = buildUpsertQueryMSSQL(dialect, "trade_categories", tradeCategoryPrimaryKeyColumns, update, insert, ret)

		whitelist := make([]string, len(tradeCategoryPrimaryKeyColumns))
		copy(whitelist, tradeCategoryPrimaryKeyColumns)
		whitelist = append(whitelist, update...)
		whitelist = append(whitelist, insert...)

		cache.valueMapping, err = queries.BindMapping(tradeCategoryType, tradeCategoryMapping, whitelist)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(tradeCategoryType, tradeCategoryMapping, ret)
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
		return errors.Wrap(err, "models: unable to upsert trade_categories")
	}

	if !cached {
		tradeCategoryUpsertCacheMut.Lock()
		tradeCategoryUpsertCache[key] = cache
		tradeCategoryUpsertCacheMut.Unlock()
	}

	return nil
}

// Delete deletes a single TradeCategory record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *TradeCategory) Delete(exec boil.Executor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no TradeCategory provided for delete")
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), tradeCategoryPrimaryKeyMapping)
	sql := "DELETE FROM [dbo].[trade_categories] WHERE [id]=$1"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	result, err := exec.Exec(sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from trade_categories")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for trade_categories")
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q tradeCategoryQuery) DeleteAll(exec boil.Executor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no tradeCategoryQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.Exec(exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from trade_categories")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for trade_categories")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o TradeCategorySlice) DeleteAll(exec boil.Executor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), tradeCategoryPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM [dbo].[trade_categories] WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, tradeCategoryPrimaryKeyColumns, len(o))

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args)
	}

	result, err := exec.Exec(sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from tradeCategory slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for trade_categories")
	}

	return rowsAff, nil
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *TradeCategory) Reload(exec boil.Executor) error {
	ret, err := FindTradeCategory(exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *TradeCategorySlice) ReloadAll(exec boil.Executor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := TradeCategorySlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), tradeCategoryPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT [dbo].[trade_categories].* FROM [dbo].[trade_categories] WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, tradeCategoryPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(nil, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in TradeCategorySlice")
	}

	*o = slice

	return nil
}

// TradeCategoryExists checks if the TradeCategory row exists.
func TradeCategoryExists(exec boil.Executor, iD int) (bool, error) {
	var exists bool
	sql := "select case when exists(select top(1) 1 from [dbo].[trade_categories] where [id]=$1) then 1 else 0 end"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, iD)
	}

	row := exec.QueryRow(sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if trade_categories exists")
	}

	return exists, nil
}