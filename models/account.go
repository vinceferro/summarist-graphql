// Code generated by SQLBoiler 4.2.0 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
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

	"github.com/friendsofgo/errors"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"github.com/volatiletech/sqlboiler/v4/queries/qmhelper"
	"github.com/volatiletech/strmangle"
)

// Account is an object representing the database table.
type Account struct {
	ID       string      `boil:"id" json:"id" toml:"id" yaml:"id"`
	Email    string      `boil:"email" json:"email" toml:"email" yaml:"email"`
	Password string      `boil:"password" json:"password" toml:"password" yaml:"password"`
	Role     null.String `boil:"role" json:"role,omitempty" toml:"role" yaml:"role,omitempty"`

	R *accountR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L accountL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var AccountColumns = struct {
	ID       string
	Email    string
	Password string
	Role     string
}{
	ID:       "id",
	Email:    "email",
	Password: "password",
	Role:     "role",
}

// Generated where

type whereHelperstring struct{ field string }

func (w whereHelperstring) EQ(x string) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.EQ, x) }
func (w whereHelperstring) NEQ(x string) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.NEQ, x) }
func (w whereHelperstring) LT(x string) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.LT, x) }
func (w whereHelperstring) LTE(x string) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.LTE, x) }
func (w whereHelperstring) GT(x string) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.GT, x) }
func (w whereHelperstring) GTE(x string) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.GTE, x) }
func (w whereHelperstring) IN(slice []string) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereIn(fmt.Sprintf("%s IN ?", w.field), values...)
}
func (w whereHelperstring) NIN(slice []string) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereNotIn(fmt.Sprintf("%s NOT IN ?", w.field), values...)
}

type whereHelpernull_String struct{ field string }

func (w whereHelpernull_String) EQ(x null.String) qm.QueryMod {
	return qmhelper.WhereNullEQ(w.field, false, x)
}
func (w whereHelpernull_String) NEQ(x null.String) qm.QueryMod {
	return qmhelper.WhereNullEQ(w.field, true, x)
}
func (w whereHelpernull_String) IsNull() qm.QueryMod    { return qmhelper.WhereIsNull(w.field) }
func (w whereHelpernull_String) IsNotNull() qm.QueryMod { return qmhelper.WhereIsNotNull(w.field) }
func (w whereHelpernull_String) LT(x null.String) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LT, x)
}
func (w whereHelpernull_String) LTE(x null.String) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LTE, x)
}
func (w whereHelpernull_String) GT(x null.String) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GT, x)
}
func (w whereHelpernull_String) GTE(x null.String) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GTE, x)
}

var AccountWhere = struct {
	ID       whereHelperstring
	Email    whereHelperstring
	Password whereHelperstring
	Role     whereHelpernull_String
}{
	ID:       whereHelperstring{field: "\"account\".\"id\""},
	Email:    whereHelperstring{field: "\"account\".\"email\""},
	Password: whereHelperstring{field: "\"account\".\"password\""},
	Role:     whereHelpernull_String{field: "\"account\".\"role\""},
}

// AccountRels is where relationship names are stored.
var AccountRels = struct {
	Books string
}{
	Books: "Books",
}

// accountR is where relationships are stored.
type accountR struct {
	Books BookSlice `boil:"Books" json:"Books" toml:"Books" yaml:"Books"`
}

// NewStruct creates a new relationship struct
func (*accountR) NewStruct() *accountR {
	return &accountR{}
}

// accountL is where Load methods for each relationship are stored.
type accountL struct{}

var (
	accountAllColumns            = []string{"id", "email", "password", "role"}
	accountColumnsWithoutDefault = []string{"email", "password", "role"}
	accountColumnsWithDefault    = []string{"id"}
	accountPrimaryKeyColumns     = []string{"id"}
)

type (
	// AccountSlice is an alias for a slice of pointers to Account.
	// This should generally be used opposed to []Account.
	AccountSlice []*Account
	// AccountHook is the signature for custom Account hook methods
	AccountHook func(boil.Executor, *Account) error

	accountQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	accountType                 = reflect.TypeOf(&Account{})
	accountMapping              = queries.MakeStructMapping(accountType)
	accountPrimaryKeyMapping, _ = queries.BindMapping(accountType, accountMapping, accountPrimaryKeyColumns)
	accountInsertCacheMut       sync.RWMutex
	accountInsertCache          = make(map[string]insertCache)
	accountUpdateCacheMut       sync.RWMutex
	accountUpdateCache          = make(map[string]updateCache)
	accountUpsertCacheMut       sync.RWMutex
	accountUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var accountBeforeInsertHooks []AccountHook
var accountBeforeUpdateHooks []AccountHook
var accountBeforeDeleteHooks []AccountHook
var accountBeforeUpsertHooks []AccountHook

var accountAfterInsertHooks []AccountHook
var accountAfterSelectHooks []AccountHook
var accountAfterUpdateHooks []AccountHook
var accountAfterDeleteHooks []AccountHook
var accountAfterUpsertHooks []AccountHook

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *Account) doBeforeInsertHooks(exec boil.Executor) (err error) {
	for _, hook := range accountBeforeInsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *Account) doBeforeUpdateHooks(exec boil.Executor) (err error) {
	for _, hook := range accountBeforeUpdateHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *Account) doBeforeDeleteHooks(exec boil.Executor) (err error) {
	for _, hook := range accountBeforeDeleteHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *Account) doBeforeUpsertHooks(exec boil.Executor) (err error) {
	for _, hook := range accountBeforeUpsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *Account) doAfterInsertHooks(exec boil.Executor) (err error) {
	for _, hook := range accountAfterInsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterSelectHooks executes all "after Select" hooks.
func (o *Account) doAfterSelectHooks(exec boil.Executor) (err error) {
	for _, hook := range accountAfterSelectHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *Account) doAfterUpdateHooks(exec boil.Executor) (err error) {
	for _, hook := range accountAfterUpdateHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *Account) doAfterDeleteHooks(exec boil.Executor) (err error) {
	for _, hook := range accountAfterDeleteHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *Account) doAfterUpsertHooks(exec boil.Executor) (err error) {
	for _, hook := range accountAfterUpsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddAccountHook registers your hook function for all future operations.
func AddAccountHook(hookPoint boil.HookPoint, accountHook AccountHook) {
	switch hookPoint {
	case boil.BeforeInsertHook:
		accountBeforeInsertHooks = append(accountBeforeInsertHooks, accountHook)
	case boil.BeforeUpdateHook:
		accountBeforeUpdateHooks = append(accountBeforeUpdateHooks, accountHook)
	case boil.BeforeDeleteHook:
		accountBeforeDeleteHooks = append(accountBeforeDeleteHooks, accountHook)
	case boil.BeforeUpsertHook:
		accountBeforeUpsertHooks = append(accountBeforeUpsertHooks, accountHook)
	case boil.AfterInsertHook:
		accountAfterInsertHooks = append(accountAfterInsertHooks, accountHook)
	case boil.AfterSelectHook:
		accountAfterSelectHooks = append(accountAfterSelectHooks, accountHook)
	case boil.AfterUpdateHook:
		accountAfterUpdateHooks = append(accountAfterUpdateHooks, accountHook)
	case boil.AfterDeleteHook:
		accountAfterDeleteHooks = append(accountAfterDeleteHooks, accountHook)
	case boil.AfterUpsertHook:
		accountAfterUpsertHooks = append(accountAfterUpsertHooks, accountHook)
	}
}

// OneG returns a single account record from the query using the global executor.
func (q accountQuery) OneG() (*Account, error) {
	return q.One(boil.GetDB())
}

// One returns a single account record from the query.
func (q accountQuery) One(exec boil.Executor) (*Account, error) {
	o := &Account{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(nil, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for account")
	}

	if err := o.doAfterSelectHooks(exec); err != nil {
		return o, err
	}

	return o, nil
}

// AllG returns all Account records from the query using the global executor.
func (q accountQuery) AllG() (AccountSlice, error) {
	return q.All(boil.GetDB())
}

// All returns all Account records from the query.
func (q accountQuery) All(exec boil.Executor) (AccountSlice, error) {
	var o []*Account

	err := q.Bind(nil, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to Account slice")
	}

	if len(accountAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// CountG returns the count of all Account records in the query, and panics on error.
func (q accountQuery) CountG() (int64, error) {
	return q.Count(boil.GetDB())
}

// Count returns the count of all Account records in the query.
func (q accountQuery) Count(exec boil.Executor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRow(exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count account rows")
	}

	return count, nil
}

// ExistsG checks if the row exists in the table, and panics on error.
func (q accountQuery) ExistsG() (bool, error) {
	return q.Exists(boil.GetDB())
}

// Exists checks if the row exists in the table.
func (q accountQuery) Exists(exec boil.Executor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRow(exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if account exists")
	}

	return count > 0, nil
}

// Books retrieves all the book's Books with an executor.
func (o *Account) Books(mods ...qm.QueryMod) bookQuery {
	var queryMods []qm.QueryMod
	if len(mods) != 0 {
		queryMods = append(queryMods, mods...)
	}

	queryMods = append(queryMods,
		qm.InnerJoin("\"book_star\" on \"book\".\"id\" = \"book_star\".\"book_id\""),
		qm.Where("\"book_star\".\"account_id\"=?", o.ID),
	)

	query := Books(queryMods...)
	queries.SetFrom(query.Query, "\"book\"")

	if len(queries.GetSelect(query.Query)) == 0 {
		queries.SetSelect(query.Query, []string{"\"book\".*"})
	}

	return query
}

// LoadBooks allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for a 1-M or N-M relationship.
func (accountL) LoadBooks(e boil.Executor, singular bool, maybeAccount interface{}, mods queries.Applicator) error {
	var slice []*Account
	var object *Account

	if singular {
		object = maybeAccount.(*Account)
	} else {
		slice = *maybeAccount.(*[]*Account)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &accountR{}
		}
		args = append(args, object.ID)
	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &accountR{}
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

	query := NewQuery(
		qm.Select("\"book\".*, \"a\".\"account_id\""),
		qm.From("\"book\""),
		qm.InnerJoin("\"book_star\" as \"a\" on \"book\".\"id\" = \"a\".\"book_id\""),
		qm.WhereIn("\"a\".\"account_id\" in ?", args...),
	)
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.Query(e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load book")
	}

	var resultSlice []*Book

	var localJoinCols []string
	for results.Next() {
		one := new(Book)
		var localJoinCol string

		err = results.Scan(&one.ID, &one.Title, &one.Author, &one.Isbn, &one.PublishedAt, &one.Publisher, &one.CoverURL, &one.Overview, &one.KeyInsights, &one.CategoryID, &localJoinCol)
		if err != nil {
			return errors.Wrap(err, "failed to scan eager loaded results for book")
		}
		if err = results.Err(); err != nil {
			return errors.Wrap(err, "failed to plebian-bind eager loaded slice book")
		}

		resultSlice = append(resultSlice, one)
		localJoinCols = append(localJoinCols, localJoinCol)
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results in eager load on book")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for book")
	}

	if len(bookAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(e); err != nil {
				return err
			}
		}
	}
	if singular {
		object.R.Books = resultSlice
		for _, foreign := range resultSlice {
			if foreign.R == nil {
				foreign.R = &bookR{}
			}
			foreign.R.Accounts = append(foreign.R.Accounts, object)
		}
		return nil
	}

	for i, foreign := range resultSlice {
		localJoinCol := localJoinCols[i]
		for _, local := range slice {
			if local.ID == localJoinCol {
				local.R.Books = append(local.R.Books, foreign)
				if foreign.R == nil {
					foreign.R = &bookR{}
				}
				foreign.R.Accounts = append(foreign.R.Accounts, local)
				break
			}
		}
	}

	return nil
}

// AddBooksG adds the given related objects to the existing relationships
// of the account, optionally inserting them as new records.
// Appends related to o.R.Books.
// Sets related.R.Accounts appropriately.
// Uses the global database handle.
func (o *Account) AddBooksG(insert bool, related ...*Book) error {
	return o.AddBooks(boil.GetDB(), insert, related...)
}

// AddBooks adds the given related objects to the existing relationships
// of the account, optionally inserting them as new records.
// Appends related to o.R.Books.
// Sets related.R.Accounts appropriately.
func (o *Account) AddBooks(exec boil.Executor, insert bool, related ...*Book) error {
	var err error
	for _, rel := range related {
		if insert {
			if err = rel.Insert(exec, boil.Infer()); err != nil {
				return errors.Wrap(err, "failed to insert into foreign table")
			}
		}
	}

	for _, rel := range related {
		query := "insert into \"book_star\" (\"account_id\", \"book_id\") values ($1, $2)"
		values := []interface{}{o.ID, rel.ID}

		if boil.DebugMode {
			fmt.Fprintln(boil.DebugWriter, query)
			fmt.Fprintln(boil.DebugWriter, values)
		}
		_, err = exec.Exec(query, values...)
		if err != nil {
			return errors.Wrap(err, "failed to insert into join table")
		}
	}
	if o.R == nil {
		o.R = &accountR{
			Books: related,
		}
	} else {
		o.R.Books = append(o.R.Books, related...)
	}

	for _, rel := range related {
		if rel.R == nil {
			rel.R = &bookR{
				Accounts: AccountSlice{o},
			}
		} else {
			rel.R.Accounts = append(rel.R.Accounts, o)
		}
	}
	return nil
}

// SetBooksG removes all previously related items of the
// account replacing them completely with the passed
// in related items, optionally inserting them as new records.
// Sets o.R.Accounts's Books accordingly.
// Replaces o.R.Books with related.
// Sets related.R.Accounts's Books accordingly.
// Uses the global database handle.
func (o *Account) SetBooksG(insert bool, related ...*Book) error {
	return o.SetBooks(boil.GetDB(), insert, related...)
}

// SetBooks removes all previously related items of the
// account replacing them completely with the passed
// in related items, optionally inserting them as new records.
// Sets o.R.Accounts's Books accordingly.
// Replaces o.R.Books with related.
// Sets related.R.Accounts's Books accordingly.
func (o *Account) SetBooks(exec boil.Executor, insert bool, related ...*Book) error {
	query := "delete from \"book_star\" where \"account_id\" = $1"
	values := []interface{}{o.ID}
	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, query)
		fmt.Fprintln(boil.DebugWriter, values)
	}
	_, err := exec.Exec(query, values...)
	if err != nil {
		return errors.Wrap(err, "failed to remove relationships before set")
	}

	removeBooksFromAccountsSlice(o, related)
	if o.R != nil {
		o.R.Books = nil
	}
	return o.AddBooks(exec, insert, related...)
}

// RemoveBooksG relationships from objects passed in.
// Removes related items from R.Books (uses pointer comparison, removal does not keep order)
// Sets related.R.Accounts.
// Uses the global database handle.
func (o *Account) RemoveBooksG(related ...*Book) error {
	return o.RemoveBooks(boil.GetDB(), related...)
}

// RemoveBooks relationships from objects passed in.
// Removes related items from R.Books (uses pointer comparison, removal does not keep order)
// Sets related.R.Accounts.
func (o *Account) RemoveBooks(exec boil.Executor, related ...*Book) error {
	var err error
	query := fmt.Sprintf(
		"delete from \"book_star\" where \"account_id\" = $1 and \"book_id\" in (%s)",
		strmangle.Placeholders(dialect.UseIndexPlaceholders, len(related), 2, 1),
	)
	values := []interface{}{o.ID}
	for _, rel := range related {
		values = append(values, rel.ID)
	}

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, query)
		fmt.Fprintln(boil.DebugWriter, values)
	}
	_, err = exec.Exec(query, values...)
	if err != nil {
		return errors.Wrap(err, "failed to remove relationships before set")
	}
	removeBooksFromAccountsSlice(o, related)
	if o.R == nil {
		return nil
	}

	for _, rel := range related {
		for i, ri := range o.R.Books {
			if rel != ri {
				continue
			}

			ln := len(o.R.Books)
			if ln > 1 && i < ln-1 {
				o.R.Books[i] = o.R.Books[ln-1]
			}
			o.R.Books = o.R.Books[:ln-1]
			break
		}
	}

	return nil
}

func removeBooksFromAccountsSlice(o *Account, related []*Book) {
	for _, rel := range related {
		if rel.R == nil {
			continue
		}
		for i, ri := range rel.R.Accounts {
			if o.ID != ri.ID {
				continue
			}

			ln := len(rel.R.Accounts)
			if ln > 1 && i < ln-1 {
				rel.R.Accounts[i] = rel.R.Accounts[ln-1]
			}
			rel.R.Accounts = rel.R.Accounts[:ln-1]
			break
		}
	}
}

// Accounts retrieves all the records using an executor.
func Accounts(mods ...qm.QueryMod) accountQuery {
	mods = append(mods, qm.From("\"account\""))
	return accountQuery{NewQuery(mods...)}
}

// FindAccountG retrieves a single record by ID.
func FindAccountG(iD string, selectCols ...string) (*Account, error) {
	return FindAccount(boil.GetDB(), iD, selectCols...)
}

// FindAccount retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindAccount(exec boil.Executor, iD string, selectCols ...string) (*Account, error) {
	accountObj := &Account{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"account\" where \"id\"=$1", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(nil, exec, accountObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from account")
	}

	return accountObj, nil
}

// InsertG a single record. See Insert for whitelist behavior description.
func (o *Account) InsertG(columns boil.Columns) error {
	return o.Insert(boil.GetDB(), columns)
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *Account) Insert(exec boil.Executor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no account provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(accountColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	accountInsertCacheMut.RLock()
	cache, cached := accountInsertCache[key]
	accountInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			accountAllColumns,
			accountColumnsWithDefault,
			accountColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(accountType, accountMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(accountType, accountMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"account\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"account\" %sDEFAULT VALUES%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			queryReturning = fmt.Sprintf(" RETURNING \"%s\"", strings.Join(returnColumns, "\",\""))
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
		return errors.Wrap(err, "models: unable to insert into account")
	}

	if !cached {
		accountInsertCacheMut.Lock()
		accountInsertCache[key] = cache
		accountInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(exec)
}

// UpdateG a single Account record using the global executor.
// See Update for more documentation.
func (o *Account) UpdateG(columns boil.Columns) (int64, error) {
	return o.Update(boil.GetDB(), columns)
}

// Update uses an executor to update the Account.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *Account) Update(exec boil.Executor, columns boil.Columns) (int64, error) {
	var err error
	if err = o.doBeforeUpdateHooks(exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	accountUpdateCacheMut.RLock()
	cache, cached := accountUpdateCache[key]
	accountUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			accountAllColumns,
			accountPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update account, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"account\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, accountPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(accountType, accountMapping, append(wl, accountPrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "models: unable to update account row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for account")
	}

	if !cached {
		accountUpdateCacheMut.Lock()
		accountUpdateCache[key] = cache
		accountUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(exec)
}

// UpdateAllG updates all rows with the specified column values.
func (q accountQuery) UpdateAllG(cols M) (int64, error) {
	return q.UpdateAll(boil.GetDB(), cols)
}

// UpdateAll updates all rows with the specified column values.
func (q accountQuery) UpdateAll(exec boil.Executor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.Exec(exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for account")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for account")
	}

	return rowsAff, nil
}

// UpdateAllG updates all rows with the specified column values.
func (o AccountSlice) UpdateAllG(cols M) (int64, error) {
	return o.UpdateAll(boil.GetDB(), cols)
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o AccountSlice) UpdateAll(exec boil.Executor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), accountPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"account\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, accountPrimaryKeyColumns, len(o)))

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}
	result, err := exec.Exec(sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in account slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all account")
	}
	return rowsAff, nil
}

// UpsertG attempts an insert, and does an update or ignore on conflict.
func (o *Account) UpsertG(updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	return o.Upsert(boil.GetDB(), updateOnConflict, conflictColumns, updateColumns, insertColumns)
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *Account) Upsert(exec boil.Executor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no account provided for upsert")
	}

	if err := o.doBeforeUpsertHooks(exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(accountColumnsWithDefault, o)

	// Build cache key in-line uglily - mysql vs psql problems
	buf := strmangle.GetBuffer()
	if updateOnConflict {
		buf.WriteByte('t')
	} else {
		buf.WriteByte('f')
	}
	buf.WriteByte('.')
	for _, c := range conflictColumns {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
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

	accountUpsertCacheMut.RLock()
	cache, cached := accountUpsertCache[key]
	accountUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			accountAllColumns,
			accountColumnsWithDefault,
			accountColumnsWithoutDefault,
			nzDefaults,
		)
		update := updateColumns.UpdateColumnSet(
			accountAllColumns,
			accountPrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("models: unable to upsert account, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(accountPrimaryKeyColumns))
			copy(conflict, accountPrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryPostgres(dialect, "\"account\"", updateOnConflict, ret, update, conflict, insert)

		cache.valueMapping, err = queries.BindMapping(accountType, accountMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(accountType, accountMapping, ret)
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
			err = nil // Postgres doesn't return anything when there's no update
		}
	} else {
		_, err = exec.Exec(cache.query, vals...)
	}
	if err != nil {
		return errors.Wrap(err, "models: unable to upsert account")
	}

	if !cached {
		accountUpsertCacheMut.Lock()
		accountUpsertCache[key] = cache
		accountUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(exec)
}

// DeleteG deletes a single Account record.
// DeleteG will match against the primary key column to find the record to delete.
func (o *Account) DeleteG() (int64, error) {
	return o.Delete(boil.GetDB())
}

// Delete deletes a single Account record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *Account) Delete(exec boil.Executor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no Account provided for delete")
	}

	if err := o.doBeforeDeleteHooks(exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), accountPrimaryKeyMapping)
	sql := "DELETE FROM \"account\" WHERE \"id\"=$1"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}
	result, err := exec.Exec(sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from account")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for account")
	}

	if err := o.doAfterDeleteHooks(exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

func (q accountQuery) DeleteAllG() (int64, error) {
	return q.DeleteAll(boil.GetDB())
}

// DeleteAll deletes all matching rows.
func (q accountQuery) DeleteAll(exec boil.Executor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no accountQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.Exec(exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from account")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for account")
	}

	return rowsAff, nil
}

// DeleteAllG deletes all rows in the slice.
func (o AccountSlice) DeleteAllG() (int64, error) {
	return o.DeleteAll(boil.GetDB())
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o AccountSlice) DeleteAll(exec boil.Executor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(accountBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), accountPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"account\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, accountPrimaryKeyColumns, len(o))

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args)
	}
	result, err := exec.Exec(sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from account slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for account")
	}

	if len(accountAfterDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterDeleteHooks(exec); err != nil {
				return 0, err
			}
		}
	}

	return rowsAff, nil
}

// ReloadG refetches the object from the database using the primary keys.
func (o *Account) ReloadG() error {
	if o == nil {
		return errors.New("models: no Account provided for reload")
	}

	return o.Reload(boil.GetDB())
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *Account) Reload(exec boil.Executor) error {
	ret, err := FindAccount(exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAllG refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *AccountSlice) ReloadAllG() error {
	if o == nil {
		return errors.New("models: empty AccountSlice provided for reload all")
	}

	return o.ReloadAll(boil.GetDB())
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *AccountSlice) ReloadAll(exec boil.Executor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := AccountSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), accountPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"account\".* FROM \"account\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, accountPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(nil, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in AccountSlice")
	}

	*o = slice

	return nil
}

// AccountExistsG checks if the Account row exists.
func AccountExistsG(iD string) (bool, error) {
	return AccountExists(boil.GetDB(), iD)
}

// AccountExists checks if the Account row exists.
func AccountExists(exec boil.Executor, iD string) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"account\" where \"id\"=$1 limit 1)"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, iD)
	}
	row := exec.QueryRow(sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if account exists")
	}

	return exists, nil
}
