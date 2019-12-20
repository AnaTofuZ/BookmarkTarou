// Code generated by SQLBoiler 3.6.1 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package record

import "testing"

// This test suite runs each operation test in parallel.
// Example, if your database has 3 tables, the suite will run:
// table1, table2 and table3 Delete in parallel
// table1, table2 and table3 Insert in parallel, and so forth.
// It does NOT run each operation group in parallel.
// Separating the tests thusly grants avoidance of Postgres deadlocks.
func TestParent(t *testing.T) {
	t.Run("BookmarkTagRelations", testBookmarkTagRelations)
	t.Run("Bookmarks", testBookmarks)
	t.Run("Entries", testEntries)
	t.Run("Tags", testTags)
	t.Run("Users", testUsers)
}

func TestDelete(t *testing.T) {
	t.Run("BookmarkTagRelations", testBookmarkTagRelationsDelete)
	t.Run("Bookmarks", testBookmarksDelete)
	t.Run("Entries", testEntriesDelete)
	t.Run("Tags", testTagsDelete)
	t.Run("Users", testUsersDelete)
}

func TestQueryDeleteAll(t *testing.T) {
	t.Run("BookmarkTagRelations", testBookmarkTagRelationsQueryDeleteAll)
	t.Run("Bookmarks", testBookmarksQueryDeleteAll)
	t.Run("Entries", testEntriesQueryDeleteAll)
	t.Run("Tags", testTagsQueryDeleteAll)
	t.Run("Users", testUsersQueryDeleteAll)
}

func TestSliceDeleteAll(t *testing.T) {
	t.Run("BookmarkTagRelations", testBookmarkTagRelationsSliceDeleteAll)
	t.Run("Bookmarks", testBookmarksSliceDeleteAll)
	t.Run("Entries", testEntriesSliceDeleteAll)
	t.Run("Tags", testTagsSliceDeleteAll)
	t.Run("Users", testUsersSliceDeleteAll)
}

func TestExists(t *testing.T) {
	t.Run("BookmarkTagRelations", testBookmarkTagRelationsExists)
	t.Run("Bookmarks", testBookmarksExists)
	t.Run("Entries", testEntriesExists)
	t.Run("Tags", testTagsExists)
	t.Run("Users", testUsersExists)
}

func TestFind(t *testing.T) {
	t.Run("BookmarkTagRelations", testBookmarkTagRelationsFind)
	t.Run("Bookmarks", testBookmarksFind)
	t.Run("Entries", testEntriesFind)
	t.Run("Tags", testTagsFind)
	t.Run("Users", testUsersFind)
}

func TestBind(t *testing.T) {
	t.Run("BookmarkTagRelations", testBookmarkTagRelationsBind)
	t.Run("Bookmarks", testBookmarksBind)
	t.Run("Entries", testEntriesBind)
	t.Run("Tags", testTagsBind)
	t.Run("Users", testUsersBind)
}

func TestOne(t *testing.T) {
	t.Run("BookmarkTagRelations", testBookmarkTagRelationsOne)
	t.Run("Bookmarks", testBookmarksOne)
	t.Run("Entries", testEntriesOne)
	t.Run("Tags", testTagsOne)
	t.Run("Users", testUsersOne)
}

func TestAll(t *testing.T) {
	t.Run("BookmarkTagRelations", testBookmarkTagRelationsAll)
	t.Run("Bookmarks", testBookmarksAll)
	t.Run("Entries", testEntriesAll)
	t.Run("Tags", testTagsAll)
	t.Run("Users", testUsersAll)
}

func TestCount(t *testing.T) {
	t.Run("BookmarkTagRelations", testBookmarkTagRelationsCount)
	t.Run("Bookmarks", testBookmarksCount)
	t.Run("Entries", testEntriesCount)
	t.Run("Tags", testTagsCount)
	t.Run("Users", testUsersCount)
}

func TestHooks(t *testing.T) {
	t.Run("BookmarkTagRelations", testBookmarkTagRelationsHooks)
	t.Run("Bookmarks", testBookmarksHooks)
	t.Run("Entries", testEntriesHooks)
	t.Run("Tags", testTagsHooks)
	t.Run("Users", testUsersHooks)
}

func TestInsert(t *testing.T) {
	t.Run("BookmarkTagRelations", testBookmarkTagRelationsInsert)
	t.Run("BookmarkTagRelations", testBookmarkTagRelationsInsertWhitelist)
	t.Run("Bookmarks", testBookmarksInsert)
	t.Run("Bookmarks", testBookmarksInsertWhitelist)
	t.Run("Entries", testEntriesInsert)
	t.Run("Entries", testEntriesInsertWhitelist)
	t.Run("Tags", testTagsInsert)
	t.Run("Tags", testTagsInsertWhitelist)
	t.Run("Users", testUsersInsert)
	t.Run("Users", testUsersInsertWhitelist)
}

// TestToOne tests cannot be run in parallel
// or deadlocks can occur.
func TestToOne(t *testing.T) {
	t.Run("BookmarkTagRelationToBookmarkUsingBookmark", testBookmarkTagRelationToOneBookmarkUsingBookmark)
	t.Run("BookmarkTagRelationToTagUsingTag", testBookmarkTagRelationToOneTagUsingTag)
	t.Run("BookmarkToUserUsingUser", testBookmarkToOneUserUsingUser)
	t.Run("BookmarkToEntryUsingEntry", testBookmarkToOneEntryUsingEntry)
}

// TestOneToOne tests cannot be run in parallel
// or deadlocks can occur.
func TestOneToOne(t *testing.T) {}

// TestToMany tests cannot be run in parallel
// or deadlocks can occur.
func TestToMany(t *testing.T) {
	t.Run("BookmarkToBookmarkTagRelations", testBookmarkToManyBookmarkTagRelations)
	t.Run("EntryToBookmarks", testEntryToManyBookmarks)
	t.Run("TagToBookmarkTagRelations", testTagToManyBookmarkTagRelations)
	t.Run("UserToBookmarks", testUserToManyBookmarks)
}

// TestToOneSet tests cannot be run in parallel
// or deadlocks can occur.
func TestToOneSet(t *testing.T) {
	t.Run("BookmarkTagRelationToBookmarkUsingBookmarkTagRelations", testBookmarkTagRelationToOneSetOpBookmarkUsingBookmark)
	t.Run("BookmarkTagRelationToTagUsingBookmarkTagRelations", testBookmarkTagRelationToOneSetOpTagUsingTag)
	t.Run("BookmarkToUserUsingBookmarks", testBookmarkToOneSetOpUserUsingUser)
	t.Run("BookmarkToEntryUsingBookmarks", testBookmarkToOneSetOpEntryUsingEntry)
}

// TestToOneRemove tests cannot be run in parallel
// or deadlocks can occur.
func TestToOneRemove(t *testing.T) {}

// TestOneToOneSet tests cannot be run in parallel
// or deadlocks can occur.
func TestOneToOneSet(t *testing.T) {}

// TestOneToOneRemove tests cannot be run in parallel
// or deadlocks can occur.
func TestOneToOneRemove(t *testing.T) {}

// TestToManyAdd tests cannot be run in parallel
// or deadlocks can occur.
func TestToManyAdd(t *testing.T) {
	t.Run("BookmarkToBookmarkTagRelations", testBookmarkToManyAddOpBookmarkTagRelations)
	t.Run("EntryToBookmarks", testEntryToManyAddOpBookmarks)
	t.Run("TagToBookmarkTagRelations", testTagToManyAddOpBookmarkTagRelations)
	t.Run("UserToBookmarks", testUserToManyAddOpBookmarks)
}

// TestToManySet tests cannot be run in parallel
// or deadlocks can occur.
func TestToManySet(t *testing.T) {}

// TestToManyRemove tests cannot be run in parallel
// or deadlocks can occur.
func TestToManyRemove(t *testing.T) {}

func TestReload(t *testing.T) {
	t.Run("BookmarkTagRelations", testBookmarkTagRelationsReload)
	t.Run("Bookmarks", testBookmarksReload)
	t.Run("Entries", testEntriesReload)
	t.Run("Tags", testTagsReload)
	t.Run("Users", testUsersReload)
}

func TestReloadAll(t *testing.T) {
	t.Run("BookmarkTagRelations", testBookmarkTagRelationsReloadAll)
	t.Run("Bookmarks", testBookmarksReloadAll)
	t.Run("Entries", testEntriesReloadAll)
	t.Run("Tags", testTagsReloadAll)
	t.Run("Users", testUsersReloadAll)
}

func TestSelect(t *testing.T) {
	t.Run("BookmarkTagRelations", testBookmarkTagRelationsSelect)
	t.Run("Bookmarks", testBookmarksSelect)
	t.Run("Entries", testEntriesSelect)
	t.Run("Tags", testTagsSelect)
	t.Run("Users", testUsersSelect)
}

func TestUpdate(t *testing.T) {
	t.Run("BookmarkTagRelations", testBookmarkTagRelationsUpdate)
	t.Run("Bookmarks", testBookmarksUpdate)
	t.Run("Entries", testEntriesUpdate)
	t.Run("Tags", testTagsUpdate)
	t.Run("Users", testUsersUpdate)
}

func TestSliceUpdateAll(t *testing.T) {
	t.Run("BookmarkTagRelations", testBookmarkTagRelationsSliceUpdateAll)
	t.Run("Bookmarks", testBookmarksSliceUpdateAll)
	t.Run("Entries", testEntriesSliceUpdateAll)
	t.Run("Tags", testTagsSliceUpdateAll)
	t.Run("Users", testUsersSliceUpdateAll)
}
