// Code generated by SQLBoiler 3.6.1 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package record

import (
	"context"
	"database/sql"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/friendsofgo/errors"
	"github.com/volatiletech/sqlboiler/boil"
	"github.com/volatiletech/sqlboiler/queries"
	"github.com/volatiletech/sqlboiler/queries/qm"
	"github.com/volatiletech/sqlboiler/queries/qmhelper"
	"github.com/volatiletech/sqlboiler/strmangle"
)

// BookmarkTagRelation is an object representing the database table.
type BookmarkTagRelation struct {
	ID         uint64 `boil:"id" json:"id" toml:"id" yaml:"id"`
	BookmarkID uint64 `boil:"bookmark_id" json:"bookmark_id" toml:"bookmark_id" yaml:"bookmark_id"`
	TagID      uint64 `boil:"tag_id" json:"tag_id" toml:"tag_id" yaml:"tag_id"`

	R *bookmarkTagRelationR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L bookmarkTagRelationL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var BookmarkTagRelationColumns = struct {
	ID         string
	BookmarkID string
	TagID      string
}{
	ID:         "id",
	BookmarkID: "bookmark_id",
	TagID:      "tag_id",
}

// Generated where

var BookmarkTagRelationWhere = struct {
	ID         whereHelperuint64
	BookmarkID whereHelperuint64
	TagID      whereHelperuint64
}{
	ID:         whereHelperuint64{field: "`bookmark_tag_relation`.`id`"},
	BookmarkID: whereHelperuint64{field: "`bookmark_tag_relation`.`bookmark_id`"},
	TagID:      whereHelperuint64{field: "`bookmark_tag_relation`.`tag_id`"},
}

// BookmarkTagRelationRels is where relationship names are stored.
var BookmarkTagRelationRels = struct {
	Bookmark string
	Tag      string
}{
	Bookmark: "Bookmark",
	Tag:      "Tag",
}

// bookmarkTagRelationR is where relationships are stored.
type bookmarkTagRelationR struct {
	Bookmark *Bookmark
	Tag      *Tag
}

// NewStruct creates a new relationship struct
func (*bookmarkTagRelationR) NewStruct() *bookmarkTagRelationR {
	return &bookmarkTagRelationR{}
}

// bookmarkTagRelationL is where Load methods for each relationship are stored.
type bookmarkTagRelationL struct{}

var (
	bookmarkTagRelationAllColumns            = []string{"id", "bookmark_id", "tag_id"}
	bookmarkTagRelationColumnsWithoutDefault = []string{"bookmark_id", "tag_id"}
	bookmarkTagRelationColumnsWithDefault    = []string{"id"}
	bookmarkTagRelationPrimaryKeyColumns     = []string{"id"}
)

type (
	// BookmarkTagRelationSlice is an alias for a slice of pointers to BookmarkTagRelation.
	// This should generally be used opposed to []BookmarkTagRelation.
	BookmarkTagRelationSlice []*BookmarkTagRelation
	// BookmarkTagRelationHook is the signature for custom BookmarkTagRelation hook methods
	BookmarkTagRelationHook func(context.Context, boil.ContextExecutor, *BookmarkTagRelation) error

	bookmarkTagRelationQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	bookmarkTagRelationType                 = reflect.TypeOf(&BookmarkTagRelation{})
	bookmarkTagRelationMapping              = queries.MakeStructMapping(bookmarkTagRelationType)
	bookmarkTagRelationPrimaryKeyMapping, _ = queries.BindMapping(bookmarkTagRelationType, bookmarkTagRelationMapping, bookmarkTagRelationPrimaryKeyColumns)
	bookmarkTagRelationInsertCacheMut       sync.RWMutex
	bookmarkTagRelationInsertCache          = make(map[string]insertCache)
	bookmarkTagRelationUpdateCacheMut       sync.RWMutex
	bookmarkTagRelationUpdateCache          = make(map[string]updateCache)
	bookmarkTagRelationUpsertCacheMut       sync.RWMutex
	bookmarkTagRelationUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var bookmarkTagRelationBeforeInsertHooks []BookmarkTagRelationHook
var bookmarkTagRelationBeforeUpdateHooks []BookmarkTagRelationHook
var bookmarkTagRelationBeforeDeleteHooks []BookmarkTagRelationHook
var bookmarkTagRelationBeforeUpsertHooks []BookmarkTagRelationHook

var bookmarkTagRelationAfterInsertHooks []BookmarkTagRelationHook
var bookmarkTagRelationAfterSelectHooks []BookmarkTagRelationHook
var bookmarkTagRelationAfterUpdateHooks []BookmarkTagRelationHook
var bookmarkTagRelationAfterDeleteHooks []BookmarkTagRelationHook
var bookmarkTagRelationAfterUpsertHooks []BookmarkTagRelationHook

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *BookmarkTagRelation) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range bookmarkTagRelationBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *BookmarkTagRelation) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range bookmarkTagRelationBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *BookmarkTagRelation) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range bookmarkTagRelationBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *BookmarkTagRelation) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range bookmarkTagRelationBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *BookmarkTagRelation) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range bookmarkTagRelationAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterSelectHooks executes all "after Select" hooks.
func (o *BookmarkTagRelation) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range bookmarkTagRelationAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *BookmarkTagRelation) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range bookmarkTagRelationAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *BookmarkTagRelation) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range bookmarkTagRelationAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *BookmarkTagRelation) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range bookmarkTagRelationAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddBookmarkTagRelationHook registers your hook function for all future operations.
func AddBookmarkTagRelationHook(hookPoint boil.HookPoint, bookmarkTagRelationHook BookmarkTagRelationHook) {
	switch hookPoint {
	case boil.BeforeInsertHook:
		bookmarkTagRelationBeforeInsertHooks = append(bookmarkTagRelationBeforeInsertHooks, bookmarkTagRelationHook)
	case boil.BeforeUpdateHook:
		bookmarkTagRelationBeforeUpdateHooks = append(bookmarkTagRelationBeforeUpdateHooks, bookmarkTagRelationHook)
	case boil.BeforeDeleteHook:
		bookmarkTagRelationBeforeDeleteHooks = append(bookmarkTagRelationBeforeDeleteHooks, bookmarkTagRelationHook)
	case boil.BeforeUpsertHook:
		bookmarkTagRelationBeforeUpsertHooks = append(bookmarkTagRelationBeforeUpsertHooks, bookmarkTagRelationHook)
	case boil.AfterInsertHook:
		bookmarkTagRelationAfterInsertHooks = append(bookmarkTagRelationAfterInsertHooks, bookmarkTagRelationHook)
	case boil.AfterSelectHook:
		bookmarkTagRelationAfterSelectHooks = append(bookmarkTagRelationAfterSelectHooks, bookmarkTagRelationHook)
	case boil.AfterUpdateHook:
		bookmarkTagRelationAfterUpdateHooks = append(bookmarkTagRelationAfterUpdateHooks, bookmarkTagRelationHook)
	case boil.AfterDeleteHook:
		bookmarkTagRelationAfterDeleteHooks = append(bookmarkTagRelationAfterDeleteHooks, bookmarkTagRelationHook)
	case boil.AfterUpsertHook:
		bookmarkTagRelationAfterUpsertHooks = append(bookmarkTagRelationAfterUpsertHooks, bookmarkTagRelationHook)
	}
}

// One returns a single bookmarkTagRelation record from the query.
func (q bookmarkTagRelationQuery) One(ctx context.Context, exec boil.ContextExecutor) (*BookmarkTagRelation, error) {
	o := &BookmarkTagRelation{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "record: failed to execute a one query for bookmark_tag_relation")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all BookmarkTagRelation records from the query.
func (q bookmarkTagRelationQuery) All(ctx context.Context, exec boil.ContextExecutor) (BookmarkTagRelationSlice, error) {
	var o []*BookmarkTagRelation

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "record: failed to assign all query results to BookmarkTagRelation slice")
	}

	if len(bookmarkTagRelationAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all BookmarkTagRelation records in the query.
func (q bookmarkTagRelationQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "record: failed to count bookmark_tag_relation rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q bookmarkTagRelationQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "record: failed to check if bookmark_tag_relation exists")
	}

	return count > 0, nil
}

// Bookmark pointed to by the foreign key.
func (o *BookmarkTagRelation) Bookmark(mods ...qm.QueryMod) bookmarkQuery {
	queryMods := []qm.QueryMod{
		qm.Where("`id` = ?", o.BookmarkID),
	}

	queryMods = append(queryMods, mods...)

	query := Bookmarks(queryMods...)
	queries.SetFrom(query.Query, "`bookmark`")

	return query
}

// Tag pointed to by the foreign key.
func (o *BookmarkTagRelation) Tag(mods ...qm.QueryMod) tagQuery {
	queryMods := []qm.QueryMod{
		qm.Where("`id` = ?", o.TagID),
	}

	queryMods = append(queryMods, mods...)

	query := Tags(queryMods...)
	queries.SetFrom(query.Query, "`tag`")

	return query
}

// LoadBookmark allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (bookmarkTagRelationL) LoadBookmark(ctx context.Context, e boil.ContextExecutor, singular bool, maybeBookmarkTagRelation interface{}, mods queries.Applicator) error {
	var slice []*BookmarkTagRelation
	var object *BookmarkTagRelation

	if singular {
		object = maybeBookmarkTagRelation.(*BookmarkTagRelation)
	} else {
		slice = *maybeBookmarkTagRelation.(*[]*BookmarkTagRelation)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &bookmarkTagRelationR{}
		}
		args = append(args, object.BookmarkID)

	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &bookmarkTagRelationR{}
			}

			for _, a := range args {
				if a == obj.BookmarkID {
					continue Outer
				}
			}

			args = append(args, obj.BookmarkID)

		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(qm.From(`bookmark`), qm.WhereIn(`bookmark.id in ?`, args...))
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load Bookmark")
	}

	var resultSlice []*Bookmark
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice Bookmark")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results of eager load for bookmark")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for bookmark")
	}

	if len(bookmarkTagRelationAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(ctx, e); err != nil {
				return err
			}
		}
	}

	if len(resultSlice) == 0 {
		return nil
	}

	if singular {
		foreign := resultSlice[0]
		object.R.Bookmark = foreign
		if foreign.R == nil {
			foreign.R = &bookmarkR{}
		}
		foreign.R.BookmarkTagRelations = append(foreign.R.BookmarkTagRelations, object)
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if local.BookmarkID == foreign.ID {
				local.R.Bookmark = foreign
				if foreign.R == nil {
					foreign.R = &bookmarkR{}
				}
				foreign.R.BookmarkTagRelations = append(foreign.R.BookmarkTagRelations, local)
				break
			}
		}
	}

	return nil
}

// LoadTag allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (bookmarkTagRelationL) LoadTag(ctx context.Context, e boil.ContextExecutor, singular bool, maybeBookmarkTagRelation interface{}, mods queries.Applicator) error {
	var slice []*BookmarkTagRelation
	var object *BookmarkTagRelation

	if singular {
		object = maybeBookmarkTagRelation.(*BookmarkTagRelation)
	} else {
		slice = *maybeBookmarkTagRelation.(*[]*BookmarkTagRelation)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &bookmarkTagRelationR{}
		}
		args = append(args, object.TagID)

	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &bookmarkTagRelationR{}
			}

			for _, a := range args {
				if a == obj.TagID {
					continue Outer
				}
			}

			args = append(args, obj.TagID)

		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(qm.From(`tag`), qm.WhereIn(`tag.id in ?`, args...))
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load Tag")
	}

	var resultSlice []*Tag
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice Tag")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results of eager load for tag")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for tag")
	}

	if len(bookmarkTagRelationAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(ctx, e); err != nil {
				return err
			}
		}
	}

	if len(resultSlice) == 0 {
		return nil
	}

	if singular {
		foreign := resultSlice[0]
		object.R.Tag = foreign
		if foreign.R == nil {
			foreign.R = &tagR{}
		}
		foreign.R.BookmarkTagRelations = append(foreign.R.BookmarkTagRelations, object)
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if local.TagID == foreign.ID {
				local.R.Tag = foreign
				if foreign.R == nil {
					foreign.R = &tagR{}
				}
				foreign.R.BookmarkTagRelations = append(foreign.R.BookmarkTagRelations, local)
				break
			}
		}
	}

	return nil
}

// SetBookmark of the bookmarkTagRelation to the related item.
// Sets o.R.Bookmark to related.
// Adds o to related.R.BookmarkTagRelations.
func (o *BookmarkTagRelation) SetBookmark(ctx context.Context, exec boil.ContextExecutor, insert bool, related *Bookmark) error {
	var err error
	if insert {
		if err = related.Insert(ctx, exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE `bookmark_tag_relation` SET %s WHERE %s",
		strmangle.SetParamNames("`", "`", 0, []string{"bookmark_id"}),
		strmangle.WhereClause("`", "`", 0, bookmarkTagRelationPrimaryKeyColumns),
	)
	values := []interface{}{related.ID, o.ID}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, updateQuery)
		fmt.Fprintln(writer, values)
	}
	if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	o.BookmarkID = related.ID
	if o.R == nil {
		o.R = &bookmarkTagRelationR{
			Bookmark: related,
		}
	} else {
		o.R.Bookmark = related
	}

	if related.R == nil {
		related.R = &bookmarkR{
			BookmarkTagRelations: BookmarkTagRelationSlice{o},
		}
	} else {
		related.R.BookmarkTagRelations = append(related.R.BookmarkTagRelations, o)
	}

	return nil
}

// SetTag of the bookmarkTagRelation to the related item.
// Sets o.R.Tag to related.
// Adds o to related.R.BookmarkTagRelations.
func (o *BookmarkTagRelation) SetTag(ctx context.Context, exec boil.ContextExecutor, insert bool, related *Tag) error {
	var err error
	if insert {
		if err = related.Insert(ctx, exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE `bookmark_tag_relation` SET %s WHERE %s",
		strmangle.SetParamNames("`", "`", 0, []string{"tag_id"}),
		strmangle.WhereClause("`", "`", 0, bookmarkTagRelationPrimaryKeyColumns),
	)
	values := []interface{}{related.ID, o.ID}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, updateQuery)
		fmt.Fprintln(writer, values)
	}
	if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	o.TagID = related.ID
	if o.R == nil {
		o.R = &bookmarkTagRelationR{
			Tag: related,
		}
	} else {
		o.R.Tag = related
	}

	if related.R == nil {
		related.R = &tagR{
			BookmarkTagRelations: BookmarkTagRelationSlice{o},
		}
	} else {
		related.R.BookmarkTagRelations = append(related.R.BookmarkTagRelations, o)
	}

	return nil
}

// BookmarkTagRelations retrieves all the records using an executor.
func BookmarkTagRelations(mods ...qm.QueryMod) bookmarkTagRelationQuery {
	mods = append(mods, qm.From("`bookmark_tag_relation`"))
	return bookmarkTagRelationQuery{NewQuery(mods...)}
}

// FindBookmarkTagRelation retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindBookmarkTagRelation(ctx context.Context, exec boil.ContextExecutor, iD uint64, selectCols ...string) (*BookmarkTagRelation, error) {
	bookmarkTagRelationObj := &BookmarkTagRelation{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from `bookmark_tag_relation` where `id`=?", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, bookmarkTagRelationObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "record: unable to select from bookmark_tag_relation")
	}

	return bookmarkTagRelationObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *BookmarkTagRelation) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("record: no bookmark_tag_relation provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(bookmarkTagRelationColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	bookmarkTagRelationInsertCacheMut.RLock()
	cache, cached := bookmarkTagRelationInsertCache[key]
	bookmarkTagRelationInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			bookmarkTagRelationAllColumns,
			bookmarkTagRelationColumnsWithDefault,
			bookmarkTagRelationColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(bookmarkTagRelationType, bookmarkTagRelationMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(bookmarkTagRelationType, bookmarkTagRelationMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO `bookmark_tag_relation` (`%s`) %%sVALUES (%s)%%s", strings.Join(wl, "`,`"), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO `bookmark_tag_relation` () VALUES ()%s%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			cache.retQuery = fmt.Sprintf("SELECT `%s` FROM `bookmark_tag_relation` WHERE %s", strings.Join(returnColumns, "`,`"), strmangle.WhereClause("`", "`", 0, bookmarkTagRelationPrimaryKeyColumns))
		}

		cache.query = fmt.Sprintf(cache.query, queryOutput, queryReturning)
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, vals)
	}
	result, err := exec.ExecContext(ctx, cache.query, vals...)

	if err != nil {
		return errors.Wrap(err, "record: unable to insert into bookmark_tag_relation")
	}

	var lastID int64
	var identifierCols []interface{}

	if len(cache.retMapping) == 0 {
		goto CacheNoHooks
	}

	lastID, err = result.LastInsertId()
	if err != nil {
		return ErrSyncFail
	}

	o.ID = uint64(lastID)
	if lastID != 0 && len(cache.retMapping) == 1 && cache.retMapping[0] == bookmarkTagRelationMapping["id"] {
		goto CacheNoHooks
	}

	identifierCols = []interface{}{
		o.ID,
	}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.retQuery)
		fmt.Fprintln(writer, identifierCols...)
	}
	err = exec.QueryRowContext(ctx, cache.retQuery, identifierCols...).Scan(queries.PtrsFromMapping(value, cache.retMapping)...)
	if err != nil {
		return errors.Wrap(err, "record: unable to populate default values for bookmark_tag_relation")
	}

CacheNoHooks:
	if !cached {
		bookmarkTagRelationInsertCacheMut.Lock()
		bookmarkTagRelationInsertCache[key] = cache
		bookmarkTagRelationInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the BookmarkTagRelation.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *BookmarkTagRelation) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	bookmarkTagRelationUpdateCacheMut.RLock()
	cache, cached := bookmarkTagRelationUpdateCache[key]
	bookmarkTagRelationUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			bookmarkTagRelationAllColumns,
			bookmarkTagRelationPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("record: unable to update bookmark_tag_relation, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE `bookmark_tag_relation` SET %s WHERE %s",
			strmangle.SetParamNames("`", "`", 0, wl),
			strmangle.WhereClause("`", "`", 0, bookmarkTagRelationPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(bookmarkTagRelationType, bookmarkTagRelationMapping, append(wl, bookmarkTagRelationPrimaryKeyColumns...))
		if err != nil {
			return 0, err
		}
	}

	values := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), cache.valueMapping)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, values)
	}
	var result sql.Result
	result, err = exec.ExecContext(ctx, cache.query, values...)
	if err != nil {
		return 0, errors.Wrap(err, "record: unable to update bookmark_tag_relation row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "record: failed to get rows affected by update for bookmark_tag_relation")
	}

	if !cached {
		bookmarkTagRelationUpdateCacheMut.Lock()
		bookmarkTagRelationUpdateCache[key] = cache
		bookmarkTagRelationUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q bookmarkTagRelationQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "record: unable to update all for bookmark_tag_relation")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "record: unable to retrieve rows affected for bookmark_tag_relation")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o BookmarkTagRelationSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	ln := int64(len(o))
	if ln == 0 {
		return 0, nil
	}

	if len(cols) == 0 {
		return 0, errors.New("record: update all requires at least one column argument")
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), bookmarkTagRelationPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE `bookmark_tag_relation` SET %s WHERE %s",
		strmangle.SetParamNames("`", "`", 0, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, bookmarkTagRelationPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "record: unable to update all in bookmarkTagRelation slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "record: unable to retrieve rows affected all in update all bookmarkTagRelation")
	}
	return rowsAff, nil
}

var mySQLBookmarkTagRelationUniqueColumns = []string{
	"id",
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *BookmarkTagRelation) Upsert(ctx context.Context, exec boil.ContextExecutor, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("record: no bookmark_tag_relation provided for upsert")
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(bookmarkTagRelationColumnsWithDefault, o)
	nzUniques := queries.NonZeroDefaultSet(mySQLBookmarkTagRelationUniqueColumns, o)

	if len(nzUniques) == 0 {
		return errors.New("cannot upsert with a table that cannot conflict on a unique column")
	}

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
	buf.WriteByte('.')
	for _, c := range nzUniques {
		buf.WriteString(c)
	}
	key := buf.String()
	strmangle.PutBuffer(buf)

	bookmarkTagRelationUpsertCacheMut.RLock()
	cache, cached := bookmarkTagRelationUpsertCache[key]
	bookmarkTagRelationUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			bookmarkTagRelationAllColumns,
			bookmarkTagRelationColumnsWithDefault,
			bookmarkTagRelationColumnsWithoutDefault,
			nzDefaults,
		)
		update := updateColumns.UpdateColumnSet(
			bookmarkTagRelationAllColumns,
			bookmarkTagRelationPrimaryKeyColumns,
		)

		if len(update) == 0 {
			return errors.New("record: unable to upsert bookmark_tag_relation, could not build update column list")
		}

		ret = strmangle.SetComplement(ret, nzUniques)
		cache.query = buildUpsertQueryMySQL(dialect, "bookmark_tag_relation", update, insert)
		cache.retQuery = fmt.Sprintf(
			"SELECT %s FROM `bookmark_tag_relation` WHERE %s",
			strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, ret), ","),
			strmangle.WhereClause("`", "`", 0, nzUniques),
		)

		cache.valueMapping, err = queries.BindMapping(bookmarkTagRelationType, bookmarkTagRelationMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(bookmarkTagRelationType, bookmarkTagRelationMapping, ret)
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

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, vals)
	}
	result, err := exec.ExecContext(ctx, cache.query, vals...)

	if err != nil {
		return errors.Wrap(err, "record: unable to upsert for bookmark_tag_relation")
	}

	var lastID int64
	var uniqueMap []uint64
	var nzUniqueCols []interface{}

	if len(cache.retMapping) == 0 {
		goto CacheNoHooks
	}

	lastID, err = result.LastInsertId()
	if err != nil {
		return ErrSyncFail
	}

	o.ID = uint64(lastID)
	if lastID != 0 && len(cache.retMapping) == 1 && cache.retMapping[0] == bookmarkTagRelationMapping["id"] {
		goto CacheNoHooks
	}

	uniqueMap, err = queries.BindMapping(bookmarkTagRelationType, bookmarkTagRelationMapping, nzUniques)
	if err != nil {
		return errors.Wrap(err, "record: unable to retrieve unique values for bookmark_tag_relation")
	}
	nzUniqueCols = queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), uniqueMap)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.retQuery)
		fmt.Fprintln(writer, nzUniqueCols...)
	}
	err = exec.QueryRowContext(ctx, cache.retQuery, nzUniqueCols...).Scan(returns...)
	if err != nil {
		return errors.Wrap(err, "record: unable to populate default values for bookmark_tag_relation")
	}

CacheNoHooks:
	if !cached {
		bookmarkTagRelationUpsertCacheMut.Lock()
		bookmarkTagRelationUpsertCache[key] = cache
		bookmarkTagRelationUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single BookmarkTagRelation record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *BookmarkTagRelation) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("record: no BookmarkTagRelation provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), bookmarkTagRelationPrimaryKeyMapping)
	sql := "DELETE FROM `bookmark_tag_relation` WHERE `id`=?"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "record: unable to delete from bookmark_tag_relation")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "record: failed to get rows affected by delete for bookmark_tag_relation")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q bookmarkTagRelationQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("record: no bookmarkTagRelationQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "record: unable to delete all from bookmark_tag_relation")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "record: failed to get rows affected by deleteall for bookmark_tag_relation")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o BookmarkTagRelationSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(bookmarkTagRelationBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), bookmarkTagRelationPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM `bookmark_tag_relation` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, bookmarkTagRelationPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "record: unable to delete all from bookmarkTagRelation slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "record: failed to get rows affected by deleteall for bookmark_tag_relation")
	}

	if len(bookmarkTagRelationAfterDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	return rowsAff, nil
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *BookmarkTagRelation) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindBookmarkTagRelation(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *BookmarkTagRelationSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := BookmarkTagRelationSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), bookmarkTagRelationPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT `bookmark_tag_relation`.* FROM `bookmark_tag_relation` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, bookmarkTagRelationPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "record: unable to reload all in BookmarkTagRelationSlice")
	}

	*o = slice

	return nil
}

// BookmarkTagRelationExists checks if the BookmarkTagRelation row exists.
func BookmarkTagRelationExists(ctx context.Context, exec boil.ContextExecutor, iD uint64) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from `bookmark_tag_relation` where `id`=? limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "record: unable to check if bookmark_tag_relation exists")
	}

	return exists, nil
}
